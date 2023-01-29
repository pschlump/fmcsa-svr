

all:
	go build

run:
	go build
	./kill-fmcsa-svr.sh
	check-json-syntax --ignore-tab-warning ./cfg.json
	./fmcsa-svr >>log/output.log 2>&1 &
	sleep 1
	./test/tgo_is_running.sh
	@echo PASS | color-cat -c green

# (base) philip@victoria gin1 % go build
# # golang.org/x/sys/unix
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/syscall_darwin.1_13.go:25:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:27:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.1_13.go:40:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:28:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:43:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:59:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:75:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:90:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:105:3: //go:linkname must refer to declared function or variable
# ../../../../../pkg/mod/golang.org/x/sys@v0.0.0-20200116001909-b77594299b42/unix/zsyscall_darwin_amd64.go:121:3: //go:linkname must
#
fix_error:
	 go get -u golang.org/x/sys




linux:
	GOOS=linux GOARCH=amd64 go build -o fmcsa-svr_linux

deploy:
	scp fmcsa-svr_linux fmcsa-systemd.service philip@45.79.53.54:/home/philip/tmp


