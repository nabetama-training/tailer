GOARCH=amd64 GOOS=linux go build -o app .
scp app logviewer:bin/
