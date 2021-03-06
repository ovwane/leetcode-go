package main

import "fmt"

func main() {
	arr := []int{4,3,2,1}
	ms(arr)
	fmt.Println(arr)
}

func ms(arr[]int) {
	if len(arr) == 0 || len(arr) == 1 { return }

	left,right := arr[:len(arr)/2], arr[len(arr)/2:]

	ms(left)
	ms(right)

	// merge
	merged := make([]int, len(arr))
	x, y, idx := 0,0,0
	for x < len(left) && y < len(right) {
		if left[x] < right[y] {merged[idx] = left[x]; x++;idx++;continue;}
		if left[x] >= right[y] {merged[idx] = right[y]; y++;idx++;continue;}
	}
	for x < len(left) { merged[idx] = left[x]; idx++; x++ }
	for y < len(right) { merged[idx] = right[y]; idx++; y++ }
	for i:=0;i<len(arr);i++{ arr[i] = merged[i] }
}

