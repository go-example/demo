package main

import "fmt"

type User1 struct {
	id   int
	name string
}

func (u *User1) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", u, u)
}

func (u User1) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &u, u)
}

func main() {
	u := User1{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User1.TestValue
	mv(u)

	mp := (*User1).TestPointer
	mp(&u)

	mp2 := (*User1).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2(&u)
}
