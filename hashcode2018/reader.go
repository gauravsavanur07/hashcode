package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Data struct {
	R, C, F, N, B, T int
	RS               []Ride
}

type Ride struct {
	S, F   [2]int
	ES, LF int
}

type Vahicle struct {
	X, Y, S int
	RS      []int
}

func (r *Ride) Lenght() int {
	return Lenght(r.S[0], r.S[1], r.F[0], r.F[1])
}

func Lenght(x1, y1, x2, y2 int) int {
	if x2 > x1 {
		x1, x2 = x2, x1
	}
	if y2 > y1 {
		y1, y2 = y2, y1
	}
	return (x1 - x2) + (y1 - y2)
}

// Fill will read pizza input information to the Data struct
func (d *Data) Fill(src io.Reader) error {
	scanner := bufio.NewScanner(src)

	if ok := scanner.Scan(); !ok {
		err := scanner.Err()
		return errors.Wrap(err, "reading first line from input source")
	}

	n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d %d", &d.R, &d.C, &d.F, &d.N, &d.B, &d.T)
	if err != nil || n != 6 {
		return errors.Wrapf(err, "first line missed some important valuest, should contain 6 separate number (%d was readed)", n)
	}

	d.RS = make([]Ride, d.R)
	i := 0 // row index
	var x1, y1, x2, y2, es, lf int
	for scanner.Scan() {
		n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d %d", &x1, &y1, &x2, &y2, &es, &lf)
		if err != nil || n != 6 {
			return errors.Wrapf(err, "ride line missed some important valuest, should contain 6 separate number (%d was readed)", n)
		}
		d.RS[i] = Ride{S: [2]int{x1, y1}, F: [2]int{x2, y2}, ES: es, LF: lf}
		i++
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "scanning input rides data")
	}

	return nil
}
