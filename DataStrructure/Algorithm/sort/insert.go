package sort

func InsertSort(arrays []int) {
	for i := range arrays {
		preIndex := i - 1
		for preIndex >= 0 && arrays[i] < arrays[preIndex] {
			arrays[preIndex+1] = arrays[preIndex]
			preIndex -= 1
		}
		arrays[preIndex+1] = arrays[i]
	}
}
