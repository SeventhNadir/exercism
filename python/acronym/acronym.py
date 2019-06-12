import re

def abbreviate(words):
    abbreviation = ""
    regex = re.compile("[^a-zA-Z']")
    words = regex.sub(" ", words)
    words = words.split()
    for word in words:
        abbreviation += word[0].capitalize()
    return abbreviation