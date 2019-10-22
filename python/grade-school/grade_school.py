class School(object):
    def __init__(self):
        self.students = []

    def add_student(self, name, grade):
        self.students.append((name, grade))
        self.students.sort()

    def roster(self):
        roster = sorted([student for student in self.students],
                        key=lambda t: (t[1], t[0]))
        return [student[0] for student in roster]
        
    def grade(self, grade_number):
        return [student[0] for student in self.students if student[1] == grade_number]
