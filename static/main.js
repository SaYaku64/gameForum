$("document").ready(() => {

    let btn1 = $("#buttonLogin");
    let btn2 = $("#buttonComment");

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
    });

    btn2.click(() => {

        var comtitle = document.getElementById('comtitle').innerHTML
        var comment = $("#comment").val();
        


        $('.close').trigger('click');



        $.post("/article/comment", {
            comtitle: comtitle,
            comment: comment,
        }, function(result){
            if (result.message == "Successfully added") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                // document.getElementById("commentPlace").innerHTML = result.message;
                $( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                //setTimeout(() => location.reload(), 1500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => location.reload(), 1500);
            }
        });
    });
});
