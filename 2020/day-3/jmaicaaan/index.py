from data import getContents

import functools

def is_tree(char):
  return char == '#'

def is_open_square(char): 
  return char == "."

def nonce_fn(args):
  pass

def get_part_1():
  def peek_location(
    initial_start_position,
    initial_moves,
    initial_text,
    peeker_fn
  ):
    moves_left = initial_moves
    char_position = initial_start_position
    text = initial_text

    if char_position == len(text):
      peek_location(0, moves_left, text + initial_text, peeker_fn)
      return

    if char_position > len(text):
      peek_location(char_position, moves_left, text + initial_text, peeker_fn)
      return

    # print(text[char_position])
    # share the value to the peeker
    # peeker_fn(text[char_position])

    if moves_left == 0:
      if peeker_fn:
        # print('calling peeker_fn...')
        peeker_fn(text[char_position])
        # print('done calling peeker_fn...')
      return

    moves_left -= 1
    char_position += 1

    peek_location(char_position, moves_left, text, peeker_fn)
    
  # peek_location(9, 3, ".4...##..#.")

  # not sure why local variable don't work but this work...
  count = {
    'val': 0,
  }
  coords = {
    'x': 0,
    'y': 1,
    'default_x': 3,
    'default_y': 1,
  }

  map = getContents()

  def peeker(val):
    if is_tree(val):
      count['val'] += 1
    return

  def recurse():
    if coords['y'] == len(map):
      print('part 1 =', count['val'])
      return

    path = map[coords['y']]
    peek_location(
      coords['x'],
      coords['default_x'],
      path,
      peeker,
    )

    coords['x'] += coords['default_x']
    coords['y'] += coords['default_y']

    recurse()

  recurse()

def get_part_2():
  # not sure why local variable don't work but this work...
  count = {
    'val': 0,
    'product': 0,
  }
  list_counter = {
    'val': 0,
  }

  coords_list = [
    {
      'x': 0,
      'y': 1,
      'step_x': 1,
      'step_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'step_x': 3,
      'step_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'step_x': 5,
      'step_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'step_x': 7,
      'step_y': 1,
    },
    {
      'x': 0,
      'y': 2,
      'step_x': 1,
      'step_y': 2,
    },
  ]

  play_map = getContents()

  def peeker(val):
    if is_tree(val):
      count['val'] += 1
    return

  def handle_count_product():
    if count['product'] == 0:
      count['product'] = count['val']
      return
    count['product'] *= count['val']
    
  def handle_count_reset():
    count['val'] = 0
    return

  def handle_counter_increment():
    list_counter['val'] += 1
    return

  def peek_location(position, number_of_step, input, peeker_fn):
    final_position = (position + number_of_step) % len(input)
    peeker_fn(input[final_position])
    return

  def recurse(slope_coordinates):
    slope_counter = list_counter['val']
    has_no_slopes_to_check = slope_counter > len(coords_list) -1
    is_on_last_slope_path_row = slope_coordinates['y'] > len(play_map) -1

    if is_on_last_slope_path_row:
      count_copy = count.copy()
      handle_count_reset()
      return count_copy

    slope_path_row = play_map[slope_coordinates['y']]

    peek_location(
      slope_coordinates['x'],
      slope_coordinates['step_x'],
      slope_path_row,
      peeker,
    )

    slope_coordinates['x'] += slope_coordinates['step_x']
    slope_coordinates['y'] += slope_coordinates['step_y']

    return recurse(slope_coordinates)

  values = map(lambda item_dict: item_dict['val'], [
    recurse(coords_list[0]),
    recurse(coords_list[1]),
    recurse(coords_list[2]),
    recurse(coords_list[3]),
    recurse(coords_list[4])
  ])
  product = functools.reduce(
    lambda accumulator, current_value: accumulator * current_value,
    values
  )
  print('part 1 =', product)

# get_part_1()
get_part_2()
