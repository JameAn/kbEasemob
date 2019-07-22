# 环信SDK
 ## Installation
 
 To install this package, you need to install Go and set your Go workspace first.
 
 1. Download and install it:
 
 ```sh
 $ go get -u  
 ```
 
 2. Import it in your code:
 
 ```go
```

## Quick start
```sh
# assume the following codes in example.go file
$ cat example.go
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
```