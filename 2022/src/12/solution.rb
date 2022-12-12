# frozen_string_literal: true

require 'amazing_print'
require 'set'

# Day 12
class PathFinder
  def initialize(shortcuts:)
    @shortcuts = shortcuts
    @map = File.readlines(ARGV[0], chomp: true).map(&:chars)
    init_adacency_list
    @distance_map = { @start => 0 }
    @seen = Set[]

    visit(@start)

    visit(find_smallest_unvisited) while find_smallest_unvisited
  end

  def shortest_path_legth
    @distance_map.select { |k, _v| k == @finish }.values.first
  end

  private

  def init_adacency_list
    @adj_map = {}
    @map.each_with_index do |line, y|
      line.each_with_index do |node, x|
        @adj_map[[y, x]] = possible_locations(y, x)
        @start = [y, x] if node == 'S'
        @finish = [y, x] if node == 'E'
      end
    end
  end

  def possible_locations(y_pos, x_pos)
    [[y_pos - 1, x_pos], [y_pos, x_pos - 1], [y_pos, x_pos + 1], [y_pos + 1, x_pos]].filter do |y1, x1|
      y1 >= 0 &&
        y1 < @map.length &&
        x1 >= 0 &&
        x1 < @map[0].length &&
        allowed_climb?(@map[y_pos][x_pos], @map[y1][x1])
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

  def visit(node)
    @adj_map[node].each do |connected_node|
      distance = @shortcuts && @map[node[0]][node[1]] == 'a' ? 1 : @distance_map[node] + 1
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

puts("Day 12, part 1: #{PathFinder.new(shortcuts: false).shortest_path_legth}")
puts("Day 12, part 2: #{PathFinder.new(shortcuts: true).shortest_path_legth}")
