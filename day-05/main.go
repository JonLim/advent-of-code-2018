package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	funcPtr := flag.Int("part", 1, "Flag for determining which part we are running!")
	flag.Parse()

	switch *funcPtr {
	case 1:
		part1()
	case 2:
		part2()
	}

	fmt.Printf("Total time: %s\n", time.Since(start))
}
