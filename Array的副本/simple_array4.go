package Array

import (
	"fmt"
	"strconv"
)
//通过转化为字符串，逆序字符串比较
func IsPalindrome(x int) bool {
	//int->string
	temp := strconv.Itoa(x)
	fmt.Println(temp)
	//字符串逆序
	var bytes []byte = []byte(temp)
	for i := 0; i < len(temp)/2; i++ {
		// 定义一个变量存放从后往前的值
		tmp := bytes[len(temp)-i-1]
		// 从后往前的值跟从前往后的值调换
		bytes[len(temp)-i-1] = bytes[i]
		// 从前往后的值跟从后往前的值进行调换
		bytes[i] = tmp
	}
	temp1 := string(bytes)


	fmt.Println(temp1)
	if temp == temp1{
		return true
	}else{
		return false
	}
}
/*
自己写的方法-----通过逆序数字判断
func isPalindrome(x int) bool {
    temp := x
    if x < 0 {
        return false
    }
    if x >= 0 {
        var res = 0
        for x!= 0{
            res = res*10 + x%10
              x /= 10
        }

        if temp == res{
            return true
        }else{
            return false
        }
    }
    return true
}
*/