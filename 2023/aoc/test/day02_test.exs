defmodule Day02Test do
  use ExUnit.Case
  doctest Day01

  test "part 1" do
    game = Day02.parse_game("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

    assert game == %{
             game: 1,
             sets: [
               %{blue: 3, red: 4},
               %{red: 1, green: 2, blue: 6},
               %{green: 2}
             ]
           }

    impossible_game =
      Day02.parse_game("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")

    assert impossible_game == %{
             game: 3,
             sets: [
               %{green: 8, blue: 6, red: 20},
               %{red: 4, green: 13, blue: 5},
               %{green: 5, red: 1}
             ]
           }

    impossible_set =
      Day02.find_impossible_set(
        impossible_game.sets,
        %{red: 12, green: 13, blue: 14}
      )

    assert impossible_set == %{green: 8, blue: 6, red: 20}

    lines = Day02.lines("test/resources/day02/test_input")
    games = Day02.parse_games(lines)

    assert games == [
             %{game: 1, sets: [%{blue: 3, red: 4}, %{blue: 6, green: 2, red: 1}, %{green: 2}]},
             %{
               game: 2,
               sets: [%{blue: 1, green: 2}, %{blue: 4, green: 3, red: 1}, %{blue: 1, green: 1}]
             },
             %{
               game: 3,
               sets: [
                 %{blue: 6, green: 8, red: 20},
                 %{blue: 5, green: 13, red: 4},
                 %{green: 5, red: 1}
               ]
             },
             %{
               game: 4,
               sets: [
                 %{blue: 6, green: 1, red: 3},
                 %{green: 3, red: 6},
                 %{blue: 15, green: 3, red: 14}
               ]
             },
             %{game: 5, sets: [%{blue: 1, green: 3, red: 6}, %{blue: 2, green: 2, red: 1}]}
           ]

    possible_games = Day02.find_possible_games(games)
    assert possible_games == [
             %{game: 1, sets: [%{blue: 3, red: 4}, %{blue: 6, green: 2, red: 1}, %{green: 2}]},
             %{
               game: 2,
               sets: [%{blue: 1, green: 2}, %{blue: 4, green: 3, red: 1}, %{blue: 1, green: 1}]
             },
             %{game: 5, sets: [%{blue: 1, green: 3, red: 6}, %{blue: 2, green: 2, red: 1}]}
           ]

    part1 = Day02.part1("test/resources/day02/test_input")
    assert part1 == 8
  end

  test "part 2" do
    # test_input = "test/resources/day02/test_input"
  end
end
