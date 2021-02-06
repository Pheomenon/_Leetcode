package main

import "fmt"

func main() {
	arr := []int{
		2, 3, 54, 2, 5, 6, 4, 3, 2, 16, 3, 523, 12, 754, 87, 432, 93, 9,
	}
	sort(arr)
	for _, v := range arr {
		fmt.Printf("%d, ", v)
	}
}

func less(x, y int) bool {
	return x < y
}
func merge(a []int, lo, mid, hi int) {
	i := lo
	j := mid + 1
	aux := make([]int, hi+1)
	for k := 0; k <= hi; k++ {
		aux[k] = a[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
func sort(a []int) {
	sortReal(a, 0, len(a)-1)
}
func sortReal(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortReal(a, lo, mid)
	sortReal(a, mid+1, hi)
	merge(a, lo, mid, hi)
}
