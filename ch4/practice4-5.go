package main

import "fmt"

func removeDup(s []string) []string {
	var out []string
	for i, v := range s {
		next := i+1
		if next < len(s)-1 && s[i] == s[next] {
			continue
		}
		out = append(out, v)
	}

	return out
}

func main() {
	s := []string{"A", "B", "CD", "CD", "CD", "C", "D"}
	s = removeDup(s)
	fmt.Println(s)

	t := []string{}
	fmt.Println(removeDup(t))
}
