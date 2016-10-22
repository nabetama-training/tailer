package main

func main() {

	nginxLog := &TargetFile{
		FilePath: "/var/log/nginx/access.log",
		Follow: true,
	}
	nginxLog.Tail()
}
