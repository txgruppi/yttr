package yttr

import "strconv"

type Size int64

func (s Size) String() string {
	return strconv.FormatInt(int64(s), 10)
}
