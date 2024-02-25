export function errorCheck(pattern: string[]): boolean {
    if (pattern.length != 5) {
      return true;
    }
    return false;
  }

export function digitHelper(digit: string): string {
    if (digit.length == 1) {
      return "0" + digit;
    }
    else {
      return digit;
    }
  } 

export function validNumberHelper(digit: string, upper: number, lower: number): boolean {
    if (isNaN(Number(digit))) {
      return false;
    }
    let num: number = Number(digit);
    if (num > upper) {
      return false;
    }
    if (num < lower) {
      return false;
    }
    return true;
  }


