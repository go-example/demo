package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//写
	str := `好像
	113322
afsadf
adsfadf11122
`
	file.WriteString(str)

	//读
	reader := bufio.NewReader(file)
	var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line = append(line, data...)
		if !prefix {
			fmt.Printf("data:%s\n", string(line))
			line = line[:]
		}

	}
}