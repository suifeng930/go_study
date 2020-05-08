package main

import (
	"go_study/chapter10/test_demo/testingDemo/testcase"
	"testing"
)

func main() {

}

//编写一个测试用例，去测试addUpper 是否正确

func TestAddUpper(t *testing.T) {
	// 调用
	res := testcase.AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望值为%v 实际值为%v", 55, res)
	}
	t.Logf("AddUpper(10) 执行正确...")
}
