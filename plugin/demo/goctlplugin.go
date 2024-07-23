package main

import (
	"fmt"

	"github.com/bendh1/goctls/plugin"
)

func main() {
	plugin, err := plugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	if plugin.Api != nil {
		fmt.Printf("api: %+v \n", plugin.Api)
	}
	fmt.Println("Enjoy anything you want.")
}
