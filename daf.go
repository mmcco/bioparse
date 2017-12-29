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

type dafAllele struct {
	allele    allele
	ancAllele base
	derAllele base
	freq      float64
}

func parseDAF(rdr *bufio.Reader) []dafAllele {
	var alleles []dafAllele

	uniqMap := make(map[allele]bool)

	for lnNum := 0; true; lnNum++ {
		var allele dafAllele

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
			fmt.Println("line", string(lnNum), "contains too few fields:")
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
			allele.ancAllele = A
		case "c":
			allele.ancAllele = C
		case "t":
			allele.ancAllele = T
		case "g":
			allele.ancAllele = G
		default:
			fmt.Println("Invalid base:", flds[2])
			os.Exit(1)
		}

		switch strings.ToLower(flds[3]) {
		case "a":
			allele.derAllele = A
		case "c":
			allele.derAllele = C
		case "t":
			allele.derAllele = T
		case "g":
			allele.derAllele = G
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

		if uniqMap[allele.allele] {
			fmt.Println("duplicate allele (chrom " + string(allele.allele.chrom) + ", pos " + string(allele.allele.pos) + " encountered")
			os.Exit(1)
		}
		uniqMap[allele.allele] = true

		alleles = append(alleles, allele)
	}

	return alleles
}
