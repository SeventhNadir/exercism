# Globals for the bearings
# Change the values as you see fit
EAST = 1
NORTH = 0
WEST = 3
SOUTH = 2


class Robot(object):
    def __init__(self, bearing=NORTH, x=0, y=0):
        self.bearing = bearing
        self.coordinates = (x, y)

    def turn_left(self):
        self.bearing = (self.bearing - 1) % 4

    def turn_right(self):
        self.bearing = (self.bearing + 1) % 4

    def advance(self):
        x, y = self.coordinates

        if self.bearing == NORTH:
            self.coordinates = x, y + 1

        if self.bearing == EAST:
            self.coordinates = x + 1, y

        if self.bearing == SOUTH:
            self.coordinates = x, y - 1

        if self.bearing == WEST:
            self.coordinates = x - 1, y

    def simulate(self, instructions):
        for instruction in instructions:
            if instruction == "L":
                self.turn_left()
            if instruction == "R":
                self.turn_right()
            if instruction == "A":
                self.advance()
