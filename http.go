package gowebshell

import (
	"github.com/zyylhn/gowebshell/crypto"
	"io"
	"io/ioutil"
	"net/http"
)

func GetRespones(webshell *Webshell,body io.Reader,rc4key string,client *http.Client) (string,error) {
	req,err:=http.NewRequest("POST",webshell.Url,body)
	if err!=nil{
		return "",err
	}
	for k,v:=range webshell.Header{
		req.Header.Set(k,v)
	}
	if client==nil{
		client=DefaultOption.client
	}
	resp,err:=client.Do(req)
	if err!=nil{
		return "",err
	}
	defer resp.Body.Close()
	resultb,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return "",err
	}
	resultb,err=crypto.Rc4Base64Decode(resultb,rc4key)
	results:=string(resultb)
	return results,nil
}
