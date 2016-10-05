package src

import (
	"github.com/LavGo/LavServer/src/model"
	"net/http"
	"github.com/LavGo/LavServer/src/logs"
	"io"
	"os"
	"io/ioutil"
)

type SysDealRequest struct{
	configInfo model.ConfigInfo
	log logs.Logs
	header *SysDealHeader
	uri *SysDealRequestURI
}

func (this *SysDealRequest)Init(){
	this.configInfo.Init()
	go this.log.Init(this.configInfo.LogFilePath)
}

func (this *SysDealRequest)dealRequest(rep http.ResponseWriter,req *http.Request){
	//Http Header
	this.header=&SysDealHeader{Request:req,Response:rep}
	this.header.Init()

	//处理uri
	this.uri=&SysDealRequestURI{uri:req.RequestURI,configInfo:this.configInfo}
	this.uri.Init()
	file,err:=os.Open(this.uri.GetURI())
	defer file.Close()
	if err != nil{
		if os.IsNotExist(err){
			this.header.SetStatusCode(404)
		}
		this.log.Error(err)
	}
	filebuf,err:=ioutil.ReadAll(file)
	if err != nil{
		this.log.Error(err)
	}

	rep=this.header.Response
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
