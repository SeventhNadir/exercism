package protein

import "errors"

//ErrStop indicates that we've reached a STOP protein in the chain
var ErrStop = errors.New("stop")

//ErrInvalidBase indicates that we've hit an invalid protein we can't translate
var ErrInvalidBase = errors.New("invalid base")

var proteinMapping = map[string]string{
	"UCU": "Serine",
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UGG": "Tryptophan",
	"UGA": "STOP",
	"UAU": "Tyrosine",
	"UCG": "Serine",
	"UGC": "Cysteine",
	"UCC": "Serine",
	"UUC": "Phenylalanine",
	"UCA": "Serine",
	"UAA": "STOP",
	"UAC": "Tyrosine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UAG": "STOP",
	"UGU": "Cysteine",
}

//FromCodon takes a codon as a string and translates it into it's corresponding protein
//It will raise an error if the protein is invalid or if it is a STOP protein.
func FromCodon(codon string) (string, error) {
	if protein, ok := proteinMapping[codon]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

//FromRNA processes a given RNA string and breaks it down into it's composite proteins
func FromRNA(rna string) ([]string, error) {
	codons := []string{}
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return codons, ErrInvalidBase
		}
		codons = append(codons, protein)
	}
	return codons, nil
}
