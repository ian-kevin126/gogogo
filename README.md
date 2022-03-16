## learn go

### 1.go mod
go 语言依赖管理的三个阶段： GOPATH ——> GOVENDOR ——> go mod

- 由go命令统一的管理，用户不必关心目录结构
- 初始化：go mod init
- 增加依赖：go get
- 更新依赖：go get [@v...]，go mod tidy
- 将旧项目迁移到 go mod：go mod init，go build ./...

### 2.包
- 如何扩充系统类型或者别人的类型
  - 使用组合(basics/tree)：最常用
  - 定义别名(basics/queue)：最简单
  - 使用内嵌的方式来扩展已有类型(basics/embedded)：需要省下许多代码
