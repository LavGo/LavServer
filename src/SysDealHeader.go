package src

import (
	"net/http"
	"strings"
	"regexp"
	"github.com/LavGo/LavServer/src/tools"
)

type SysDealHeader struct {
	Request *http.Request
	Response http.ResponseWriter
	RequestAccept []string
	ResponseContentType string
}
func (this *SysDealHeader)Init(){
	this.initHeader()
	//this.SetResponseContentType()
}
func (this *SysDealHeader)initHeader(){
	this.RequestAccept=strings.Split(this.Request.Header.Get("Accept"),",")
}
func (this *SysDealHeader)getClientAcceptFile(ftype string,req *http.Request)bool{
	ftypes:=strings.Split(req.Header.Get("Accept"),",")
	for i:=0;i<len(ftypes);i++{
		if ok,_:=regexp.MatchString(strings.ToLower(ftypes[i]),strings.ToLower(ftype));ok{
			return true;
		}
	}
	return false;
}
func (this *SysDealHeader)SetResponseContentType(ftype string){
	//accpetedtype:=[]string{"text/css"}
	if tools.Contains(this.RequestAccept,ftype){
		if this.ResponseContentType==""{
			this.ResponseContentType=ftype;
		}else{
			this.ResponseContentType+=(","+ftype)
		}
	}

}
func (this *SysDealHeader)SetResponseHeader(){
	if this.ResponseContentType != ""{
		this.Response.Header().Set("Content-Type",this.ResponseContentType)
	}
}
func (this *SysDealHeader)SetStatusCode(code int){
	if code == 404 {
		http.NotFound(this.Response, this.Request)
	}
	if code == 500{
		http.Error(this.Response,"something failed!",http.StatusInternalServerError)

	}
}