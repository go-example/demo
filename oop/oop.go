package main

import (
	"fmt"
)

type Poem struct {
	Title  string
	Author string
	intro  string
}

func (poem *Poem) publish() {
	fmt.Println("poem publish")
}

func (poem *Poem) recite(v ...interface{}) {
	fmt.Println(v)
}

func NewPoem(author string) (poem *Poem) {
	poem = &Poem{}
	poem.Author = author
	return
}

type ProsePoem struct {
	Poem
	Author string
}

func (prosePoem *ProsePoem) publish() {
	fmt.Println("prose poem publish")
}

func main() {
	//poem实例化
	poem := &Poem{}
	poem.Author = "Heine"
	poem2 := &Poem{Author: "Heine"}
	poem3 := new(Poem)
	poem3.Author = "Heine"
	poem4 := Poem{}
	poem4.Author = "Heine"
	poem5 := Poem{Author: "Heine"}

	fmt.Println(poem)
	fmt.Println(poem2)
	fmt.Println(poem3)
	fmt.Println(poem4)
	fmt.Println(poem5)

	//使用方法实例化poem
	poem6 := NewPoem("Heine")
	fmt.Println(poem6)

	//调用print方法
	poem6.recite(3)

	//ProsePoem组合Poem
	prosePoem := &ProsePoem{}
	prosePoem.Author = "Shelley"
	prosePoem.Poem.Author = "Heine"
	prosePoem.Poem.publish()
	prosePoem.publish()

	fmt.Println(prosePoem)

}
