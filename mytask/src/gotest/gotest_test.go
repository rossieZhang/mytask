package gotest

import(
    "testing"
)


//常规输入测试一：存在只出现一次的字符 
func TestTrain1(t *testing.T){
    if i := Train("abccda"); i != 'b'{
        t.Error("测试一失败")
    } else {
        t.Log("测试一成功")
    }
}


//常规输入测试二：不存在只出现一次的字符
func TestTrain2(t *testing.T){
    if i := Train("aabb"); i != '0'{
        t.Error("测试二失败")
    } else {
        t.Log("测试二成功")
    }
}


//常规输入测试三：字符串长度为1
func TestTrain3(t *testing.T){
    if i := Train("a"); i != 'a'{
        t.Error("测试三失败")
    } else {
        t.Log("测试三成功")
    }
}


//常规输入测试四：空字符
func TestTrain4(t *testing.T){
    if i := Train(""); i != '0'{
        t.Error("测试四失败")
    } else {
        t.Log("测试四成功")
    }
}


//非常规测试：非法字符
func TestTrain5(t *testing.T){
    var str string = "123ab"
    var bs []byte = []byte(str)
    for i := 0;i<len(bs);i++{
        if(bs[i]<'a'||bs[i]>'z'){
            t.Error("非法字符，请重新输入")
            break
        }
    }
}
