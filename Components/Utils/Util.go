package Utils

//Removes a channel at a particular index from a channel slice
func RemoveChannel(slice []chan float64, idx int) []chan float64 {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
