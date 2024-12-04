defmodule Day01 do
  def run do
    IO.puts("Day01, part 1: #{part1("test/resources/day01/input")}")
    IO.puts("Day01, part 2: #{part2("test/resources/day01/input")}")
  end

  def part1(file) do
    lines(file)
    |> pairs()
    |> sorted()
    |> Enum.zip()
    |> Enum.map(fn {a, b} -> abs(a - b) end)
    |> Enum.sum()
  end

  def part2(file) do
    [locations, popularities] = lines(file) |> pairs()

    Enum.reduce(locations, 0, fn location, acc ->
      count =
        Enum.filter(popularities, fn popularity -> popularity == location end)
        |> Enum.count()

      acc + location * count
    end)
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
