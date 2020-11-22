package main


import (
	"cryptology/Base64"
	"fmt"
)

func main() {
	str := "人心不足蛇吞象"
	//使用base64对str进行编码
	cipherText :=Base64.Base64EncodeString(str)
	fmt.Println("base64 编码后:",cipherText)
	fmt.Println("base64 解码后：",Base64.Base64DecodeString(cipherText))
}

//func Base64EncodeString(str string)string  {
//	return base64.StdEncoding.EncodeToString([]byte(str))
//
//}
//func Base64DecodeString(str string)string  {
//	result, _ :=base64.StdEncoding.DecodeString(str)
//	return  string(result)
//
//}
