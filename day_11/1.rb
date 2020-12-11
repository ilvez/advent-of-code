require 'amazing_print'
require 'pry'

def columns_count
  @map.first.size
end

def rows_count
  @map.size
end

def print_map(map)
  map.each { puts(_1.join) }
end

def occupied?(map, row, column)
  map[row][column] == '#'
end

def free?(map, row, column)
  map[row][column] == 'L'
end

def floor?(map, row, column)
  map[row][column] == '.'
end

def adjacent_occupied?(map, row, column)
  (row - 1..row + 1).flat_map { |r|
    next if r < 0 || r >= columns_count
    (column - 1..column + 1).map { |c|
      next if c < 0 || c >= columns_count
      next if c == column && r == row
      occupied?(map, r, c)
    }
  }.compact.any?
end

def can_sit?(map, row, column)
  free?(map, row, column) && !adjacent_occupied?(map, row, column)
end

def sit(current_round, next_round, row, column)
  next_round[row][column] = '#' if can_sit?(current_round, row, column)
end

def get_map_copy(map)
  Marshal.load(Marshal.dump(map))
end

def take_a_seat(current_round)
  next_round = get_map_copy(current_round)
  rows_count.times.each { |row|
    columns_count.times.each { |column|
      sit(current_round, next_round, row, column)
    }
  }
  next_round
end

@map =
  begin
    map = []
    IO.readlines('test-input', chomp: true).each do
      map << _1.chars
    end
    map
  end

current_round = @map

loop do
  previous_round = get_map_copy(current_round)
  current_round = take_a_seat(previous_round)
  break if current_round == previous_round
end

print_map(current_round)
