package utils

type MapCallback[Input any, Output any] func(Input, int) Output
type FilterCallback[Input any, Output any] func(Input, int) Output
type ReduceCallback[Accumulator any, Input any] func(Accumulator, Input) Accumulator
type SortCallback[Input any] func(Input, Input) int

func SliceMap[Input any, Output any](input []Input, callback MapCallback[Input, Output]) []Output {
	var newSlice []Output

	for index, val := range input {
		newSlice = append(newSlice, callback(val, index))
	}

	return newSlice
}

func SliceFilter[Input any](input []Input, callback FilterCallback[Input, bool]) []Input {
	var newSlice []Input

	for index, val := range input {
		if callback(val, index) {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}

func SliceReduce[Input any, Output any](input []Input, reducerFn ReduceCallback[Output, Input]) Output {
	var result Output

	for _, val := range input {
		result = reducerFn(result, val)
	}

	return result
}

// Sort a slice with the outcome of a callback. A number less than 0 means that the first element should be sorted before the second element.
func SliceSort[Input any](input []Input, sortFn SortCallback[Input]) []Input {
	var result []Input

	for _, val := range input {
		result = append(result, val)

		for i := len(result) - 1; i > 0; i-- {
			if sortFn(result[i], result[i-1]) < 0 {
				// Swap elements at the current index and the previous index
				result[i], result[i-1] = result[i-1], result[i]
			}
		}
	}

	return result
}

func SliceJoin(input []string) string {
	var result string

	for _, val := range input {
		result += val
	}

	return result
}
