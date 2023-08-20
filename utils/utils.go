package utils

import "os"

type Callback[Input any, Output any] func(Input) Output
type ReduceCallback[Accumulator any, Input any] func(Accumulator, Input) Accumulator

func SliceMap[Input any, Output any](input []Input, callback Callback[Input, Output]) []Output {
	var newSlice []Output

	for _, val := range input {
		newSlice = append(newSlice, callback(val))
	}

	return newSlice
}

func SliceFilter[Input any](input []Input, callback Callback[Input, bool]) []Input {
	var newSlice []Input

	for _, val := range input {
		if callback(val) {
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

func FindFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	files := SliceFilter(entries, func(file os.DirEntry) bool {
		return !file.IsDir()
	})

	fileNames := SliceMap(files, func(file os.DirEntry) string {
		return path + file.Name()
	})

	return fileNames, nil
}
