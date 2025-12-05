path = 'test_input'
# path = 'input'

ranges = []
ids = []
for line in open(path).readlines():
    line = line.strip()
    if "-" in line:
        ranges += [[int(x) for x in line.split("-")]]
    elif len(line) > 0:
        ids += [int(line)]

def in_range(id, range):
    return id >= range[0] and id <= range[1]

result = 0
for id in ids:
    for r in ranges:
        if in_range(id, r):
            result += 1
            break
print(f"Part 1: {result}")

result = 0
for r in ranges:
    result += r[1] - r[0]

print(f"Part 2: {result}")
