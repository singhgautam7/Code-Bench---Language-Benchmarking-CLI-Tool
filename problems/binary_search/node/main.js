"use strict";

/**
 * Perform binary search on a sorted array.
 * Returns the index of target, or -1 if not found.
 */
function binarySearch(arr, target) {
  let low = 0;
  let high = arr.length - 1;

  while (low <= high) {
    const mid = low + ((high - low) >> 1);
    if (arr[mid] === target) return mid;
    else if (arr[mid] < target) low = mid + 1;
    else high = mid - 1;
  }

  return -1;
}

let input = "";
process.stdin.setEncoding("utf8");
process.stdin.on("data", (chunk) => (input += chunk));
process.stdin.on("end", () => {
  const tokens = input.trim().split(/\s+/).map(Number);
  if (tokens.length < 2) {
    process.stderr.write("Error: expected array and target on stdin\n");
    process.exit(1);
  }

  const target = tokens.pop();
  const result = binarySearch(tokens, target);
  console.log(result);
});
