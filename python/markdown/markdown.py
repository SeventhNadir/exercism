import re


def headers(line):
    if re.match("######", line):
        line = f"<h6>{line[7:]}</h6>"
    if re.match("##", line):
        line = f"<h2>{line[3:]}</h2>"
    if re.match("#", line):
        line = f"<h1>{line[2:]}</h1>"
    return line


def bold(line):
    match = re.match("(.*)__(.*)__(.*)", line)
    if match:
        line = f"{match.group(1)}<strong>{match.group(2)}</strong>{match.group(3)}"
    return line


def italic(line):
    match = re.match("(.*)_(.*)_(.*)", line)
    if match:
        line = f"{match.group(1)}<em>{match.group(2)}</em>{match.group(3)}"
    return line


def unordered_list(unord_list, in_list):
    subline = unord_list.group(1)

    line = bold(subline)
    line = italic(subline)

    if in_list:
        line = f"<li>{subline}</li>"
    else:
        in_list = True
        line = f"<ul><li>{subline}</li>"
    return line, in_list


def parse(markdown):
    lines = markdown.split("\n")
    res = ""
    in_list = False
    in_list_append = False

    for line in lines:

        # Check for bold
        line = bold(line)

        # Check for italics
        line = italic(line)

        # Check for headers
        line = headers(line)

        # The absence of tags indicates we're in a paragraph
        if not re.match("\\* (.*)|<h|<ul|<p|<li", line):
            line = "<p>" + line + "</p>"

        # Check to see if we're inside an unordered list
        unord_list = re.match(r"\* (.*)", line)
        if unord_list:
            line, in_list = unordered_list(unord_list, in_list)
        else:
            if in_list:
                line = f"</ul>{line}"
                in_list = False

        res += line
    if in_list:
        res += "</ul>"
    return res
