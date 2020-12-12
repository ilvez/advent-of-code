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
  end

  def move(instruction)
    direction = instruction[:direction]
    distance = instruction[:distance]
    case direction
    when F
      move_direction(@direction, distance)
    when N, E, S, W
      move_direction(direction, distance)
    when R, L
      turn(current_direction_pos(direction, distance))
    end
  end

  def to_s
    "Boat(direction: #{@direction}, ns_pos: #{@north_south_position}, ew_pos: #{@east_west_position})"
  end

  def manhatten_distance
    @north_south_position.abs + @east_west_position.abs
  end

  private

  CARDINAL_DIRECTIONS = [N, E, S, W]

  def move_direction(direction, distance)
    case direction
    when N
      @north_south_position += distance
    when E
      @east_west_position += distance
    when S
      @north_south_position -= distance
    when W
      @east_west_position -= distance
    end
  end

  def turn(current_direction_pos)
    @direction = CARDINAL_DIRECTIONS[current_direction_pos % 4]
  end

  def current_direction_pos(turn_direction, angle)
    current_direction_pos = turn_direction_multiplier(turn_direction) *
      direction_multiplier(angle) +
      CARDINAL_DIRECTIONS.find_index(@direction)
  end

  def turn_direction_multiplier(turn_direction)
    turn_direction == R ? 1 : - 1
  end

  def direction_multiplier(angle)
    angle / 90
  end
end

instructions = parse_input('input')

boat = Boat.new

instructions.each { |instruction| boat.move(instruction) }
ap(boat.manhatten_distance)
