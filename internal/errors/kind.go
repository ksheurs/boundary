package errors

// Kind specifies the kind of error (unknown, parameter, integrity, etc).
type Kind uint32

const (
	Other Kind = iota
	Parameter
	Integrity
	Search
)

func (e Kind) String() string {
	return map[Kind]string{
		Other:     "unknown",
		Parameter: "parameter violation",
		Integrity: "integrity violation",
		Search:    "search issue",
	}[e]
}
