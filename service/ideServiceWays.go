package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//调用链码向fabric账本中添加ide对象的信息,返回交易id和错误信息
func (t *ServiceSetup) SaveIde(ide Identity) (string, error) {

	eventID := "eventAddIde"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将ide对象序列化成为字节数组
	b, err := json.Marshal(ide)
	if err != nil {
		return "", fmt.Errorf("指定的ide对象序列化时发生错误")
	}
	//ChaincodeID：链码名称；Fcn：链码层中添加信息到账本的方法（addIde）；Args：链码中的addIde方法需要的两个参数（一个是序列化后的ide对象，另一个是eventName）
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addIde", Args: [][]byte{b, []byte(eventID)}}

	//业务层调用链码，只有Execute方法（执行操作）和Query方法（查询操作）
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

//根据证书编号和姓名查询ide对象的状态信息,由于链码中返回peer.response，其实就是字节数组
func (t *ServiceSetup) FindIdeByCertNoAndName(certNo, name string) ([]byte, error){
	//ChaincodeID：链码名称；Fcn：链码层中根据证书编号和姓名查询ide对象的状态信息的方法（queryIdeByCertNoAndName）；Args：链码层中的queryIdeByCertNoAndName方法需要的两个参数（一个是证书编号，另一个是学生的名字）
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryIdeByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}

	//业务层调用链码，只有Execute方法（执行操作）和Query方法（查询操作）
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err//返回空的字节数组
	}

	return respone.Payload, nil
}

//根据身份证号码查询详细信息
func (t *ServiceSetup) FindIdeInfoByIDNumber(IDNumber string) ([]byte, error){
	//ChaincodeID：链码名称；Fcn：链码层中根据身份证号码查询详细信息的方法（queryIdeInfoByIDNumber）；Args：链码层中的queryIdeInfoByIDNumber方法需要的一个参数（身份证号码）
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryIdeInfoByIDNumber", Args: [][]byte{[]byte(IDNumber)}}

	//业务层调用链码，只有Execute方法（执行操作）和Query方法（查询操作）
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}


//对账本中的数据进行修改,返回交易id
func (t *ServiceSetup) ModifyIde(ide Identity) (string, error) {

	eventID := "eventModifyIde"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将ide对象序列化成为字节数组
	b, err := json.Marshal(ide)
	if err != nil {
		return "", fmt.Errorf("指定的ide对象序列化时发生错误")
	}

	//ChaincodeID：链码名称；Fcn：链码层中对账本中的数据进行修改的方法（updateIdeByCertNo）；Args：链码层中的updateIdeByCertNo方法需要的两个个参数（一个是序列化后的ide对象，一个是eventname）
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateIdeByCertNo", Args: [][]byte{b, []byte(eventID)}}

	//业务层调用链码，只有Execute方法（执行操作）和Query方法（查询操作）
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) DelIde(IDNumber string) (string, error) {

	eventID := "eventDelIde"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "delIdeByCertNo", Args: [][]byte{[]byte(IDNumber), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

