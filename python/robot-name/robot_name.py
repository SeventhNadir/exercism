from string import ascii_uppercase
from random import sample

import cProfile


class Robot(object):
    def __init__(self):
        try:
            self.name = next(Robot.random_names)
        except StopIteration as e:
            raise Exception("We've run out of unique names to assign to robots")

    def reset(self):
        self.__init__()

    namespace = sample(
        [
            f"{i}{j}{k:03d}"
            for i in ascii_uppercase
            for j in ascii_uppercase
            for k in range(0, 1000)
        ],
        26 * 26 * 1000,
    )

    random_names = (name for name in namespace)