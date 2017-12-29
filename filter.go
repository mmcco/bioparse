package main

type alleleInfo struct {
	ctr  int
	ancs []base
}

func filterAlleles(alleleSlices [][]dafAllele) []allele {

	alleleInfo := make(map[allele]alleleInfo)

	for _, alleleSlice := range alleleSlices {
		for _, allele := range alleleSlice {
			x := alleleInfo[allele.allele]

			x.ctr++
			if len(x.ancs) == 0 || allele.ancAllele != x.ancs[0] {
				x.ancs = append(x.ancs, allele.ancAllele)
			}

			alleleInfo[allele.allele] = x
		}
	}

	var interesting_alleles []allele
	for allele, alleleInfo := range alleleInfo {
		if alleleInfo.ctr > 1 && len(alleleInfo.ancs) == 1 {
			interesting_alleles = append(interesting_alleles, allele)
		}
	}

	return interesting_alleles
}
