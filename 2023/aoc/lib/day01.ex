defmodule Day01 do
  def run do
    IO.puts("Day 01, part 1: #{part1("test/resources/day01/input")}")
    IO.puts("Day 01, part 2: #{part2("test/resources/day01/input")}")
  end

  def part1(file) do
    lines(file)
    |> numbers1()
    |> Enum.sum()
  end

  def part2(file) do
    lines(file)
    |> numbers2()
    |> Enum.sum()
  end

  def numbers1(lines) do
    lines
    |> Stream.map(&String.replace(&1, ~r/[^\d]/, ""))
    |> Stream.map(&reduce_string/1)
    |> Stream.map(&String.to_integer/1)
    |> Enum.to_list()
  end

  def numbers2(lines) do
    lines
    |> Stream.map(&word_replace/1)
    |> Stream.map(&String.replace(&1, ~r/[^\d]/, ""))
    |> Stream.map(&reduce_string/1)
    |> Stream.map(&String.to_integer/1)
    |> Enum.to_list()
  end

  def lines(file) do
    File.stream!(file)
      |> Stream.map(&String.trim/1)
      |> Enum.to_list()
      |> List.flatten()
  end

  defp reduce_string(string) when byte_size(string) <1 do
    "00"
  end

  defp reduce_string(string) do
    String.first(string) <> String.last(string)
  end

  def word_replace(string) do
    parse_both(string)
    |> Enum.map(&Integer.to_string/1)
    |> Enum.join()
  end

  def parse_both(""), do: []
  def parse_both(<<c>>   <> rest) when c in ?0..?9, do: [c - ?0 | parse_both(rest)]
  def parse_both("zero"  <> rest), do: [0 | parse_both("o" <> rest)]
  def parse_both("one"   <> rest), do: [1 | parse_both("e" <> rest)]
  def parse_both("two"   <> rest), do: [2 | parse_both("o" <> rest)]
  def parse_both("three" <> rest), do: [3 | parse_both("e" <> rest)]
  def parse_both("four"  <> rest), do: [4 | parse_both("r" <> rest)]
  def parse_both("five"  <> rest), do: [5 | parse_both("e" <> rest)]
  def parse_both("six"   <> rest), do: [6 | parse_both("x" <> rest)]
  def parse_both("seven" <> rest), do: [7 | parse_both("n" <> rest)]
  def parse_both("eight" <> rest), do: [8 | parse_both("t" <> rest)]
  def parse_both("nine"  <> rest), do: [9 | parse_both("e" <> rest)]
  def parse_both(<<_>>   <> rest), do: parse_both(rest)
end
