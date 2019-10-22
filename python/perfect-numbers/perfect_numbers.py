from math import sqrt, ceil


def classify(number):
    aliquot_sum = sum(find_factors(number))
    if aliquot_sum == number:
        return "perfect"
    elif aliquot_sum < number:
        return "deficient"
    elif aliquot_sum > number:
        return "abundant"


def find_factors(number):
    if number <= 0:
        raise ValueError("Number must be greater than or equal to zero")
    if number == 1:
        return {}
    factors = {1}
    upperbound = ceil(sqrt(number))
    for factor in range(2, upperbound):
        if number % factor == 0 and number != factor:
            factors.add(factor)
            factors.add(number // factor)
    # 1 is a factor for all numbers
    return factors
