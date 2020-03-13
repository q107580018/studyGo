package cal

import "testing"

func TestAddUpper(t *testing.T) {

	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误,期望值=%v，实际值为=%v", 55, res)
	}
	t.Log("AddUpper(10)执行正确...")
}

