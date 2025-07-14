package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	// t.Run("collection of 5 numbers", func(t *testing.T) {

	// 	numbers := []int{1,2,3,4,5}

	// 	got := Sum(numbers)

	// 	want := 15

	// 	if got != want{
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := Sum(numbers) 
		want := 6

		if got != want{
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// We need a new function called SumAll which will take a varying number of slices, 
// returning a new slice containing the totals for each slice passed in. 
// For example 
// `SumAll([]int{1,2}, []int{0,9})` would return `[]int{3,9}`
//  or
//  `SumAll([]int{1,1,1})` would return `[]int{3}`


func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestSumAllTails(t *testing.T){

	checkSums := func (t testing.TB, got, want []int)  {
		t.Helper()

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got , want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
	
}

// We could've created a new function checkSums like we normally do, 
// but in this case, we're showing a new technique, assigning a function to a variable.
//  It might look strange but, it's no different to assigning a variable to a string, or an int,
//  functions in effect are values too.

// It's not shown here, but this technique can be useful when you want to bind a function to other 
// local variables in "scope" (e.g between some {}). It also allows you to reduce the surface area 
// of your API.

// By defining this function inside the test, it cannot be used by other functions in this package.
//  Hiding variables and functions that don't need to be exported is an important design consideration.

// A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly
//  adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.