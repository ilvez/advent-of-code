require 'amazing_print'

def parse_answer_groups
  answer_groups = []
  group_answers = []
  IO.readlines('input', chomp: true).each do |line|
    if line.strip.empty?
      answer_groups << group_answers
      group_answers = []
    else
      group_answers << line
    end
  end
  answer_groups << group_answers # :KLUDGE: Because line.strip.empty? wont include last entry
  answer_groups
end

def find_common_answers(answer_group)
  all_answers = answer_group.join.scan(/\w/).sort.uniq
  all_answers.select { |answer| answer_group.all? { |specific_answer| specific_answer.include?(answer) } }
end

answer_groups = parse_answer_groups
ap(answer_groups.map { find_common_answers(_1).size }.sum)
