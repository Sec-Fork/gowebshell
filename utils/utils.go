package utils

import (
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"strings"
	"time"
	"unsafe"
)


const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdBits = 6
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandStr(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

//加密和url编码
func GetIoBody(data map[string]string) io.Reader {
	var stringBody []string
	if data!=nil{
		for k,v:=range data{
			if k!="data"{   //不对文件加密
				//在这对v加密
			}
			v=url.QueryEscape(v)
			stringBody=append(stringBody,fmt.Sprintf("%v=%v",k,v))
		}
		return strings.NewReader(strings.Join(stringBody,"&"))
	}

	return nil
}

