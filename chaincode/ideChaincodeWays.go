package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const DOC_TYPE = "ideObj"

// 保存ide
// args: identity对象
func PutIde(stub shim.ChaincodeStubInterface, ide Identity) ([]byte, bool) {

	ide.ObjectType = DOC_TYPE

	b, err := json.Marshal(ide)
	if err != nil {
		return nil, false
	}

	// 保存ide状态
	err = stub.PutState(ide.IDNumber, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根据身份证号码查询信息状态
// args: IDNumber
func GetIdeInfo(stub shim.ChaincodeStubInterface, IDNumber string) (Identity, bool)  {
	var ide Identity
	// 根据身份证号码查询信息状态
	b, err := stub.GetState(IDNumber)
	if err != nil {
		return ide, false
	}

	if b == nil {
		return ide, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &ide)
	if err != nil {
		return ide, false
	}

	// 返回结果
	return ide, true
}

//couchDB富查询需要
// 根据指定的查询字符串实现富查询
func getIdeByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	//根据指定字符串查询，返回一个迭代器对象
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer  resultsIterator.Close()//延迟关闭迭代器对象


	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false

	//将查询结果从resultsIterator迭代器对象中查询出来，并使用逗号分割
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		//向buffer中写入一个对象后，加一个逗号。kv和kv之间用逗号相隔
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	//将buffer中的字符串转换成字节数组返回
	return buffer.Bytes(), nil

}

// 添加信息
// args[0]:ideObj,args[1]:eventName
func (t *IdentityChaincode) addIde(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var ide  Identity
	err := json.Unmarshal([]byte(args[0]), &ide)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 身份证号码必须唯一
	_, exist := GetIdeInfo(stub, ide.IDNumber)
	if exist {
		return shim.Error("要添加的身份证号码已存在")
	}

	_, bl := PutIde(stub, ide)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据证书编号及姓名查询信息
// args: CertNo, name
func (t *IdentityChaincode) queryIdeByCertNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	CertNo := args[0]
	name := args[1]

	// 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"ideObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)

	// 查询数据
	result, err := getIdeByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根据证书编号及姓名查询信息时发生错误")
	}
	if result == nil {
		return shim.Error("根据指定的证书编号及姓名没有查询到相关的信息")
	}
	return shim.Success(result)
}

// 根据身份证号码查询详情（溯源）
// args[0]: IDNumber
func (t *IdentityChaincode) queryIdeInfoByIDNumber(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询ide状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var ide Identity
	err = json.Unmarshal(b, &ide)
	if err != nil {
		return  shim.Error("反序列化ide信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(ide.IDNumber)
	if err != nil {
		return shim.Error("根据指定的身份证号码查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem//多个历史对象组成的数组
	var hiside Identity //历史变更对象中的一个对象详细信息

	for iterator.HasNext() {//判断iterator迭代器对象中下一条是否还有数据
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取ide的历史变更数据失败")
		}

		//一个历史记录对象
		var historyItem HistoryItem

		//hisData.TxId就是上传一次信息返回的交易ID
		historyItem.TxId = hisData.TxId

		//hisData.value就是上传的Identity对象
		json.Unmarshal(hisData.Value, &hiside)

		if hisData.Value == nil {
			var empty Identity
			historyItem.Identity = empty
		}else {
			historyItem.Identity = hiside
		}

		//将所有的历史记录historyItem组装进historys数组中
		historys = append(historys, historyItem)

	}

	//将historys历史所有数据赋值给Identity对象的Historys字段中
	ide.Historys = historys

	// 将ide对象序列化成字节数组返回
	result, err := json.Marshal(ide)
	if err != nil {
		return shim.Error("序列化ide信息时发生错误")
	}
	return shim.Success(result)
}

// 根据身份证号更新信息
// // args[0]:ideObj,args[1]:eventName
func (t *IdentityChaincode) updateIdeByCertNo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var ideinfo Identity//用户上传的新的信息

	//将用户上传上来的新信息反序列化
	err := json.Unmarshal([]byte(args[0]), &ideinfo)
	if err != nil {
		return  shim.Error("反序列化ide信息失败")
	}

	// 根据身份证号码查询信息
	result, bl := GetIdeInfo(stub, ideinfo.IDNumber)
	if !bl{
		return shim.Error("根据身份证号码查询信息时发生错误")
	}

	//只允许修改学历信息
	result.EnrollDate = ideinfo.EnrollDate
	result.GraduationDate = ideinfo.GraduationDate
	result.SchoolName = ideinfo.SchoolName
	result.Major = ideinfo.Major
	result.QuaType = ideinfo.QuaType
	result.Length = ideinfo.Length
	result.Mode = ideinfo.Mode
	result.Level = ideinfo.Level
	result.Graduation = ideinfo.Graduation
	result.CertNo = ideinfo.CertNo;

	//将修改后的信息存入账本中
	_, bl = PutIde(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

// 根据身份证号删除信息（暂不对外提供）
// args: IDNumber
func (t *IdentityChaincode) delIdeByCertNo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}
