pi:
	GOOS=linux GOARCH=arm GOARM=5 go build . && upx -9 piper
