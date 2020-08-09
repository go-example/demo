package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	c := flag.Int("c",0,"this is c")
	d := flag.String("d","dd","this is d")
	e := flag.Bool("e",false,"this is e")
	f := flag.String("f","","this is f")
	flag.Parse()
	a := flag.Arg(0)
	b := flag.Arg(1)

	env := os.Args
	fmt.Println(a,b,*c,*d,*e,*f,env)
	//go run test.go -c=99 --d=88 -e=true dev sub
}
