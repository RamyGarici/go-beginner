package main
import("os"
"fmt"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"io")



func main(){


















	// key := []byte("secretKey")
	// // Encoding with xor:
	// file, err := os.ReadFile("secret.txt")
	// if err!=nil{
	// 	fmt.Printf("Error:%v\n",err)
	// 	os.Exit(0)

	// }
	// encryptedFile := xorEncoding(key,file)
	// os.WriteFile("secret.txt.enc", encryptedFile,0644)

	// //Decoding with xor:
	// file, err = os.ReadFile("secret.txt.enc")
	// if err!=nil{
	// 	fmt.Printf("Error:%v\n",err)
	// 	os.Exit(0)

	// }
	// decryptedFile := xorEncoding(key,file)
	// os.WriteFile("secret.txt", decryptedFile,0644)


}


func aesEncryption(key []byte, data []byte)([]byte,error){
	aesBlock,err := aes.NewCipher(key)
	if err != nil{
		fmt.Printf("Error: %v\n",err)
		return nil,err
	}
	aesGCM,err := cipher.NewGCM(aesBlock)
	if err != nil{
		fmt.Printf("Error: %v\n",err)
		return nil,err
	}
	nonce := make([]byte,12)
	_,err=io.ReadFull(rand.Reader, nonce)
	if err != nil{
		fmt.Printf("Error: %v\n",err)
		return nil,err
	}
	encryptedFile := aesGCM.Seal(nonce,nonce,data,nil)
	return encryptedFile, nil
	
	




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

