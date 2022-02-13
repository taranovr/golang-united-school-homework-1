package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return "Hello :world_map:!"
}

func main() {
	var text string = emoji.Sprint(GetMessage())
	fmt.Println(text)
}
