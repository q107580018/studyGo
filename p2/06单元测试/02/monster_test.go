package monster

import "testing"

func TestReStore(t *testing.T) {
	m := &Monster{
		Name:  "牛魔王",
		Age:   18,
		Skill: "打架",
	}
	m.Store()
	m2 := new(Monster)
	m2.ReStore()
	if m2.Name != m.Name {
		t.Fatalf("测试失败，期待值：%v, 实际值:%v\n", m.Name, m2.Name)
	}
	t.Log("测试成功")
}
