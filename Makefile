mac:
#告知Go编译器生成二进制文件的目标环境：amd64CPU的Linux系统
	GOOS=darwin GOARCH=amd64 go build main.go
linux:
	GOOS=linux GOARCH=amd64 go build main.go
windows:
	GOOS=windows GOARCH=amd64 go build main.go