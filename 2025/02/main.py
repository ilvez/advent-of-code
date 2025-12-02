path = 'test_input'
path = 'input'
input = open(path).read().strip().split(',')

input_list = [x.split('-') for x in input]

result = 0
for pair in input_list:
    for i in range(int(pair[0]), int(pair[1]) + 1):
        istr = str(i)
        if len(istr) % 2 == 0:
            if istr[:len(istr) // 2] == istr[len(istr) // 2:]:
                result += i
print(result)
