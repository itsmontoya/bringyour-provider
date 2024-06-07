package plugin

import "strconv"

func parseInt(in string) (out int, err error) {
	if len(in) == 0 {
		return
	}

	return strconv.Atoi(in)
}
