package controllers

import (
	"path"
	"os"
	"fmt"
	"strings"
)

type ApkController struct {
	BaseController
}

/*下载apk文件*/
func (this *ApkController) DownloadFile() {
	pluginName := this.Ctx.Input.Param(":pluginName")
	pluginRepositoryRoot := "static/plugin"
	fileSuffix := "apk"
	filePath := pluginRepositoryRoot+"/"+pluginName+"/"+pluginName+"."+fileSuffix
	this.Ctx.Output.Download(filePath)
}


func (this *ApkController) UpLoadForm(){
	this.TplName = "upload.tpl"
}

/*下载apk文件*/
func (this *ApkController) UploadFile() {

	//获取上传的文件
	f, h, _ := this.GetFile("fileName")
	fileNameWithSuffix := h.Filename
	//获取文件后缀
	fileSuffix := path.Ext(fileNameWithSuffix)

	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".apk":true,
	}
	if _,ok:=AllowExtMap[fileSuffix];!ok{
		this.Ctx.WriteString( "文件不符合上传要求" )
		return
	}
	//创建目录
	//获取文件名，去掉后缀
	onFileName := strings.TrimSuffix(fileNameWithSuffix, fileSuffix)
	uploadDir := "static/plugin/" +onFileName+"/"
	err := os.MkdirAll(uploadDir, 0777)
	if err != nil {
		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
		return
	}

	fileName := h.Filename
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = this.SaveToFile("fileName", fpath)
	if err != nil {
		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
	}
	this.Ctx.WriteString( "上传成功~！！！！！！！" )
}




