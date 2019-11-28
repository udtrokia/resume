//route
function findHandler(method, url) {
  const base = "http://www.xxx.com/";
  let target = url.substring(base.length + 1);
  let uc = url.charAt(base.length).toUpperCase();
  target = uc + target;
  
  if (method == "Get") {
    return "query" + target;
  } else if (method == "Post") {
    return "add" + target;
  } else if (method == "Delete") {
    return "delete" + target;
  } else if (method == "Update") {
    return "update" + target;
  }
}

//testing
(function(){
  console.log(findHandler("Get", "http://www.xxx.com/weibo"));
})();
