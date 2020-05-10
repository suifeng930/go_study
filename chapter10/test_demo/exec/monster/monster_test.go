package monster

import "testing"

func TestMonster_Store(t *testing.T) {

	// 先创建一个对象实例
	monster := &Monster{
		Name:  "jack",
		Age:   0,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.store() 错误，期望是%v,实际为%v", true, res)
	}
	t.Log("monster.store() 执行完成")

}
