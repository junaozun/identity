package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type IdentityChaincode struct {

}

func (t *IdentityChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *IdentityChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addIde"{
		return t.addIde(stub, args)        // 添加信息
	}else if fun == "queryIdeByCertNoAndName" {
		return t.queryIdeByCertNoAndName(stub, args)        // 根据证书编号及姓名查询信息
	}else if fun == "queryIdeInfoByIDNumber" {
		return t.queryIdeInfoByIDNumber(stub, args)    // 根据身份证号码及姓名查询详情
	}else if fun == "updateIdeByCertNo" {
		return t.updateIdeByCertNo(stub, args)        // 根据证书编号更新信息
	}else if fun == "delIdeByCertNo"{
		return t.delIdeByCertNo(stub, args)    // 根据证书编号删除信息
	}

	return shim.Error("指定的函数名称错误")

}

func main(){
	err := shim.Start(new(IdentityChaincode))
	if err != nil{
		fmt.Printf("启动链码时发生错误: %s", err)
	}
}
