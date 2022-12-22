require 'amazing_print'
class Monkey
  attr_reader :name

  def initialize(line)
    @name = line.split(': ')[0]
    @calculation = line.split(': ')[1]
  end

  def to_s
    "#{@name}: #{number?}"
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

  def b
    @calculation.split[2]
  end
end

monkies = File.readlines(ARGV[0], chomp: true)
              .to_h do |line|
                monkey = Monkey.new(line)
                [monkey.name, monkey]
              end

def find_value(monkies, monkey)
  if monkey.number?
    monkey.number
  else
    a = find_value(monkies, monkies[monkey.a])
    b = find_value(monkies, monkies[monkey.b])

    a.public_send(monkey.operation, b)
  end
end

puts "Day 22 part 1: #{find_value(monkies, monkies['root'])}"
