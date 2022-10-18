"use strict";
exports.__esModule = true;
var fs = require("fs");
var Product_1 = require("./Product");
function startServer() {
    console.log("started");
    var str = fs.readFileSync("products.txt", "utf8");
    var splitted = str.split(";");
    switch (splitted[0]) {
        case "ball": {
            var factory = new Product_1.BallFactory();
            var item = factory.createBall({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            });
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            };
        }
        case "shoe": {
            var factory = new Product_1.ShoeFactory();
            var item = factory.createShoe({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            });
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            };
        }
        case "racket": {
            var factory = new Product_1.ShoeFactory();
            var item = factory.createRacket({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            });
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            };
        }
        default: {
            return {
                name: "",
                price: "item.price",
                discount: "",
                discountPrice: 0
            };
        }
    }
}
exports["default"] = startServer;
