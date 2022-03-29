package crypto

import (
	"crypto/rc4"
	"encoding/base64"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

//先base64解码，然后rc4解密
func Rc4Base64Decode(data []byte,key string) ([]byte,error) {
	lengthpadding:=strings.Count(string(data),"=")
	n := base64.StdEncoding.DecodedLen(len(data))
	bytecode:=make([]byte,n-lengthpadding)
	base64.StdEncoding.Decode(bytecode,data)    //base64解码
	ketb:=[]byte(key)
	c,err:=rc4.NewCipher(ketb)
	if err!=nil{
		return nil,err
	}
	c.XORKeyStream(bytecode,bytecode)
	return bytecode,nil
}

func SpecialEncoding(entype string,data []byte) ([]byte,error) {
	switch entype {
	case "GBK":
		return simplifiedchinese.GBK.NewDecoder().Bytes(data)
	default:
		return data,nil
	}
}