## 图

<img src="http://t-blog-images.aijs.top/img/202210311134749.webp" />

## 参考

<a href="https://zhuanlan.zhihu.com/p/501477368" target="_blank" >见</a>

## 初始化项目

```
mkdir src
touch main.go
go mod init go-localize
go mod tidy
mkdir config controller dao middleware model router service tmp util constant internet
...

```

## air

### 安装

<a href="https://github.com/cosmtrek/air/issues/135" target="_blank" >见 issues 135</a>

```
$ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
$ air -v

---

/ /\ | | | |_)
/_/--\ |_| |_| \_ 1.40.4, built with Go 1.18.3

```

### 配置 alias

```
alias air='$(go env GOPATH)/bin/air'
```

注意：比较奇怪的是 air -v 可输出， air 命令颜色为红色， 非绿色

### 使用

注意：终端需要重启

```
最简单的方法是执行

# 优先在当前路径查找 `.air.toml` 后缀的文件，如果没有找到，则使用默认的

air -c .air.toml

您可以运行以下命令初始化，把默认配置添加到当前路径下的.air.toml 文件。

air init

在这之后，你只需执行 air 命令，无需添加额外的变量，它就能使用 .air.toml 文件中的配置了。

air


```

## 出错了

写了 you.go 文件，空文件，没有写 package

```
♠ /Users/chenhailong/code/go/go-localize $ air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ 1.40.4, built with Go 1.18.3

watching .
watching config
watching controller
watching controller/v1
watching dao
watching middleware
watching model
watching router
watching service
!exclude tmp
watching util
building...
controller/v1/you.go:1:1: expected 'package', found 'EOF'
failed to build, error: exit status 1
```

### 起来了

```go
♠ /Users/chenhailong/code/go/go-localize $ air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ 1.40.4, built with Go 1.18.3

watching .
watching config
watching controller
watching controller/v1
watching dao
watching middleware
watching model
watching router
watching service
!exclude tmp
watching util
building...
running...
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/me                   --> go-localize/controller/v1.GetMe (3 handlers)
[GIN-debug] POST   /api/me                   --> go-localize/controller/v1.PostMe (3 handlers)
[GIN-debug] PUT    /api/me                   --> go-localize/controller/v1.PutMe (3 handlers)
[GIN-debug] DELETE /api/me                   --> go-localize/controller/v1.DeleteMe (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8001
```

## 重启耗时如何查看？

监听文件有点多，应该需要限制下范围

```go
[GIN-debug] Listening and serving HTTP on :8001
event: "/Users/chenhailong/code/go/go-localize/README.md": CHMOD
event: "/Users/chenhailong/code/go/go-localize/README.md": WRITE
event: "/Users/chenhailong/code/go/go-localize/README.md": CHMOD
event: "/Users/chenhailong/code/go/go-localize/README.md": WRITE
event: "/Users/chenhailong/code/go/go-localize/README.md": CHMOD
```

### 文件范围缩小

创建文件夹，将所有代码移动到 src 目录下，监听 src 目录

```go
// .air.toml
root = "./src"
```

## git

```
// 空文件夹上传
find ./ -type d -empty -execdir touch {}/.gitkeep {} \;
git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:841660202/go-localize.git
git push -u origin main
```

## gin

## swag

```go
// 访问地址
// http://localhost:8080/swagger/index.html

//


//
```

## knife

注意：必须要有@ID 注解，生成 operationId，点击菜单才能正常打开
id 用于标识操作的唯一字符串。在所有 API 操作中必须唯一。

### 安装 swag

```go
// 1. 安装 swagger
// go install github.com/swaggo/swag/cmd/swag@latest

// 2. 配置~/.zshrc  注意：引号 github 上引号没写
// <a href="https://github.com/swaggo/swag/issues/197" target="_blank" >issues197</a>
// export PATH="$(go env GOPATH)/bin:$PATH"

// $ swag -v
// swag version v1.8.7

// swag init  // 注意，一定要和main.go处于同一级目录

```

<a href="https://gitee.com/youbeiwuhuan/knife4go#https://gitee.com/xiaoym/knife4j" target="_blank" >见</a>

## 例子

<a href="https://github.com/swaggo/swag/blob/master/example/celler/main.go" target="_blank" >/swaggo/swag</a>
