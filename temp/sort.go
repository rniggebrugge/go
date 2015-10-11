package main 

import (
	"fmt"
	"strings"
	"sort"
)

func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

type FoldedStrings []string 

func (slice FoldedStrings) Len() int { return len(slice) }

func (slice FoldedStrings) Less(i,j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i,j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func main() {
	files := []string{"Remco","clara","Willem","pieter"}
	target:="Clara"
	fmt.Printf("Looking for %s in %q\n", target, files)
	sort.Strings(files)

	compareSimple := func(i int) bool { return files[i]>=target}
	
	i := sort.Search(len(files), compareSimple)
	if i <len(files) && files[i] == target {
		fmt.Printf("Found \"%s\" at files[%d]\n", files[i], i)
	} else {
		fmt.Printf("Did not find \"%s\".\n", target)
	}

	SortFoldedStrings(files)	
	fmt.Printf("Looking for %s in %q\n", target, files)
	betterCompare := func(i int) bool {
		return strings.ToLower(files[i]) >= strings.ToLower(target)
	}
	i = sort.Search(len(files), betterCompare)
	if i <len(files) && strings.EqualFold(files[i], target) {
		fmt.Printf("Found \"%s\" at files[%d]\n", files[i], i)
	} else {
		fmt.Printf("Did not find \"%s\".\n", target)
	}
}