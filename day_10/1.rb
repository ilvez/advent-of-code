require 'amazing_print'
require 'pry'

adapters_list = IO.readlines('input', chomp: true).map(&:to_i).sort
parsed_output = {
  last_adapter: 0
}
adapters_list.each do |adapter|
  diff = adapter - parsed_output[:last_adapter]
  parsed_output = parsed_output.merge(
    "#{diff}": parsed_output.fetch(:"#{diff}", 0) + 1,
    last_adapter: adapter
  )
end

parsed_output[:'3'] += 1 # Finally, your device's built-in adapter is always 3 higher than the highest adapter
ap(parsed_output[:'1'] * parsed_output[:'3'])
