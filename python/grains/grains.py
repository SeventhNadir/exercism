def square(number):
    if number > 64 or number < 1:
        raise ValueError("Number must be greater than one and less than or equal to 64")
    return 2 ** (number - 1)


def total(number):
    if number > 64 or number < 1:
        raise ValueError("Number must be greater than one and less than or equal to 64")
    return sum([square(i+1) for i in range(number)])
