$("document").ready(() => {

    let btn = $("#adminButton");

    btn.click(() => {

        var element = $("#delElem").val();
        var name = $("#delName").val();
        
        $.post("/admin/panel", {
            delElem: element,
            delName: name,
        }, function(result){
            if (result.message == "Successfully deleted") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => location.reload(), 1500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
            }
        });
    });
});
