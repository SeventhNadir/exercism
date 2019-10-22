class Matrix(object):
    def __init__(self, matrix_string):

        # Convert a string into a list of lists
        matrix_rows = matrix_string.splitlines()
        matrix_rows = [row.split() for row in matrix_rows]
        matrix_rows = [[int(i) for i in row] for row in matrix_rows]

        if len(set([len(row) for row in matrix_rows])) != 1:
            raise Exception("Unbalanced matrix, rows not uniform length")
        
        self.matrix = matrix_rows

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return [row[index - 1] for row in self.matrix]
