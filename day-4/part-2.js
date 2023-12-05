const fs = require('node:fs');
const readline = require('node:readline');

const getParsedLines = (lines) => {
  const parsedLines = []

  for (const lineIndex in lines) {
    const line = lines[lineIndex]
    const [cardNrStr, gamesStr] = line.split(":")
    const [winningNrsStr, myNrsStr] = gamesStr.split("|")
    const numbers = ['0','1','2','3','4','5','6','7','8','9']
    let cardMatchingNumbers = 0
    // const cardNr = parseInt(cardNrStr.split("Card ")[1])

    // Find winning numbers
    const winningNrs = []
    for (let i = 0; i < winningNrsStr.length; i++) {
      if (numbers.includes(winningNrsStr[i])) {
        beginIndex = i;
        endIndex = null;
        // Find end index of the number
        for (let j = beginIndex; j < winningNrsStr.length; j++) {
          if (winningNrsStr[j] === " ") {
            endIndex = j - 1;
            break;
          } else if (j + 1 === winningNrsStr.length) {
            endIndex = j;
          }
        }

        if (endIndex == null) throw new Error("Could not find end index of winning number")

        const winningNr = winningNrsStr.substring(beginIndex, endIndex + 1)
        winningNrs.push(winningNr)

        // Reset the outer loop that iterates over the whole winning number line
        i = endIndex + 1
      }
    }

    // Find my numbers
    const myNrs = []
    for (let i = 0; i < myNrsStr.length; i++) {
      if (numbers.includes(myNrsStr[i])) {
        beginIndex = i
        endIndex = null
        // Find end index of the number
        for (let j = beginIndex; j < myNrsStr.length; j++) {
          if (myNrsStr[j] === " ") {
            endIndex = j - 1
            break;
          } else if (j + 1 === myNrsStr.length) {
            endIndex = j
          }
        }

        if (endIndex == null) throw new Error("Could not find end index of my number")
        
        const myNr = myNrsStr.substring(beginIndex, endIndex + 1)
        myNrs.push(myNr)

        // Reset the outer loop that iterates over the whole winning number line
        i = endIndex + 1
      }
    }

    const myWinningNrs = []
    // Find my numbers that are winning
    for (const winningNr of winningNrs) {
      for (const myNr of myNrs) {
        if (winningNr === myNr) {
          myWinningNrs.push(myNr)
          break;
        }
      }
    }

    cardMatchingNumbers = myWinningNrs.length

    parsedLines.push({
      lineIndex: parseInt(lineIndex),
      matchingNumbers: cardMatchingNumbers,
    })
  }

  return parsedLines
}

const main = async () => {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  const lines = []

  for await (const line of rl) {
    lines.push(line)
  }

  const parsedLines = getParsedLines(lines)

  let totalCards = 0
  let backlog = [...parsedLines]

  while(backlog.length > 0) {
    const backlogLen = backlog.length
    
    // Go through the backlog
    for (const backlogItemIndex in backlog) {
      const backlogItem = backlog[backlogItemIndex]

      if (backlogItem.matchingNumbers > 0) {
        // For the current backlog item find all the copies of cards
        for (let i = backlogItem.lineIndex + 1; i < backlogItem.lineIndex + 1 + backlogItem.matchingNumbers; i++) {
          // Add the copy of card to backlog
          backlog.push(parsedLines[i])
        }
      }
    }

    // Remove the cards from backlog that were just processed
    backlog.splice(0, backlogLen)
     
    // Add to total card since we processed it
    totalCards = totalCards + backlogLen
  }

  console.log(totalCards);
}

main()