package kata

func FindOdd(seq []int) int {
	intCountMap := make(map[int]int)

	for _,val := range seq {

		_,exists := intCountMap[val]


		if exists {
			intCountMap[val] += 1
		} else {
			intCountMap[val] = 1
		}
	}

	oddOccurenceKey := 0
	for key, val := range intCountMap {
		if val % 2 != 0 {
			oddOccurenceKey = key
		}
	}
	return oddOccurenceKey
}






/*
Other interesting solution I found on there. ^= operator

package kata

func FindOdd(seq []int) int {
    res := 0
    for _, x := range seq {
        res ^= x
    }
    return res
}
 */