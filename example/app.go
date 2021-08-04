package main

import (
	"math/rand"

	"github.com/bcowtech/gorand"
)

func main() {
	r := gorand.New(rand.NewSource(9527))
	r.Seed(9527)

	println("When seed is 9527. Int63r(0,10) return", r.Int63r(0, 10), ".")
}
