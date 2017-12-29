package main

type alleleInfo struct {
	ctr  int
	ancs []base
	ders []base
}

func uniq(x []base) bool {
	for i := 1; i < len(x); i++ {
		if x[i-1] != x[i] {
			return false
		}
	}
	return true
}

func filterAlleles(alleleSlices [][]dafAllele) []allele {

	alleleInfo := make(map[allele]alleleInfo)

	for _, alleleSlice := range alleleSlices {
		for _, allele := range alleleSlice {
			x := alleleInfo[allele.allele]

			x.ctr++
			x.ancs = append(alleleInfo[allele.allele].ancs, allele.ancAllele)
			x.ders = append(alleleInfo[allele.allele].ancs, allele.derAllele)

			alleleInfo[allele.allele] = x
		}
	}

	var interesting_alleles []allele
	for allele, alleleInfo := range alleleInfo {
		if alleleInfo.ctr > 1 && uniq(alleleInfo.ancs) && uniq(alleleInfo.ders) {
			interesting_alleles = append(interesting_alleles, allele)
		}
	}

	return interesting_alleles
}
