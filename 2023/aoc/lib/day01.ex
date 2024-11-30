defmodule Day01 do
  def part1(file) do
    lines(file)
    |> numbers()
    |> Enum.sum()
  end

  def numbers(lines) do
    lines
    |> Stream.map(&String.replace(&1, ~r/[^\d]/, ""))
    |> Stream.map(&reduce_string/1)
    |> Stream.map(&String.to_integer/1)
    |> Enum.to_list()
  end

  def lines(file) do
    File.stream!(file)
      |> Stream.map(&String.trim/1)
      |> Stream.map(&String.split(&1, ","))
      |> Enum.to_list()
      |> List.flatten()
  end

  defp reduce_string(string) do
    String.first(string) <> String.last(string)
  end
end
