defmodule Day01Test do
  use ExUnit.Case
  doctest Day01

  test "part 1" do
    test_input = "test/resources/day01/test_input"
    lines = Day01.lines(test_input)
    assert lines == ["1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"]
    assert Day01.numbers1(lines) == [12, 38, 15, 77]
    part1 = Day01.part1(test_input)
    assert part1 == 142
  end

  test "part 2" do
    test_input = "test/resources/day01/test_input2"
    lines = Day01.lines(test_input)

    assert lines == [
             "two1nine",
             "eightwothree",
             "abcone2threexyz",
             "xtwone3four",
             "4nineeightseven2",
             "zoneight234",
             "7pqrstsixteen"
           ]

    assert Day01.word_replace("two1nine") == "219"
    assert Day01.word_replace("eightwothree") == "823"
    assert Day01.word_replace("abcone2threexyz") == "123"
    assert Day01.word_replace("eightnineseventwo1seven") == "897217"
    assert Day01.word_replace("9h1xcrcggtwo38") == "91238"
    assert Day01.word_replace("bnjqlftwobvsvjqptdp1two94twonej") == "2129421"

    # IO.puts(Day01.numbers2(lines))
    assert Day01.numbers2(lines) == [29, 83, 13, 24, 42, 14, 76]
    assert Day01.part2(test_input) == 281
  end
end
