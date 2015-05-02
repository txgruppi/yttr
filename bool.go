package yttr

type Bool bool

func (b Bool) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}
