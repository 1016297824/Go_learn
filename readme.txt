Go汇总

1.启动：go run *.go
2.编译为可执行文件：go build *.go
3.运行可执行文件：./*
4.无法下载工具解决方案
    更改服务器地址：go env -w GOPROXY=https://goproxy.cn
5.代码提示工具gopls：go install golang.org/x/tools/gopls@latest
6.debug工具dlv：go install github.com/go-delve/delve/cmd/dlv@latest
7.设置go环境[module]：go env -w GO111MODULE=auto