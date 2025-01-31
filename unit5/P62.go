package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误：", err)
		panic(err)
	}
	fmt.Println("上面的除法操作成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("can't use 0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
