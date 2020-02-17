package mathematics

func Min(list []int) int {
	if len(list) <= 1 {
		return list[0]
	}

	return bubbleSort(list)[0]
}

func Max(list []int) int {
	if len(list) <= 1 {
		return list[0]
	}

	return bubbleSort(list)[len(list)-1]
}

func bubbleSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	walking := true

	for walking {
		walking = false
		for index, value := range list {
			if index == len(list)-1 {
				break
			}

			if value > list[index+1] {
				list[index] = list[index+1]
				list[index+1] = value
				walking = true
			}
		}
	}

	return list
}
