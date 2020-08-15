package Utils

func RemoveChannel(slice []chan float64, idx int) []chan float64 {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
