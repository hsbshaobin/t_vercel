set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-w -s" -o .\bin\adCheck .\cmd\api\
