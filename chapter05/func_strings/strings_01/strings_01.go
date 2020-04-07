package main

import (
	"fmt"
	"strconv"
	"strings"
)

// go  字符串常用函数
func main() {

	//1.  内建函数len返回 v 的长度，这取决于具体类型
	str := "hello"
	fmt.Println(" str len=", len(str))

	//2. 字符串遍历存在中文的时候
	str2 := "hello北京"
	runes := []rune(str2)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("字符 #%c \n", runes[i])
	}
	//3. 字符串转整数   ：
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败，", err)

	} else {
		fmt.Println("num =", num)
	}
	//4. 整数转字符串
	itoa := strconv.Itoa(12345)
	fmt.Println("itoa =", itoa)
	//5. 字符串转成 []byte
	var bytes = []byte("hello go demo")
	fmt.Printf("bytes==%v \n", bytes) //打印的是ascii

	//6. [] byte转换成string
	str3 := string([]byte{98, 99, 100})
	fmt.Printf("str3= %v \f", str3)
	//7. 10进制转换成2 8 16进制

	formatInt := strconv.FormatInt(123, 2)
	fmt.Printf("123 转成2进制的value= %v \n", formatInt)
	//8. 查找子串中是否包含指定的字符串
	str4 := "hello chong qin"
	contains := strings.Contains(str4, "qin")
	fmt.Printf("str= %v is contains'qin' \n", contains)

	//9. 统计一个字符串中包含几个不重复的字符串子串
	countNum := strings.Count(str4, "o")
	fmt.Printf("countNum=%v \n", countNum)
	//10. 不区分大小写的字符串比较  返回bool
	fold := strings.EqualFold("abd", "Abd")
	fmt.Printf("fold=%v \n", fold)

	//11. 返回子串在字符串最后一次出现的index,如果没有返回 -1
	lastIndex := strings.LastIndex("go goland", "go")
	fmt.Printf("index= %v \n ", lastIndex)

	//12. 将指定的子串替换成，另外一个子串，
	replaceStr := strings.Replace("go go chongqing", "go", "重庆", 1)
	fmt.Println("replaceStr", replaceStr)

	//13. 按照指定的字符串，以 "，"为分隔符，将一个字符串拆分成字符串数组
	splitString := strings.Split("Hello,world,ni,hao,chongqing", ",")
	for i := 0; i < len(splitString); i++ {
		fmt.Printf("str[%v]=%v \n", i, splitString[i])
	}

	//14. 将字符串的字母进行大小写的转换
	strLower := "goland Hello"
	toLower := strings.ToLower(strLower)
	fmt.Printf("toLower= %v \n ", toLower)

	//15. 将字符串左右两边的空格去掉，
	trimSpace := strings.TrimSpace(" Hello go developer go goland ")
	fmt.Printf("trimSpace= %q \n ", trimSpace)

	//16.  将字符串左右两边的指定字符去掉
	trim := strings.Trim(" !hello world ! ", " !")
	fmt.Printf("trim= %q \n ", trim)

	//17. 判断字符串是否以执行的字符串开头
	hasPrefix := strings.HasPrefix("Http://127.0.0.1", "Http")
	fmt.Printf("hasPrefix= %v \n ", hasPrefix)

}
