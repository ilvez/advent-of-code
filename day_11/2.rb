require 'amazing_print'
require 'pry'

OCCUPIED = '#'
FREE = 'L'
FLOOR = '.'

def columns_count(map)
  map.first.size
end

def rows_count(map)
  map.size
end

def print_map(map)
  map.each { puts(_1.join) }
  ''
end

def occupied?(map, row, column)
  map[row][column] == OCCUPIED
end

def free?(map, row, column)
  map[row][column] == FREE
end

def on_map?(map, row, column)
  row >= 0 &&
    column >= 0 &&
    row < rows_count(map) &&
    column < columns_count(map)
end

def direction_multipliers(direction)
  {
    0 => [-1, -1],
    1 => [-1, 0],
    2 => [-1, 1],
    3 => [0, -1],
    4 => [nil, nil],
    5 => [0, 1],
    6 => [1, -1],
    7 => [1, 0],
    8 => [1, 1],
  }[direction]
end

def visible_occupied?(map, row, column, direction)
  row_multiplier, column_multiplier = direction_multipliers(direction)
  return if row_multiplier.nil? && column_multiplier.nil?
  (1..).each do
    spot_row = row + _1 * row_multiplier
    spot_col = column + _1 * column_multiplier

    return false unless on_map?(map, spot_row, spot_col)
    return false if free?(map, spot_row, spot_col)
    return true if occupied?(map, spot_row, spot_col)
  end
end

def occupied_around(map, row, column)
  (0..8).map { visible_occupied?(map, row, column, _1) }.compact
end

def adjacent_occupied?(map, row, column)
  occupied_around(map, row, column).any?
end

def can_sit?(map, row, column)
  free?(map, row, column) && !adjacent_occupied?(map, row, column)
end

def must_leave?(map, row, column)
  occupied?(map, row, column) && occupied_around(map, row, column).select { true == _1 }.size >= 5
end

def get_map_copy(map)
  Marshal.load(Marshal.dump(map))
end

def resolve_seat(current_round, next_round, row, column)
  if can_sit?(current_round, row, column)
    next_round[row][column] = OCCUPIED
  elsif must_leave?(current_round, row, column)
    next_round[row][column] = FREE
  end
end

def resolve_round(current_round)
  next_round = get_map_copy(current_round)
  rows_count(current_round).times.each { |row|
    columns_count(current_round).times.each do |column|
      resolve_seat(current_round, next_round, row, column)
    end
  }
  next_round
end

def import_map(file_name)
  map = []
  IO.readlines(file_name, chomp: true).each do
    map << _1.chars
  end
  map
end

def main
  current_round = import_map('input')

  loop do
    previous_round = get_map_copy(current_round)
    current_round = resolve_round(previous_round)
    break if current_round == previous_round
  end

  puts(current_round.flatten.select { _1 == OCCUPIED }.count)
end

main
