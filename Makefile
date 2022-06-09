configure:
	@mkdir -p ./bin/darwin/amd
	@mkdir -p ./bin/darwin/arm
	@mkdir -p ./bin/windows/amd
	@mkdir -p ./bin/windows/arm
	@mkdir -p ./bin/linux/amd
	@mkdir -p ./bin/linux/arm

darwin: ./main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/amd/kaiji ./main.go

darwin.arm: ./main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin/arm/kaiji ./main.go

win: ./main.go
	@GOOS=windows GOARCH=amd64 go build -o ./bin/windows/amd/kaiji.exe ./main.go

win.arm: ./main.go
	@GOOS=windows GOARCH=arm go build -o ./bin/windows/arm/kaiji.exe ./main.go

linux: ./main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd/kaiji ./main.go

linux.arm: ./main.go
	@GOOS=linux GOARCH=arm go build -o ./bin/linux/arm/kaiji ./main.go

build: configure
	$(MAKE) darwin
	$(MAKE) darwin.arm
	$(MAKE) win
	$(MAKE) win.arm
	$(MAKE) linux
	$(MAKE) linux.arm

clean: ./bin
	@rm -rf ./bin
