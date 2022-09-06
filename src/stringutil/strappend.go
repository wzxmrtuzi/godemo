package stringutil

import (
	"fmt"
	"regexp"
	"strings"
)

func Append(source, appendStr string) string {
	return source + appendStr
}

func StringTest() {
	str1 := "123a啊"
	arr1 := []rune(str1)
	fmt.Println(len(str1))
	for _, v := range arr1 {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	// 查找字串
	fmt.Println("aaa", strings.Index(str1, "1啊"))
	fmt.Println(strings.IndexRune(str1, '啊'))
	// 将字符拆开查找,有一个找到了就返回
	fmt.Println("b", strings.IndexAny(str1, "2啊"))

	fmt.Println(strings.Contains(str1, "a啊"))
	// 比较两个字符串大小, 会逐个字符地进行比较ASCII值
	// 第一个参数 >  第二个参数 返回 1
	// 第一个参数 <  第二个参数 返回 -1
	// 第一个参数 == 第二个参数 返回 0
	fmt.Println(strings.Compare("abc", "bcd"))

	// 判断两个字符串是否相等, 可以判断字符和中文
	// 判断时会忽略大小写进行判断
	fmt.Println(strings.EqualFold("Abc", "abc"))

	strArr := [3]string{"r", "g", "h"}
	fmt.Println(strings.Split(str1, "a"))
	fmt.Println(strings.Join(strArr[0:], "格"))

	fmt.Println(strings.Replace(str1, "a", "有有有", -1))

	// 正则表达式
	reg := regexp.MustCompile("a")
	fmt.Println(reg.FindAllString(str1, -1))
}
