import re
from collections import Counter


def count_words(sentence):
    sentence = re.sub("_", " ", sentence.lower())
    words = re.findall(r"([\w]+[']?[\w]+|\d|[a-z])", sentence)
    return Counter(words)
