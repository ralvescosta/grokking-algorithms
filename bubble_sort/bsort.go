package bsort

func BSort(v *[]int, n int) {
	if n < 1 {
		return
	}

	for i := 0; i < n; i++ {
		if (*v)[i] > (*v)[i+1] {
			swap(&(*v)[i], &(*v)[i+1])
		}
	}

	BSort(v, n-1)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
