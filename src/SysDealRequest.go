package src

import (
	"github.com/LavGo/LavServer/src/model"
	"net/http"
	"github.com/LavGo/LavServer/src/logs"
	"io"
	"os"
	"io/ioutil"
	"fmt"
)

type SysDealRequest struct{
	configInfo model.ConfigInfo
	log logs.Logs
	header *SysDealHeader
	uri *SysDealRequestURI
}

func (this *SysDealRequest)Init(){
	this.configInfo.Init()
	this.log.Init(this.configInfo.LogFilePath)
	fmt.Println("LavServer init...")
}

func (this *SysDealRequest)dealRequest(rep http.ResponseWriter,req *http.Request){
	defer func(){
		if r:=recover();r!=nil{
			this.log.Error(r)
			this.header.SetStatusCode(500)
		}
	}()
	//Http Header
	this.header=&SysDealHeader{Request:req,Response:rep}
	this.header.Init()
	//处理uri
	this.uri=&SysDealRequestURI{uri:req.RequestURI,configInfo:this.configInfo,header:this.header}
	this.uri.Init()
	file,err:=os.Open(this.uri.GetURI())
	defer file.Close()
	if err != nil{
		if os.IsNotExist(err){
			this.header.SetStatusCode(404)
		}
		return
		panic(err)
	}
	filebuf,err:=ioutil.ReadAll(file)
	if err != nil{
		panic(err)
	}

	//rep=this.header.Response
	//设置返回http报头
	this.header.SetResponseHeader()
	io.WriteString(rep,string(filebuf))
}
func (this *SysDealRequest)DealRequest(){
	http.HandleFunc("/",this.dealRequest)
}
func (this *SysDealRequest)Start(){
	this.DealRequest()
	fmt.Println("LavServer Started.")
	err:=http.ListenAndServe(":"+this.configInfo.Port,nil)
	if err != nil{
		this.log.Error(err)
	}
}
