$("document").ready(() => {

    let btn1 = $("#buttonLogin");
    let check = "false";
    // Handle the checkbox
    $("#checkLogin").click(function (e) {
        if ($(this).is(':checked')) {
            check = "true";
        } else {
            check = "false";
        }
    });

    btn1.click(() => {

        var name = $("#usernameLogin").val();
        var password = $("#passwordLogin").val();
        
        $.post("/u/login", {
            usernameLogin: name,
            passwordLogin: password,
            checkLogin: check,
        }, function(result){
            if (result.message == "Successful login") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => location.reload(), 1500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
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
