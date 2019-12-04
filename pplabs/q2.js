// q2
;(function() {
  function sum(arr, target) {
    let map = {};
    for (let i = 0; i < arr.length; i++) {
      if (map[target - arr[i]] != null) {
        return true
      }
      map[arr[i]] = i;
    }

    return false;
  }
  console.log(sum([1, 2, 3, 4, 5], 9));
})();
/** LeetCode 第 1 题
 * + O(n): arr 越长效率越低
 * + 确保相加为 target 的两个 element 同时存在于数组即可
 **/
