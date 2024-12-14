defmodule Day11Test do
  use ExUnit.Case
  doctest Day11

  test "part 1 simple blink" do
    simple_example = "0 1 10 99 999"
    assert Day11.parse_stones(simple_example) == [0, 1, 10, 99, 999]
    stones = Day11.parse_stones(simple_example)
    assert Day11.blink(stones) == [1, 2024, 1, 0, 9, 9, 2_021_976]
  end

  test "part 1 blink n times" do
    stones = [125, 17]

    assert Day11.blink_n_times(stones, 6) == [
             2_097_446_912,
             14168,
             4048,
             2,
             0,
             2,
             4,
             40,
             48,
             2024,
             40,
             48,
             80,
             96,
             2,
             8,
             6,
             7,
             6,
             0,
             3,
             2
           ]
  end

  test "part 1" do
    test_file = "test/resources/day11/test_input"
    assert Day11.part1(test_file) == 55312
  end

  test "blink_count" do
    # assert Day11.blink_count(0, 0) == 1 # 0 -> 1 -> count 1
    # assert Day11.blink_count(0, 1) == 1 # 0 -> 1 -> 2024 -> count 1
    # assert Day11.blink_count(0, 2) == 2 # 0 -> 1 -> 2024 -> 20, 24 -> count 2
    # assert Day11.blink_count(0, 3) == 4 # 0 -> 1 -> 2024 -> 20, 24 -> 2, 0, 2, 4 -> count 4
    # assert Day11.blink_stone(1000) == [10, 0]
    # IO.puts(Day11.blink_count(0, 0))
    # IO.puts(Day11.blink_count(0, 1))
    # IO.puts(Day11.blink_count(0, 2))
    # IO.puts(Day11.blink_count(0, 3))
    # IO.puts(Day11.blink_count(0, 4))
    # IO.puts(Day11.blink_count(0, 5))
    # IO.puts(Day11.blink_count(0, 6))
    # IO.puts(Day11.blink_count(0, 7))
    # IO.puts(Day11.blink_count(0, 8))
    # IO.puts(Day11.blink_count(0, 9))
    # IO.puts(Day11.blink_count(0, 10))
    # IO.puts(Day11.blink_count(0, 11))
    # IO.puts(Day11.blink_count(0, 12))
    # IO.puts(Day11.blink_count(0, 13))
    IO.puts(Day11.blink_count(0, 14))
  end

  # test "part 2" do
  #   test_file = "test/resources/day11/test_input"
  # end
end
