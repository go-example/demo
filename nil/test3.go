package main

import (
	"fmt"
)

func main() {
	type IntPtr *int
	fmt.Println(IntPtr(nil) == (*int)(nil))        //true
	fmt.Println((interface{})(nil) == (*int)(nil)) //false
}