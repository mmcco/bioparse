package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type base int8

const (
	A base = iota
	B
	T
	G
)

type daf_allele struct {
	chrom      uint8 // the various formats specify this as a string
	pos        uint64
	anc_allele base
	der_allele base
	freq       float64
}

func parseDAF(rdr *bufio.Reader) {
	for ln_num := 0; true; ln_num++ {
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
	}
}

// https://wiki.nci.nih.gov/display/TCGA/TCGA+Variant+Call+Format+%28VCF%29+1.1.1+Specification

/*
func parseVCF(rdr *bufio.Reader) {
	in_hdr := true
	hdr_vals := make(map[string]string)

	for ln_num := 0; true; ln_num++ {
		ln, err := rdr.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("reading line from VCF file failed:", err)
			os.Exit(1)
		}
		ln = strings.TrimSpace(ln)
		if len(ln) == 0 {
			continue
		}

		if in_hdr {
			if !strings.HasPrefix(ln, "##") {
				in_hdr = false
				continue
			}
			kv = strings.Split(ln[2:], "=")
			if len(kv) < 2 {
				fmt.Println("HEADER lines in VCF must use ##key=val format")
				os.Exit(1)
			}
			hdr_vals[kv[0]] = kv[1]
			continue
		}

		// logic for BODY
	}
}
*/
