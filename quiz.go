/*Write a program that finds the smallest number
in this list:
x := []int{
48,96,86,68,
57,82,63,70,
37,34,83,27,
19,97, 9,17,
}*/

package main

import (
	"fmt"
	"sort"
)

func main()  {

	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	sort.Ints(x)
	fmt.Println(x)
}



























/*package main

import "fmt"

func main(){
		var temp int
		var i int

	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	for i=0;i<=len(x);i++{
		if x[i]<x[i+1]{
			temp =x[i]
			x[i]= x[i+1]
			x[i+1]=temp


		}
		for i=0;i<=len(x);i++{
			fmt.Println(x[i])
		}
	}



}*/


