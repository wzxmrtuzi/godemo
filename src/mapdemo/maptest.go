package mapdemo

import (
	"fmt"
)

/*
map翻译过来就是字典或者映射, 可以把map看做是切片的升级版
切片是用来存储一组相同类型的数据的, map也是用来存储一组相同类型的数据的
在切片中我们可以通过索引获取对应的元素, 在map中我们可以通过key获取对应的元素
切片的索引是系统自动生成的,从0开始递增. map中的key需要我们自己指定
只要是可以做==、!=判断的数据类型都可以作为key(数值类型、字符串、数组、指针、结构体、接口)
map的key的数据类型不能是:slice、map、function
map和切片一样容量都不是固定的, 当容量不足时底层会自动扩容
map格式:var dic map[key数据类型]value数据类型
*/
func MapTest() {
	var dict map[int]string = map[int]string{0: "zhangsan", 2: "lisi"}
	dict[0] = "222"
	// 添加,如果key不存在就会自动添加
	dict[4] = "当当当"
	// 删除
	delete(dict, 2)
	fmt.Println(dict, dict[2])
	// 查询,判断是否存储
	value, ok := dict[2]
	fmt.Println("find", value, ok)

	createMapByMake()

	foreachMap()
}

func createMapByMake() {
	// 语法糖
	// var dict map[int]string = map[int]string{0: "zhangsan", 2: "lisi"}
	// make(map[key类型]value类型,容量)
	// var dict = make(map[int]string, 3)
	// make(map[key类型]value类型)
	var dict = make(map[int]string)
	dict[2] = "啊啊啊"
	fmt.Println(dict)
}

func foreachMap() {
	var dict map[int]string = map[int]string{0: "zhangsan", 2: "lisi"}
	for key, value := range dict {
		fmt.Println("foreach", key, value)
	}
}
