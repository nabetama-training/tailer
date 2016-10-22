GOARCH=amd64 GOOS=linux go build -o tailer .
scp tailer logviewer:bin/
