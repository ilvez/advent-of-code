from collections import Counter

path = 'test_input'
path = 'input'

input_list = [x.split('-') for x in open(path).read().strip().split(',')]

result = 0
for pair in input_list:
    for i in range(int(pair[0]), int(pair[1]) + 1):
        istr = str(i)
        if len(istr) % 2 == 0:
            if istr[:len(istr) // 2] == istr[len(istr) // 2:]:
                result += i
print(f"Part 1: {result}")


def group(s, n):
    return Counter(s[i:i + n] for i in range(0, len(s), n))


def is_illegal(string):
    for i in range(1, len(string) // 2 + 1):
        groups = group(string, i)
        if len(groups) == 1:
            return True


result = 0
for pair in input_list:
    for number in range(int(pair[0]), int(pair[1]) + 1):
        if is_illegal(str(number)):
            result += number
print(f"Part 2: {result}")
