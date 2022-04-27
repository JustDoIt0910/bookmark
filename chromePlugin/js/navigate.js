document.getElementById("space").addEventListener("click", function() {
    chrome.tabs.create({
        url: "http://asilentboy.cn/bookmark/#/space"
    })
})

document.getElementById("login").addEventListener("click", function() {
    chrome.tabs.create({
        url: "../pages/login.html"
    })
})