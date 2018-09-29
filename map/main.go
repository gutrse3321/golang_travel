package main

import "fmt"

func main() {
	// 1.创建一个简单的映射
	firstMap()
	// 输出一个不存在的键和值，0对应int类型初始值
	// console: 0 false

	// 2.len的使用与delete的使用
	secondMap()
	// 声明创建一个映射，第一次输出他的len
	// delete(lookup, "goku"),第二次输出删除映射中的键的len
	// console: 0 回车 1

	// 3. 设置初始化大小
	thirdMap()
	// 设置第二个参数后，输出len，无创建映射，故输出0的长度
	// "定义时指定一个初始大小可以获得一定的性能提升"
	// 说的是这样的...
	// console: 0

	// 4. 映射作为一个结构体的字段
	fouthMap()
	// console: Krillin

	// 5. 复合初始化和遍历
	fifthMap()
	// 编译映射不是有序的，每次遍历返回的键值对都是随机的
	//console: key: goku value: 9001
	//console: key: gohan value: 2044

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("aswwwwke"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefg"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func firstMap() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	fmt.Println(power, exists)
}

func secondMap() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	fmt.Println(len(lookup))

	delete(lookup, "goku")
	fmt.Println(len(lookup))
}

func thirdMap() {
	lookup := make(map[string]int, 100)
	fmt.Println(len(lookup))
}

type Saiyan struct {
	name    string
	Friends map[string]*Saiyan
}

func fouthMap() {
	goku := &Saiyan{
		name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	goku.Friends["krillin"] = &Saiyan{
		name:    "Krillin",
		Friends: make(map[string]*Saiyan),
	}

	fmt.Println(goku.Friends["krillin"].name)
}

func fifthMap() {
	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	for key, value := range lookup {
		fmt.Println("key:", key, "value:", value)
	}
}
