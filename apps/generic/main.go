package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type integer int

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"SDA-ASD458A-ASDASD", "Zapatos", 50}

	fmt.Println(product1, product2)
}
