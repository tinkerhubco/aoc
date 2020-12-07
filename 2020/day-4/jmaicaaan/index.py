from data import getContents

import functools
import re

def range_validator(min_value, max_value, inclusive = True):
  def validator(value):
    range_min = min_value - 1 if inclusive else min_value
    range_max = max_value + 1 if inclusive else max_value
    return int(value) in range(range_min, range_max)
  return validator

def height_validator(value):
  is_cm_unit = re.search('^[\dcm]*$', value)
  is_inches_unit = re.search('^[\din]*$', value)
  has_metric_unit = is_cm_unit or is_inches_unit

  if has_metric_unit == False:
    return False

  value_without_metric_unit = re.search('^[\d]*', value).group()

  if is_inches_unit:
    return range_validator(59, 76)(value_without_metric_unit)
  return range_validator(150, 193)(value_without_metric_unit)

def hair_color_validator(value):
  res = re.search('^#[0-9a-f]*$', value)
  if res is None:
    return False
  return True

def eye_color_validator(value):
  return value in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']

def passport_validator(value):
  return len(value) == 9 and value.isnumeric()


# Validator function runs after it passes the `required` check
form_field = {
  'BirthYear': {
    'key': 'byr',
    'required': True,
    'validator': range_validator(1920, 2002)
  },
  'IssueYear': {
    'key': 'iyr',
    'required': True,
    'validator': range_validator(2010, 2020)
  },
  'ExpirationYear': {
    'key': 'eyr',
    'required': True,
    'validator': range_validator(2020, 2030)
  },
  'Height': {
    'key': 'hgt',
    'required': True,
    'validator': height_validator,
  },
  'HairColor': {
    'key': 'hcl',
    'required': True,
    'validator': hair_color_validator,
  },
  'EyeColor': {
    'key': 'ecl',
    'required': True,
    'validator': eye_color_validator,
  },
  'PassportId': {
    'key': 'pid',
    'required': True,
    'validator': passport_validator,
  },
  'CountryId': {
    'key': 'cid',
    'required': False,
  },
}

# Transform it to array of dicts
def get_normalized_form_field():
  return list(
    map(
      lambda val: form_field[val],
      form_field
    )
  )

def get_required_form_field_keys():
  normalized_form_field = get_normalized_form_field()
  return list(
    map(
      lambda val: val['key'],
      filter(
        lambda val: val['required'] == True,
        normalized_form_field,
      )
    )
    
  )

def get_optional_form_field_keys():
  normalized_form_field = get_normalized_form_field()
  return list(
    map(
      lambda val: val['key'],
      filter(
        lambda val: val['required'] == False,
        normalized_form_field,
      )
    )
  )
def get_field_by_key(key):
  normalized_form_field = get_normalized_form_field()

  for f in normalized_form_field:
    if f['key'] == key:
      return f

  return None

def parse_content():
  data = getContents()

  def remap_items(sequence):
    def reducer(accumulator, current_value):
      accumulator = [*accumulator, *current_value.split(" ")]
      return accumulator

    return functools.reduce(
      reducer,
      sequence,
      []
    )

  def organize_structure_to_2d(sequence):
    grouped = []
    temp_group = []

    # Tried functional approach like reduce however there's no index provided
    # and can't get the last line
    for i, line in enumerate(data):
      is_last_line = i == len(data) -1
      is_line_terminator = line == ""

      if is_line_terminator:
        grouped.append(remap_items(temp_group))
        temp_group = []
        continue

      temp_group.append(line)

      if is_last_line:
        grouped.append(remap_items(temp_group))
        break        

    return grouped.copy()

  return organize_structure_to_2d(data)

def get_part_1():
  def get_valid_passports(sequence):
    required_form_field_keys = get_required_form_field_keys()

    def get_field_key(value):
      key, _ = value.split(":")
      return key

    def filterer(val):
      # Check if required
      field_keys = list(map(get_field_key, val))
      required_result = [required_key in field_keys for required_key in required_form_field_keys]
      has_pass_required_requirement = all(required_result)
      return has_pass_required_requirement

    filtered_sequence = list(filter(filterer, sequence))
    return filtered_sequence

  
  passports = parse_content()
  valid_passports = get_valid_passports(passports)

  print('part 1 =', len(valid_passports))
  pass


def get_part_2():
  def get_valid_passports(sequence):
    required_form_field_keys = get_required_form_field_keys()

    def get_field_key(value):
      key, _ = value.split(":")
      return key

    def get_field_value(value):
      _, field_value = value.split(":")
      return field_value

    def filterer(val):
      # Check if required
      field_keys = list(map(get_field_key, val))
      required_result = [required_key in field_keys for required_key in required_form_field_keys]
      has_pass_required_requirement = all(required_result)

      if has_pass_required_requirement:
        
        def handle_reducer(accumulator, current_value):
          key, value = current_value.split(":")
          field_key = get_field_by_key(key)

          if field_key['required'] == False:
            # Skip
            accumulator.append(True)
            return accumulator
          
          if 'validator' in field_key:
            validator_func = field_key['validator']
            res = validator_func(value)
            accumulator.append(res)

          return accumulator

        is_valid = all(functools.reduce(
          handle_reducer,
          val,
          [],
        ))

        return is_valid

      return has_pass_required_requirement

    filtered_sequence = list(filter(filterer, sequence))
    return filtered_sequence

  
  passports = parse_content()
  valid_passports = get_valid_passports(passports)

  # the correct answer is 103
  # my answer is 104. After submitting 104, it says the answer is too high so I tried going down
  # still figuring out why 103
  print('part 2 =', len(valid_passports))
  pass

get_part_1()
get_part_2()
