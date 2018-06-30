package main

import (

    "bufio"
    "fmt"
    "os"
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

    //非法字符处理
    var bs []byte = []byte(str_in)
    for i := 0;i<len(bs);i++{
        if(bs[i]<'a'||bs[i]>'z'){
            panic("非法字符，请重新输入")
        }
    }

    //寻找目标字符
    var CharCountMap map[byte]int
    CharCountMap = make(map[byte]int)
    for i := 0;i<len(str_in);i++{
        count, ok := CharCountMap[byte(str_in[i])]
        if (ok) {
            count=-1
            CharCountMap[byte(str_in[i])] = count
        } else {
            CharCountMap[byte(str_in[i])] = i
        }
    }
    minindex := len(str_in)
    for char_key := range CharCountMap {
        if CharCountMap[char_key] != -1 {
            if CharCountMap[char_key]<minindex{
                minindex=CharCountMap[char_key]
            }
        }
    }
    if minindex == len(str_in){
        return '0'
    }
    return str_in[minindex]
}
