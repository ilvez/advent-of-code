require 'amazing_print'
require 'pry'

def parse_input
  unparsed_target_pattern = '13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23'

  {
    bus_ids: unparsed_bus_ids.split(',').select { _1 != 'x' }.map { Integer(_1) }.sort,
    pattern: unparsed_bus_ids
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

def parse_pattern(pattern)
end

class CommonTimeFinder
  def initialize(pattern)
    offset = 0
    times = {}
    pattern.split(',').each {
      if _1 == 'x'
        offset += 1
        next
      end

      times = times.merge(_1.to_i => offset)
      offset +=1
    }
    @times = times
  end

  def find_common_time
    (1..).each { |i|
      time_to_match = i * highest_bus_id - @times[highest_bus_id]
      return time_to_match if time_suits?(time_to_match)
    }
  end

  def time_suits?(time_to_match)
    @times.all? { |bus_id, offset| (time_to_match + offset) % bus_id == 0 }
  end

  def highest_bus_id
    @highest_bus_id ||= @times.keys.max
  end
end

tests.map {
  finder = CommonTimeFinder.new(_1[:pattern])
  ap(finder.find_common_time == _1[:result])
}

binding.pry
ap(
  CommonTimeFinder
    .new('13,x,x,41,x,x,x,x,x,x,x,x,x,467,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,353,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23')
    .find_common_time
)
