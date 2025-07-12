package arraysslices

// Arrays have a fixed capacity which you define when you declear the variable. We can initialize an array into ways:
// [N]type{value1, vlaue2,...,valueN} e.g numbers := [5]int{1,2,3,4,5} 
// [...]type{value1, value3,...,valueN} e.g numbers := [...]int{1,2,3,4,5}
// To get the value out of an array a particular index, just use `array[index] syntax`
func Sum(numbers []int) int{
	sum := 0
	// for i := 0; i < 5; i++{
	// 	sum += numbers[i]
	// }

	for _, number := range numbers{
		sum += number
	}
	// `range` lets you iterate over an array. On each iteration, 
	// range returns two values - the index and the value.
	//  We are choosing to ignore the index value by using _
	return sum
}

// Arrays and their type 
// An interesting property of arrays is that the size is encoded in its type. 
// If you try to pass an [4]int into a function that expects [5]int, it won't compile.
// They are different types so it's just the same as trying to pass a string into a
// function that wants an int.

// Go has slices which do not encode the size of trhe collection and instead can have any size

