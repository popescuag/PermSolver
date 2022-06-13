package utils

func AddToSortedArray(arr *[]string, elem string) {
	if len(*arr) == 0 {
		*arr = append(*arr, elem)
		return
	}
	for i := 0; i < len(*arr); i++ {
		if elem == (*arr)[i] {
			// Element already in array
			return
		}
		if elem > (*arr)[i] {
			if i == (len(*arr) - 1) {
				// make room for the new element
				*arr = append(*arr, "")
				(*arr)[i+1] = elem
				return
			} else if elem < (*arr)[i+1] {
				// make room for the new element
				*arr = append(*arr, "")
				copy((*arr)[(i+1):], (*arr)[i:])
				(*arr)[i+1] = elem
				return
			}
		} else {
			if len(*arr) == 1 {
				tmp := (*arr)[0]
				(*arr)[0] = elem
				*arr = append(*arr, tmp)
				return
			} else if i == 0 {
				//add at the beginning
				*arr = append(*arr, "")
				copy((*arr)[(i+1):], (*arr)[i:])
				(*arr)[0] = elem
				return
			}
		}
	}
	// Element is last
	*arr = append(*arr, elem)
}
