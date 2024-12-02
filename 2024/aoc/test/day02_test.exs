defmodule Day02Test do
  use ExUnit.Case
  doctest Day02

  test "part 1" do
    test_file = "test/resources/day02/test_input"
    lines = Day02.lines(test_file)
    levels = Day02.levels(lines)

    assert levels == [
             [7, 6, 4, 2, 1],
             [1, 2, 7, 8, 9],
             [9, 7, 6, 2, 1],
             [1, 3, 2, 4, 5],
             [8, 6, 4, 4, 1],
             [1, 3, 6, 7, 9]
           ]

    assert Day02.direction([1, 2]) == :ascending
    assert Day02.direction([2, 1]) == :descending
    assert Day02.direction([1, 1]) == :same

    assert Day02.find_direction([[1, 2], [2, 3], [3, 4]]) == :ascending
    assert Day02.find_direction([[4, 3], [3, 2], [2, 1]]) == :descending
    assert Day02.find_direction([[1, 1], [1, 1], [1, 1]]) == :same

    assert Day02.safe?([7, 6, 4, 2, 1]) == true

    # big jump
    assert Day02.safe?([1, 2, 7, 8, 9]) == false

    # not ascending
    assert Day02.safe?([1, 3, 2, 4, 5]) == false

    # not descending
    assert Day02.safe?([9, 7, 8, 6, 5]) == false
    assert Day02.part1(test_file) == 2
  end

  test "part 2" do
    assert Day02.reject([1, 2, 3], 1) == [1, 3]

    assert Day02.levels_with_one_missing([1, 2, 3, 4]) == [
             [2, 3, 4],
             [1, 3, 4],
             [1, 2, 4],
             [1, 2, 3]
           ]

    assert Day02.safe_one_removed?([7, 6, 4, 2, 1]) == true
    assert Day02.safe_one_removed?([1, 2, 7, 8, 9]) == false
    assert Day02.safe_one_removed?([9, 7, 6, 2, 1]) == false
    assert Day02.safe_one_removed?([1, 3, 2, 4, 5]) == true
    assert Day02.safe_one_removed?([8, 6, 4, 4, 1]) == true
    assert Day02.safe_one_removed?([1, 3, 6, 7, 9]) == true

    test_file = "test/resources/day02/test_input"
    assert Day02.part2(test_file) == 4
  end
end
