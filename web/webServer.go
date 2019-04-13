package web

import (
	"fmt"
	"identity/web/controller"
	"net/http"
)

// 启动Web服务并指定路由信息
//controller目录下userInfo.go代码中有个Application结构体
func WebStart(app controller.Application)  {

	//将web/static和/static/进行映射
	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由信息(匹配请求)
	http.HandleFunc("/", app.LoginView) //       		 /-->LoginView方法-->login.html
	http.HandleFunc("/login", app.Login) //       	 /login-->Login方法-->登录成功进入index.html，登录失败进入login.html
	http.HandleFunc("/loginout", app.LoginOut)// 		 /loginout-->LoginOut方法-->login.html

	http.HandleFunc("/index", app.Index)//       		 /index-->Index方法-->index.html
	http.HandleFunc("/help", app.Help)//        	     /help-->Help方法-->help.html

	http.HandleFunc("/addIdeInfo", app.AddIdeShow)// 显示添加信息页面
	http.HandleFunc("/addIde", app.AddIde)// 提交信息请求

	http.HandleFunc("/queryPage", app.QueryPage)	// 转至根据证书编号与姓名查询信息页面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 根据证书编号与姓名查询信息

	http.HandleFunc("/queryPage2", app.QueryPage2)	// 转至根据身份证号码查询信息页面
	http.HandleFunc("/query2", app.FindByID)	// 根据身份证号码查询信息

	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改信息页面
	http.HandleFunc("/modify", app.Modify)	//  修改信息

	http.HandleFunc("/upload", app.UploadFile)//上传照片

	fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}

}



