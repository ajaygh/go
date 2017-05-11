package main

import (
	"fmt"
	//	"github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
	"strings"
)

func main() {
	flow.New().TextFile(
		"/etc/passwd", 3,
	).Filter(func(line string) bool {
		return !strings.HasPrefix(line, "#")
	}).Map(func(line string, ch chan string) {
		for _, token := range strings.Split(line, ":") {
			ch <- token
		}
	}).Map(func(key string) int {
		return 1
	}).Reduce(func(x, y int) int {
		return x + y
	}).Map(func(x int) {
		fmt.Println("count: ", x)
	}).Run()

}
