require 'amazing_print'
require 'pry'

def parse_input
  unparsed_bus_ids = '13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23'

  {
    timestamp: 1008713,
    bus_ids: unparsed_bus_ids.split(',').select { _1 != 'x' }.map { Integer(_1) }.sort
  }
end

def next_departure(timestamp, bus_id)
  bus_id - (timestamp % bus_id)
end

input = parse_input
timestamp = input[:timestamp]
closest_bus_id = input[:bus_ids].min { |a, b| next_departure(timestamp, a) <=> next_departure(timestamp, b) }
next_departure = next_departure(timestamp, closest_bus_id)
ap(closest_bus_id * next_departure)
