package main
import("os"
"fmt")



func main(){
	key := []byte("secretKey")
	// Encoding:
	file, err := os.ReadFile("secret.txt")
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(0)

	}
	encryptedFile := xorEncoding(key,file)
	os.WriteFile("secret.txt.enc", encryptedFile,0644)

	//Decoding:
	file, err = os.ReadFile("secret.txt.enc")
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(0)

	}
	decryptedFile := xorEncoding(key,file)
	os.WriteFile("secret.txt", decryptedFile,0644)


}


func xorEncoding(key []byte, data []byte)[]byte{
	encryptedResult := make([]byte,len(data))
	var xoredValue byte 
	for i:=0;i<len(data);i++{
		xoredValue = data[i] ^ key[i%len(key)]
		encryptedResult[i] = xoredValue
	}



   return encryptedResult
}

