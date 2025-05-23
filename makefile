# for updating the version of the project and pushing the tag to the repository 
# NEVER GO BEYOND 1.0.0 (breaking change use 0.x.0 ) 
# keep the 0.x.x -> 1.x.x versioning in the v0 directory from 1.0.0 -> 2.0.0 make v2 folder and new go.mod file in it

VERSION = 0.0.3

MODULE_NAME := github.com/fengdotdev/golibs-future

#my GOLIBS For all projects 
TRAITS = github.com/fengdotdev/golibs-traits
TESTING = github.com/fengdotdev/golibs-testing
KWOW = github.com/fengdotdev/golibs-kwowledge
NDRIVE = github.com/fengdotdev/golibs-nativedrive
VDRIVE = github.com/fengdotdev/golibs-vdrive
COMMON = github.com/fengdotdev/golibs-commontypes
FUNCS = github.com/fengdotdev/golibs-funcs
ONEDRIVE = github.com/fengdotdev/golibs-1driveclient
BRIDGE = github.com/fengdotdev/golibs-bridge
ML = github.com/fengdotdev/golibs-ml
STAT = github.com/fengdotdev/golibs-statistics
LA = github.com/fengdotdev/golibs-linealalgebra
DUMMY = github.com/fengdotdev/golibs-dummy
DC = github.com/fengdotdev/golibs-datacontainer
SPAGE = github.com/fengdotdev/golibs-staticpages



.PHONY: sand init helloworld fix get tag test

sand: 
	go run cmd/sandbox/main.go


# create a new go module
init:
	go mod init $(MODULE_NAME)
	make get
	make helloworld


# create some directories for the project
helloworld:
# create  cmd directory
	mkdir -p cmd
	mkdir -p cmd/sandbox
# keep the 0.x.x -> 1.x.x versioning in the v1 directory from 1.0.0 -> 2.0.0 make v2 folder and new go.mod file in it
	mkdir -p v1
# create the main.go file
	echo "package main" > cmd/sandbox/main.go
	echo "import \"fmt\"" >> cmd/sandbox/main.go
	echo "func main() {" >> cmd/sandbox/main.go
	echo "fmt.Println(\"Hello, world!\")" >> cmd/sandbox/main.go
	echo "}" >> cmd/sandbox/main.go

fix:	
# fix the go.mod file
	go mod tidy
	go mod vendor
	go mod verify
	

# update the go.mod file with the latest dependencies
get:
	go get $(TRAITS)@latest
	go get $(TESTING)@latest

# update the version of the project
tag:
		git tag v${VERSION} && git push origin v${VERSION}

# run all tests
test:
	go clean -testcache
	go test -count=1 -v ./...



