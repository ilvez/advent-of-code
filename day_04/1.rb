require 'amazing_print'

class Passport
  MANDATORY_FIELDS = %i[ecl pid eyr hcl byr iyr hgt].freeze
  ALL_FIELDS = (MANDATORY_FIELDS + [:cid]).freeze

  attr_reader :unparsed_input, :values

  def initialize(unparsed_input)
    @unparsed_input = unparsed_input
    @values = unparsed_input.split(' ')
                            .map { _1.split(':') }
                            .to_h
                            .transform_keys(&:to_sym)
  end

  def full_passport?
    (ALL_FIELDS - values.keys).empty?
  end

  def nop_passport?
    full_passport? || (MANDATORY_FIELDS - values.keys).empty?
  end
end

passports = []
passport_string = ''
IO.readlines('input').each do |line|
  if line.strip.empty?
    passports << Passport.new(passport_string)
    passport_string = ''
  else
    passport_string.concat(' ' + line)
  end
end
passports << Passport.new(passport_string) # :KLUDGE: Because line.strip.empty? wont include last passport

ap(passports.select(&:nop_passport?).size)
