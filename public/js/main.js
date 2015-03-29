/**
 The Main.js module is the starting point of the client side JavaScript application.
 It starts and stops modules depending on the site the user is visiting.
 **/
Core.DomManipulation.ready(document, function() {

    // get the current site the user is visiting
    var url = window.location.href.split("/"),
        site = url[url.length - 1].split("?")[0];

    // stop all modules
    Core.stopAll();

    // start modules depending on which site the user is visiting
    if (site == "Todo") {
        Core.start("todo");
    }

    if (site == "Todo") {
        Core.start("todo");
    }
});