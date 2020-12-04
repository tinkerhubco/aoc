from data import getContents

import sys

sys.setrecursionlimit(10**6) 

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
    'product': 0,
  }
  list_counter = {
    'val': 0,
  }

  coords_list = [
    {
      'x': 0,
      'y': 1,
      'default_x': 1,
      'default_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'default_x': 3,
      'default_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'default_x': 5,
      'default_y': 1,
    },
    {
      'x': 0,
      'y': 1,
      'default_x': 7,
      'default_y': 1,
    },
    {
      'x': 0,
      'y': 2,
      'default_x': 1,
      'default_y': 2,
    },
  ]

  map = getContents()

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

  def recurse():
    counter = list_counter['val']

    if counter == len(coords_list):
      print('done')
      print('part 2 =', count['product'])
      return

    coord = coords_list[counter]

    if coord['y'] > len(map):
      handle_count_product()
      handle_counter_increment()
      handle_count_reset()

      recurse()
      return


    if coord['y'] == len(map):

      handle_count_product()
      handle_counter_increment()
      handle_count_reset()

      
      recurse()
      return

    path = map[coord['y']]
    peek_location(
      coord['x'],
      coord['default_x'],
      path,
      peeker,
    )

    coord['x'] += coord['default_x']
    coord['y'] += coord['default_y']

    recurse()

  recurse()

# get_part_1()
get_part_2()
