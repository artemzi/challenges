/**
 * A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.



Example 1:

Input: "(()())(())"
Output: "()()()"
Explanation:
The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
Example 2:

Input: "(()())(())(()(()))"
Output: "()()()()(())"
Explanation:
The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
Example 3:

Input: "()()"
Output: ""
Explanation:
The input string is "()()", with primitive decomposition "()" + "()".
After removing outer parentheses of each part, this is "" + "" = "".


Note:

S.length <= 10000
S[i] is "(" or ")"
S is a valid parentheses string
 */

/**
 * @param {string} S
 * @return {string}
 */
let removeOuterParentheses = function (S) {
    let res = '';
    let c = 0;
    for (let s of S) {
        if (s == '(') {
            c++;
            if (c > 1) res += '(';
        } else {
            c--;
            if (c >= 1) res += ')';
        }
    }
    return res;
};

/**
 * Result:
 * 
 * Runtime: 80 ms, faster than 18.38 % of JavaScript online submissions
    for Remove Outermost Parentheses.
    Memory Usage: 36.2 MB, less than 100.00 % of JavaScript online submissions
    for Remove Outermost Parentheses.
 */

/**
 * @param {string} S
 * @return {string}
 */
let removeOuterParentheses = function (S) {
    let res = '';
    let c = 0;
    let i;
    let l = S.length;
    for (i = 0; i < l; i++) {
        if (S[i] == '(') {
            c++;
            if (c > 1) res += '(';
        } else {
            c--;
            if (c >= 1) res += ')';
        }
    }
    return res;
};

/**
 * Result:
 * 
 * Runtime: 60 ms, faster than 100.00 % of JavaScript online submissions
 for Remove Outermost Parentheses.
 Memory Usage: 35.9 MB, less than 100.00 % of JavaScript online submissions
 for Remove Outermost Parentheses.
 */