package kata

func Parse(data string) []int{
	//TODO: write your code here

	countVar := 0

	countSli := []int{}

	for _, val := range data {
		switch string(val) {
		case "i":
			countVar += 1
		case "d":
			countVar -= 1
		case "s":
			countVar = countVar * countVar
		case "o":
			countSli = append(countSli, countVar)
		}
	}
	return countSli
}