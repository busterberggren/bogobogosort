package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	// If you wanna see this ever finishing you shouldn't use more
	// than five elements.
	a := []int{4, 50, 300, 6000}//, 2, 1}
	g := bogobogosort(a)
	fmt.Println(g)

}

func Shuffle(a []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for h := 0; h < 4; h++ {
		for i := range a {
			j := rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}
	}
}

func bogobogosort(list []int) []int {
	if len(list) == 1 {
		return list
	}
	listcopy := make([]int, len(list))
	copy(listcopy, list)
	n := len(list) - 1
	for {
		Shuffle(list)
		Shuffle(listcopy)
		z := bogobogosort(listcopy[:n])
		if listcopy[n] >= z[n-1] {
			if reflect.DeepEqual(list, listcopy) {
				break
			}
		}
	}
	return list
}
