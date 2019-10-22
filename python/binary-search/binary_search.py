def find(search_list, value, start=0, end=None):
    if end == None:
        end = len(search_list) - 1
    if start > end:
        raise ValueError("Value not found")

    midpoint = (start+end) // 2
    if value == search_list[midpoint]:
        return midpoint
    if value < search_list[midpoint]:
        return find(search_list, value, start, midpoint - 1)
    
    return find(search_list, value, midpoint+1, end)
  