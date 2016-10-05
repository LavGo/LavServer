package model

import (
	"strconv"
)

type ConfigInfo struct {
	StaticBasePath string
	HtmlBasePath string
	CssBasePath  string
	JsBasePath string
	LogFilePath string
	Port string
        MaxThread int32
	Config *InitConfig
}

func (this *ConfigInfo)Init(){
	this.Config=new(InitConfig)
	this.Config.Init()

	this.StaticBasePath="./view/"
	this.HtmlBasePath = "/static/html"
	this.CssBasePath="/static/css"
	this.JsBasePath="/static/js"
	//this.LogFilePath="LavServer.log"
	this.Port="50101"
	this.MaxThread = 10
	if this.Config.GetKey("staticbasepath")!=""{
		this.StaticBasePath=this.Config.GetKey("staticbasepath")
	}
	if this.Config.GetKey("port")!=""{
		this.Port=this.Config.GetKey("port")
	}
	if this.Config.GetKey("htmlbasepath")!=""{
		this.HtmlBasePath=this.Config.GetKey("htmlbasepath")
	}
	if this.Config.GetKey("cssbasepath")!=""{
		this.CssBasePath=this.Config.GetKey("cssbasepath")
	}
	if this.Config.GetKey("jsbasepath")!=""{
		this.JsBasePath=this.Config.GetKey("jsbasepath")
	}
	if this.Config.GetKey("logfile")!=""{
		this.LogFilePath=this.Config.GetKey("logfile")
	}
	if this.Config.GetKey("maxthread")!=""{
		v,_:=strconv.Atoi(this.Config.GetKey("maxthread"))
		this.MaxThread=int32(v)
	}
	this.HtmlBasePath=this.StaticBasePath+this.HtmlBasePath
	this.CssBasePath=this.StaticBasePath+this.CssBasePath
	this.JsBasePath=this.StaticBasePath+this.JsBasePath
}
