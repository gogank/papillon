# Papillon
[![Go Report Card](https://goreportcard.com/badge/github.com/gogank/papillon?ver=0.1)](https://goreportcard.com/report/github.com/gogank/papillon)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gogank/papillon)
[![GitHub stars](https://img.shields.io/github/stars/gogank/papillon.svg?style=social&label=Stars)]()

A distributed static blog publish tool on IPFS


## Prepare 
[install ipfs](https://ipfs.io/docs/install/)

### start up ipfs node
**Note: alpha version need a local IPFS node**

```bash
# first run ipfs
ipfs init

# start ipfs daemon
ipfs daemon

```

## Alpha Test

**Note: Alpha version please use those command**

```bash
go get -u github.com/gogank/papillon

cd $GOPATH/src/github.com/gogank/papillon 

make

cd $GOPATH/src/github.com/gogank/papillon/build/blog

# generate new post
./papi new mypost

# edit it
vim $GOPATH/src/github.com/gogank/papillon/build/blog/source/posts/mypost.md

# generate whole website
./papi gen 

# publish IPFS, and get your blog URL 
./papi pub
```


## install Papillon
Note: beta version command

```bash
go get -u github.com/gogank/papillon
cd $GOPATH/src/github.com/gogank/papillon && go build -o $GOPATH/bin/papi
```

## init (todo)

```bash
cd blog_dir
papi init 
```

## new post
```bash
papi new my_post_name 
```

## genreate blog files

```bash
papi gen
```

## publish your blog onto IPFS
```bash
papi pub 
```


> This is a project for [Go Hack 2017](http://gohack2017.golangfoundation.org/)
