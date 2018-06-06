$(document).ready(function () {
    $(".fullscreen").css({ height: window.innerHeight - 62 });

    var html = $("#section-html").text();
    $("#section-preview").html(html);

    var li = $(".hp-li");
    var li0 = $(".hp-li:eq(0)");
    var li1 = $(".hp-li:eq(1)");
    var li2 = $(".hp-li:eq(2)");
    var li3 = $(".hp-li:eq(3)");
    li0.click(function(){
        setAllDisplayNone(li);
        li0.addClass("selected");
        $("#section-preview").css("display", "block");
    });
    li1.click(function(){
        setAllDisplayNone(li);
        li1.addClass("selected");
        $("#section-html").css("display", "block");
    });
    li2.click(function(){
        setAllDisplayNone(li);
        li2.addClass("selected");
        $("#section-css").css("display", "block");
    });
    li3.click(function(){
        setAllDisplayNone(li);
        li3.addClass("selected");
        $("#section-js").css("display", "block");
    });      
});

function setSelectedDiv(element){
    $(".td-box").removeClass('border-blue');
    $(element).toggleClass('border-blue');
    var id = element.id;
    document.getElementById("path").value = id;
}

function setAllDisplayNone(li){
    if (li.hasClass("selected")) {
        li.removeClass("selected");
    }
    $("#section-preview").css("display", "none");
    $("#section-html").css("display", "none");
    $("#section-css").css("display", "none");
    $("#section-js").css("display", "none");
}

function openSidebar() {
    document.getElementById("hm-menu").style.width = "250px";
    document.getElementById("close").style.display = "block";
    document.getElementById("open").style.display = "none";
    disableScroll();
}

function closeSidebar() {
    document.getElementById("hm-menu").style.width = "0";
    document.getElementById("close").style.display = "none";
    document.getElementById("open").style.display = "block";
    enableScroll();
}

var keys = { 37: 1, 38: 1, 39: 1, 40: 1 };

function preventDefault(e) {
    e = e || window.event;
    if (e.preventDefault)
        e.preventDefault();
    e.returnValue = false;
}

function preventDefaultForScrollKeys(e) {
    if (keys[e.keyCode]) {
        preventDefault(e);
        return false;
    }
}

function disableScroll() {
    if (window.addEventListener)
        window.addEventListener('DOMMouseScroll', preventDefault, false);
    window.onwheel = preventDefault;
    window.onmousewheel = document.onmousewheel = preventDefault;
    window.ontouchmove = preventDefault;
    document.onkeydown = preventDefaultForScrollKeys;
}

function enableScroll() {
    if (window.removeEventListener)
        window.removeEventListener('DOMMouseScroll', preventDefault, false);
    window.onmousewheel = document.onmousewheel = null;
    window.onwheel = null;
    window.ontouchmove = null;
    document.onkeydown = null;
}