package main

import "fmt"
//11,9,2,8,3,7,4,6,5,10
//9,11,2,8,3,7,4,6,5,10//
//9,2,11,8 3 7 4 6 5 10
//9 2 8 11 3 7 4 6 5 10
//两两比较每次取最大值，最后把最大的数往后移
func BubbleFindMax(arr []int)int{
	length:=len(arr)
	if length==1{
		return arr[0]
	}else{
		for i:=0;i<length-1;i++{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
			}
		}
		return arr[length-1]
	}

}

func Bubblesort(arr []int)[]int{

	length:=len(arr)
	if length==1{
		return arr
	}else{
		for i:=0;i<length-1;i++{//只剩一个不需要冒泡了
		isneedexchange:=false
		for j:=0;j<length-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
				isneedexchange=true
			}
		}
		fmt.Println(isneedexchange)
		if !isneedexchange{
			break
		}
		fmt.Println(arr)
		}
		return arr
	}
}


func main(){
	arr:=[]int{11,9,2,8,3,7,4,6,5,10}
	fmt.Println(BubbleFindMax(arr))
	fmt.Println(Bubblesort(arr))
}
