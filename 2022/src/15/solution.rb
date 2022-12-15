require 'amazing_print'

class Point
  attr_reader :x, :y

  def self.from_line(x, y)
    new(x.split('=')[1].to_i, y.split('=')[1].to_i)
  end

  def initialize(x, y)
    @x = x
    @y = y
  end

  def relative_to(other)
    if other.x.negative?
      Point.new(x + other.x, other.y)
    else
      Point.new(other.x, other.y)
    end
  end

  def ==(other)
    x == other.x && y == other.y
  end

  def to_s
    "x=#{@x},y=#{@y}"
  end
end

class Coverage
  attr_reader :sensor, :beacon

  def initialize(line)
    @sensor = Point.from_line(line[2], line[3])
    @beacon = @sensor.relative_to(Point.from_line(line[-2], line[-1]))
  end

  def to_s
    "#{max_x}, sensor(#{@sensor}), beacon(#{@beacon}): #{distance}"
  end

  def slice_on_line(y)
    diff = (y - sensor.y).abs
    line_length = distance - diff
    ((sensor.x - line_length)..(sensor.x + line_length))
      .map { _1 }
      .filter { |x_pos| self != Point.new(x_pos, y) }
  end

  private

  def distance_between(a, b)
    (a.x - b.x).abs + (a.y - b.y).abs
  end

  def distance
    distance_between(@sensor, @beacon)
  end
end

input = File.readlines(ARGV[0], chomp: true)
            .map(&:split)
coverages = input.map { |line| Coverage.new(line) }
no_beacons_on_line = coverages
                     .map { |c| c.slice_on_line(ARGV[1].to_i) }
                     .flatten
                     .uniq
                     .length - 1
puts("Day 15 part 1: #{no_beacons_on_line}")
