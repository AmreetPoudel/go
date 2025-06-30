package main

import (
	"fmt"
	"strings"
)

func main() {

	//using builder since string are immutable
	var amritbuilder strings.Builder
	amritbuilder.WriteString("my ")
	amritbuilder.WriteString("name is ")
	amritbuilder.WriteString("amrit poudel")

	str := amritbuilder.String()
	fmt.Println(str)
}
