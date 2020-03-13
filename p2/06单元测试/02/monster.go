package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m Monster) Store() {
	buf, _ := json.MarshalIndent(m, "", " ")
	f, err := os.Create("1.json")
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f.Close()
	_, err = f.Write(buf)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
}
func (m *Monster) ReStore() {
	buf, err := ioutil.ReadFile("1.json")
	if err != nil {
		fmt.Println("read fail:", err)
		return
	}
	json.Unmarshal(buf, m)
}
