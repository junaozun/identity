package controller

import (
	"encoding/json"
	"fmt"
	"identity/service"
	"net/http"
)

//全局变量，当前的登录用户
var cuser User

//用户输入网址后进入的登录页面
func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}

//用户登录成功后的页面
func (app *Application) Index(w http.ResponseWriter, r *http.Request)  {
	ShowView(w, r, "index.html", nil)
}

//帮助页面
func (app *Application) Help(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
	}{
		CurrentUser:cuser,
	}
	ShowView(w, r, "help.html", data)//将data数据（也就是当前登录用户的名字）传入help.html页面
}

// 用户登录
func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool//用于判断是否登录成功
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user//将在页面输入的user赋值给当前登录用户
			flag = true//账号和密码输入正确，flag置为true，表示登录成功
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:false,//用于判断是否在页面显示提示信息
	}

	if flag {
		// 登录成功
		ShowView(w, r, "index.html", data)//将data数据（也就是当前登录用户的名字）传入index.html页面
	}else{
		// 登录失败
		data.Flag = true//Flag为true表示显示登录失败的提示信息
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)//将data数据（也就是当前登录用户的名字）传入login.html页面。数据回填
	}
}

// 用户登出
func (app *Application) LoginOut(w http.ResponseWriter, r *http.Request)  {
	cuser = User{}
	ShowView(w, r, "login.html", nil)
}

// 显示添加信息页面，页面有个添加信息按钮，点击后进入添加信息页面
func (app *Application) AddIdeShow(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "addEdu.html", data)
}

// 添加信息的页面填入添加信息
func (app *Application) AddIde(w http.ResponseWriter, r *http.Request)  {

	ide := service.Identity{
		Name:r.FormValue("name"),
		Sex:r.FormValue("sex"),
		Nation:r.FormValue("nation"),
		IDNumber:r.FormValue("idNumber"),
		NativePlace:r.FormValue("nativePlace"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Length:r.FormValue("length"),
		Mode:r.FormValue("mode"),
		Level:r.FormValue("level"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
	}

	app.Setup.SaveIde(ide)//调用业务层的SaveIde方法将信息保存至账本中

	r.Form.Set("certNo", ide.CertNo)
	r.Form.Set("name", ide.Name)
	app.FindCertByNoAndName(w, r)//添加完信息后查询添加的信息
}

//点击根据证书编号和姓名查询的按钮
func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query.html", data)//跳入输入证书编号和姓名查询的页面
}

// 根据证书编号与姓名查询信息的方法
func (app *Application) FindCertByNoAndName(w http.ResponseWriter, r *http.Request)  {
	certNo := r.FormValue("certNo")
	name := r.FormValue("name")
	result, err := app.Setup.FindIdeByCertNoAndName(certNo, name)
	var ide = service.Identity{}
	json.Unmarshal(result, &ide)

	fmt.Println("根据证书编号与姓名查询信息成功：")
	fmt.Println(ide)

	data := &struct {
		Ide service.Identity
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Ide:ide,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:false,//根据证书编号和姓名查询，不会显示历史记录
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)//跳入查询结果页面
}

//点击根据身份证号码查询的按钮
func (app *Application) QueryPage2(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query2.html", data)//跳入输入身份证号码查询的页面
}

// 根据身份证号码查询信息的方法
func (app *Application) FindByID(w http.ResponseWriter, r *http.Request)  {
	IDNumber := r.FormValue("idNumber")
	result, err := app.Setup.FindIdeInfoByIDNumber(IDNumber)
	var ide = service.Identity{}
	json.Unmarshal(result, &ide)

	data := &struct {
		Ide service.Identity
		CurrentUser User
		Msg string
		Flag bool
		History bool
	}{
		Ide:ide,
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
		History:true,
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "queryResult.html", data)
}

// 在添加信息的结果页面，才会有修改信息的按钮,点击这个按钮触发本方法，然后跳入修改信息的页面
func (app *Application) ModifyShow(w http.ResponseWriter, r *http.Request)  {
	// 根据证书编号与姓名查询信息
	certNo := r.FormValue("certNo")
	name := r.FormValue("name")
	result, err := app.Setup.FindIdeByCertNoAndName(certNo, name)

	var Ide = service.Identity{}
	json.Unmarshal(result, &Ide)

	data := &struct {
		Ide service.Identity
		CurrentUser User
		Msg string
		Flag bool
	}{
		Ide:Ide,
		CurrentUser:cuser,
		Flag:true,
		Msg:"",
	}

	if err != nil {
		data.Msg = err.Error()
		data.Flag = true
	}

	ShowView(w, r, "modify.html", data)//进入修改信息的页面
}

// 在修改信息的页面修改信息
func (app *Application) Modify(w http.ResponseWriter, r *http.Request) {
	Ide := service.Identity{
		Name:r.FormValue("name"),
		Sex:r.FormValue("sex"),
		Nation:r.FormValue("nation"),
		IDNumber:r.FormValue("idNumber"),
		NativePlace:r.FormValue("nativePlace"),
		BirthDay:r.FormValue("birthDay"),
		EnrollDate:r.FormValue("enrollDate"),
		GraduationDate:r.FormValue("graduationDate"),
		SchoolName:r.FormValue("schoolName"),
		Major:r.FormValue("major"),
		QuaType:r.FormValue("quaType"),
		Length:r.FormValue("length"),
		Mode:r.FormValue("mode"),
		Level:r.FormValue("level"),
		Graduation:r.FormValue("graduation"),
		CertNo:r.FormValue("certNo"),
		Photo:r.FormValue("photo"),
	}


	app.Setup.ModifyIde(Ide)//调用业务层方法

	r.Form.Set("certNo", Ide.CertNo)
	r.Form.Set("name", Ide.Name)
	app.FindCertByNoAndName(w, r)//将修改后的信息显示在页面

	//r.Form.Set("idNumber",Ide.IDNumber)
	//app.FindByID(w,r)
}
