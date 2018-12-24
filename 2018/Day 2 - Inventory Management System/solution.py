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


def get_common_letters_of_prototypes(ids):
    for id in ids:
        for other_id in ids:
            if id != other_id:
                matching_characters = [id[index] for index in range(0, len(id)) if id[index] == other_id[index]]
                if len(matching_characters) == len(id) - 1: return "".join(matching_characters)
    return ""


print("Part 1: " + str(get_checksum(box_ids)))
print("Part 2: " + get_common_letters_of_prototypes(box_ids))
