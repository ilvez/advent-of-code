defmodule Day03 do
  def run do
    IO.puts("Day 01, part 1: #{part1("test/resources/day03/input")}")
    IO.puts("Day 01, part 2: #{part2("test/resources/day03/input")}")
  end

  def part1(file) do
    file
    |> memory()
    |> arguments()
    |> multiply()
  end

  def part2(file) do
    file
    |> memory()
    |> selected_arguments()
    |> multiply()
  end

  def memory(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.join("")
  end

  def arguments(memory) do
    Regex.scan(~r"mul\((\d+),(\d+)\)", memory, capture: :all_but_first)
  end

  def selected_arguments(memory) do
    {:ok, memory} =
      Regex.scan(~r"mul\((\d+),(\d+)\)|don't\(\)|do\(\)", memory)
      |> Enum.reduce(%{on: true, memory: []}, fn
        ["don't()"], acc -> %{on: false, memory: acc.memory}
        ["do()"], acc -> %{on: true, memory: acc.memory}
        [_, a, b], %{on: true} = acc -> %{on: true, memory: acc.memory ++ [[a, b]]}
        _, acc -> acc
      end)
      |> Map.fetch(:memory)

    memory
  end

  def multiply(arguments) do
    arguments
    |> Enum.map(fn [a, b] -> String.to_integer(a) * String.to_integer(b) end)
    |> Enum.sum()
  end
end
