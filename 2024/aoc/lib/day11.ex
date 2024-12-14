defmodule Day11 do
  def run do
    IO.puts("Day11, part 1: #{part1("test/resources/day11/input")}")
    # IO.puts("Day11, part 2: #{part2("test/resources/day11/input")}")
  end

  def part1(file) do
    stones = Enum.at(lines(file), 0) |> parse_stones()
    blink_n_times(stones, 25) |> Enum.count()
  end

  def part2(file) do
    stones = Enum.at(lines(file), 0) |> parse_stones()
    blink_n_times(stones, 75) |> Enum.count()
  end

  def lines(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> List.flatten()
  end

  def blink_count(stone, n) do
    blinked = blink_stone(stone)
    cond do
      n == 0 -> blinked |> Enum.count()
      length(blinked) == 1 -> blink_count(Enum.at(blinked, 0), n - 1)
      true -> blink_count(Enum.at(blinked, 0), n - 1) + blink_count(Enum.at(blinked, 1), n - 1)
    end
  end

  def blink_stone(stone) do
    stone_str = Integer.to_string(stone)
    stone_str_len = String.length(stone_str)

    cond do
      stone == 0 ->
        [1]

      rem(stone_str_len, 2) == 0 ->
        str_chunk(stone_str, round(stone_str_len / 2))
        |> Enum.map(&String.to_integer/1)

      true ->
        [stone * 2024]
    end
  end

  def blink_n_times(stones, n) do
    Enum.reduce(1..n, stones, fn _, acc -> blink(acc) end)
  end

  def blink(stones) do
    Enum.map(stones, fn stone -> blink_stone(stone) end) |> flatten()
  end

  def str_chunk(str, chunk_at) do
    String.codepoints(str)
    |> Enum.chunk_every(chunk_at)
    |> Enum.map(&Enum.join/1)
  end

  def flatten(list) do
    Enum.flat_map(list, fn
      sublist when is_list(sublist) -> flatten(sublist)
      element -> [element]
    end)
  end

  def parse_stones(input) do
    input
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
