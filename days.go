package yttr

import "strconv"

type Days int

func (d Days) String() string {
	return strconv.Itoa(int(d))
}
