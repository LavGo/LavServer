package src

import (
	"github.com/LavGo/LavServer/src/model"
	"net/http"
	"github.com/LavGo/LavServer/src/logs"
	"io"
	"os"
	"io/ioutil"
	"regexp"
)

type SysDealRequest struct{
	configInfo model.ConfigInfo
	log logs.Logs
	header *SysDealHeader
}

func (this *SysDealRequest)Init(){
	go this.log.Init()
	this.configInfo.Init()
}
func (this *SysDealRequest)PreDealRequestPath(path string)string{
	if ok,_:=regexp.MatchString(".[cC][sS]{2}|.[jJ][sS]|.ico|(?i).+?\\.(jpg|gif|bmp).*",path);!ok{
		return this.configInfo.HtmlBasePath+path
	}
	return path
}

func (this *SysDealRequest)dealRequest(rep http.ResponseWriter,req *http.Request){
	//Http Header
	this.header=&SysDealHeader{Request:req,Response:rep}
	this.header.Init()
	rep=this.header.Response

	path:=req.RequestURI
	path = this.PreDealRequestPath(path)
	file,err:=os.Open("./"+this.configInfo.StaticBasePath+path)
	defer file.Close()
	if err != nil{
		this.log.Error(err)
	}
	filebuf,err:=ioutil.ReadAll(file)
	if err != nil{
		this.log.Error(err)
	}
	io.WriteString(rep,string(filebuf))
}
func (this *SysDealRequest)DealRequest(){
	http.HandleFunc("/",this.dealRequest)
}
func (this *SysDealRequest)Start(){
	this.DealRequest()
	err:=http.ListenAndServe(":"+this.configInfo.Port,nil)
	if err != nil{
		this.log.Error(err)
	}
}
