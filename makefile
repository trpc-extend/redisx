APP        := redisx
TARGET     := redisx_demo
ENV        := test
INS        :=

export CGO_CXXFLAGS_ALLOW:=.*
export CGO_LDFLAGS_ALLOW:=.*
export CGO_CFLAGS_ALLOW:=.*

app:="redisx_demo"

.PHONY: all test clean

all: build
check:
ifeq ($(strip $(BK_CI_PIPELINE_NAME)),)
	@echo "\033[32m <============== 合法性校验 app ${app} =============> \033[0m"
	goimports -format-only -w -local git.code.oa.com .
	gofmt -s -w .
	golangci-lint run
endif

build:
	@echo "\033[32m <============== making app ${app} =============> \033[0m"
	go build -ldflags='-w -s' $(FLAGS) -o ./${app} ./demo
	
unit-test: $(DEPENDENCIES)
	@echo -e "\033[32m ============== making unit test =============> \033[0m"
	go test `go list ./... |grep -v api_test` -v -run='^Test' -covermode=count -gcflags=all=-l ./...


redisx-test: $(DEPENDENCIES)
	@echo -e "\033[32m ============== making redisx test =============> \033[0m"
	go test `go list ./... |grep -v api_test` -v -run='^Test' -covermode=count -gcflags=all=-l ./...

clean:
	@echo -e "\033[32m ============== cleaning files =============> \033[0m"
	rm -fv ${TARGET}