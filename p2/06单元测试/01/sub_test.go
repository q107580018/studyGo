package cal

import "testing"

func TestGetSub(t *testing.T){
res := GetSub(10,3)
if res != 7{
	t.Fatalf("GetSub(10,3)执行错误,期望值%v，实际值%v", 7,res)
	
}
t.Logf("GetSub(10,3)执行正确！！！")
}
