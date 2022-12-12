# frozen_string_literal: true

require 'set'

# Day 12
class PathFinder
  def initialize
    @map = Map.new
  end

  def shortest_path_length(shortcuts:)
    @seen = Set[]
    @shortcuts = shortcuts
    @distance_map = { @map.start => 0 }
    visit(@map.start)
    visit(find_smallest_unvisited) while find_smallest_unvisited

    @distance_map.select { |k, _v| k == @map.finish }.values.first
  end

  private

  def visit(node)
    @map.adjancency_map[node].each do |connected_node|
      distance = @shortcuts && @map.map[node[0]][node[1]] == 'a' ? 1 : @distance_map[node] + 1
      if !@distance_map.key?(connected_node) || distance < @distance_map[connected_node]
        @distance_map[connected_node] = distance
      end
    end
    @seen.add(node)
  end

  def find_smallest_unvisited
    @distance_map.filter { |k, _v| !@seen.include?(k) }.min_by { |_, v| v }&.first
  end
end

# Input for Dijkstra's shortest path
class Map
  attr_reader :map, :adjancency_map, :start, :finish

  def initialize
    @map = File.readlines(ARGV[0], chomp: true).map(&:chars)
    @adjancency_map = {}
    @map.each_with_index do |line, y|
      line.each_with_index do |node, x|
        @adjancency_map[[y, x]] = possible_locations(y, x, map)
        @start = [y, x] if node == 'S'
        @finish = [y, x] if node == 'E'
      end
    end
  end

  private

  def possible_locations(y_pos, x_pos, map)
    [[y_pos - 1, x_pos], [y_pos, x_pos - 1], [y_pos, x_pos + 1], [y_pos + 1, x_pos]].filter do |y, x|
      y >= 0 &&
        y < map.length &&
        x >= 0 &&
        x < map[0].length &&
        allowed_climb?(map[y_pos][x_pos], map[y][x])
    end
  end

  def allowed_climb?(current, test)
    if current == 'S'
      test.ord - 1 < 'a'.ord
    elsif test == 'E'
      'z'.ord - 1 <= current.ord
    else
      test.ord - 1 <= current.ord
    end
  end
end

pathfinder = PathFinder.new

puts("Day 12, part 1: #{pathfinder.shortest_path_length(shortcuts: false)}")
puts("Day 12, part 2: #{pathfinder.shortest_path_length(shortcuts: true)}")
