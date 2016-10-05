package model

import (
	"os"
	"bufio"
	"io"
	"strings"
)

type InitConfig struct {
	value map[string]string
}

func (this *InitConfig)Init(){
	this.value=make(map[string]string)
	f,err:=os.Open("conf/config.ini")
	defer f.Close()
	if err!=nil{
		panic(err)
	}
	r:=bufio.NewReader(f)
	for{
		b,_,err:=r.ReadLine()
		if err == io.EOF{
			break
		}else if err !=nil{
			panic(err)
		}
		if strings.Index(string(b),"#") >=0{
			continue
		}
		maps:=strings.Split(string(b),"=")
		if len(maps) ==2 {
			k := strings.ToLower(strings.Trim(maps[0], " "))
			v := strings.Trim(maps[1], " ")
			this.value[k] = v
		}
	}
}
func (this *InitConfig)GetKey(key string)string{
	if this.value[key]!=""{
		return this.value[key]
	}
	return ""
}