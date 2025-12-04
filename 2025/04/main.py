path = 'test_input'
path = 'input'

floor_plan = [[x for x in y.strip()] for y in open(path).readlines()]

def safe_coord(coord, x, y):
    x1, y1 = coord[0], coord[1]
    return not (x1 < 0) and not (y1 < 0) and (y1 < len(floor_plan) and x1 < len(floor_plan[y1])) and not (x1 == x and y1 == y)


def safe_coords(x, y):
    return [c for c in [(i, j) for i in [x - 1, x, x + 1] for j in [y - 1, y, y + 1]] if safe_coord(c, x, y)]


def neighbour_count(coords):
    return len([paper for paper in [floor_plan[y][x] for (x, y) in coords] if paper == '@'])


result = 0
for y in range(0, len(floor_plan)):
    for x in range(0, len(floor_plan[y])):
        if floor_plan[y][x] == '@':
            if neighbour_count(safe_coords(x, y)) < 4:
                result += 1
print(f"Part 1: {result}")
