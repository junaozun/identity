package main

import (
	"fmt"
	"identity/sdkInit"
	"identity/service"
	"identity/web"
	"identity/web/controller"
	"os"
)

const (
	configFile = "config.yaml"
	initialized = false
	IdeCC = "idecc"//链码名称
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "mychannel",//在channel.tx中查看当初创建时使用的通道名称
		ChannelConfig: os.Getenv("GOPATH") + "/src/identity/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.sxf.bjut.cn",//在docker-compose.yaml中查看orderer名称

		ChaincodeID: IdeCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "identity/chaincode/",
		UserName:"User1",
	}

	//创建sdk
	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	//创建通道
	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//安装实例化链码
	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//测试调用业务层方法//
	serviceSetup := service.ServiceSetup{
		ChaincodeID:IdeCC,
		Client:channelClient,
	}
//.........................................测试代码...................................................//
	/*ide := service.Identity{
		Name: "张三",
		Sex: "男",
		Nation: "汉",
		IDNumber: "101",
		NativePlace: "北京",
		BirthDay: "1991年01月01日",
		EnrollDate: "2009年9月",
		GraduationDate: "2013年7月",
		SchoolName: "中国政法大学",
		Major: "社会学",
		QuaType: "普通",
		Length: "四年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "毕业",
		CertNo: "111",
		Photo: "/static/photo/11.png",
	}

	ide2 := service.Identity{
		Name: "李四",
		Sex: "男",
		Nation: "汉",
		IDNumber: "102",
		NativePlace: "上海",
		BirthDay: "1992年02月01日",
		EnrollDate: "2010年9月",
		GraduationDate: "2014年7月",
		SchoolName: "中国人民大学",
		Major: "行政管理",
		QuaType: "普通",
		Length: "四年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "毕业",
		CertNo: "222",
		Photo: "/static/photo/22.png",
	}

	//保存ide对象状态至账本中
	msg, err := serviceSetup.SaveIde(ide)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveIde(ide2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	// 根据证书编号和姓名查询ide对象的状态信息
	result, err := serviceSetup.FindIdeByCertNoAndName("222","李四")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var ide service.Identity
		json.Unmarshal(result, &ide)
		fmt.Println("根据证书编号与姓名查询信息成功：")
		fmt.Println(ide)
	}

	// 根据身份证号码IDNumber查询详细信息
	result, err = serviceSetup.FindIdeInfoByIDNumber("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var ide service.Identity
		json.Unmarshal(result, &ide)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(ide)
	}

	// 修改更新信息
	info := service.Identity{
		Name: "张三",
		Sex: "男",
		Nation: "汉",
		IDNumber: "101",
		NativePlace: "北京",
		BirthDay: "1991年01月01日",
		EnrollDate: "2013年9月",
		GraduationDate: "2015年7月",
		SchoolName: "中国政法大学",
		Major: "社会学",
		QuaType: "普通",
		Length: "两年",
		Mode: "普通全日制",
		Level: "研究生",
		Graduation: "毕业",
		CertNo: "333",
		Photo: "/static/photo/11.png",
	}
	msg, err = serviceSetup.ModifyIde(info)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	// 修改后，根据身份证号码查询最新的信息
	result, err = serviceSetup.FindIdeInfoByIDNumber("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Identity
		json.Unmarshal(result, &edu)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(edu)
	}

	// 修改后，根据证书编号与姓名查询最新信息
	result, err = serviceSetup.FindIdeByCertNoAndName("333","张三")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Identity
		json.Unmarshal(result, &edu)
		fmt.Println("根据证书编号与姓名查询信息成功：")
		fmt.Println(edu)
	}

	// 删除信息
	msg, err = serviceSetup.DelIde("101")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息删除成功, 交易编号为: " + msg)
	}

	// 删除后，根据身份证号码查询信息
	result, err = serviceSetup.FindIdeInfoByIDNumber("101")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("根据身份证号码查询信息失败，指定身份证号码的信息不存在或已被删除...")
	} else {
		var edu service.Identity
		json.Unmarshal(result, &edu)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(edu)
	}*/
	//......................................测试代码完毕.........................................//

	//调用web层
	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)
}
