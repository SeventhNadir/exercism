class Clock(object):
    def __init__(self, hour, minute):
        self.minutes = ((hour * 60) + minute) % 1440

    def __repr__(self):
        return f"{self.minutes//60:02}:{self.minutes%60:02}"

    def __eq__(self, other):
        return self.minutes == other.minutes
        
    def __add__(self, minutes):
        self.minutes = (self.minutes + minutes) % 1440
        return self

    def __sub__(self, minutes):
        self.minutes = (self.minutes - minutes) % 1440
        return self
