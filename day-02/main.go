package main

import "flag"

func main() {
	funcPtr := flag.Int("part", 1, "Flag for determining which part we are running!")
	flag.Parse()

	switch *funcPtr {
	case 1:
		part1()
	case 2:
		part2()
	}
}
