package flagT


import (
	"flag"
	"fmt"
	"os"
)

var (
	n int
	h bool
	q *bool
	s string
)

func init(){
	q = flag.Bool("q", false, "Exit")
	flag.BoolVar(&h, "h", false, "Show help")
	flag.IntVar(&n, "n", 0, "set number")
	flag.StringVar(&s, "s", "default string", "set string")
}

// flag 說明
// example 下指令 go run main.go -h 或是 go run main.go -s abc
// flag.Parse 解析 -h 參數
// 依照參數值 進行初始化動作 一般來說建置在 init()

func GoflagT(){

	flag.Parse()

	if h {
		//usage 印出預設說明
		flag.Usage()
	} else {
		if *q {
			fmt.Println("q is", *q)
			os.Exit(0)
		}
		fmt.Println("Number is ", n)
		fmt.Println("String is ", s)
	}

}