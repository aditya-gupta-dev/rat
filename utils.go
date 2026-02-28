package main 

func LinearSearch[T comparable](arr []T, item T) int {
	for index, element := range arr { 
		if element == item { 
			return index 
		}
	} 
	return -1 
}
