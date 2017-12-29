package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	test_filenames := []string{"test_input/denisova.daf", "test_input/Mgenomes3.daf", "test_input/altai.daf"}
	for _, fn := range test_filenames {
		f, err := os.Open(fn)
		if err != nil {
			fmt.Println("Opening DAF file", fn, "failed:", err)
			os.Exit(1)
		}

		alleles := parseDAF(bufio.NewReader(f))
		fmt.Println(fn, len(alleles))

		err = f.Close()
		if err != nil {
			fmt.Println("Closing DAF file", fn, "failed:", err)
			os.Exit(1)
		}
	}
}
