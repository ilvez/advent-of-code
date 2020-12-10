require 'amazing_print'
require 'pry'

def connects?(adapters)
  adapters.chunk_while { _2 - _1 <= 3 }.count == 1
end

adapters_list = IO.readlines('test-input', chomp: true)
                  .map(&:to_i)
                  .sort
