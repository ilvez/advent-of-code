defmodule Day05 do
  def run do
    IO.puts("Day05, part 1: #{part1("test/resources/day05/input")}")
    # IO.puts("Day05, part 2: #{part2("test/resources/day05/input")}")
  end

  def part1(file) do
    parsed_input = parse_input(file)

    Enum.reduce(parsed_input.updates, 0, fn update, acc ->
      if valid_update?(update, parsed_input.rules) do
        acc + String.to_integer(Enum.at(update, round(length(update) / 2) - 1))
      else
        acc
      end
    end)
  end

  def valid_update?(update, rules) do
    update_rules = update_to_rules(update)

    !Enum.any?(rules, fn [a, b] -> [b, a] in update_rules end)
  end

  def update_to_rules(update) do
    for i <- 0..(length(update) - 2),
        j <- (i + 1)..(length(update) - 1),
        do: [Enum.at(update, i), Enum.at(update, j)]
  end

  def parse_input(file) do
    output = %{rules: [], updates: []}

    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> Enum.reduce(output, fn line, acc ->
      cond do
        String.contains?(line, "|") ->
          Map.put(acc, :rules, acc.rules ++ [parse_rule(line)])

        String.contains?(line, ",") ->
          Map.put(acc, :updates, acc.updates ++ [parse_update(line)])

        true ->
          acc
      end
    end)
  end

  def parse_rule(rule) do
    String.split(rule, "|")
  end

  def parse_update(update) do
    String.split(update, ",")
  end
end
