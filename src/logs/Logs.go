package logs

import (
	"log"
	"os"
	"fmt"
)

type Logs struct{
	sysLog *log.Logger

}
func (this *Logs)Init(path string){
	//file,_:=os.Create("LavServer.log")
	fmt.Println(path)
	if path != ""{
		f,err:=os.Create(path)
		defer f.Close()
		if err ==nil {
			this.sysLog = log.New(f, "", log.LstdFlags)
		}else{
			fmt.Println(err)
		}
		return
	}else{
		this.sysLog = log.New(os.Stdout,"",log.LstdFlags)
	}
}
func (this *Logs)LogsPreoutPut()(string){
     return ""
}
func (this *Logs)Error(msg interface{}){
	this.sysLog.Println("[Error] : ",msg)
}

func (this *Logs)Info(msg interface{}){
	this.sysLog.Println("[Info] : ",msg)
}

func (this *Logs)Warn(msg interface{}){
	this.sysLog.Println("[Warn] : ",msg)
}