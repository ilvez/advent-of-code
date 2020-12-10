require 'amazing_print'

answer_groups = []
group_string = ''
IO.readlines('input', chomp: true).each do |line|
  if line.strip.empty?
    answer_groups << group_string
    group_string = ''
  else
    group_string.concat(line)
  end
end
answer_groups << group_string # :KLUDGE: Because line.strip.empty? wont include last entry

ap(answer_groups.map { _1.scan(/\w/).sort.uniq.size }.sum)
