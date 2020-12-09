package main


//
//
func Search(strings []string) string {
	length := len(strings)
	if length == 1{
		return ""
	}
	if length  ==2 && strings[0] == ")"{
		 return ""
	}
	var leftArr,rightArr = make([]string,length),make([]string,length)
	for i := 0;i<length;i++ {
		if strings[i] == ")" && strings[i-1] == "("{
			leftArr = strings[:i-2]
			rightArr = strings[i+1:length]

			break
		}
	}
	newString := append(leftArr,rightArr...)
	return Search(newString)
}
