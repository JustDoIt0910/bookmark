document.getElementById("space").addEventListener("click", function() {
    chrome.tabs.create({
        url: "http://localhost:8080/#/space"
    })
})

document.getElementById("login").addEventListener("click", function() {
    chrome.tabs.create({
        url: "../pages/login.html"
    })
})