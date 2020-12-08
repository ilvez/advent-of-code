require 'amazing_print'
require 'pry'

parsed_input = IO.readlines('input', chomp: true).map do |input|
  parts = input.split(' ')
  min, max = parts[0].split('-').map(&:to_i)
  char = parts[1][0]
  password = parts[2]
  { pos_min: min, pos_max: max, char: char, password: password }
end

def valid?(input)
  (input[:password][input[:pos_min] - 1] == input[:char]) ^
    (input[:password][input[:pos_max] - 1] == input[:char])
end

ap(parsed_input.select { valid?(_1) }.size)
