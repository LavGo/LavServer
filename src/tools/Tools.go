package tools


func Contains(strs []string,str string)bool{
	for _,value:= range strs{
		if value == str{
			return true
		}
	}
	return false
}
