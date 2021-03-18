package main

import (
	"context"
	"github.com/docker/docker/client"
)

func main() {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}


	c.ImageBuild(context.TODO(), )
}