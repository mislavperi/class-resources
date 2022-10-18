"use strict";
exports.__esModule = true;
var Server_1 = require("./Server");
var timer = function (ms) { return new Promise(function (res) { return setTimeout(res, ms); }); };
var i = 0;
// while (i < 10) {
//     await timer(3000)
//     i++
// }
var data = (0, Server_1["default"])();
console.log("Name: ".concat(data.name));
console.log("Price: ".concat(data.price));
console.log("Discount: ".concat(data.discount));
console.log("Discounted Price: ".concat(data.discountPrice));
