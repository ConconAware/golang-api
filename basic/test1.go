package test

import (
	"sort"

	"golang.org/x/exp/slices"
)

func MainUniqueAnimals() []string {

	ani1 := []string{"Dog", "Dog", "Cat", "Goose"}
	ani2 := []string{"Dog", "Cat", "Rabbit"}

	res := uniqueAnimals(ani1, ani2)

	return res
}

func uniqueAnimals(a, b []string) []string {
	var result []string
	//{"Dog", "Dog", "Cat", "Goose", "Dog", "Cat", "Rabbit"}
	a = append(a, b...)

	for _, v := range a {
		if slices.Index(result, v) < 0 {
			result = append(result, v)
		}
	}
	// slices.Sort[string](result)
	sort.Strings(sort.StringSlice(result))
	// sort.Sort(sort.StringSlice(result))
	return result
}

func uniqueAnimals1(a, b []string) []string {
	var result []string
	//{"Dog", "Dog", "Cat", "Goose", "Dog", "Cat", "Rabbit"}
	m := make(map[string]string)

	for _, v := range a {
		m[v] = v
	}

	for _, v := range b {
		m[v] = v
	}

	for _, v := range m {
		result = append(result, v)
	}
	// slices.Sort[string](result)
	sort.Strings(sort.StringSlice(result))
	// sort.Sort(sort.StringSlice(result))
	return result
}
