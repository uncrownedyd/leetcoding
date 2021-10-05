package findAllAnagramsInAString

func findAnagrams(s string, p string) []int {
	var ret []int
	if len(s) < len(p) {
		return ret
	}

	markS := [26]int{}
	markP := [26]int{}

	for _, l := range p {
		markP[l - 'a']++
	}

	winLen := len(p)
	for i := 0; i < winLen; i++ {
		markS[s[i] - 'a']++
	}

	if markS == markP {
		ret = append(ret, 0)
	}

	for i := winLen; i < len(s); i++ {
		markS[s[i - winLen] - 'a']--
		markS[s[i] - 'a']++
		if markS == markP {
			ret = append(ret, i - winLen + 1)
		}
	}

	return ret
}
