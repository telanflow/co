# co
Go common libraries （Go语言公共库）

# Func

### file.go
```
FileExists()        文件是否存在
CreatFileNested()   给定path创建文件，如果目录不存在就递归创建
GetPathDirs()       获取目录所有文件夹
GetPathFiles()      获取目录所有文件
ReadLine()          按行读取文件
```

### strings.go
```
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
```

### bytes.go
```
Bytes2Str()         将字节切片转换为字符串
Byte2Float64()      字节切片转64位浮点数
Byte2Int64()        字节切片转64位整型
Byte2Hex()          字节切片转16进制字符串
Byte2HexSlice()     字节切片转16进制切片
```

### slice.go
```
ContainsUint()      返回list中是否包含
ContainsInt()       返回list中是否包含
ContainsString()    返回list中是否包含
SliceIntersect()    求两个切片交集
SliceDifference()   求两个切片差集
```

### number.go
```
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

### md5.go
```
Md5()               获取字符串md5值, length指定结果长度32/16
Md5Str()            字符串生成md5摘要
```

### sha1.go
```
Sha1Str()           计算字符串的sha1散列值
ShaXStr()           计算字符串的shaX散列值，x为1/256/512
```