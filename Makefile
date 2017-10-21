all: build
.PHONY: build

build:
	@echo "building papi"
	@echo "-----"
	cp -R ${GOPATH}/src/github.com/gogank/papillon/configuration/blog  ${GOPATH}/src/github.com/gogank/papillon/build/
	govendor build -ldflags -s -o ${GOPATH}/src/github.com/gogank/papillon/build/papi -tags=embed

