filename = 'input.txt'

MAX_CUBES_RED = 12
MAX_CUBES_GREEN = 13
MAX_CUBES_BLUE = 14

with open(filename) as file:
  lines = [line.rstrip() for line in file]

possible_game_ids = []

for line in lines:
  [gameStr, setStr] = line.split(":")
  gameId = gameStr.replace("Game ", "")
  gameSets = setStr.split(";")
  gamePossible = True

  for gameSet in gameSets:
    gameSetColors = gameSet.split(",")
    for gameSetColor in gameSetColors:
      trimmedColor = gameSetColor.strip()
      [amount, color] = trimmedColor.split(" ")
      
      if color == 'red' and int(amount) > MAX_CUBES_RED:
        gamePossible = False
      elif color == 'green' and int(amount) > MAX_CUBES_GREEN:
        gamePossible = False
      elif color == 'blue' and int(amount) > MAX_CUBES_BLUE:
        gamePossible = False

  if gamePossible is True:
    possible_game_ids.append(int(gameId))

print(sum(possible_game_ids))