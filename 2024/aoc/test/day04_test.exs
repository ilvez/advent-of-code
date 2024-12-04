defmodule Day04Test do
  use ExUnit.Case
  doctest Day04

  test "part 1" do
    assert Day04.count_xmas(["S", "A", "M", "X", "A", "X", "M", "A", "S", "X"]) == 2
    assert Day04.count_xmas(["S", "A", "M", "X"]) == 1
    assert Day04.count_xmas(["X", "M", "A"]) == 0

    test_file = "test/resources/day04/test_input"
    lines = Day04.lines(test_file)

    assert lines == [
             ["M", "M", "M", "S", "X", "X", "M", "A", "S", "M"],
             ["M", "S", "A", "M", "X", "M", "S", "M", "S", "A"],
             ["A", "M", "X", "S", "X", "M", "A", "A", "M", "M"],
             ["M", "S", "A", "M", "A", "S", "M", "S", "M", "X"],
             ["X", "M", "A", "S", "A", "M", "X", "A", "M", "M"],
             ["X", "X", "A", "M", "M", "X", "X", "A", "M", "A"],
             ["S", "M", "S", "M", "S", "A", "S", "X", "S", "S"],
             ["S", "A", "X", "A", "M", "A", "S", "A", "A", "A"],
             ["M", "A", "M", "M", "M", "X", "M", "M", "M", "M"],
             ["M", "X", "M", "X", "A", "X", "M", "A", "S", "X"]
           ]

    vertical_lines = Day04.vertical_lines(lines)

    assert vertical_lines == [
             ["M", "M", "A", "M", "X", "X", "S", "S", "M", "M"],
             ["M", "S", "M", "S", "M", "X", "M", "A", "A", "X"],
             ["M", "A", "X", "A", "A", "A", "S", "X", "M", "M"],
             ["S", "M", "S", "M", "S", "M", "M", "A", "M", "X"],
             ["X", "X", "X", "A", "A", "M", "S", "M", "M", "A"],
             ["X", "M", "M", "S", "M", "X", "A", "A", "X", "X"],
             ["M", "S", "A", "M", "X", "X", "S", "S", "M", "M"],
             ["A", "M", "A", "S", "A", "A", "X", "A", "M", "A"],
             ["S", "S", "M", "M", "M", "M", "S", "A", "M", "S"],
             ["M", "A", "M", "X", "M", "A", "S", "A", "M", "X"]
           ]

    assert Day04.diagonals_from(lines, 0, 2) == ["M", "M", "X", "S", "X", "A", "S", "A"]
    assert Day04.diagonals_from(lines, 2, 0) == ["A", "S", "A", "M", "S", "A", "M", "A"]

    assert Day04.reverse_lines(lines) == [
             ["M", "S", "A", "M", "X", "X", "S", "M", "M", "M"],
             ["A", "S", "M", "S", "M", "X", "M", "A", "S", "M"],
             ["M", "M", "A", "A", "M", "X", "S", "X", "M", "A"],
             ["X", "M", "S", "M", "S", "A", "M", "A", "S", "M"],
             ["M", "M", "A", "X", "M", "A", "S", "A", "M", "X"],
             ["A", "M", "A", "X", "X", "M", "M", "A", "X", "X"],
             ["S", "S", "X", "S", "A", "S", "M", "S", "M", "S"],
             ["A", "A", "A", "S", "A", "M", "A", "X", "A", "S"],
             ["M", "M", "M", "M", "X", "M", "M", "M", "A", "M"],
             ["X", "S", "A", "M", "X", "A", "X", "M", "X", "M"]
           ]

    assert Day04.diagonal_lines(lines) == [
             ["M", "S", "X", "M", "A", "X", "S", "A", "M", "X"],
             ["M", "A", "S", "A", "M", "X", "X", "A", "M"],
             ["M", "M", "X", "S", "X", "A", "S", "A"],
             ["S", "X", "M", "M", "A", "M", "S"],
             ["X", "M", "A", "S", "M", "A"],
             ["X", "S", "A", "M", "M"],
             ["M", "M", "M", "X"],
             ["M", "M", "A", "S", "M", "A", "S", "M", "S"],
             ["A", "S", "A", "M", "S", "A", "M", "A"],
             ["M", "M", "A", "M", "M", "X", "M"],
             ["X", "X", "S", "A", "M", "X"],
             ["X", "M", "X", "M", "A"],
             ["S", "A", "M", "X"]
           ]

    assert Day04.part1(test_file) == 18
  end
end
