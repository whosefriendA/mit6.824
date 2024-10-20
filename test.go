// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }
// func main() {
// 	fmt.Println("helloworld")
// 	for i:=0; i<5; i++{

// 	}
// }
// func main() {
// 	i:=10
// for {
// 	i++
// }
// fmt.Println(i);
// }
// func Sqrt(x float64) float64 {
// 	z:=1.0
// 	for z1:=float64(0);z1!=z;{
// 		z1=z
// 		z-=(z*z-x)/(2*z)
// 		fmt.Println(z)
// 		fmt.Println(z1)
// 	}
// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }
// func main() {
// 	i, j := 42, 2701
// 	p := &i         // 指向 i
// 	fmt.Println(*p) // 通过指针读取 i 的值
// 	*p = 21         // 通过指针设置 i 的值
// 	fmt.Println(i)  // 查看 i 的值

// 	p = &j         // 指向 j
// 	*p = *p / 37   // 通过指针对 j 进行除法运算
// 	fmt.Println(j) // 查看 j 的值
// }
// func main(){
// 	var a []int = []int{1,2,3,4}
// 	var b []int = a[1:2]
// 	fmt.Print(a)
// 	fmt.Print(b)
// }
// func main(){
// 	a:=make([]int,0,5)
// 	b:=make([]int ,3,4)
// 	fmt.Println(a,b)
// }
package main

import ("fmt")

func main(){
	a :=[4]int{
		1,2,3,4,
	}
	b:=a[1:2]
	c:=a[:1]
	d:=a[:3]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	for i,v :=range(a){
		fmt.Println(i)
		fmt.Println(v)
	}
}