from string import digits


def is_valid(isbn):
    cleaned_isbn = [10 if i == "X" else int(i) for i in isbn if i in digits or i == "X"]
    if len(cleaned_isbn) != 10:
        return False
    if 10 in cleaned_isbn and cleaned_isbn[-1] != 10:
        return False

    return sum([i * (10 - index) for index, i in enumerate(cleaned_isbn)]) % 11 == 0
