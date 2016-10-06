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
	fmt.Println("Init Log System...")
	if path != ""{
		f,err:=os.Create(path)
		if err ==nil {
			fmt.Println("Log write To : ",path)
			this.sysLog = log.New(f, "", log.LstdFlags)
		}else {
			fmt.Println(err)
			goto stderrout
		}
		defer f.Close()
		return
	}
	stderrout:
	fmt.Println("Log write To : StdOut.")
	this.sysLog = log.New(os.Stdout,"",log.LstdFlags)
	fmt.Println("Log System Started.")
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