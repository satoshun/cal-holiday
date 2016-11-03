package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type calSt struct {
	year  int
	month int
	days  int

	origin []byte
}

func (cal *calSt) format() string {
	pc := color.New(color.BgBlue).SprintFunc()
	d := string(cal.origin)
	for i := 1; i <= cal.days; i++ {
		if IsHoliday(cal.year, cal.month, i) {
			s := fmt.Sprintf("%2d", i)
			d = strings.Replace(d, s, pc(s), 1)
		}
	}
	// TODO: this code should move read function
	return strings.Trim(d, "\n")
}

func main() {
	// read
	c := read()

	// parse
	cal := parse(c)
	fmt.Println(cal.format())
}

func read() []byte {
	reader := bufio.NewReader(os.Stdin)
	var c []byte
	for {
		data := make([]byte, 4096)
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		c = append(c, data...)
	}
	return bytes.Trim(c, "\x00")
}

func parse(data []byte) *calSt {
	data = normalize(data)
	s := strings.Split(string(data), "\n")
	s[0] = strings.Trim(s[0], " ")

	t := strings.Split(s[0], " ")
	m := monthNormalize(t[0])
	y, _ := strconv.Atoi(t[1])

	var days int
	i := 2
	for {
		if len(s) <= i+2 {
			break
		}

		t := s[i]
		j := 0
		for {
			if len(t) <= j+1 {
				break
			}
			a := t[j : j+2]
			day, _ := strconv.Atoi(a)
			if days < day {
				days = day
			}
			j += 3
		}
		i++
	}

	return &calSt{year: y, month: m, days: days, origin: data}
}

func monthNormalize(m string) int {
	switch m {
	case "April":
		return 4
	case "November":
		return 11
	}
	return 0
}

// TODO: it's a hacky code
func normalize(o []byte) []byte {
	n := make([]byte, 0, len(o))
	for i := 0; i < len(o); i++ {
		// TODO: check index out of range
		if o[i] == byte(95) && o[i+1] == byte(8) {
			i++ // skip
			continue
		}
		n = append(n, o[i])
	}
	return n
}