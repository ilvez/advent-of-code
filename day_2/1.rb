require 'amazing_print'

parsed_input = IO.readlines('input', chomp: true).map do |input|
  parts = input.split(' ')
  min, max = parts[0].split('-').map(&:to_i)
  char = parts[1][0]
  password = parts[2]
  { min: min, max: max, char: char, password: password }
end

def valid?(input)
  count = input[:password].count(input[:char])
  count >= input[:min] && count <= input[:max]
end

ap(parsed_input.select { valid?(_1) }.size)
