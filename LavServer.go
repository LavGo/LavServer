package main

import (
	"github.com/LavGo/LavServer/src"
)
func init(){

}
func main(){
	sys:=new(src.SysDealRequest)
	sys.Init()
	sys.Start()
}
