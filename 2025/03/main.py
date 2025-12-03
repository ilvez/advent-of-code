path = 'test_input'
path = 'input'

def find_max(head, tail):
    if len(tail) == 0:
        return 0
    return max([head * 10 + max(tail), find_max(tail[0], tail[1:])])


def to_list(str):
    return [int(i) for i in list(str)]


banks_list = [to_list(c) for c in [x.strip() for x in open(path).readlines()]]

result = 0
for battery_list in banks_list:
    result += find_max(battery_list[0], battery_list[1:])

print(f"Part 1: {result}")


def into_value(battery_list):
    value = int(''.join(map(str, battery_list)))
    return value


def survivor(current_list):
    return current_list if len(current_list) == 12 else survivor(max([current_list[:i] + current_list[i + 1:] for i in range(0, len(current_list))]))


result = 0

for battery_list in banks_list:
    s = into_value(survivor(battery_list))
    result += s

print(f"Part 2: {result}")
