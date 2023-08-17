// package main
//
// import "fmt"
//
//	func main() {
//		fmt.Print("Hello, World!\n" + "Dapao")
//		var username, sex string
//		username = "老张"
//		sex = "男"
//		fmt.Println(username, sex)
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"runtime"
//
// )
//
//	func main() {
//		fmt.Printf("%s", runtime.Version())
//
// }
// package main
//
// import (
//
//	"fmt"
//	"os"
//	"runtime"
//
// )
//
//	func main() {
//		var goos string = runtime.GOOS
//		fmt.Printf("The operating system is : %s\n", goos)
//		path := os.Getenv("PATH")
//		fmt.Printf("Path is %s\n", path)
//
//		//var username string
//		//username = "哈喽"
//		//a := 5.0
//		//b := int(a)
//		//const Pi = 3.1415926
//		//fm.Println(Pi + 1)
//		//
//		//fm.Println("Hello,Go is Good")
//		//fm.Println(username)
//		//fm.Println(b)
//		//const (
//		//	a = iota
//		//	b
//		//	c
//		//	d = 5
//		//	e
//		//)
//
// }
// package main
//
// var a string
//
//	func main() {
//		a = "G"
//		print(a)
//		f1()
//	}
//
//	func f1() {
//		a := "X"
//		print(a)
//		f2()
//
// }
//
//	func f2() {
//		print(a)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		var a int
//		var b int32
//		var c int
//		a = 15
//		b = int32(a + a)
//		c = int(b + 5)
//		fmt.Println(a, b, c)
//
// }
// package main
//
// import "fmt"
//
//	func main() {
//		var n int16 = 34
//		var m int32
//		m = int32(n)
//		fmt.Printf("32 bit int is:%d\n", m)
//		fmt.Printf("16 bit is :%d\n", n)
//
// }
// package main
//
// import (
//
//	"fmt"
//	"math/rand"
//	"time"
//
// )
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			a := rand.Int()
//			fmt.Printf("%d /", a)
//		}
//		for i := 0; i < 5; i++ {
//			r := rand.Int()
//			fmt.Printf("%d /", r)
//		}
//		fmt.Println()
//		timens := int64(time.Now().Nanosecond())
//		rand.Seed(timens)
//		for i := 0; i < 10; i++ {
//			fmt.Printf("%2.2f /", 100*rand.Float32())
//		}
//	}
//
// package main
//
// import "fmt"
//
// type TZ int
//
//	func main() {
//		var a, b TZ = 3, 4
//		c := a + b
//		fmt.Printf("c has the value:%d", c)
//	}
package main

import "fmt"

var ch int = '\u0041'
var ch2 int = '\u03B2'
var ch3 int = '\U00101234'

func main() {
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
