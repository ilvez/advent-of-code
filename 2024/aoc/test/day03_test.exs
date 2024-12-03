defmodule Day03Test do
  use ExUnit.Case
  doctest Day03

  test "part 1" do
    test_file = "test/resources/day03/test_input"
    memory = Day03.memory(test_file)
    assert memory == "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    assert Day03.multiply(Day03.arguments(memory)) == 161
  end

  test "part 2" do
    memory = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    assert Day03.selected_arguments(memory) == [["2", "4"], ["8", "5"]]
  end
end
