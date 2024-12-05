defmodule Day05Test do
  use ExUnit.Case
  doctest Day05

  test "part 1" do
    test_file = "test/resources/day05/test_input"
    lines = Day05.parse_input(test_file)

    assert lines == %{
             rules: [
               ["47", "53"],
               ["97", "13"],
               ["97", "61"],
               ["97", "47"],
               ["75", "29"],
               ["61", "13"],
               ["75", "53"],
               ["29", "13"],
               ["97", "29"],
               ["53", "29"],
               ["61", "53"],
               ["97", "53"],
               ["61", "29"],
               ["47", "13"],
               ["75", "47"],
               ["97", "75"],
               ["47", "61"],
               ["75", "61"],
               ["47", "29"],
               ["75", "13"],
               ["53", "13"]
             ],
             updates: [
               ["75", "47", "61", "53", "29"],
               ["97", "61", "53", "29", "13"],
               ["75", "29", "13"],
               ["75", "97", "47", "61", "53"],
               ["61", "13", "29"],
               ["97", "13", "75", "29", "47"]
             ]
           }

    assert Day05.update_to_rules(["11", "22", "33"]) == [["11", "22"], ["11", "33"], ["22", "33"]]
    assert Day05.valid_update?(["11", "22", "33"], [["11", "33"], ["22", "33"]]) == true
    assert Day05.valid_update?(["11", "22", "33"], [["11", "33"], ["33", "22"]]) == false

    assert Day05.valid_update?(["75", "47", "61", "53", "29"], lines.rules) == true
    assert Day05.valid_update?(["97", "61", "53", "29", "13"], lines.rules) == true
    assert Day05.valid_update?(["75", "29", "13"], lines.rules) == true
    assert Day05.valid_update?(["75", "97", "47", "61", "53"], lines.rules) == false
    assert Day05.valid_update?(["61", "13", "29"], lines.rules) == false
    assert Day05.valid_update?(["97", "13", "75", "29", "47"], lines.rules) == false
    assert Day05.part1(test_file) == 143
  end
end
