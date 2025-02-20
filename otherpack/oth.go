package otherpack

import "fmt"

var Name = "zzwsec"

func PrName(na string) {
	fmt.Println(na)
}

func init() {
	fmt.Println("这是init函数")
}
