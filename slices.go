package main

import "fmt"

func main(){
	x:=make([]int,6,8)
	fmt.Println(x)
	//fmt.Println(sliceFunc)
	//APPENDING
	slice1:=[]int{5,5,6,7}
	slice2:=append(slice1,4,5,6)
	fmt.Println(slice1,slice2)

	//COPYING

	slice3:=[]int{2,3,4,5}
	slice4:=make([]int ,3)
	copy(slice4,slice3)
	fmt.Println(slice3,slice4)

}

//func sliceFunc() int {
////	append
//slice1:=[]int{5,5,6,7}
//slice2:=append(slice1,4,5,6)
////fmt.Println(slice1,slice2)
//
//	return slice1,slice2
//
//}