#!/bin/sh

cp -v lib/day01.ex lib/day$1.ex
cp -v test/day01_test.exs test/day$1_test.exs
mkdir -p test/resources/day$1
touch test/resources/day$1/test_input
sed -i '' -e "s/01/$1/g" lib/day$1.ex
sed -i '' -e "s/01/$1/g" test/day$1_test.exs
