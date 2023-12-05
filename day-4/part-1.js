const fs = require('node:fs');
const readline = require('node:readline');

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

  const allCardPoints = []

  for (const line of lines) {
    const [cardNrStr, gamesStr] = line.split(":")
    const [winningNrsStr, myNrsStr] = gamesStr.split("|")
    const numbers = ['0','1','2','3','4','5','6','7','8','9']
    
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

    let cardPoints = 0

    for (const myWinningNrIndex in myWinningNrs) {
      if (myWinningNrIndex === '0') {
        cardPoints = 1  
      } else {
        cardPoints = cardPoints * 2
      }
    }
    
    allCardPoints.push(cardPoints)
  }

  let allPoints = 0;

  for (const cardPoint of allCardPoints) {
    allPoints += cardPoint
  }

  console.log(allPoints);
}

main()