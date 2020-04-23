package main

import (
	"fmt"
	"strings"

	"github.com/bombsimon/perl-go/grep"
	"github.com/bombsimon/perl-go/qw"
	"github.com/bombsimon/perl-go/warn"
)

func main() {
	fmt.Println("=== Debugging")
	debugMe(true, true, true)
	fmt.Println("")

	fmt.Println("=== Lists and grep")
	lists()
	fmt.Println("")

	fmt.Println("=== Creating lists")
	quoteWords()
}

func debugMe(maybe1, maybe2, maybe3 bool) {
	if maybe1 {
		warn.Here()
	}

	if maybe2 {
		warn.Here("In maybe2!")
	}

	if maybe3 {
		warn.Here(fmt.Sprintf("Uhh, I'm having an error: %s", "UNKNOWN"))
	}
}

func lists() {
	newList := grep.List("string-literal", []string{"foo", "string-literal", "bar"})
	fmt.Printf("found %d matches: %s\n", len(newList), newList)

	newList = grep.List("/re(gexp)?/", []string{"re", "redo", "do"})
	fmt.Printf("found %d matches with regexp: %s\n", len(newList), newList)

	value := grep.First("item", []string{"not item", "item", "item"})
	fmt.Printf("found item at least once: %s\n", value)

	isFound := grep.Find([]string{"a", "b", "c"}, func(s string) bool {
		return s == "b"
	})

	if isFound {
		fmt.Println("yep, 'b' found in the list")
	}

	type User struct {
		Name string
		Age  int
	}

	users := []interface{}{
		User{Name: "simon", Age: 32},
		User{Name: "jenny", Age: 14},
	}

	isSimonInList := grep.FindI(users, func(item interface{}) bool {
		return item.(User).Name == "simon"
	})

	fmt.Println("simon was in list:", isSimonInList)
}

func quoteWords() {
	for _, s := range qw.S("quick way", "to write", "a range") {
		fmt.Println(s)
	}

	cats := grep.List("/cat/", qw.S("cat: rufus", "cat: leopold", "dog: fido"))
	fmt.Printf("found %d cats: %s\n", len(cats), strings.Join(cats, ", "))

	fmt.Println("just som digits", qw.I(10, 20, 50, 100))
}
