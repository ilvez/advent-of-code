require 'amazing_print'

# input = IO.readlines('input', chomp: true)
# ap(input.size)

def calc_seat(input)
  row_steps = input[..6]
  col_steps = input[7..]

  row = position(0, max_position(row_steps), row_steps)
  col = position(0, max_position(col_steps), col_steps)

  row * 8 + col
end

def position(lower, upper, steps)
  ap("upper #{lower}, #{upper}, #{steps[0]}, #{steps[1..]}")
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

def max_position(input)
  2**input.size - 1
end

seat = "FBFBBFFRLR"
ap(calc_seat(seat))
