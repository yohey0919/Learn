//Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var graph = make(map[string]map[string]bool) //等价于 JAVA中的 Map<String,Map<String,Boolean>>

// 惰性初始化map是一个惯用方式，也就是说在每个值首次作为key时才初始化
func addEdge(from string, to string) { // (from string, to string) <=> (from, to string)
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

// slice并不是一个纯粹的引用类型，它实际上是一个类似下面结构体的聚合类型
type IntSlice struct {
	ptr      *int // slice值包含指向第一个slice的指针，向函数传递slice将允许修改底层数组的元素
	len, cap int
}

// 结构体的值不能包含它自身，但是结构体可以包含 *自身指针类型的成员
type tree struct {
	value       int
	left, right *tree
}

// 结构体中可以包含其他结构体
type Employee struct {
	ID        int //成员名称:选择用大写字母开头
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
	//it        IntSlice
}

// 数组反转
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Println("-----------复合类型")
	fmt.Println("1) 数组--------------------------------------------------------")
	var q [3]int = [3]int{1, 2, 3}
	// q= [4]int{1, 2, 3,4}	 //编译失败，数组是固定长度的
	for _, value := range q {
		fmt.Print(strconv.Itoa(value) + ",")
	}
	fmt.Println()
	symbol := [...]string{1: "$", 2: "€", 3: "￡", 7: "￥"} // ... 通过{}里的值固定数组长度,此时len为8
	fmt.Println(symbol[7])                                //symbol[7] ="￥", symbol[4] = ""
	// 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)                                          // "true false false"
	fmt.Println("2) 切片--------------------------------------------------------") //slice的语法和数组很像，只是没有固定长度而已
	sli := symbol[1:3]                                                           //symbol[:]切片操作则是引用整个数组
	fmt.Println(sli)                                                             //[$ €]
	a2 := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a2[0:3])
	fmt.Println(a2)                 // "[2 1 0 3 4 5]"
	fmt.Println(len(sli) == 0)      //slice是否是空的，使用len(s) == 0来判断
	withoutCap := make([]string, 8) //内置的make函数创建一个指定元素类型、长度和容量的slice
	withCap := make([]string, 6, 8) // 等价于make([]T, 8)[:6]。额外的元素是留给未来的增长用的。
	fmt.Println(withoutCap, withCap)
	withCap = append(withCap, "1", "2", "3", "4") //内置的append函数用于向slice追加元素
	fmt.Println("after append", withCap)          //append可以扩展容量
	fmt.Println(cap(withCap))                     //超过之前容量,则复制一个cap()*2的切片
	stack := make([]int32, 4, 16)                 //一个slice可以用来模拟一个stack
	stack = append(stack, 1)                      // push
	stack = stack[:len(stack)-1]                  // pop
	fmt.Println("3) Map--------------------------------------------------------")
	mp := make(map[string]int, 16) // 内置的make函数可以创建一个map
	mp["alice"] = 31
	mp["charlie"] = 34
	fmt.Println(mp)           // map[alice:31 charlie:34]
	fmt.Println(len(mp))      //2
	ageMap := map[string]int{ // 直接声明
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ageMap["charlie"])      //34
	delete(ageMap, "alice")             // 删除key
	fmt.Println("after delete", ageMap) // after delete map[charlie:34]
	ageMap["bob"]++
	fmt.Println("after bob", ageMap) // after bob map[bob:1 charlie:34]
	fmt.Println(ageMap["haha"])      //0 ,如果key不存在，那么将得到value对应类型的零值
	//禁止对“可能随着元素数量的增长而重新分配更大的内存空间”的类型的元素进行取址，比如"map","slice"
	for name, age := range ageMap {
		fmt.Printf("%s\t%d\n", name, age) // 循环“无序”输出ageMap
	}
	value, ok := ageMap["bob"] // map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在
	fmt.Println(value, ok)     // 1 true
	value, ok = ageMap["haha"]
	fmt.Println(value, ok) // 0 false
	// Go语言中并没有提供一个set类型，但是map中的key也是不相同的，可以用map实现类似set的功能
	strSet := make(map[string]bool)
	strSet["111"] = false
	strSet["222"] = false
	strSet["111"] = false
	for key := range strSet { // key的循环输出可以不用 key, _
		fmt.Println("key = ", key) //key =  111 key =  222
	}
	fmt.Println("4) 结构体(struct)--------------------------------------------------------")
	var dilbert Employee
	dilbert.ID = 111 //dilbert是一个变量，它所有的成员也同样是变量
	dilbert.Name = "zhangSan"
	salaryAdd := &dilbert.Salary
	*salaryAdd = 2000
	fmt.Println(dilbert) //{111 zhangSan  0001-01-01 00:00:00 +0000 UTC  2000 0}
	var emp *Employee = &dilbert
	emp.Salary = 1000
	fmt.Println(*emp) //{111 zhangSan  0001-01-01 00:00:00 +0000 UTC  1000 0}
	(*emp).Salary = 3000
	fmt.Println(dilbert) //{111 zhangSan  0001-01-01 00:00:00 +0000 UTC  3000 0}
	Salary2(&dilbert)    // 传入指针 或者 Salary2(emp) 效果一致
	fmt.Println(dilbert) //{111 zhangSan  0001-01-01 00:00:00 +0000 UTC  6000 0}

	//结构体类型的零值是每个成员都是零值
	type Point struct{ X, Y int }
	p := Point{1, 2} //结构体值也可以用“字面值表示”
	fmt.Println(p)
	type Circle struct {
		Point      //匿名成员,只写类型 不写变量名
		Radius int `json:"released"` //结构体成员Tag
	} // 声明了 匿名成员 无法用“字面值表示”

	c3 := Circle{
		Point:  Point{X: 8, Y: 8},
		Radius: 5,
	}
	c3.X = 44 // <=> c3.Point.X = 1
	fmt.Println("4) JSON--------------------------------------------------------")
	// 结构体slice转为JSON的过程叫编组（marshaling）。编组通过调用json.Marshal函数完成
	data, err := json.Marshal(c3)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(c3)          //{{44 8} 5}
	fmt.Printf("%s\n", data) //{"X":44,"Y":8,"released":5}  有带属性名(输出紧凑)
	// 编码的逆操作是解码,通过json.Unmarshal函数完成
	var po Point
	if err := json.Unmarshal(data, &po); err != nil { //选择性地解码JSON中感兴趣的成员
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(po) //{44 8}

	data2, err2 := json.MarshalIndent(c3, "", "    ") //后面两个参数，表示每一行输出的前缀和每一个层级的缩进
	if err2 != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)
	// 整齐输出：
	// {
	// 	"X": 44,
	// 	"Y": 8,
	// 	"released": 5	//不是Radius的原因：一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串
	// }
	fmt.Println("5) 文本和HTML模板--------------------------------------------------------")
	const templ = `
	E_Employee---------
	E_Name:   {{.Name}}
	E_Address: {{.Address}}
	Title:  通过模板生成`
	var report = template.Must(template.New("ep").
		Parse(templ))
	if err := report.Execute(os.Stdout, dilbert); err != nil { // 模板一般配合结构体使用
		log.Fatal(err)
	}
	fmt.Println()
	// E_Employee---------
	// E_Name:   zhangSan
	// E_Address:
	// Title:  通过模板生成
	var htmlTemp = template.Must(template.New("ht").Parse(`<h1>{{.Name}} issues</h1>`))
	if err := htmlTemp.Execute(os.Stdout, dilbert); err != nil { // 模板一般配合结构体使用
		log.Fatal(err)
	}
	// <h1>zhangSan issues</h1>
}

func Salary2(e *Employee) {
	e.Salary = e.Salary * 2
}
