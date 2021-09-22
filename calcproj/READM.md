# 新运行方式

1. 查看go环境变量 `go env` , 保证`GO111MODULE="on"`
2. 项目calcproj目录执行 `go mod init calcproj`
3. cd到bin目录执行 `go build calcproj/src/calc` 就能看到一个calc二进制可执行文件了
