var $ = require('jquery');
alert("OK");
var button = $("<button>").html("click me").on("click", function() {
	alert("test");
});

module.exports = button;