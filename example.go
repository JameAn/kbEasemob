package main

import (
	"fmt"
	"github.com/kbEasemob"
)

func main() {
	iu := kbEasemob.NewImUser("test1", "123456")
	err := iu.Register()
	if err != nil {
	    fmt.Println(err)
	}
	
}
