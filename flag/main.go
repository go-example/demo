// 这些示例演示了对标志包的更复杂的使用。
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// 示例1：名为“species”的单个字符串标志，默认值为“gopher”。
var species = flag.String("species", "gopher", "the species we are studying")

// 示例2：共享变量的两个标志，因此我们可以使用速记。
// 初始化的顺序是未定义的，因此请确保两者都使用
// 相同的默认值。 必须使用init函数设置它们。
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// 示例3：用户定义的标志类型，持续时间片。
type interval []time.Duration

// String是格式化标志值的方法，是flag.Value接口的一部分。
// String方法的输出将用于诊断。
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set是设置标志值的方法，flag.Value接口的一部分。
// Set的参数是要解析以设置标志的字符串。
// 这是一个以逗号分隔的列表，因此我们将其拆分。
func (i *interval) Set(value string) error {
	// 如果我们想允许多次设置标志，
	// 累积值，我们将删除此if语句。
	// 这将允许诸如此类的用法
	//	-deltaT 10s -deltaT 15s
	// 和其他组合。
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// 定义一个标志来累积持续时间。 因为它有特殊的类型，
// 我们需要使用Var函数，因此在期间创建标志
// init。

var intervalFlag interval

func init() {
	// 将命令行标志绑定到intervalFlag变量和
	// 设置用法消息。
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// 所有有趣的部分都是上面声明的变量，但是
	// 要使标志包能够看到那里定义的标志，就必须这样做
	// 执行，通常在main（不是init！）的开头执行：
	//	flag.Parse()
	// 我们不在这里运行它，因为这不是主要功能
	//测试套件已经解析了标志。
}