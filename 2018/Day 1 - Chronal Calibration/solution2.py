def get_repeated_total():
    file = open("input.txt", "r")
    lines = file.readlines()
    current_total = int(0)
    previous_totals = {0}

    while True:
        for x in lines:
            current_total += int(x)
            if current_total in previous_totals:
                return current_total
            else:
                previous_totals.add(current_total)



print(get_repeated_total())
