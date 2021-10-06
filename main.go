package main

import "fmt"

func main()  {
	n:=5
	var zbs []int
	for i:=0;i<n;i++{
		zb:=i
		fmt.Scanf("%d",&zb)
		zbs=append(zbs, zb)
	}
	fmt.Println(zbs)
}
