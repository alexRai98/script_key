package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)
// https://gist.github.com/wklken/d9eba02e40adbc876ea61704a28ba92b

func decrytText(key string, encryptedText string, isEcb bool) ([]byte, error){

	key2, err := hex.DecodeString(key)
	crypted, _ := hex.DecodeString(encryptedText)
	if err != nil{
		fmt.Printf("Error DecodeString")
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key2)
	if err != nil {
		fmt.Printf("Error NewTripleDESCipher")
		return nil, err
	}
	iv := []byte("00000000")

	blockMode := cipher.NewCBCDecrypter(block, iv)

	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	return origData, nil

}
func main(){
	khakiKey := "e5fd7a94ecdfd3df703badf8fbb95702"
	data :=  "7067F4C2C60DCD067562444B35E1E20A"
	decrytText(khakiKey, data, true)
}
