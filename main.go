package main

import (
	"ChildrenMath/routers"
)

func main() {
	r := routers.InitRouter()

	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
