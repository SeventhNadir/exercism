day_and_gift = [
    None,
    ["first", "and a Partridge in a Pear Tree."],
    ["second", "two Turtle Doves"],
    ["third", "three French Hens"],
    ["fourth", "four Calling Birds"],
    ["fifth", "five Gold Rings"],
    ["sixth", "six Geese-a-Laying"],
    ["seventh", "seven Swans-a-Swimming"],
    ["eighth", "eight Maids-a-Milking"],
    ["ninth", "nine Ladies Dancing"],
    ["tenth", "ten Lords-a-Leaping"],
    ["eleventh", "eleven Pipers Piping"],
    ["twelfth", "twelve Drummers Drumming"],
]


def verse(verse):
    if verse == 1:
        gifts = ["a Partridge in a Pear Tree."]
    else:
        v = verse
        gifts = []
        while v > 0:
            gifts.append(day_and_gift[v][1])
            v -= 1

    return f"On the {day_and_gift[verse][0]} day of Christmas my true love gave to me: {', '.join(gifts)}"


def recite(start_verse, end_verse):
    return [verse(i) for i in range(start_verse, end_verse + 1)]
