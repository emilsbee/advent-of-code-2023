filename = 'input.txt'

with open(filename) as file:
  lines = [line.rstrip() for line in file]

powersOfSets = []

for line in lines:
  [gameStr, setStr] = line.split(":")
  gameId = gameStr.replace("Game ", "")
  gameSets = setStr.split(";")

  highestRed = 0
  highestGreen = 0
  highestBlue = 0

  for gameSet in gameSets:
    gameSetColors = gameSet.split(",")
    for gameSetColor in gameSetColors:
      trimmedColor = gameSetColor.strip()
      [amountStr, color] = trimmedColor.split(" ")
      amountNum = int(amountStr)

      if color == 'red' and amountNum > highestRed:
        highestRed = amountNum
      elif color == 'green' and amountNum > highestGreen:
        highestGreen = amountNum
      elif color == 'blue' and amountNum > highestBlue:
        highestBlue = amountNum
  
  powerOfSet = highestRed * highestGreen * highestBlue
  powersOfSets.append(powerOfSet)

print(sum(powersOfSets))
  
