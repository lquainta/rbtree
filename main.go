package main

import (
	"fmt"
	"math"
	"math/rand"

	"redblack/src"
)

func check(tree *src.RedblackTree[int], key int, expected map[int]int, total int) {
	if err := src.Validate(tree); err != nil {
		panic(fmt.Sprintf("after changing %d: %v", key, err))
	}
	if src.Search(key, tree) != (expected[key] > 0) {
		panic(fmt.Sprintf("wrong search result for %d", key))
	}
	if src.Size(tree) != int(total) {
		panic(fmt.Sprintf("size is %v, expected %d", src.Size(tree), total))
	}
}

func main() {
	tree := src.NewRedblackTree[int]()
	expected := map[int]int{}
	total := 0
	rng := rand.New(rand.NewSource(42))

	add := func(value int) {
		src.Insert(value, tree)
		expected[value]++
		total++
		check(tree, value, expected, total)
	}

	remove := func(value int) {
		src.Delete(value, tree)
		expected[value]--
		total--
		check(tree, value, expected, total)
	}

	values := []int{
		// Ascending run: repeated left rotations
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,

		// Descending run: repeated right rotations
		-1, -2, -3, -4, -5, -6, -7, -8, -9, -10,

		// Zigzag: double rotations
		100, 50, 75, 60, 70, 65, 68, 66, 67,

		// Duplicates
		5, 5, 5, 0, 0, 67, 67, -10, -10,

		// Extreme values
		math.MaxInt, math.MinInt, math.MaxInt - 1, math.MinInt + 1,
	}

	for _, value := range values {
		add(value)
	}
	src.PrintTree(tree)

	// Delete every value in random order
	rng.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
	for _, value := range values {
		remove(value)
	}

	// Random inserts and deletes
	for i := 0; i < 20000; i++ {
		key := rng.Intn(1001) - 500

		if rng.Intn(2) == 0 {
			add(key)
		} else if expected[key] > 0 {
			remove(key)
		}
	}

	fmt.Println("All tests passed")
}
