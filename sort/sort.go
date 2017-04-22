package sort

func Merge(a []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m
	la := make([]int, n1)
	ra := make([]int, n2)
	for i := 0; i < n1; i++ {
		la[i] = a[l+i]
	}
	for j := 0; j < n2; j++ {
		ra[j] = a[m+1+j]
	}
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if la[i] <= ra[j] {
			a[k] = la[i]
			i++
		} else {
			a[k] = ra[j]
			j++
		}
		k++
	}
	for i < n1 {
		a[k] = la[i]
		i++
		k++
	}
	for j < n2 {
		a[k] = ra[j]
		j++
		k++
	}
}

func MergeSort(a []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2
		MergeSort(a, l, m)
		MergeSort(a, m+1, r)
		Merge(a, l, m, r)
	}
}

func Partition(a []int, l int, h int) int {
	p := a[h]
	i := l - 1
	for j := l; j < h; j++ {
		if a[j] <= p {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}

func QuickSort(a []int, l int, h int) {
	if l < h {
		i := Partition(a, l, h)
		QuickSort(a, l, i-1)
		QuickSort(a, i+1, h)
	}
}
