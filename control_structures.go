//Go If statement
//package main
//
//import "fmt"
//
//func main(){
//	var a int= 40
//	if (a>20){
//		fmt.Println("greater than 20 ")
//
//	}
//	//fmt.Println("mmmh")
//}
/*package main

import "fmt"

func main(){
	var a string ="k"
	if a=="hello"{
		fmt.Println("the word is",a)
	}else{
		fmt.Println("the word is not Hello")
	}*/


//FOR LOOP

/*package main

import "fmt"

func main(){
	var x int
	for x=0;x<=50;x++{
		if x%2==0{
			fmt.Println("The even number is :=",x)
		}else {
			fmt.Println("The odd number is := ",x)
		}
	}
}*/

//SWITCH CASE STATEMENT
package main

import "fmt"

func main(){
	var i int =5
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}
}