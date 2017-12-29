package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type base int8

const (
	A base = iota
	C
	T
	G
)

type allele struct {
	chrom uint8
	pos   uint64
}

type daf_allele struct {
	allele     allele
	anc_allele base
	der_allele base
	freq       float64
}

func parseDAF(rdr *bufio.Reader) []daf_allele {
	var alleles []daf_allele

	for ln_num := 0; true; ln_num++ {
		var allele daf_allele

		ln, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("reading line from DAF file failed:", err)
			os.Exit(1)
		}
		ln = strings.TrimSpace(ln)
		if len(ln) == 0 || ln[0] == '#' {
			continue
		}

		flds := strings.Fields(ln)
		if len(flds) < 5 {
			fmt.Println("line", string(ln_num), "contains too few fields:")
			fmt.Println(ln)
			os.Exit(1)
		}

		chrom, err := strconv.ParseUint(flds[0], 10, 8)
		if err != nil {
			fmt.Println("Invalid chromosome:", flds[0])
			os.Exit(1)
		}
		allele.allele.chrom = uint8(chrom)

		pos, err := strconv.ParseUint(flds[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid position:", flds[1])
			os.Exit(1)
		}
		allele.allele.pos = pos

		switch strings.ToLower(flds[2]) {
		case "a":
			allele.anc_allele = A
		case "c":
			allele.anc_allele = C
		case "t":
			allele.anc_allele = T
		case "g":
			allele.anc_allele = G
		default:
			fmt.Println("Invalid base:", flds[2])
			os.Exit(1)
		}

		switch strings.ToLower(flds[3]) {
		case "a":
			allele.der_allele = A
		case "c":
			allele.der_allele = C
		case "t":
			allele.der_allele = T
		case "g":
			allele.der_allele = G
		default:
			fmt.Println("Invalid base:", flds[3])
			os.Exit(1)
		}

		freq, err := strconv.ParseFloat(flds[4], 64)
		if err != nil || freq > 1.0 || freq < 0.0 {
			fmt.Println("Invalid allele frequency:", flds[4])
			os.Exit(1)
		}
		allele.freq = freq

		alleles = append(alleles, allele)
	}

	return alleles
}
