def slices(series, length):
    if length < 1:
        raise ValueError("Length must be a positive integer")
    if length > len(series):
        raise ValueError("Slice length is larger than the length of the series")

    return [series[i : i + length] for i in range(0, len(series) - length + 1)]
