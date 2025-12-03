path = 'test_input'
path = 'input'

def find_max(head, tail):
    if len(tail) == 0:
        return 0
    return max([head * 10 + max(tail), find_max(tail[0], tail[1:])])

banks_list = [[int(i) for i in list(c)] for c in [x.strip() for x in open(path).readlines()]]

result = 0
for battery_list in banks_list:
    result += find_max(battery_list[0], battery_list[1:])

print(f"Part 1: {result}")
