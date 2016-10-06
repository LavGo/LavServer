package main

import (
	"github.com/LavGo/LavServer/src"
	"sync"
)
var once sync.Once
func init(){

}
func main(){
	sys:=new(src.SysDealRequest)
	once.Do(sys.Init)
	sys.Start()
}
