package LessCommonWords

import (
	"fmt"
)

func ExampleKeywords() {
	result := Keywords("I'm from Iceland and I make goat cheese. How about you? Do you work?")

	fmt.Println(result)
	// Output: [iceland goat cheese]
}

func ExampleFilter() {
	result := Filter([]string{"from",
		"Iceland",
		"and",
		"I",
		"make",
		"goat",
		"cheese",
		"How",
		"about",
		"you",
		"Do",
		"you",
		"work"})

	fmt.Println(result)
	// Output: [iceland goat cheese]
}
