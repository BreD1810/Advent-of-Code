from string import ascii_lowercase

file = open("input.txt", "r")
box_ids = file.readlines()


def get_checksum(ids):
    double = 0
    triple = 0
    for id in ids:
        if contains_double_letter(id): double += 1
        if contains_triple_letter(id): triple += 1
    return double * triple


def contains_double_letter(id):
    for letter in ascii_lowercase:
        if id.count(letter) == 2:
            return True
    return False


def contains_triple_letter(id):
    for letter in ascii_lowercase:
        if id.count(letter) == 3:
            return True
    return False


print(get_checksum(box_ids))
