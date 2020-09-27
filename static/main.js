$("document").ready(() => {

    let btn = $("#b44");

    btn.click(() => {

        var name = $("#usernameLogin").val();
        var password = $("#passwordLogin").val();
        $.post("/ping", {
            usernameLogin: name,
            passwordLogin: password,
        }, function(result){
            //$("span").html(result);
            //alert(result.message)
            if (result.message == "Login Failed: Invalid login or password") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
            } else {
                location.reload();
            }
        });

        // $.post("demo_test_post.asp",
        // {
        //   name: "Donald Duck",
        //   city: "Duckburg"
        // },
        // function(data,status){
        //   alert("Data: " + data + "\nStatus: " + status);
        // });

        // $.ajax({
        //     //type: "POST",
        //     url: "/ping",
        //     content: "application/json",
        //     success: (result) => {
        //         console.log(result)
        //         document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                
        //     },
        // });
    });
});
