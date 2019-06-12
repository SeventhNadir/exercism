class Luhn(object):
    def __init__(self, card_num):
        self.card_num = card_num.replace(" ", "")

    def is_valid(self):
        if len(self.card_num) <= 1:
            return False
        try:
            rev_card_num = [int(i) for i in self.card_num[::-1]]
        except:
            return False
        total = 0
        is_second_digit = False
        for i in rev_card_num:
            if is_second_digit:
                i *= 2
                if i >= 10:
                    i -= 9
                total += i
                is_second_digit = False
            else:
                total += i
                is_second_digit = True
        if total % 10 == 0:
            return True
        return False


