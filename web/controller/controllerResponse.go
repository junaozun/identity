package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

//w：将结果写入客户端；r：请求对象；templateName：模板文件名称；data：页面动态数据
func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{})  {

	// 指定视图所在路径
	pagePath := filepath.Join("web", "tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("创建模板实例错误: %v", err)
		return
	}

	//将页面动态的数据传入模板tpl文件中,如给页面中{{.ide.}}传入实际的数据
	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("在模板中融合数据时发生错误: %v", err)
		//fmt.Fprintf(w, "显示在客户端浏览器中的错误信息")
		return
	}

}
