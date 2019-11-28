// nums
function topKFrequent(nums, k) {
  let exist = [];
  let stack = [];

  for (let i in nums) {
    let e = nums[i];
    if (exist.indexOf(e) == -1) {
      exist.push(e);
      stack.push(1);
      continue;
    }

    stack[exist.indexOf(e)] += 1;
  }

  let tree = {};
  for (let k in stack) {
    tree[stack[k]] = exist[k];
  }

  let _res = stack.sort(function(a, b) {return (b-a)});
  let res = [];

  for (let i = 0; i<k; i++) {
    res.push(tree[_res[i].toString()]);
  }

  return res;
}

(function() {
  console.log(topKFrequent([3, 3, 1, 2, 2, 3], 3));
})();
