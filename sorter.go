package tf_vars_sorter

type TFVarSorter struct {
}

func NewTFVarSorter() *TFVarSorter {
	return &TFVarSorter{}
}

func (this *TFVarSorter) Sort(unsortedArray []*Var) []*Var {
	newArr := make([]*Var, len(unsortedArray))
	for i, v := range unsortedArray {
		newArr[i] = v
	}
	this.recursiveSort(newArr, 0, len(newArr)-1)
	return newArr
}

func (this *TFVarSorter) recursiveSort(arr []*Var, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i].Name < pivot.Name {
			if splitIndex != i {
				temp := arr[splitIndex]
				arr[splitIndex] = arr[i]
				arr[i] = temp
			}

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	this.recursiveSort(arr, start, splitIndex-1)
	this.recursiveSort(arr, splitIndex+1, end)
}
