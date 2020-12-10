from data import get_contents

import functools
import string

def parse_content(reducer_fn):
  data = get_contents()

  def organize_structure_to_2d(sequence):
    grouped = []
    temp_group = []

    # Tried functional approach like reduce however there's no index provided
    # and can't get the last line
    for i, line in enumerate(data):
      is_last_line = i == len(data) -1
      is_line_terminator = line == ""

      if is_line_terminator:
        grouped.append(reducer_fn(temp_group))
        temp_group = []
        continue

      temp_group.append(line)

      if is_last_line:
        grouped.append(reducer_fn(temp_group))
        break        

    return grouped.copy()

  return organize_structure_to_2d(data)

def get_part_1():
  def remap_items(sequence):
    def reducer(accumulator, current_value):
      accumulator += current_value
      return accumulator

    return functools.reduce(
      reducer,
      sequence,
      '',
    )
  data = parse_content(remap_items)

  unique_data = {}
  sum_count = 0

  for index, characters in enumerate(data):
    if index not in unique_data:
      unique_data[index] = []

    for char in characters:
      if char not in unique_data[index]:
        unique_data[index].append(char)
    
   # TODO - use sum() 
  for i in unique_data:
    sum_count += len(unique_data[i])
    
  print('part 1 =', sum_count)

def get_part_2():
  def remap_items(sequence):
    def reducer(accumulator, current_value):
      accumulator = [*accumulator, current_value]
      return accumulator

    return functools.reduce(
      reducer,
      sequence,
      [],
    )
    
  def create_alphabet_counter_dict():
    alphabet_counter_dict = {}
    alphabets = list(string.ascii_lowercase)
    for letter in alphabets:
      alphabet_counter_dict[letter] = 0
    return alphabet_counter_dict

  def get_all_answered_letters(alphabet_counter_dict):
    return list(filter(
      lambda val: alphabet_counter_dict[val] > 0,
      alphabet_counter_dict
    ))

  sum_count = 0
  data = parse_content(remap_items)
  # Tienshiao gave me a new perspective how to handle this and that is through this dict of alphabets <string, int> counter
  # I took a new approach here compared on part 1
  alphabet_counter_dict = create_alphabet_counter_dict()

  for index, answers in enumerate(data):
    total_number_of_people = len(answers)
    for answer in answers:
      for char in answer:
        alphabet_counter_dict[char] += 1

    for i in get_all_answered_letters(alphabet_counter_dict):
      count = alphabet_counter_dict[i]
      if count == 0:
        continue
      if count == total_number_of_people:
        sum_count += 1

    alphabet_counter_dict = create_alphabet_counter_dict()

      
  print('part 2 =', sum_count)
  pass

# get_part_1()
get_part_2()
