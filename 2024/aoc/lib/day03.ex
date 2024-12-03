defmodule Day03 do
  def run do
    IO.puts("Day 01, part 1: #{part1("test/resources/day03/input")}")
    # IO.puts("Day 01, part 2: #{part2("test/resources/day03/input")}")
  end

  def part1(file) do
    file
    |> memory()
    |> multiply()
  end

  def memory(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.join("")
  end

  def multiply(memory) do
    Regex.scan(~r"mul\((\d+),(\d+)\)", memory, capture: :all_but_first)
    |> Enum.map(fn [a, b] -> String.to_integer(a) * String.to_integer(b) end)
    |> Enum.sum()
  end
end
