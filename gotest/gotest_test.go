package gotest

import(
	"testing"

)

//test one: normal input 
func TestTrain1(t *testing.T){
	if i := Train("abccda"); i != 'b'{ 
        	t.Error("fail") 
    	} else {
		t.Log("ok") 
    }
}

//test two: normal input, each character appeared once 
func TestTrain2(t *testing.T){
	if i := Train("abcde"); i != 'a'{ 
        	t.Error("fail") 
    	} else {
		t.Log("ok") 
    }
}

//test three: normal input, only one character
func TestTrain3(t *testing.T){
	if i := Train("a"); i != 'a'{ 
        	t.Error("fail") 
    	} else {
		t.Log("ok") 
    }
}

//test four: abnormal input, ""
func TestTrain4(t *testing.T){
	if i := Train(""); i != '0'{ 
        	t.Error("fail") 
    	} else {
		t.Log("ok") 
    }
}

//test five: abnormal input,"123"
func TestTrain5(t *testing.T){
	var str string = "123ab"
	var bs []byte = []byte(str)
	for i := 0;i<len(bs);i++{
		if(bs[i]<'a'||bs[i]>'z'){
			t.Error("abnormal input,please type again")
			break
		}else{
			t.Log("input is ok") 
    		}
	}
}

