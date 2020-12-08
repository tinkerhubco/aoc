from data import get_contents

import functools

def is_front(value):
  return value == "F"

def is_left(value):
  return value == "L"

def get_char_position(value):
  return 0 if (is_front(value) or is_left(value)) else 1

def compute_seat_id(row, column):
  # 8 = based on the instruction
  return row * 8 + column

def get_part_1():
  rows = 128
  columns = 8

  seat_ids = []
  
  def get_seat(number_of_rows, characters):
    half_row_count = int(number_of_rows / 2)
    ranges = []

    for char in characters:
      if len(ranges) == 0:
        half_row_count = int(number_of_rows / 2)
        ranges = [
          range(0, half_row_count),
          range(half_row_count, number_of_rows)
        ]

      position = get_char_position(char)
      current_range = ranges[position]

      half_row_count = int(len(current_range) / 2)
      
      range_start, range_end = [
        current_range.start,
        current_range.stop
      ]

      ranges = [
        range(range_start, range_start + half_row_count),
        range(range_start + half_row_count, range_end),
      ]

      if half_row_count == 0:
        return ranges[position].start

  for seat in get_contents():
    seat_rows_characters = seat[0:7] # first 7 characters
    seat_columns_characters = seat[7:10] # the last 3 characters

    row = get_seat(rows, seat_rows_characters)
    column = get_seat(columns, seat_columns_characters)
    seat_ids.append(
      compute_seat_id(row, column)
    )
  print('part 1 =', max(seat_ids))
  return seat_ids

def get_part_2(seat_ids = []):
  my_sitting_id = 0
  for i, elem in enumerate(range(min(seat_ids), max(seat_ids))):
    if i < min(seat_ids):
      continue
    if elem not in seat_ids:
      my_sitting_id = elem
      break
  print('part 2 =', my_sitting_id)
  pass

seat_ids = get_part_1()
get_part_2(seat_ids)
