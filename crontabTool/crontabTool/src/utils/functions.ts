import { MonthEnum } from "./MonthEnum";
import { WeekEnum } from "./WeekEnum";

export function errorCheck(pattern: string[]): boolean {
  let extensions: string[] = ["@yearly", "@annually", "@monthly", "@hourly", "@weekly", "@daily"];
  if (pattern.length == 1 && extensions.includes(pattern[0])) {
    return false;
  }
  else if (pattern.length != 5) {
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

export function stringMonthToEnum(month: string): MonthEnum {
  let upperMonth:string = month.toUpperCase();
  if (upperMonth == "JAN" || upperMonth == "JANUARY" || month == "1") {
    return MonthEnum.Jan;
  }
  else if (upperMonth == "FEB" || upperMonth == "FEBRUARY" || month == "2") {
    return MonthEnum.Feb;
  }
  else if (upperMonth == "MARCH" || upperMonth == "MAR" || month == "3") {
    return MonthEnum.Mar;
  }
  else if (upperMonth == "APR" || upperMonth == "APRIL" || month == "4") {
    return MonthEnum.April;
  }
  else if (upperMonth == "May" || month == "5") {
    return MonthEnum.May;
  }
  else if (upperMonth == "JUNE" || month == "6") {
    return MonthEnum.June;
  }
  else if (upperMonth == "JULY" || month == "7") {
    return MonthEnum.July;
  }
  else if (upperMonth == "AUG" || upperMonth == "AUGUST" || month == "8") {
    return MonthEnum.Aug;
  }
  else if (upperMonth == "SEPT" || upperMonth == "SEPTEMBER" || month == "9") {
    return MonthEnum.Sept;
  }
  else if (upperMonth == "OCT" || upperMonth == "OCTOBER" || month == "10") {
    return MonthEnum.Oct;
  }
  else if (upperMonth == "NOV" || upperMonth == "NOVEMBER" || month == "11") {
    return MonthEnum.Nov;
  }
  else if (upperMonth == "DEC" || upperMonth == "DECEMBER" || month == "12") {
    return MonthEnum.Dec;
  }
  else {
    return MonthEnum.NA;
  }
}

export function stringDayOfWeekToEnum(dayOfWeek: string): WeekEnum {
  let upperDOW: string = dayOfWeek.toUpperCase();
  if (upperDOW == "SUN" || upperDOW == "SUNDAY" || dayOfWeek == "1") {
    return WeekEnum.Sunday;
  }
  else if (upperDOW == "M" || upperDOW == "MON" || upperDOW == "MONDAY" || dayOfWeek == "2") {
    return WeekEnum.Monday;
  }
  else if (upperDOW == "TUES" || upperDOW == "TUESDAY" || dayOfWeek == "3") {
    return WeekEnum.Tuesday;
  }
  else if (upperDOW == "W" || upperDOW == "WED" || upperDOW == "WEDNESDAY" || dayOfWeek == "4") {
    return WeekEnum.Wednesday;
  }
  else if (upperDOW == "THURS" || upperDOW == "THURSDAY" || dayOfWeek == "5") {
    return WeekEnum.Thursday;
  }
  else if (upperDOW == "F" || upperDOW == "FRI" || upperDOW == "FRIDAY" || dayOfWeek == "6") {
    return WeekEnum.Friday;
  }
  else if (upperDOW == "SAT" || upperDOW == "SATURDAY" || dayOfWeek == "7") {
    return WeekEnum.Saturday;
  }
  else {
    return WeekEnum.NA;
  }
}
