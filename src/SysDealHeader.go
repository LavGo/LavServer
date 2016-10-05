package src

import (
	"net/http"
	"strings"
	"regexp"
)

type SysDealHeader struct {
	Request *http.Request
	Response http.ResponseWriter
}
func (this *SysDealHeader)Init(){
	this.SetResponseContentType()
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
func (this *SysDealHeader)SetResponseContentType(){
	accpetedtype:=[]string{"text/css"}
	contenttype:=""
	for i:=0;i<len(accpetedtype);i++ {
		if this.getClientAcceptFile(accpetedtype[i], this.Request) {
			if contenttype==""{
				contenttype=accpetedtype[i];
			}else{
				contenttype+=(","+accpetedtype[i])
			}
		}
	}
	this.Response.Header().Set("Content-type", contenttype)
}

func (this *SysDealHeader)SetStatusCode(code int){
	if code == 404 {
		http.NotFound(this.Response, this.Request)
	}
	if code == 500{
	}
}