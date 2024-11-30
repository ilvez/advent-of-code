defmodule Day02 do
  def run do
    IO.puts("Day 02, part 1: #{part1("test/resources/day02/input")}")
    IO.puts("Day 02, part 2: #{part2("test/resources/day02/input")}")
  end

  def part1(file) do
    lines(file)
    |> parse_games()
    |> find_possible_games()
    |> Enum.map(&(&1.game))
    |> Enum.sum()
  end

  def part2(file) do
    lines(file)
    |> parse_games()
    |> Enum.map(&(&1.sets))
    |> Enum.map(&find_maximums/1)
    |> Enum.map(&pow/1)
    |> Enum.sum()
  end

  def find_maximums(sets) do
    # IO.puts(sets)
    sets
    |> Enum.flat_map(&Map.to_list/1)
    |> Enum.group_by(fn {key, _value} -> key end, fn {_key, value} -> value end)
    |> Enum.map(fn {_, counts} -> Enum.max(counts) end)
  end

  def pow(maximums) do
    Enum.reduce(maximums, 1, &Kernel.*/2)
  end

  def find_possible_games(games) do
    maximums = %{red: 12, green: 13, blue: 14}

    Enum.filter(games, fn %{sets: sets} -> find_impossible_set(sets, maximums) == nil end)
  end

  def find_impossible_set(sets, maximums) do
    Enum.find_value(sets, fn set ->
      cond do
        Enum.all?(set, fn {color, count} -> count <= Map.get(maximums, color) end) -> nil
        true -> set
      end
    end)
  end

  def parse_games(lines) do
    lines |> Enum.map(&parse_game/1)
  end

  def parse_game(line) do
    [game_part, sets_part] = line |> String.split(":")

    game = Enum.at(game_part |> String.split(" "), 1) |> String.to_integer()

    sets = sets_part |> String.split(";") |> Enum.map(&parse_set/1)

    %{game: game, sets: sets}
  end

  def parse_set(set) do
    set
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&parse_color/1)
    |> Enum.reduce(%{}, fn {color, count}, acc -> Map.put(acc, color, count) end)
  end

  def parse_color(color_line) do
    [count, color] = color_line |> String.trim() |> String.split(" ")

    {String.to_atom(color), String.to_integer(count)}
  end

  def lines(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> List.flatten()
  end
end
