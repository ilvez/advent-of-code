file = open("test_input2", "r")
file = open("input", "r")

position = 50
count = 0
for line in file.readlines():
    line = line.strip()
    direction, amount = line[0], int(line[1:][-2:])
    if direction == "L":
        position = position - amount
    else:
        position += amount

    if position < 0:
        position = 100 + position
    elif position > 100:
        position = position - 100
    elif position == 100:
        position = 0
    if position == 0:
        count += 1
print(count)
