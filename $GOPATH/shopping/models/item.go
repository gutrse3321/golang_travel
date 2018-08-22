// 创建共享包，解决循包报错
// 并且不导入任何东西
package models

// 可见性，golang规定类型或函数是否对外部的包可见
// 大小写设置，大写可见
type Item struct {
	Price float64
}
