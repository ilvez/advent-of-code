defmodule Day01Test do
  use ExUnit.Case
  doctest Day01

  test "gives lines" do
    test_input = "test/resources/day01/test_input"
    lines = Day01.lines(test_input)
    assert lines == ["1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"]
    assert Day01.numbers(lines) == [12, 38, 15, 77]
    assert Day01.part1(test_input) == 142
    IO.puts("Day01 part1: #{Day01.part1("test/resources/day01/input")}")
  end
end
