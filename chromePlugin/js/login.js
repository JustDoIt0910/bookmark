document.getElementById("login-btn").addEventListener("click", function() {
    var username = document.getElementById("username").value
    var password = document.getElementById("password").value
    $.ajax({ 
        type: "POST",
        url: "http://asilentboy.cn:9999/bookmark/api/v1/login",
        data: JSON.stringify({
            username: username,
            password: password
          }),
        success: function(result){
            if(result.code == 200) {
                window.localStorage.setItem("token", result.token)
                alert("登录成功")
            }
            else {
                alert("登录失败")
            }
          }
      })
})