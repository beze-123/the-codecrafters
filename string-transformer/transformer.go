package main

import (
	"fmt"
	"strings"
)

func upper(me string) string {
	return strings.ToUpper(me)
}

func cap(is string) string {
	return strings.ToUpper(is[1:]) + strings.ToLower(is[:1])
}

func lower(If string) string {
	return strings.ToLower(If)
}

func reverse(two string) string {
	r := []rune(two)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

}

func title(one string) string {
	the := map[string]bool{
		"a":   true,
		"an":  true,
		"the": true,
		"and": true,
		"but": true,
		"or":  true,
		"for": true,
		"nor": true,
		"on":  true,
		"at":  true,
		"to":  true,
		"by":  true,
		"in":  true,
		"of":  true,
		"up":  true,
		"as":  true,
		"is":  true,
		"it":  true,
	}
	text := strings.Fields(one)
	for i, w := range text {
		if i == 0 || !the[w] {
			text[i] = strings.ToUpper(string(w[0])) + w[1:]
		}

	}
	return strings.Join(text, " ")

}
func snake(lolipop string) string {
	w := strings.ToLower(lolipop)
	return strings.ReplaceAll(w, " ", "_")
}

func main() {
	Reader := ""
	operation := ""
	// index := 0
	for {
		fmt.Println("Enter word")
		fmt.Scan(&Reader)
		fmt.Print("select operation:\n", "upper:\n", "lower:\n", "title:\n", "reverse:\n", "snake:\n", "cap:\n", "Quit:\n")
		fmt.Scan(&operation)
		if operation == "upper" {
			fmt.Println(upper(Reader))
		} else if operation == "lower" {
			fmt.Println(lower(Reader))
		} else if operation == "title" {
			fmt.Println(title(Reader))
		} else if operation == "reverse" {
			fmt.Println(reverse(Reader))
		} else if operation == "snake" {
			fmt.Println(snake(Reader))
		} else if operation == "cap" {
			fmt.Println(cap(Reader))
		} else if operation == "Quit" {
			fmt.Println("goodbye")
			break
		} else {
			fmt.Println("No text provided")

		}
	}
}
