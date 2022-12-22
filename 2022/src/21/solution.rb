require 'amazing_print'
class Monkey
  attr_reader :name

  def initialize(line)
    @name = line.split(': ')[0]
    @calculation = line.split(': ')[1]
  end

  def to_s
    "#{@name}: #{a}#{operation}#{b}"
  end

  def number?
    number

    true
  rescue ArgumentError
    false
  end

  def number
    Integer(@calculation)
  end

  def a
    @calculation.split[0]
  end

  def operation
    @calculation.split[1].to_sym
  end

  def humn_operation
    case operation
    when :+
      :-
    when :-
      :+
    when :*
      :/
    when :/
      :*
    else
      raise ArgumentError
    end
  end

  def b
    @calculation.split[2]
  end

  def humn?
    @name == 'humn'
  end
end

@monkies = File.readlines(ARGV[0], chomp: true)
               .to_h do |line|
                 monkey = Monkey.new(line)
                 [monkey.name, monkey]
               end

def find_value(monkey)
  if monkey.number?
    monkey.number
  else
    a = find_value(@monkies[monkey.a])
    b = find_value(@monkies[monkey.b])

    a.public_send(monkey.operation, b)
  end
end

def contains_humn?(monkey)
  return false if monkey.number?

  a = @monkies[monkey.a]
  b = @monkies[monkey.b]
  a.humn? || b.humn? || contains_humn?(a) || contains_humn?(b)
end

root = @monkies['root']
root_a = @monkies[root.a]
root_b = @monkies[root.b]

def humn_value(monkey, value)
  a = @monkies[monkey.a]
  b = @monkies[monkey.b]

  if a.humn?
    a_value(find_value(b), value, monkey)
  elsif b.humn?
    b_value(find_value(a), value, monkey)
  end

  if contains_humn?(a)
    humn_value(a, a_value(find_value(b), value, monkey))
  elsif contains_humn?(b)
    humn_value(b, b_value(find_value(a), value, monkey))
  end
end

def a_value(b_value, value, monkey)
  value.public_send(monkey.humn_operation, b_value)
end

def b_value(a_value, value, monkey)
  case monkey.operation
  when :/
    a_value / value
  when :-
    a_value - value
  else
    value.public_send(monkey.humn_operation, a_value)
  end
end

humn_value = if contains_humn?(root_a)
               humn_value(root_a, find_value(root_b))
             else
               humn_value(root_b, find_value(root_a))
             end

puts "Day 22 part 1: #{find_value(root)}"
puts "Day 22 part 2: #{humn_value}"
