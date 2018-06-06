$(document).ready(function () {
    $(".d-sidebar").css({ height: window.innerHeight - 58 });
    $(".d-content").css({ height: window.innerHeight - 58 });
});

document.getElementById("d-mini").onclick = function (e) {
    e.stopPropagation();
    document.getElementById("d-sidebar").classList.toggle("d-sidebar-shrink");
    document.getElementById("d-content").classList.toggle("d-content-expand");

};

var accItem = document.getElementsByClassName("accordionItem");
var accHD = document.getElementsByClassName("accordionItemHeading");
for (i = 0; i < accHD.length; i++) {
    accHD[i].addEventListener("click", toggleItem, false);
}
function toggleItem() {
    var itemClass = this.parentNode.className;
    for (i = 0; i < accItem.length; i++) {
        accItem[i].className = "accordionItem close";
    }
    if (itemClass == "accordionItem close") {
        this.parentNode.className = "accordionItem open";
    }
}

function notifShower() {
    document.getElementById("notif-content").classList.toggle("show");
}

window.onclick = function (e) {
    if (!e.target.matches('.d-notifications')) {
        var dropdown = document.getElementById("notif-content");
        if (dropdown.classList.contains('show')) {
            dropdown.classList.remove('show');
        }
    }
}