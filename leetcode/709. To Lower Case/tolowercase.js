/**
 * Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"
 */
/**
 * @param {string} str
 * @return {string}
 */
const toLowerCase = function(str) {
  return [...str]
    .map((s) => {
      let code = s.charCodeAt(0);
      return code <= 90 && code >= 65 ? String.fromCharCode(code + 32) : s;
    })
    .join("");
};
// Runtime: 68 ms
// Memory Usage: 14.6 MB
