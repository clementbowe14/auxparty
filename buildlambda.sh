cd $1
rm $(find ./cmd -type f ! -name "main.go")
rm $(find ./build -type f -name "*")
GOARCH=amd64 GOOS=linux go build -o build/main cmd/main.go
cd build
zip main.zip main