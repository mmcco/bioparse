package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testFilenames := []string{"test_input/denisova.daf", "test_input/Mgenomes3.daf", "test_input/altai.daf"}
	var dafAllelePops [][]dafAllele
	var testMaps []map[allele]dafAllele

	for _, fn := range testFilenames {
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

		dafAllelePops = append(dafAllelePops, alleles)
		testMaps = append(testMaps, dafMap(alleles))
	}

	sharedAlleles := filterAlleles(dafAllelePops)

	for _, sharedAllele := range sharedAlleles {
		fmt.Printf("%d : %d\n", sharedAllele.chrom, sharedAllele.pos)

		for _, m := range testMaps {
			if x, ok := m[sharedAllele]; ok {
				fmt.Printf("\t%d %d\n\n", x.ancAllele, x.derAllele)
			}
		}
	}
	fmt.Println("number of alleles of interest:", len(sharedAlleles))
}
