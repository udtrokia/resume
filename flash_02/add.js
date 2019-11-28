// add
// function add() {
//   let res = 0;
//   let args = Object.values(arguments);
// 
//   for (let i in args) {
//     res = args[i] + res;
//   }
//   
//   return function(x) {
//     return res + x;
//   }
// }

function add(m) {
  var tmp = function (n) {
    return add(m + n);
  }
  
  tmp.toString = function () {
    return m;
  }
  
  return tmp;
};

(function() {
  console.log(add(3)(2).toString());
})();
