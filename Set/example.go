package main

import "fmt"

// type struct{} will not occupy any memory
type Set map[string]struct{}

func getKeys(s Set) []string{
	r := []string{}
	for k, _ :=  range s {
		r = append(r, k)
	}

	return r
}

func inSet(s Set, key string) bool {

	if _, ok := s[key]; ok {
		return true
	}
	return false
}
func main() {
	s := make(Set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}


	fmt.Println(getKeys(s))
	//println(getKeys(s))
	println(inSet(s,"item1"))
	println(inSet(s,"item3"))


}
