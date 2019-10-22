from string import ascii_lowercase as lowercase


def rotate(text, key):
    rotated_mapping = lowercase[key:] + lowercase[:key]
    translation = str.maketrans(
        lowercase + lowercase.upper(), rotated_mapping + rotated_mapping.upper()
    )
    return text.translate(translation)
