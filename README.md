# GO快速上手

## 简单工程搭建

目录结构如下：

```shell
➜  learn git:(main) ✗ tree
.
├── basic
│   ├── go.mod
│   ├── main.go
│   └── types
│       ├── impl.go
│       ├── types.go
│       └── utils.go
├── concurrency
├── go.mod
├── go.work
├── interfaces
├── pkg_mamagement
└── testing
```

```shell
# 初始化新模块（每个子目录独立模块）
cd learn/basic
go mod init learn/basic
# main.go 应该也在basic目录下，而不是types目录

# 工作区管理（顶层go.work）
cd learn
go work init basic concurrency interfaces testing pkg_management

# 代码写好之后，可以命令行跑 
cd learn/basic
go run .
# 也可以先编译，再运行
go build -o basic
./basic
# 也可以在vscode中直接运行
# 在launch.json中添加配置
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "typestest",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "learn/basic"
    },
  ]
}
```
