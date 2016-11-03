package main

import "time"

var holidays []holiday

type holiday interface {
	isHoliday(d time.Time) bool
}

type gantan struct{}
type kenkoku struct{}
type shouwa struct{}
type kenpou struct{}
type midori struct{}
type kodomo struct{}
type yama struct{}
type bunka struct{}
type kinrou struct{}
type tennou struct{}

// TODO: umi, shunbun, shuubun, taiiku, keirou, seijin

func (g *gantan) isHoliday(d time.Time) bool {
	return d.Month() == 1 && d.Day() == 1
}
func (g *kenkoku) isHoliday(d time.Time) bool {
	return d.Month() == 2 && d.Day() == 11
}
func (g *shouwa) isHoliday(d time.Time) bool {
	return d.Month() == 4 && d.Day() == 29
}
func (g *kenpou) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 3
}
func (g *midori) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 4
}
func (g *kodomo) isHoliday(d time.Time) bool {
	return d.Month() == 5 && d.Day() == 5
}
func (g *yama) isHoliday(d time.Time) bool {
	return d.Month() == 8 && d.Day() == 11
}
func (g *bunka) isHoliday(d time.Time) bool {
	return d.Month() == 11 && d.Day() == 3
}
func (g *kinrou) isHoliday(d time.Time) bool {
	return d.Month() == 11 && d.Day() == 23
}
func (g *tennou) isHoliday(d time.Time) bool {
	return (d.Month() == 12 && d.Day() == 23) || (d.Month() == 4 && d.Day() == 29)
}

// IsHoliday returns holiday or not
func IsHoliday(y, m, d int) bool {
	dd := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	for _, h := range holidays {
		if h.isHoliday(dd) {
			return true
		}
	}
	return false
}

func init() {
	holidays = make([]holiday, 0)
	holidays = append(holidays,
		&gantan{}, &kenkoku{}, &bunka{},
		&shouwa{}, &kenpou{}, &midori{}, &kodomo{},
		&yama{}, &kinrou{}, &tennou{},
	)
}
vv