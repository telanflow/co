# co
[![GoDoc](https://godoc.org/github.com/telanflow/co?status.svg)](https://godoc.org/github.com/telanflow/co)

Go common libraries （Go语言公共库）

# 函数列表

### file.go
```text
FileExists()        文件是否存在
FileCreateMust()    给定path创建文件，如果目录不存在就递归创建
ReaderMimeType()    获取Reader的MimeType
FileMimeType()      获取文件的MimeType
GetDirs()           获取目录所有文件夹
GetFiles()          获取目录所有文件
GetFilesSuffix()    获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤
ReadLine()          按行读取文件
```

### path.go
```text
PathExists()        路径是否存在
PathFillSlash()     给路径补全 `/`
PathRemoveSlash()   移除路径最后的 `/`
PathFormSlash()     将path中的反斜杠'\'替换为'/'
IsDir()             判断路径是否为目录
IsFile()            判断路径是否为文件
```

### strings.go
```text
Str2Bytes()         将字符串转换为字节切片
Str2Int()           将字符串转换为int
Str2IntStrict()     严格将字符串转换为有符号整型
Str2Int8()          将字符串转换为int8
Str2Int16()         将字符串转换为int16
Str2Int32()         将字符串转换为int32
Str2Int64()         将字符串转换为int64
Str2UintStrict()    严格将字符串转换为无符号整型
Str2Uint()          将字符串转换为uint
Str2Uint8()         将字符串转换为uint8
Str2Uint16()        将字符串转换为uint16
Str2Uint32()        将字符串转换为uint32
Str2Uint64()        将字符串转换为uint64
Str2FloatStrict()   严格将字符串转换为浮点型
Str2Float32()       将字符串转换为float32
Str2Float64()       将字符串转换为float64
Str2Bool()          将字符串转换为布尔值
Runes2Bytes()       将[]rune转为[]byte
SplitStr()          分割字符串（该方法会过滤空字符串）
SubStr()            字符串截取
RepeatStr()         重复字符串
RepeatRune()        重复字符串 rune
Padding()           字符串填充
PadLeft()           字符串填充 - 左侧
PadRight()          字符串填充 - 右侧
```

### bytes.go
```text
Bytes2Str()         将字节切片转换为字符串
Byte2Float64()      字节切片转64位浮点数
Byte2Int64()        字节切片转64位整型
Byte2Hex()          字节切片转16进制字符串
Byte2HexSlice()     字节切片转16进制切片
```

### slice.go
```text
ContainsUint()      返回list中是否包含
ContainsInt()       返回list中是否包含
ContainsString()    返回list中是否包含
SliceIntersect()    求两个切片交集
SliceDifference()   求两个切片差集
```

### number.go
```text
Int2Str()           将整数转换为字符串
Float2Str()         将浮点数转换为字符串,decimal为小数位数
Bool2Str()          将布尔值转换为字符串
Bool2Int()          将布尔值转换为整型
Dec2Bin()           将十进制转换为二进制
Bin2Dec()           将二进制转换为十进制
Hex2Bin()           将十六进制字符串转换为二进制字符串
Bin2Hex()           将二进制字符串转换为十六进制字符串
Dec2Hex()           将十进制转换为十六进制
Hex2Dec()           将十六进制转换为十进制
Dec2Oct()           将十进制转换为八进制
Oct2Dec()           将八进制转换为十进制
BaseConvert()       进制转换,在任意进制之间转换数字
Ip2Long()           将 IPV4 的字符串互联网协议转换成长整型数字
Long2Ip()           将长整型转化为字符串形式带点的互联网标准格式地址(IPV4)
ToStr()             强制将变量转换为字符串
ToBool()            强制将变量转换为布尔值
ToInt()             强制将变量转换为整型; 其中true或"true"为1
ToFloat()           强制将变量转换为浮点型;其中true或"true"为1.0
Float64ToByte()     64位浮点数转字节切片
Int64ToByte()       64位整型转字节切片
Hex2Byte()          16进制字符串转字节切片
HexSlice2Byte()     16进制切片转byte切片
```

### map.go
```text
MergeStringMap()    合并给定的2个map
KeyToLower()        将给定的map键名转换为小写
Keys()              获取给定map的所有键名
Values()            获取给定map的所有值
GetByPath()         给定key获取value, 支持递归获取 （eg "top" "top.sub"）
```

### md5.go
```text
Md5()               获取字符串md5值, length指定结果长度32/16
Md5Str()            字符串生成md5摘要
```

### sha1.go
```text
Sha1Str()           计算字符串的sha1散列值
ShaXStr()           计算字符串的shaX散列值，x为1/256/512
```

### value.go
```text
IsEmpty()           检查变量是否为空
IsNil()             检查变量是否空值
IsBool()            是否布尔值
IsHex()             是否十六进制字符串
IsByte()            变量是否字节切片
IsStruct()          变量是否结构体
IsInterface()       变量是否接口
IsString()          变量是否字符串
IsBinary()          字符串是否二进制
```

### os.go
```text
IsWindows()         当前操作系统是否Windows
IsLinux()           当前操作系统是否Linux
IsDarwin()          当前操作系统是否Mac OS/X
IsMSys()            msys(MINGW64) env，不一定支持颜色
IsConsole()         check out is in stderr/stdout/stdin
GetCwd()            获取当前执行文件的真实目录（不受 os.Chdir 函数影响） 
```

### exec.go
```text
QuickExec()         快速执行cli命令
ExecCmd()           执行命令并返回输出
ShellExec()         通过shell执行exec命令
```

### random.go
```text
RandInt()                   随机获取 min - max 之间的数值
RandInt64()                 随机获取 min - max 之间的数值
RandIntSlice()              随机获取 min - max 之间的数值切片
RandInt64Slice()            随机获取 min - max 之间的数值切片
RandStr()                   生成随机字符串
RandBytes()                 生成随机字节切片
RandStrForBytes()           生成随机字符串
RandStrToLower()            字符串随机字符转小写
RandStrToUpper()            字符串随机字符转大写
RandStrToUpperWithIndex()   指定字符索引转大写
RandStrToUpperWithIndex()   指定字符索引转小写
```

### mist.go [薄雾算法](https://github.com/asyncins/mist)
```text
MistGenerate()      生成全局唯一ID
```