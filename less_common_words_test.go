package lessCommonWords

import (
	"fmt"
)

var BlackList1 = map[string]bool{
	"iceland": true,
}

var BlackList2 = map[string]bool{
	"study":  true,
	"cheese": true,
}

func ExampleKeywords() {
	result := Keywords("I'm from Iceland and I make goat cheese. How about you? Do you work?")

	fmt.Println(result)
	// Output: [iceland goat cheese]
}

func ExampleKeywordsWithCustomBlacklists() {
	result := Keywords("I'm from Iceland and I make goat cheese. How about you? Do you work or study?", BlackList1, BlackList2)

	fmt.Println(result)
	// Output: [goat]
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

func ExampleFilterWithCustomBlackLists() {
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
		"work",
		"or",
		"study",
	}, BlackList1, BlackList2)

	fmt.Println(result)
	// Output: [goat]
}
