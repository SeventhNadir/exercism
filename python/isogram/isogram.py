import re
    
def is_isogram(string):
    string = re.sub(r'\W+', '', string).lower()
    if len(string) == len(set(string)):
        return True
    return False