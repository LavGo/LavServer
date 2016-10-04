package model

type ConfigInfo struct {
	StaticBasePath string
	HtmlBasePath string
	CssBasePath  string
	JsBasePath string
	LogFilePath string
	Port string
        MaxThread int32

}

func (this *ConfigInfo)Init(){
	this.StaticBasePath="view"
	this.HtmlBasePath = "/static/html"
	this.CssBasePath="view/css"
	this.JsBasePath="view/js"
	this.LogFilePath="LavServer.log"
	this.Port="50101"
	this.MaxThread = 10
}
