package cracked

type void struct{}

// Write a function that tests a string to determine if all characters therein
// are unique.
func rAllUniq(s string) bool {
	ra := []rune(s)
	set := make(map[rune]void)
	var member void

	cnt := 0
	for idx, v := range ra {
		cnt++
		set[v] = member
		// Given the max len of a unique string this is likely more costly in most cases...
		if len(set) < idx+1 {
			break
		}
	}
	//fmt.Printf("cnt: %d - len(ra): %d\n", cnt, len(ra)) // for funzies...

	return len(ra) == len(set)
}
