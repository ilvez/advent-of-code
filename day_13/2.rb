require 'amazing_print'
require 'pry'

def parse_input
  unparsed_bus_ids = '13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23'

  {
    timestamp: 1008713,
    bus_ids: unparsed_bus_ids.split(',').select { _1 != 'x' }.map { Integer(_1) }.sort
    target_pattern: unparsed_bus_ids
  }
end

def next_departure(timestamp, bus_id)
  bus_id - (timestamp % bus_id)
end

tests = [
  { pattern: '17,x,13,19', result: 3417 },
  { pattern: '67,7,59,61', result: 754018 },
  { pattern: '67,x,7,59,61', result: 779210 },
  { pattern: '67,7,x,59,61', result: 1261476 },
  { pattern: '1789,37,47,1889', result: 1202161486 },
]
