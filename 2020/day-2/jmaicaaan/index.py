from data import getContents

def parse_content(content):
  min_occur, max_occur = content.split(" ")[0].split("-")
  checker_char = content.split(" ")[1].split(":")[0]
  txt = content.split(" ")[2].split(":")[0]
  return [int(min_occur), int(max_occur), checker_char, txt]

# part 1

def get_part_1():
  def is_valid_password_part_1(min_occur, max_occur, checker_char):
    def validator(txt):
      counter = 0
      for char in txt:
        if char == checker_char:
          counter += 1
      
      is_between = min_occur <= counter <= max_occur
      if is_between:
        return True
      return False
    return validator

  def handle_filter(content):  
    min_occur, max_occur, checker_char, txt = parse_content(content)
    is_valid = is_valid_password_part_1(min_occur, max_occur, checker_char)(txt)
    return is_valid

  valid_passwords = filter(handle_filter, getContents().copy())
  num_of_valid_passwords = len(list(valid_passwords))

  print("part 1 =", num_of_valid_passwords)

def get_part_2():
  def is_valid_password_part_1(min_occur, max_occur, checker_char):
    def validator(txt):
      first_position_char = txt[min_occur - 1]
      second_position_char = txt[max_occur - 1]

      is_first_position_valid = first_position_char == checker_char
      is_second_position_valid = second_position_char == checker_char

      if is_first_position_valid and is_second_position_valid:
        return False
      if is_first_position_valid and not is_second_position_valid:
        return True
      if is_second_position_valid and not is_first_position_valid:
        return True
      return False
   
    return validator

  def handle_filter(content):  
    min_occur, max_occur, checker_char, txt = parse_content(content)
    is_valid = is_valid_password_part_1(min_occur, max_occur, checker_char)(txt)
    return is_valid

  valid_passwords = filter(handle_filter, getContents().copy())
  num_of_valid_passwords = len(list(valid_passwords))
  print("part 2 =", num_of_valid_passwords)

get_part_1()
get_part_2()
