package letterCombinationsOfAPhoneNumber

var (
	dict = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) <= 1 {
		return dict[digits]
	}

	if v, ok := dict[digits]; ok {
		return v
	}

	pre := letterCombinations(string(digits[0]))
	sub := letterCombinations(digits[1:])

	var ret []string
	for _, p := range pre {
		for _, s := range sub {
			ret = append(ret, p+s)
		}
	}

	dict[digits] = ret
	return ret
}
