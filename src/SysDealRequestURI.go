package src

import (
	"github.com/LavGo/LavServer/src/model"
	"regexp"
	"strings"
)

type SysDealRequestURI struct {
	uri string
	configInfo model.ConfigInfo
	flag bool
	header *SysDealHeader
}

func (this *SysDealRequestURI)Init(){
	this.flag = true
	this.FilteHtmlpath()
	this.FiltStaticPath()
	if this.flag{
		this.FilteDefault()
	}
}

func (this *SysDealRequestURI)GetURI()string{
	return this.uri
}

func (this *SysDealRequestURI)FiltStaticPath(){
	if ok,_:=regexp.MatchString(".[cC][sS]{2}",this.uri);ok{
		this.flag=false
		this.uri=this.configInfo.StaticBasePath+this.uri
		this.header.SetResponseContentType("text/css")
	}
	if ok,_:=regexp.MatchString(".[jJ][sS]",this.uri);ok{
		this.flag=false
		this.uri=this.configInfo.StaticBasePath+this.uri
	}
	if ok,_:=regexp.MatchString(".ico|(?i).+?\\.(jpg|gif|bmp).*",this.uri);ok{
		this.flag=false
		this.uri=this.configInfo.StaticBasePath+this.uri
	}
}

func (this *SysDealRequestURI)FilteHtmlpath(){
	if strings.HasSuffix(this.uri,".html")||strings.HasSuffix(this.uri,".htm"){
		this.flag=false
		this.uri=this.configInfo.HtmlBasePath+this.uri
	}
}
func (this *SysDealRequestURI)FilteDefault(){
	this.uri=this.configInfo.StaticBasePath+this.uri
}