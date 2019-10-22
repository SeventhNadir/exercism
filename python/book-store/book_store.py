def determine_grouping(basket):
    grouping = []
    for item in list(basket):
        if type(item) == list:
            continue
        if item not in grouping:
            grouping.append(item)
            basket.remove(item)
    basket.insert(0, grouping)
    if len([item for item in basket if type(item) != list]) > 0:
        determine_grouping(basket)

    return basket


def groups_three_five(total_groups):
    """
    Rather than fiddle with any advanced math, why not try a simpler approach.
    We know that any two groups of 3 and 5 books can be trivially converted into two groups of 4.
    We know that the difference between these two groupings is 40c
    Why not subtract 40c every time we have one group of 3 and one group of 5?
    """
    correction = 0
    threes = len([group for group in total_groups if len(group) == 3])
    fives = len([group for group in total_groups if len(group) == 5])
    if threes and fives:
        correction += min(threes, fives) * 40
    return correction


def determine_discount(group_size):
    discount = {0: 0, 1: 100, 2: 95, 3: 90, 4: 80, 5: 75}
    return discount[group_size]


def total(basket):
    total = 0
    total_groups = determine_grouping(basket)
    correction = groups_three_five(total_groups)
    for item in total_groups:
        total += 8 * len(item) * determine_discount(len(item))
    return total - correction
