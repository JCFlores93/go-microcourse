package utils

// []int {9,8,7,6,5,4,3,2,1}
// []int {1,2,3,4,5,6,7,8,9}
func BubbleSort(elements []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
	return elements
}
