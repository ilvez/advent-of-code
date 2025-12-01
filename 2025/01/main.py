file = open("input", "r")

position = 50
count = 0
for line in file.readlines():
    line = line.strip()
    amount = int(line[1:])

    prev_pos = position
    count += amount // 100
    position = position + amount % 100 if line[0] != "L" else position - amount % 100

    if position < 0:
        position = 100 + position
        if prev_pos > 0:
            count += 1
    elif position > 100:
        position = position - 100
        count += 1
    elif position == 100:
        position = 0
    if position == 0:
        count += 1
print(count)
