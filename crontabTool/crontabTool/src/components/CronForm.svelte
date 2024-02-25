<script lang=ts>
  
  import { MonthEnum } from "../utils/MonthEnum";
  import {errorCheck, digitHelper, validNumberHelper} from "../utils/functions";

  var showResults = false;
  var hasError = false;
  var results = "";
  var standardSpecialCharacters: string[] = ["*", ",", "-"];

  function processMinutes(value: string): string {
    let humanMinute = "";
    if (value == "*") {
      humanMinute = "Every minute";
    }
    else if (1 <= Number(value) && Number(value) <= 59) {
      humanMinute = "At :" + digitHelper(value) + " minute";
    }
    else if (value.includes(",")) {
      let multipleValues: string[] = value.split(",");
      for (let i: int = 0; i < multipleValues.length; i++) {
        if (i == (multipleValues.length - 1)) {
          humanMinute = humanMinute + ", and :" + digitHelper(multipleValues[i]);
        }
        else if (i == 0) {
          humanMinute = ":" + digitHelper(multipleValues[i]);
        }
        else {
          humanMinute = humanMinute + ", :" + digitHelper(multipleValues[i]);
        }
      }
    }
    else if (value.includes("-")) {
      let multipleValues: string[] = value.split("-");
      if (multipleValues.length != 2 || !validNumberHelper(multipleValues[0], 59, 0) || !validNumberHelper(multipleValues[1], 59, 0)) {
        return "";
      }
      humanMinute = "Each minute from :" + digitHelper(multipleValues[0]) + " through :" + digitHelper(multipleValues[1]); 
    }
    else if (value.includes("/")) {
      let multipleValues: string[] = value.split("/");
      if (multipleValues.length == 2 && (!validNumberHelper(multipleValues[0], 59, 0) || !validNumberHelper(multipleValues[1], 59, 0))) {
        return "";
      }
      else if (multipleValues.length == 1 && !validNumberHelper(multipleValues[0], 59, 0)) {
        return "";
      }
      
      if (multipleValues.length == 1 || multipleValues[0] == "*") {
        humanMinute = "Every "+ multipleValues[0] + " minutes";
      }
      else if (multipleValues.length == 2) {
        humanMinute = "Ever " + multipleValues[1] + " minutes starting at :" + digitHelper(multipleValues[0]);
      }
    }

    return humanMinute;
  }

  function processHour(value: string): string {
    let humanHour = "";
    if (value == "*") {
      humanHour = "Every hour";
    }
    else if (1 <= Number(value) && Number(value) <= 23) {
      humanHour = "At " + digitHelper(value) + " hour";
    }
    else if (value.includes(",")) {
      let multipleValues: string[] = value.split(",");
      for (let i: int = 0; i < multipleValues.length; i++) {
        if (i == (multipleValues.length - 1)) {
          humanHour = humanHour + ", and :" + digitHelper(multipleValues[i]);
        }
        else if (i == 0) {
          humanHour = ":" + digitHelper(multipleValues[i]);
        }
        else {
          humanHour = humanHour + ", :" + digitHelper(multipleValues[i]);
        }
      }
    }
    else if (value.includes("-")) {
      let multipleValues: string[] = value.split("-");
      if (multipleValues.length != 2 || !validNumberHelper(multipleValues[0], 23, 0) || !validNumberHelper(multipleValues[1], 23, 0)) {
        return "";
      }
      humanHour = "Each hour from :" + digitHelper(multipleValues[0]) + " through :" + digitHelper(multipleValues[1]); 
    }
    else if (value.includes("/")) {
      let multipleValues: string[] = value.split("/");
      if (multipleValues.length == 2 && (!validNumberHelper(multipleValues[0], 23, 0) || !validNumberHelper(multipleValues[1], 23, 0))) {
        return "";
      }
      else if (multipleValues.length == 1 && !validNumberHelper(multipleValues[0], 23, 0)) {
        return "";
      }
      
      if (multipleValues.length == 1 || multipleValues[0] == "*") {
        humanHour = "Every "+ multipleValues[0] + " hours";
      }
      else if (multipleValues.length == 2) {
        humanHour = "Ever " + multipleValues[1] + " hour starting at :" + digitHelper(multipleValues[0]);
      }
    }

    return humanHour;
  }

  function processDayOfMonth(value: string): string {
    let humanDayOfMonth = "";
    if (value == "*") {
      humanDayOfMonth = "every day of the month";
    }
    else if (1 <= Number(value) && Number(value) <= 31) {
      humanDayOfMonth = "On " + digitHelper(value) + " of the month";
    }
    else if (value.includes(",")) {
      let multipleValues: string[] = value.split(",");
      for (let i: int = 0; i < multipleValues.length; i++) {
        if (i == (multipleValues.length - 1)) {
          humanDayOfMonth = humanDayOfMonth + ", and " + digitHelper(multipleValues[i]);
        }
        else if (i == 0) {
          humanDayOfMonth = digitHelper(multipleValues[i]);
        }
        else {
          humanDayOfMonth = humanDayOfMonth + ", " + digitHelper(multipleValues[i]);
        }
      }
    }
    else if (value.includes("-")) {
      let multipleValues: string[] = value.split("-");
      if (multipleValues.length != 2 || !validNumberHelper(multipleValues[0], 31, 0) || !validNumberHelper(multipleValues[1], 31, 0)) {
        return "";
      }
      humanHour = "Each day of the month from " + digitHelper(multipleValues[0]) + " through " + digitHelper(multipleValues[1]); 
    }
    else if (value.includes("/")) {
      let multipleValues: string[] = value.split("/");
      if (multipleValues.length == 2 && (!validNumberHelper(multipleValues[0], 31, 0) || !validNumberHelper(multipleValues[1], 31, 0))) {
        return "";
      }
      else if (multipleValues.length == 1 && !validNumberHelper(multipleValues[0], 31, 0)) {
        return "";
      }
      
      if (multipleValues.length == 1 || multipleValues[0] == "*") {
        humanDayOfMonth = "Every "+ multipleValues[0] + " days of the month";
      }
      else if (multipleValues.length == 2) {
        humanDayOfMonth = "Ever " + multipleValues[1] + " day of the month starting at " + digitHelper(multipleValues[0]);
      }
    }

return humanDayOfMonth;
  }

  function processMonth(value: string): string {
    let humanMonth = "";
    switch (value) {
      case "*":
        humanMonth = "every month";
    }
    return humanMonth;
  }

  function processDayOfWeek(value: string): string {
    let humanDayOfWeek = "";
    switch (value) {
      case "*":
        humanDayOfWeek = "every day of the week";
    }
    return humanDayOfWeek;
  }

  function onSubmit(e: any) {
    hasError = false;
    showResults = false;
    let rawInput: string = e.target[0].value;
    let pattern: string[] = rawInput.split(" ");
    hasError = errorCheck(pattern);
    if (hasError) {
      return;
    }
    let currentVal = "";
    for (let i: int = 0; i < pattern.length; i++) {
      currentVal = pattern[i];
      switch (i) {
        case 0:
          let minuteResults: string = processMinutes(currentVal);
          if (minuteResults == "") {
            hasError = true;
            return;
          }
          else {
            results = minuteResults;
          }
          break;
        case 1:
          let hourResults: string = processHour(currentVal);
          if (hourResults == "") {
            hasError = true;
            return;
          }
          else {
            results = results + ", " + hourResults;
          }
          break;
        case 2:
          let dayOfMonthResults: string = processDayOfMonth(currentVal);
          if (dayOfMonthResults == "") {
            hasError = true;
            return;
          }
          else {
            results = results + ", " + dayOfMonthResults;
          }
          break;
        case 3:
          let monthResults: string = processMonth(currentVal);
          if (monthResults == "") {
            hasError = true;
            return;
          }
          else {
            results = results + ", " + monthResults;
          }
          break;
        case 4: 
          let dayOfWeekResults: string = processDayOfWeek(currentVal);
          if (dayOfWeekResults == "") {
            hasError = true;
            return;
          }
          else {
            results = results + ", " + dayOfWeekResults;
          }
          break;
        default:
          console.log("Error");
          hasError = true;
          break;
      }
    }
    if (!hasError) {
      showResults = true;
    }
  }
</script>

<style>
  .container {
    padding: 3em;
  }
</style>

<div class="container">
  <div class="form">
    <form on:submit|preventDefault={onSubmit}>
      <div class="mx-auto">
        <div>
          <div class="relative">
            <input type="text" id="cronPattern" placeholder=" " class="peer block w-full rounded-md border-black shadow-sm focus:border-primary-400 focus:ring focus:ring-primary-200 focus:ring-opacity-50 disabled:cursor-not-allowed disabled:bg-gray-50 disabled:text-gray-500" />
            <label for="cronPattern" class="peer-focus:base absolute left-2 top-0 z-10 -translate-y-2 transform bg-white px-1 text-xs text-gray-500 transition-all peer-placeholder-shown:translate-y-3 peer-placeholder-shown:text-sm peer-focus:-translate-y-2 peer-focus:text-xs peer-disabled:bg-transparent">Cron Pattern</label>
          </div>
        </div>
      </div>
    </form>
  </div>

  <br>

  {#if showResults}
    <div class="mx-auto">
      <input type="text" class="block w-full rounded-md border-black shadow-sm focus:border-primary-400 focus:ring focus:ring-primary-200 focus:ring-opacity-50 disabled:cursor-not-allowed disabled:bg-gray-50 disabled:text-gray-500" placeholder="{ results }" />
    </div>
  {:else if hasError}
     <div class="rounded-md bg-red-50 p-4 text-sm text-red-500"><b>Error alert</b> Incorrect cron pattern.</div>
  {/if}
</div>
