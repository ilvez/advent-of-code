def find_2020(a, tail)
  tail.each do |b|
    return [a, b] if a + b == 2020
  end
  next_a, *next_tail = tail
  find_2020(next_a, next_tail)
end

input_nums = IO.readlines('input_1', chomp: true).map { Integer(_1) }

a, *tail = input_nums
x, y = find_2020(a, tail)
puts x * y
