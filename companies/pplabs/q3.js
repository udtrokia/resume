// q3
/// ----------------
/// setTimeout(function() {
///   console.log(1);
/// }, 100);
/// setTimeout(function () {
///   console.log(2);
/// }, 0)
/// Promise.resolve(console.log(3)).then(() => { console.log(4) })
/// async function async1(){
///   console.log(5)
///   await async2()
///   console.log(6)
/// }
/// async function async2(){
///   console.log(7)
/// }
/// async1()
/// console.log(8)
/// ----------------
/** promise 执行顺序及 setTimeout 队列问题
 * @result: 3, 5, 7, 8, 4, 6, 2, 1
 * + Promise 定义即执行, 并优先于 setTimeout
 * + await 阻塞线程
 **/
