Go汇总

1.启动：go run *.go
2.编译为可执行文件：
    (1)go build *.go
    (2)go build -o <可执行文件名称> [编译的文件1 编译的文件2 ...]
3.运行可执行文件：./*
4.无法下载工具解决方案
    更改服务器地址：
        go env -w GOPROXY=https://goproxy.cn,direct
        go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
5.代码提示工具gopls：go install golang.org/x/tools/gopls@latest
6.debug工具dlv：go install github.com/go-delve/delve/cmd/dlv@latest
7.设置go环境[module]：go env -w GO111MODULE=auto
8.go module 命令
    (1)go mod init -> 生成 go.mod 文件
    (2)go mod download -> 下载 go.mod 文件中指明的所有依赖
    (3)go mod tidy -> 整理现有的依赖
    (4)go mod graph -> 查看现有的依赖结构
    (5)go mod edit -> 编辑go.mod 文件
    (6)go mod vendor -> 导出项目所有的依赖到vendor目录
    (7)go mod verify -> 校验一个模块是否被篡改过
    (8)go mod why -> 查看为什么需要依赖某模块
9.替换版本命令：go mod edit -replace= <旧版本> = <新版本>
10.测试客户端连接
    (1)Windows:Test-NetConnection 127.0.0.1 -port 8888
    (2)linux:nc 127.0.0.1 8888