package main

import "fmt"

var count int =0

func sum(num int) int  {
	sum := 1
	for i := 1;i<=num ;i++{
		sum = sum * i
	}
	return sum
}
//交换数据
func  swap(a *int,b *int)  {
	var temp = *a

	*a = *b
	*b = temp

}
func ReturnFunc (nums []int,m int,max int,re [][]int) [][]int  {
	if m == max-1{
		var temp = make([]int,max)
		for i := 0;i<max;i++{
			temp[i] = nums[i]
		}
		re[count] = temp
		count ++
	}else{
		for i := m; i < max ;i++{
			swap(&nums[i],&nums[m])
			ReturnFunc(nums,m+1,max,re)
			swap(&nums[i],&nums[m])
		}
	}
	return re
}
func permute (nums []int) [][]int {
	 temp := make([][]int,sum(len(nums)))
	 re := ReturnFunc(nums,0,len(nums),temp)
	 return re
 }
func main(){
	var n int
	fmt.Scanf("%d",&n)
	TestSlice := make([]int, n)
	for i := 0;i < n;i++{
		fmt.Scan(&TestSlice[i])
	}
	res := permute(TestSlice)
	fmt.Println(res)
}