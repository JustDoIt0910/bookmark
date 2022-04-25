function getBookmarks() {
  chrome.bookmarks.getTree(function(bookmarkArray){
    console.log(bookmarkArray)
  })
}

chrome.contextMenus.create({
	"type" : "normal",
	"title" : "收藏此网页",
	"checked" : false,
	"onclick" : addBookmark
})

function addBookmark() {
  chrome.tabs.getSelected(null, function (tab) {
    var token = window.localStorage.getItem('token')
    var Bearertoken = `Bearer ${token}`
    $.ajax({ 
      type: "POST",
      url: "http://localhost:3000/api/v1/bookmark",
      data: JSON.stringify({
        url: tab.url
      }),
      beforeSend: function(xhr) { 
            xhr.setRequestHeader("Authorization", Bearertoken)  
        },
      success: function(result){
          if(result.code == 1007 || result.code == 1008 || result.code == 1009) {
              chrome.tabs.create({
                url: "../pages/login.html"
            })
          }
          else if(result.code == 200) {
            alert("网页收藏成功")
          }
        }
    })
  })
}
