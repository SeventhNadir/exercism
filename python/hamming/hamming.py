def distance(strand_a, strand_b):
    if not len(strand_a) == len(strand_b):
        raise ValueError("strands of unequal length")

    hamming_score = 0
    for i in range(len(strand_a)):
        if strand_a[i] != strand_b[i]:
            hamming_score += 1

    return hamming_score
