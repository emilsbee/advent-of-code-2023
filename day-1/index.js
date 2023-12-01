const fs = require('node:fs');
const readline = require('node:readline');

const NUMBERS_AND_WORDS = false; // determines whether to consider only numbers like 6 or also words like six

const digitMap = {
  'one': '1',
  'two': '2',
  'three': '3',
  'four': '4',
  'five': '5',
  'six': '6',
  'seven': '7',
  'eight': '8',
  'nine': '9',
};
const DIGITS_SPELLED =  Object.keys(digitMap);
const DIGITS = Object.values(digitMap);

async function main() {
  const fileStream = fs.createReadStream('input.txt');

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  const lineNumbers = [];

  for await (const line of rl) {
    let lastDigit;
    let firstDigit;
    
    // Find first digit
    for (const lineCharIndex in line) {
      const lineChar = line[lineCharIndex];

      if (NUMBERS_AND_WORDS) {
        const possibleSpelledDigits = DIGITS_SPELLED.filter((spelledDigit) => spelledDigit.startsWith(lineChar));
  
        for (const possibleSpelledDigit of possibleSpelledDigits) {
          let matches = true;
  
          for (let i = lineCharIndex; i < parseInt(lineCharIndex)+possibleSpelledDigit.length; i++) {
            if (possibleSpelledDigit[i - lineCharIndex] !== line[i]) {
              matches = false;
              break;
            };
          }
  
          if (matches) {
            firstDigit = digitMap[possibleSpelledDigit];
            break;
          }
        }

        if (firstDigit) break;
      }


      if (DIGITS.includes(lineChar)) {
        firstDigit = lineChar;
        break;
      }
    }

    // Find last digit
    for (let i = line.length; i >= 0; i--) {
      const lineCharIndex = i;
      const lineChar = line[lineCharIndex];

      if (NUMBERS_AND_WORDS) {
        const possibleSpelledDigits = DIGITS_SPELLED.filter((spelledDigit) => spelledDigit.startsWith(lineChar));
      
        for (const possibleSpelledDigit of possibleSpelledDigits) {
          let matches = true;
          for (let i = lineCharIndex; i < parseInt(lineCharIndex)+possibleSpelledDigit.length; i++) {
            if (possibleSpelledDigit[i - lineCharIndex] !== line[i]) {
              matches = false;
              break;
            };
          }
  
          if (matches) {
            lastDigit = digitMap[possibleSpelledDigit];
            break;
          }
        }
  
        if (lastDigit) break;
      }

      if (DIGITS.includes(lineChar)) {
        lastDigit = lineChar;
        break;
      }
    }

    let lineCombined = firstDigit + lastDigit;
    lineNumbers.push(lineCombined)
  }

  const total = lineNumbers.reduce((prev, current) => {
    const currNumber = parseInt(current);
    return prev + currNumber;
  }, 0)

  console.log(total);
}

main(); 