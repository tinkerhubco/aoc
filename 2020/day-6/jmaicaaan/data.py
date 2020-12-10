import os

filename = "data.txt"
cwd_path = os.path.dirname(__file__)
abs_path = os.path.join(cwd_path, filename)


def get_contents():
  contents = []
  with open(abs_path, 'r') as f:
      contents = [line.rstrip() for line in f]

  return contents