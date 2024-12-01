defmodule Day01Test do
  use ExUnit.Case
  doctest Day01

  test "part 1" do
    test_file = "test/resources/day01/test_input"
    lines = Day01.lines(test_file)
    assert lines == ["3   4", "4   3", "2   5", "1   3", "3   9", "3   3"]

    pairs = Day01.pairs(lines)
    assert pairs == [[3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]]
    sorted = Day01.sorted(pairs)

    assert sorted == [
             [
               1,
               2,
               3,
               3,
               3,
               4
             ],
             [
               3,
               3,
               3,
               4,
               5,
               9
             ]
           ]

    assert Day01.part1(test_file) == 11
  end
end
