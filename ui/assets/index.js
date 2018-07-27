function getRandomNameView() {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            var element = document.querySelector("#random-name");
            element.className = "";
            element.innerHTML = this.responseText;
            element.className = "show-name";
        }
    }
    request.open("GET", "/api/v1/name", true);
    request.send();
}
function setupListeners() {
    document.addEventListener('click', function (ev) {
        getRandomNameView();
    });
}
function init() {
    setupListeners();
    getRandomNameView();
}
init();
