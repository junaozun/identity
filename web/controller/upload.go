package controller

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

//照片上传
func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request)  {

	start := "{"	//字符串的开始
	content := ""   //字符串内容
	end := "}"		//字符串的结束

	//获取照片文件
	file, _, err := r.FormFile("file")//addEdu.html中type="file"
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"指定了无效的文件\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer file.Close()

	//将照片文件转换为字节数组
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"无法读取文件内容\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	//照片文件类型
	filetype := http.DetectContentType(fileBytes)
	//log.Println("filetype = " + filetype)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		content = "\"error\":1,\"result\":{\"msg\":\"文件类型错误\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	fileName := randToken(12)	// 指定照片文件的名称
	fileEndings, err := mime.ExtensionsByType(filetype)	// 获取文件扩展名

	// 指定照片文件存储路径；文件名+扩展名
	newPath := filepath.Join("web", "static", "photo", fileName + fileEndings[0])

	newFile, err := os.Create(newPath)
	if err != nil {
		log.Println("创建文件失败：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"创建文件失败\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		log.Println("写入文件失败：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"保存文件内容失败\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	path := "/static/photo/" + fileName + fileEndings[0]
	//error为0表示成功
	content = "\"error\":0,\"result\":{\"fileType\":\"image/png\",\"path\":\"" + path + "\",\"fileName\":\""+fileName+fileEndings[0]+"\"}"
	w.Write([]byte(start + content + end))
	return
}

//随机生成字符串作为照片文件的名称
func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

