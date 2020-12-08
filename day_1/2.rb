require 'amazing_print'

def find_2020(base, tail)
  return nil if tail.empty?

  answer = find_for_base(base, tail)
  return answer unless answer.nil?

  next_a, *next_tail = tail
  find_2020(next_a, next_tail)
end

def find_for_base(base, tail)
  a, *bs = tail
  return nil if tail.empty?

  bs.each do |b|
    return [base, a, b] if base + a + b == 2020
  end

  find_for_base(base, bs)
end

input_nums = IO.readlines('input_1', chomp: true).map { Integer(_1) }

base, *tail = input_nums
x, y, z = find_2020(base, tail)
ap(x: x, y: y, z: z, answer: x * y * z)
