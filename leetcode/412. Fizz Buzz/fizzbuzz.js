/**
 * Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
 */
/**
 * @param {number} n
 * @return {string[]}
 */
const fizzBuzz = (n) => {
    return Array(n).fill(1).map((_, i) => {
        return ((i + 1) % 3 ? '' : 'Fizz') + ((i + 1) % 5 ? '' : 'Buzz') || '' + (i + 1);
    });
};
// Runtime: 80 ms, faster than 43.89 % of JavaScript online submissions for Fizz Buzz.
// Memory Usage: 36.8 MB, less than 100.00% of JavaScript online 