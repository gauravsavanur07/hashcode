package main

import (
	"fmt"
	"os"
)

func main() {
	fi, _ := os.Open("a_example.in")
	d := Data{}
	d.Fill(fi)
	fmt.Println("% v", d)
	println("----")
	vs := make([]Vahicle, d.F)

	rs := make([]int, len(d.RS))
	for j := range d.RS {
		rs[j] = j
	}
	j := 0
	var r Ride
	var v Vahicle
	for len(rs)-j > 0 { //j, r = range d.RS { //
		//fmt.Printf("%#v %#v\n", d.RS, rs)
		fmt.Println(len(rs), j, len(rs)-j == 0)
		r = d.RS[rs[j]]
		vf := false
		for i := 0; i < d.F; i++ {
			v = vs[i]
			tol := Lenght(v.X, v.Y, r.S[0], r.S[1])
			//println("to ", tol)
			if v.S+tol <= r.ES {
				continue
			}
			tol += r.Lenght()
			println("total ", tol)
			if v.S+tol > r.LF {
				continue
			}
			//println("finish ", v.S+tol)
			if v.S+tol >= d.T {
				continue
			}
			vs[i].S += tol
			vs[i].RS = append(vs[i].RS, rs[j])
			vs[i].X, vs[i].Y = r.F[0], r.F[1]
			rs = append(rs[:j], rs[j+1:]...)
			//fmt.Printf("%d %v %d\n", j, rs, len(rs))
			j--
			vf = true
			break
		}
		if !vf {
			j++
		}

	}
	for i, v := range vs {
		fmt.Println(i, v.RS)
	}
	/*
		fo, _ := os.Create("output.txt")
		scanner := bufio.NewScanner(fi)
		writer := bufio.NewWriter(fo)
		defer fi.Close()
		defer fo.Close()
		defer writer.Flush()
		scanner.Scan()
		writer.Write(scanner.Bytes())
	*/
}
