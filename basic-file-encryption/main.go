package main



func main(){

}


func xorEncoding(key []byte, data []byte)[]byte{
	encryptedResult := make([]byte,0)
	var xoredValue byte 
	for i:=0;i<len(data);i++{
		xoredValue = data[i] ^ key[i%len(key)]
		encryptedResult = append(encryptedResult, xoredValue)
	}



   return encryptedResult
}