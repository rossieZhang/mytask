
package gotest

import(
    "fmt"
    "strings"
)


func Train(str_in string) byte {
    fmt.Println(str_in)
    if len(str_in) == 0{
        return '0'
    }
    if len(str_in) == 1 {
        return str_in[0]
    }
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
}

