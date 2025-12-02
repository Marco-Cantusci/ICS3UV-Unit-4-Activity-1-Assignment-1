/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-02
 * @fileoverview Calculate average mark.
 */

let total: number = 0;

// input for number of marks
const marksString: string = prompt("How many marks will you enter today? \n") || ('\n');
const marksNumber: number = parseInt(marksString);

// for statement
for (let markAmount = 0; markAmount < marksNumber; markAmount++) {

  const gradeString: string = prompt(`Enter mark ${markAmount}: `) || ('\n');
  const gradeNumber: number = parseInt(gradeString);
  total += gradeNumber;
};

// average calculation
const average: number = total / marksNumber;
console.log(`\nAverage mark: ${average}\n`);

// statements
if (average <= 49) {
  console.log("The student is failing.");
} else if (average >= 50 && average <= 69) {
  console.log("The student's performance is just below average.");
} else if (average >= 70 && average <= 79) {
  console.log("The student's performance is average.");
} else if (average >= 80) {
  console.log("The student is on the honour roll.");
}

console.log("\nDone.")