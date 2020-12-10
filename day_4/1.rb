require 'amazing_print'

MANDATORY_FIELDS = %w[ecl pid eyr hcl byr iyr hgt].freeze
ALL_FIELDS = (MANDATORY_FIELDS + ['cid']).freeze

def full_passport?(passport)
  (ALL_FIELDS - passport.keys).empty?
end

def nop_passport?(passport)
  full_passport?(passport) || (MANDATORY_FIELDS - passport.keys).empty?
end

passports = []
passport = {}
IO.readlines('input', chomp: true).each do |line|
  if line.strip.empty?
    passports << passport
    passport = {}
    next
  end
  passport = passport.merge(line.split(' ').map { _1.split(':') }.to_h)
end

ap(passports.select { |p| nop_passport?(p) }.size)
