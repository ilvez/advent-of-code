from copy import deepcopy

def safe_coord(coord, x, y):
    x1, y1 = coord[0], coord[1]
    return not (x1 < 0) and not (y1 < 0) and (y1 < len(floor_plan) and x1 < len(floor_plan[y1])) and not (x1 == x and y1 == y)


def safe_coords(x, y):
    return [c for c in [(i, j) for i in [x - 1, x, x + 1] for j in [y - 1, y, y + 1]] if safe_coord(c, x, y)]


def neighbour_count(coords, floors):
    return len([paper for paper in [floors[y][x] for (x, y) in coords] if paper == '@'])


def papers_accessible(floor_plan):
    result = 0
    for y in range(0, len(floor_plan)):
        for x in range(0, len(floor_plan[y])):
            if floor_plan[y][x] == '@':
                if neighbour_count(safe_coords(x, y), floor_plan) < 4:
                    result += 1
    return result


def count_papers(floor_plan):
    result = 0
    for y in range(0, len(floor_plan)):
        for x in range(0, len(floor_plan[y])):
            if floor_plan[y][x] == '@':
                result += 1
    return result


def print_map(floors):
    for i in floors:
        print(''.join(map(str, i)))


def modify_floor_plan(floors):
    new_floor = deepcopy(floors)
    for y in range(0, len(floors)):
        for x in range(0, len(floors[y])):
            if floors[y][x] == '@':
                if neighbour_count(safe_coords(x, y), floors) < 4:
                    new_floor[y][x] = '.'
    return new_floor


path = 'test_input'
path = 'input'

floor_plan = [[x for x in y.strip()] for y in open(path).readlines()]

print(f"Part 1: {papers_accessible(floor_plan)}")

map_copy = deepcopy(floor_plan)
count = -1

while True:
    prev_count = count
    map_copy = modify_floor_plan(map_copy)
    count = count_papers(map_copy)
    if count == prev_count:
        break

print(f"Part 2: {count_papers(floor_plan) - count_papers(map_copy)}")
