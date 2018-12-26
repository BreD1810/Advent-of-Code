# 1 @ 1,3: 4x4
import re
import numpy


def count_overlap(claims):
    overlap = numpy.count_nonzero(create_claims_array(claims) == "x")
    return overlap


def create_claims_array(claims):
    claims_array = get_blank_claims_array(claims)
    for claim in claims:
        matcher = re.match("#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", claim)
        # redfine the bounds
        min_x = int(matcher.group(2))
        max_x = min_x + int(matcher.group(4))
        min_y = int(matcher.group(3))
        max_y = min_y + int(matcher.group(5))
        for x in range(min_x, max_x):
            for y in range(min_y, max_y):
                if claims_array[x][y] != "":
                    claims_array[x][y] = "x"
                else:
                    claims_array[x][y] = matcher.group(1)
    return claims_array


def get_blank_claims_array(claims):
    max_x = 0
    max_y = 0
    for claim in claims:
        matcher = re.match("#\d+ @ (\d+),(\d+): (\d+)x(\d+)", claim)
        x = int(matcher.group(1)) + (int(matcher.group(3)))
        y = int(matcher.group(2)) + (int(matcher.group(4)))
        if x > max_x: max_x = x
        if y > max_y: max_y = y
    # Initialise the array
    claims_array = numpy.empty((max_x, max_y), dtype=str)
    return claims_array


def get_no_overlap(claims):
    claims_array = create_claims_array(claims)
    for claim in claims:
        matcher = re.match("#(\d+) @ (\d+),(\d+): (\d+)x(\d+)", claim)
        # redfine the bounds
        min_x = int(matcher.group(2))
        max_x = min_x + int(matcher.group(4))
        min_y = int(matcher.group(3))
        max_y = min_y + int(matcher.group(5))
        if not check_overlap(claims_array, min_x, max_x, min_y, max_y): return matcher.group(1)
    return "None"


def check_overlap(claims_array, min_x, max_x, min_y, max_y):
    for x in range(min_x, max_x):
        for y in range(min_y, max_y):
            if claims_array[x][y] == "x":
                return True
    return False


file = open("input.txt", "r")
lines = file.readlines()
print("Part 1: " + str(count_overlap(lines)))
print("Part 2: " + get_no_overlap(lines))
