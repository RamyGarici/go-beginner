package main
import("os"
"fmt"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"io")



func main(){
	key := []byte("12345678901234567890123456789012")
	file, err := os.ReadFile("secret.txt")
	if err!=nil{
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	encryptedFile,err := aesEncryption(key,file)
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(1)
	}
	err=os.WriteFile("secret.txt.enc",encryptedFile,0644)
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(1)
	}
	decryptedFile,err := aesDecryption(key, encryptedFile)
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(1)
	}
	err=os.WriteFile("secret.txt",decryptedFile,0644)
	if err!=nil{
		fmt.Printf("Error:%v\n",err)
		os.Exit(1)
	}


















	// key := []byte("secretKey")
	// // Encoding with xor:
	// file, err := os.ReadFile("secret.txt")
	// if err!=nil{
	// 	fmt.Printf("Error:%v\n",err)
	// 	os.Exit(1)

	// }
	// encryptedFile := xorEncoding(key,file)
	// os.WriteFile("secret.txt.enc", encryptedFile,0644)

	// //Decoding with xor:
	// file, err = os.ReadFile("secret.txt.enc")
	// if err!=nil{
	// 	fmt.Printf("Error:%v\n",err)
	// 	os.Exit(1)

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

func aesDecryption(key []byte, data []byte)([]byte,error){
	if len(data)<12{
		return nil,fmt.Errorf("File data too short")
	}
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

	decryptedFile,err := aesGCM.Open(nil,data[:12],data[12:],nil)
	if err != nil{
		fmt.Printf("Error: %v\n",err)
		return nil,err
	}
	return decryptedFile, nil
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

