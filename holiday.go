package holiday

import "time"

var holidays []Holiday

// Holiday represents a holiday
type Holiday interface {
	isHoliday(d time.Time) bool

	Name() string
}

type gantan struct{}
type seijin struct{}
type kenkoku struct{}
type shouwa struct{}
type kenpou struct{}
type midori struct{}
type kodomo struct{}
type umi struct{}
type yama struct{}
type keirou struct{}
type bunka struct{}
type taiiku struct{}
type kinrou struct{}
type tennou struct{}

// TODO: shunbun, shuubune

func (g *gantan) isHoliday(d time.Time) bool {
	return d.Month() == 1 && d.Day() == 1
}
func (g *gantan) Name() string {
	return "gantan"
}
func (g *seijin) isHoliday(d time.Time) bool {
	return d.Month() == 1 &&
		(d.Weekday() == time.Monday && d.Day() >= 8 && d.Day() <= 14)
}
func (g *seijin) Name() string {
	return "seijin"
}
func (g *kenkoku) isHoliday(d time.Time) bool {
	return d.Month() == 2 && d.Day() == 11
}
func (g *kenkoku) Name() string {
	return "kenkoku"
}
func (g *shouwa) isHoliday(d time.Time) bool {
	return d.Month() == 4 && d.Day() == 29
}
func (g *shouwa) Name() string {
	return "shouwa"
}
func (g *kenpou) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 3
}
func (g *kenpou) Name() string {
	return "kenpou"
}
func (g *midori) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 4
}
func (g *midori) Name() string {
	return "midori"
}
func (g *kodomo) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 5
}
func (g *kodomo) Name() string {
	return "kodomo"
}
func (g *umi) isHoliday(d time.Time) bool {
	return d.Month() == 7 &&
		(d.Weekday() == time.Monday && d.Day() >= 15 && d.Day() <= 21)
}
func (g *umi) Name() string {
	return "umi"
}
func (g *yama) isHoliday(d time.Time) bool {
	return d.Month() == 8 && d.Day() == 11
}
func (g *yama) Name() string {
	return "yama"
}
func (g *keirou) isHoliday(d time.Time) bool {
	return d.Month() == 9 &&
		(d.Weekday() == time.Monday && d.Day() >= 15 && d.Day() <= 21)
}
func (g *keirou) Name() string {
	return "keirou"
}
func (g *taiiku) isHoliday(d time.Time) bool {
	return d.Month() == 10 &&
		(d.Weekday() == time.Monday && d.Day() >= 8 && d.Day() <= 14)
}
func (g *taiiku) Name() string {
	return "taiiku"
}
func (g *bunka) isHoliday(d time.Time) bool {
	return d.Month() == 11 && d.Day() == 3
}
func (g *bunka) Name() string {
	return "bunka"
}
func (g *kinrou) isHoliday(d time.Time) bool {
	return d.Month() == 11 && d.Day() == 23
}
func (g *kinrou) Name() string {
	return "kinrou"
}
func (g *tennou) isHoliday(d time.Time) bool {
	return (d.Month() == 12 && d.Day() == 23) || (d.Month() == 4 && d.Day() == 29)
}
func (g *tennou) Name() string {
	return "tennou"
}

// Get returns holiday or nil
func Get(y, m, d int) Holiday {
	dd := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	for _, h := range holidays {
		if h.isHoliday(dd) {
			return h
		}
	}
	return nil
}

func init() {
	holidays = make([]Holiday, 0)
	holidays = append(holidays,
		&gantan{}, &seijin{},
		&kenkoku{},
		&bunka{},
		&shouwa{},
		&kenpou{}, &midori{}, &kodomo{},
		&umi{},
		&yama{},
		&keirou{},
		&taiiku{},
		&kinrou{}, &tennou{},
	)
}
