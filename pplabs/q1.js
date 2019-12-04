// q1
;(function() {
  for (var i = 0; i < 5; i++) {
    let x = i;
    setTimeout(() => {
      console.log(x + 1);
    }, i * 1000);
  }
})();
/** js 单线程异步问题
 * + for / setTimeout - 同步
 * + setTimeout delay 被插入任务队列
 * + 打印 i 受作用域问题影响，故创建 temp 变量 x
 **/
