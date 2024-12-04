defmodule Day04 do
  def run do
    IO.puts("Day04, part 1: #{part1("test/resources/day04/input")}")
    IO.puts("Day04, part 2: #{part2("test/resources/day04/input")}")
  end

  def part1(file) do
    lines = lines(file)
    vertical_lines = lines |> vertical_lines()

    (lines ++ vertical_lines ++ diagonal_lines(lines) ++ diagonal_lines(reverse_lines(lines)))
    |> Enum.map(&count_xmas/1)
    |> Enum.sum()
  end

  def part2(file) do
    lines(file)
    |> indexed_lines()
    |> find_mas()
  end

  def indexed_lines(lines) do
    diagonal_lines(lines, 3, true) ++ diagonal_lines(reverse_lines(lines), 3, true, true)
  end

  def find_mas(indexed_lines) do
    indexed_lines
    |> Enum.map(fn line -> Enum.chunk_every(line, 3, 1, :discard) end)
    |> Enum.map(fn line ->
      Enum.filter(line, fn [[a, _], [b, _], [c, _]] ->
        (a == "M" && b == "A" && c == "S") || (a == "S" && b == "A" && c == "M")
      end)
    end)
    |> Enum.filter(fn mas -> length(mas) > 0 end)
    |> Enum.map(fn mas ->
      Enum.map(mas, fn submas -> Enum.filter(submas, fn sss -> Enum.at(sss, 0) == "A" end) end)
    end)
    |> flatten()
    |> Enum.filter(fn a -> is_tuple(a) end)
    |> Enum.group_by(fn {a, b} -> {a, b} end)
    |> Map.values()
    |> Enum.map(&Enum.count/1)
    |> Enum.filter(fn count -> count == 2 end)
    |> Enum.count()
  end

  def flatten(list) do
    Enum.flat_map(list, fn
      sublist when is_list(sublist) -> flatten(sublist)
      element -> [element]
    end)
  end

  def vertical_lines(lines) do
    lines
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
  end

  def reverse_lines(lines) do
    lines |> Enum.map(&Enum.reverse/1)
  end

  def diagonal_lines(lines, diagonal_size \\ 4, locations \\ false, interpolate_pos \\ false) do
    rows = length(lines)
    cols = length(Enum.at(lines, 0))

    col_diags =
      for col <- 0..(cols - 1), do: diagonals_from(lines, 0, col, locations, interpolate_pos)

    row_diags =
      for row <- 1..(rows - 1), do: diagonals_from(lines, row, 0, locations, interpolate_pos)

    (col_diags ++ row_diags) |> Enum.filter(&(length(&1) >= diagonal_size))
  end

  def diagonals_from(lines, start_row, start_col, locations \\ false, interpolate_pos \\ false) do
    row_count = length(lines)

    Enum.reduce_while(0..row_count, [], fn offset, acc ->
      current_row = start_row + offset
      current_col = start_col + offset

      col_count = length(Enum.at(lines, current_row, []))

      if current_row < length(lines) && current_col < col_count do
        char = Enum.at(Enum.at(lines, current_row), current_col)

        if locations do
          if interpolate_pos do
            {:cont, acc ++ [[char, {current_row, col_count - current_col - 1}]]}
          else
            {:cont, acc ++ [[char, {current_row, current_col}]]}
          end
        else
          {:cont, acc ++ [char]}
        end
      else
        {:halt, acc}
      end
    end)
  end

  def count_xmas(line) do
    forward = line
    backward_chunked = Enum.reverse(forward) |> Enum.chunk_every(4, 1, :discard)

    combined_chunked =
      forward
      |> Enum.chunk_every(4, 1, :discard)
      |> Enum.concat(backward_chunked)

    Enum.count(combined_chunked, fn chunk -> chunk == ["X", "M", "A", "S"] end)
  end

  def lines(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Enum.to_list()
    |> Enum.map(&String.codepoints/1)
  end
end
