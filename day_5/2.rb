require 'amazing_print'


def calc_seat(input)
  row_steps = input[..6]
  col_steps = input[7..]

  row = position(0, max_position(row_steps), row_steps)
  col = position(0, max_position(col_steps), col_steps)

  row * 8 + col
end

def position(lower, upper, steps)
  return lower if steps.empty?

  mid = lower + ((upper - lower) / 2.0).ceil
  if %w[F L].include?(steps[0])
    position(lower, mid - 1, steps[1..])
  elsif %w[B R].include?(steps[0])
    position(mid, upper, steps[1..])
  else
    raise ArgumentError
  end
end

def find_seat(array)
  ap(array)
  return array[0] + 1 if array.size == 1

  mid = array.size / 2
  mid_in_array = array[mid]
  mid_by_order = array[0] + mid

  if mid_in_array <= mid_by_order
    find_seat(array[mid..])
  else
    find_seat(array[..(mid - 1)])
  end
end

def max_position(input)
  2**input.size - 1
end

input = IO.readlines('input', chomp: true)
ap(find_seat(input.map{ calc_seat(_1) }.sort))
