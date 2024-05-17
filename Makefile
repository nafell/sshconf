BINARY_NAME=sshconf
OUT_PATH=./out
LINUX_BINARY=${OUT_PATH}/${BINARY_NAME}
WIN_BINARY=${OUT_PATH}/${BINARY_NAME}.exe

build:
	go mod tidy
	mkdir -p ${OUT_PATH}
	go build -o ${LINUX_BINARY} main.go
	GOOS=windows GOARCH=amd64 go build -o ${WIN_BINARY} main.go

clean:
	go clean
	rm ${LINUX_BINARY}
	rm ${WIN_BINARY}
