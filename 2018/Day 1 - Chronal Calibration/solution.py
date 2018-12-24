file = open("input.txt", "r")
lines = file.readlines()

total = int(0)

for x in lines:
    total += int(x)

print(total)