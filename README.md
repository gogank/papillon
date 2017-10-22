# Papillon
IPFS网络上的静态博客发布系统

## 准备工作 
[安装IPFS节点](https://ipfs.io/docs/install/)

### 启动ipfs节点
**Note: alpha 版本需要本地启一个IPFS节点**

```bash
# 第一次运行需要执行
ipfs init

# 第一次运行需要执行
ipfs daemon

```

## Alpha 试用

**Note: Alpha 版本 请直接使用本说明进行试用**

```bash
go get -u github.com/gogank/papillon

cd $GOPATH/src/github.com/gogank/papillon 

make

cd $GOPATH/src/github.com/gogank/papillon/build
# 生成新文章
./papi new mypost

# 生成静态资源
./papi gen 

# 发布到IPFS网络上，取得一个固定链接 
./papi pub
```


Note: 下面的说明还无法使用(2017/10/22)
## 安装 Papillon

```bash
go get -u github.com/gogank/papillon
cd $GOPATH/src/github.com/gogank/papillon && go build -o $GOPATH/bin/papi
```

## 初始化(todo)

```bash
cd blog_dir
papi init 
```

## 新建文章
```bash
papi new my_post_name 
```

## 生成文章

```bash
papi gen
```

## 发布网站
```bash
papi pub 
```


> This is a project for [Go Hack 2017](http://gohack2017.golangfoundation.org/)
