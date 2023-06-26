package main

import "github.com/robert-min/performance_go_python_api/go_gin/server/handler"

func main() {
	r := handler.Register()
	r.Run(":8000")
}
