defmodule Day04 do
  def run do
    IO.puts("Day04, part 1: #{part1("test/resources/day04/input")}")
    # IO.puts("Day04, part 2: #{part2("test/resources/day04/input")}")
  end

  def part1(file) do
    lines = lines(file)
    vertical_lines = lines |> vertical_lines()

    (lines ++ vertical_lines ++ diagonal_lines(lines) ++ diagonal_lines(reverse_lines(lines)))
    |> Enum.map(&count_xmas/1)
    |> Enum.sum()
  end

  def vertical_lines(lines) do
    lines
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
  end

  def reverse_lines(lines) do
    lines |> Enum.map(&Enum.reverse/1)
  end

  def diagonal_lines(lines) do
    rows = length(lines)
    cols = length(Enum.at(lines, 0))
    col_diags = for col <- 0..(cols - 1), do: diagonals_from(lines, 0, col)
    row_diags = for row <- 1..(rows - 1), do: diagonals_from(lines, row, 0)

    (col_diags ++ row_diags) |> Enum.filter(&(length(&1) >= 4))
  end

  def diagonals_from(lines, start_row, start_col) do
    Enum.reduce_while(0..length(lines), [], fn offset, acc ->
      current_row = start_row + offset
      current_col = start_col + offset

      if current_row < length(lines) && current_col < length(Enum.at(lines, current_row, [])) do
        {:cont, acc ++ [Enum.at(Enum.at(lines, current_row), current_col)]}
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
