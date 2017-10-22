all: build
.PHONY: build

build:
	@echo "building papi"
	@echo "-----"
	mkdir -p ${GOPATH}/src/github.com/gogank/papillon/build/
	cp -R ${GOPATH}/src/github.com/gogank/papillon/configuration/blog  ${GOPATH}/src/github.com/gogank/papillon/build/
	govendor build -ldflags -s -o ${GOPATH}/src/github.com/gogank/papillon/build/blog/papi -tags=embed
	cd ${GOPATH}/src/github.com/gogank/papillon/build/blog


