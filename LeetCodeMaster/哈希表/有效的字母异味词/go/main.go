package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for _, v := range s {
		count, ok := mapS[v]
		if ok {
			count++
			mapS[v] = count
		} else {
			mapS[v] = 1
		}
	}
	fmt.Printf("mapS %v\n", mapS)

	for _, v := range t {
		count, ok := mapT[v]
		if ok {
			count++
			mapT[v] = count
		} else {
			mapT[v] = 1
		}
	}
	fmt.Printf("mapT %v\n", mapT)

	for k, v := range mapS {
		count, ok := mapT[k]
		if !ok {
			return false
		}

		if count != v {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	b := isAnagram(s, t)
	fmt.Println(b)
}
