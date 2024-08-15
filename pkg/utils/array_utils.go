package utils

import "reflect"

func RemoveAll[T any](slice []T, element T) []T {
	var result []T
	for _, item := range slice {
		if !reflect.DeepEqual(item, element) {
			result = append(result, item)
		}
	}
	return result
}

func RemoveAt[T any](slice []T, index int) []T {
	var result []T
	for i := 0; i < len(slice); i++ {
		if i != index {
			result = append(result, slice[i])
		}
	}
	return result
}
