# EntropyCalc_Go

[English](https://github.com/yutianqaq/EntropyCalc_Go/readme_en.md) [简体中文]

[MalDev Academy](https://maldevacademy.com/) 中 Binary Entropy Reduction 章节的二进制文件熵计算器的 Golang 实现

用于计算二进制文件熵值

根据 https://practicalsecurityanalytics.com/file-entropy/ 可得知合法软件与恶意软件熵值的分布

合法软件熵值为 4.8 - 7.2 之间

恶意软件熵值大于 7.2 (用红色标识）

可以配合 https://github.com/yutianqaq/Supernova_CN 来加密 Shellcode，降低熵值

## 安装

## 从源码构建

```
git clone https://github.com/yutianqaq/EntropyCalc_Go
cd EntropyCalc_Go
go build
```

## 下载二进制版本

https://github.com/yutianqaq/EntropyCalc/releases

## 使用方法

```
./EntropyCalc -file filename
```



恶意软件

熵值用红色标识

![alt text](Pictures/image.png)

合法软件

熵值用绿色标识

![alt text](Pictures/image-1.png)

# 参考
- https://practicalsecurityanalytics.com/file-entropy/ 
- https://rosettacode.org/wiki/Entropy#Go 
- [MalDev Academy](https://maldevacademy.com/)
