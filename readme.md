go mod init github.com/phamngocquang0072/rssagg

go build

$env:PORT = "8080"

echo $env:PORT

go get github.com/joho/godotenv

go mod vendor  
go mod tidy

go mod vendor // refresh
