package main

import (

	"bufio"

	"fmt"

	"os"

	"strings"
)

var num int = 10
var cnum chan int

func main(){

	fmt.Println("输入一行字符串,程序执行后将会输出其中第一个只出现一次的字符")

	reader := bufio.NewReader(os.Stdin)

	data, _, _ := reader.ReadLine()

	str_in := string(data)


	//多进程并发,num为进程数量

	cnum = make (chan int,num)
	for i := 1;i<num;i++{
		go routine(str_in,i)

	}

	fmt.Println("====主线程====")

	fmt.Println("输出字符为:",string(Find(str_in)))
	fmt.Println("--------------------------")

	//主进程等待子进程

	for i := 1;i<num;i++{
		<-cnum
	}
	fmt.Println("Done")
}


func routine(str_in string, num int){

	fmt.Println("====线程:",num)

	fmt.Println("输出字符为:",string(Find(str_in)))
	fmt.Println("--------------------------")
	cnum<-1

}



func Find(str_in string) byte {

	fmt.Println(str_in)


	//空字符串处理
	if len(str_in) == 0{
		return '0'
	}

	//字符串长度为1处理
	if len(str_in) == 1 {

		return str_in[0]
	}


	//非法字符处理
	var bs []byte = []byte(str_in)
	for i := 0;i<len(bs);i++{
		if(bs[i]<'a'||bs[i]>'z'){
			panic("非法字符，请重新输入")
		}
	}

	//寻找目标字符
	c :=str_in[0]

	str_tmp := str_in[1:len(str_in)]

	str_last :=""

	for {

		if (strings.Contains(str_tmp,string(c))) {


			str_last = strings.Replace(str_tmp,string(c),"",-1)

			if len(str_last)>0{

				str_tmp = str_last[1:len(str_last)]

			}else{
				return '0'
			}

			c = str_last[0]


		}else {

			return c
		}


	}
	return '0'
}
