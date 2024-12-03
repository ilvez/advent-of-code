defmodule Day03 do
  def run do
    # IO.puts("Day 01, part 1: #{part1("test/resources/day03/input")}")
    # IO.puts("Day 01, part 2: #{part2("test/resources/day03/input")}")
  end

  def memory(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.join("")
  end

  def multiplications(memory) do
    Regex.scan(~r"mul(.*)", memory)
  end
end
