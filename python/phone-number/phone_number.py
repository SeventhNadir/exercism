from string import digits


class Phone(object):
    def __init__(self, phone_number):

        cleaned = [i for i in phone_number if i in digits]


        self.number = "".join(cleaned)
        self.country_code
        self.area_code = self.number[0:3]
        self.exchange_code = self.number[3:6]
        if self.area_code[0] != '2':
            raise ValueError("Invalid area code")
        if self.exchange_code[0] in '01':
            raise ValueError("Invalid exchange code")
        if len(cleaned) != 10 and (len(cleaned) == 11 and cleaned[0] != '1'):
            raise ValueError("Not a valid phone number")


    def pretty(self):
        return f"({self.number[0:3]}) {self.number[3:6]}-{self.number[6:]}"
