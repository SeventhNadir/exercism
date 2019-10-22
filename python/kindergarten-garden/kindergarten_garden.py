class Garden(object):

    default_students = [
        "Alice",
        "Bob",
        "Charlie",
        "David",
        "Eve",
        "Fred",
        "Ginny",
        "Harriet",
        "Ileana",
        "Joeseph",
        "Kincaid",
        "Larry",
    ]

    symbol_to_plant = {"V": "Violets", "R": "Radishes", "C": "Clover", "G": "Grass"}

    def __init__(self, diagram, students=default_students):
        self.students = sorted(students)
        self.diagram = diagram.split('\n')

    def plants(self, student):
        plant_array = []
        index = self.students.index(student)
        for row in self.diagram:
            plant_array.append(row[2*index])
            plant_array.append(row[2 * index + 1])
        return [Garden.symbol_to_plant[symbol] for symbol in plant_array]
