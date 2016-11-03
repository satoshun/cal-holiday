package main

var holidays []holiday

type holiday interface {
	isHoliday(y, m, d int) bool
}

type gantan struct{}
type bunka struct{}
type kinrou struct{}

func (g *gantan) isHoliday(y, m, d int) bool {
	return m == 1 && d == 1
}

func (g *bunka) isHoliday(y, m, d int) bool {
	return m == 11 && d == 3
}

func (g *kinrou) isHoliday(y, m, d int) bool {
	return m == 11 && d == 23
}

// IsHoliday returns holiday or not
func IsHoliday(y, m, d int) bool {
	for _, h := range holidays {
		if h.isHoliday(y, m, d) {
			return true
		}
	}
	return false
}

func init() {
	holidays = make([]holiday, 0)
	holidays = append(holidays,
		&gantan{},
		&bunka{},
		&kinrou{},
	)
}
