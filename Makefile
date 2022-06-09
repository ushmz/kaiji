configure:
	@mkdir -p ./bin/darwin/amd64
	@mkdir -p ./bin/darwin/arm64
	@mkdir -p ./bin/windows/amd64
	@mkdir -p ./bin/windows/arm
	@mkdir -p ./bin/linux/amd64
	@mkdir -p ./bin/linux/arm

darwin: ./main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/amd64/kaiji ./main.go

darwin.arm: ./main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin/arm64/kaiji ./main.go

win: ./main.go
	@GOOS=windows GOARCH=amd64 go build -o ./bin/windows/amd64/kaiji.exe ./main.go

win.arm: ./main.go
	@GOOS=windows GOARCH=arm go build -o ./bin/windows/arm/kaiji.exe ./main.go

linux: ./main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/kaiji ./main.go

linux.arm: ./main.go
	@GOOS=linux GOARCH=arm go build -o ./bin/linux/arm/kaiji ./main.go

build: configure
	@$(MAKE) darwin
	@$(MAKE) darwin.arm
	@$(MAKE) win
	@$(MAKE) win.arm
	@$(MAKE) linux
	@$(MAKE) linux.arm

define compress
	@tar -C bin/$(1)/$(2) -cf dest/kaiji-$(1)-$(2).tar $(3)
	@tar -rf dest/kaiji-$(1)-$(2).tar README.md
	@gzip dest/kaiji-$(1)-$(2).tar
endef

compress: build
	@mkdir -p dest

	$(call compress,darwin,amd64,kaiji)
	$(call compress,darwin,arm64,kaiji)
	$(call compress,windows,amd64,kaiji.exe)
	$(call compress,windows,arm,kaiji.exe)
	$(call compress,linux,amd64,kaiji)
	$(call compress,linux,arm,kaiji)


clean: ./bin
	@rm -rf ./bin
	@rm -rf ./dest
