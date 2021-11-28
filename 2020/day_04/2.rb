require 'amazing_print'

class Passport
  attr_reader :values

  def initialize(unparsed_input)
    @values = unparsed_input.split(' ')
                            .map { _1.split(':') }
                            .to_h
                            .transform_keys(&:to_sym)
  end

  def valid?
    nop_passport? && all_fields_valid?
  end

  private

  def nop_passport?
    (FIELD_VALIDATIONS.keys - values.keys).empty?
  end

  def all_fields_valid?
    FIELD_VALIDATIONS.all? { |key, validation| validation.call(@values.fetch(key)) }
  end

  EYE_COLORS = %w[amb blu brn gry grn hzl oth].freeze
  FIELD_VALIDATIONS = {
    byr: ->(value) { value.to_i >= 1920 && value.to_i <= 2002 }, # Birth year
    iyr: ->(value) { value.to_i >= 2010 && value.to_i <= 2020 }, # Issue year
    eyr: ->(value) { value.to_i >= 2020 && value.to_i <= 2030 }, # Expiration Year
    hgt: ->(value) { height_validation(value) },
    hcl: ->(value) { value.match?(/^#[a-f 0-9]{6}$/) },
    pid: ->(value) { value.match?(/^[0-9]{9}$/) },
    ecl: ->(value) { EYE_COLORS.include?(value) }
  }.freeze

  def self.height_validation(value)
    (
      value.match(/cm/) &&
        (
          cm = value.gsub(/(\d.*)cm/, '\1').to_i
          cm >= 150 && cm <= 193
        )
    ) || (
      value.match(/in/) &&
      (
        inch = value.gsub(/(\d.*)in/, '\1').to_i
        inch >= 59 && inch <= 76
      )
    )
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

ap(passports.select(&:valid?).size)
