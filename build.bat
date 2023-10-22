set VERSION=v1.0.0
wails build -clean -obfuscated -trimpath -ldflags "-w -h -H windowsgui" -platform "windows/arm64" -upx -o LdifGraph_amd64_%VERSION%.exe
