package checkforsum

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// поиск по упорядоченному массиву
func findInSortedArray(a []int, sum int) bool {
	var x int
	for last, i := len(a)-1, 0; i < last; {
		x = a[i] + a[last]
		switch {
		case x == sum:
			return true
		case x > sum:
			last--
		case x < sum:
			i++
		}
	}
	return false
}

// поиск по неупорядоченному массиву
func findInUnsortedArray1(a []int, sum int) bool {
	x := make(map[int]struct{})
	for _, i := range a {
		if _, ok := x[i]; ok {
			return true
		}
		x[sum-i] = struct{}{}
	}
	return false
}

// предварительная сортировка неупорядоченного массива штатными средствами
// и дальнейший поиск по упорядоченному массиву
// работает медленнее предыдущего способа, но выигрывает по памяти
func findInUnsortedArray2(a []int, sum int) bool {
	sort.Ints(a)
	return findInSortedArray(a, sum)
}
