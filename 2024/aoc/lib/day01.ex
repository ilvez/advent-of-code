defmodule Day01 do
  def run do
    IO.puts("Day 01, part 1: #{part1("test/resources/day01/input")}")
    # IO.puts("Day 01, part 2: #{part2("test/resources/day02/input")}")
  end

  def part1(file) do
    lines(file)
    |> pairs()
    |> sorted()
    |> Enum.zip()
    |> Enum.map(fn {a, b} -> Kernel.abs(a - b) end)
    |> Enum.sum()
  end

  def pairs(lines) do
    lines
    |> Stream.map(&String.split(&1, "   "))
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
    |> Enum.map(fn sublist -> Enum.map(sublist, &String.to_integer/1) end)
  end

  def sorted(pairs) do
    Enum.map(pairs, fn sublist -> Enum.sort(sublist) end)
  end

  def lines(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> List.flatten()
  end
end
