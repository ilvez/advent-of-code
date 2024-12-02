defmodule Day02 do
  def run do
    IO.puts("Day 01, part 1: #{part1("test/resources/day02/input")}")
    # IO.puts("Day 01, part 2: #{part2("test/resources/day01/input")}")
  end

  def part1(file) do
    lines(file)
    |> levels()
    |> Enum.filter(&safe?/1)
    |> Enum.count()
  end

  def lines(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> List.flatten()
  end

  def levels(lines) do
    lines
    |> Enum.map(&String.split/1)
    |> Enum.map(fn sublist -> Enum.map(sublist, &String.to_integer/1) end)
  end

  def safe?(level) do
    chunks = Enum.chunk_every(level, 2, 1, :discard)
    direction = find_direction(chunks)

    chunks
    |> Enum.all?(fn [a, b] ->
      case direction do
        :ascending -> a < b && b - a < 4
        :descending -> a > b && a - b < 4
        :same -> raise "Same direction"
      end
    end)
  end

  def find_direction(chunks) do
    chunks
    |> Enum.reduce_while(:same, fn chunk, _ ->
      direction = direction(chunk)

      if direction == :same do
        {:cont, :same}
      else
        {:halt, direction}
      end
    end)
  end

  def direction(chunk) do
    case chunk do
      [a, b] when a < b -> :ascending
      [a, b] when a > b -> :descending
      [a, b] when a == b -> :same
    end
  end
end
