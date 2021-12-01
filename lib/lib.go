package lib

import (
	"bufio"
	"os"
)

// ForLine will read the contents of fname and call fn for each line.
func ForLine(fname string, fn func(line string)) error {
	// Open input file.
	f, err := os.Open("./prob1.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Setup scanner for line reading.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Process each line.
	for scanner.Scan() {
		fn(scanner.Text())
	}

	return scanner.Err()
}
