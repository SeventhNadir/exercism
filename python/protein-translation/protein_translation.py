prot_translate = {
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

def proteins(strand):
    proteins = []
    protein_candidates = [strand[i:i+3] for i in range(0, len(strand), 3)]
    for protein in protein_candidates:
        if prot_translate[protein] == "STOP":
            break
        else:
            proteins.append(prot_translate[protein])
    return proteins