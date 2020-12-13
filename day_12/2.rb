require 'amazing_print'
require 'pry'

def parse_input(file_name)
  IO.readlines(file_name, chomp: true).map do |line|
    {
      direction: line[0],
      distance: line[1..].to_i
    }
  end
end

class Boat
  N = 'N'
  E = 'E'
  S = 'S'
  W = 'W'
  R = 'R'
  L = 'L'
  F = 'F'

  def initialize(location: 0, direction: E)
    @north_south_position = 0
    @east_west_position = 0
    @direction = E
    @waypoint = Waypoint.new(1, 10)
  end

  def move(instruction)
    direction = instruction[:direction]
    distance = instruction[:distance]
    case direction
    when F
      n, e = @waypoint.forward(distance)
      @north_south_position += n
      @east_west_position += e
    when N, E, S, W
      @waypoint.move(direction, distance)
    when R, L
      @waypoint.rotate(direction, distance)
    end
  end

  def to_s
    "Boat(direction: #{@direction}, waypoint: #{@waypoint})"
  end

  def manhatten_distance
    @north_south_position.abs + @east_west_position.abs
  end

  private

  class Waypoint
    def initialize(n, e)
      @n = n
      @e = e
    end

    def to_s
      "Waypoint(n: #{@n}, e: #{@e})"
    end

    def forward(times)
      [@n * times, @e * times]
    end

    def move(direction, distance)
      case direction
      when N
        @n += distance
      when E
        @e += distance
      when S
        @n -= distance
      when W
        @e -= distance
      end
    end

    def rotate(direction, angle)
      (angle / 90).times.each { direction == R ? turn_right : turn_left }
    end

    private

    def turn_right
      old_n = @n
      old_e = @e
      @n = -old_e
      @e = old_n
    end

    def turn_left
      old_n = @n
      old_e = @e
      @n = old_e
      @e = -old_n
    end
  end
end

instructions = parse_input('input')

boat = Boat.new

instructions.each { |instruction| boat.move(instruction) }
ap(boat.manhatten_distance)
