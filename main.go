package main

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return "Hello :world_map:!"
}

func main() {
	var text string = GetMessage()
	emoji.Println(text)
}
