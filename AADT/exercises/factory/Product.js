"use strict";
exports.__esModule = true;
exports.BallFactory = exports.ShoeFactory = void 0;
var Ball = /** @class */ (function () {
    function Ball(ball) {
        this.name = ball.name;
        this.price = ball.price;
        this.discount = ball.discount;
        this.discountedPrice = ball.discountedPrice;
    }
    return Ball;
}());
var Racket = /** @class */ (function () {
    function Racket(racket) {
        this.name = racket.name;
        this.price = racket.price;
        this.discount = racket.discount;
        this.discountedPrice = racket.discountedPrice;
    }
    return Racket;
}());
var Shoe = /** @class */ (function () {
    function Shoe(shoe) {
        this.name = shoe.name;
        this.price = shoe.price;
        this.discount = shoe.discount;
        this.discountedPrice = shoe.discountedPrice;
    }
    return Shoe;
}());
var BallFactory = /** @class */ (function () {
    function BallFactory() {
    }
    BallFactory.prototype.createBall = function (ball) {
        return new Ball(ball);
    };
    BallFactory.prototype.createRacket = function (racket) {
        return new Racket(racket);
    };
    BallFactory.prototype.createShoe = function (shoe) {
        return new Shoe(shoe);
    };
    return BallFactory;
}());
exports.BallFactory = BallFactory;
var ShoeFactory = /** @class */ (function () {
    function ShoeFactory() {
    }
    ShoeFactory.prototype.createBall = function (ball) {
        return new Ball(ball);
    };
    ShoeFactory.prototype.createRacket = function (racket) {
        return new Racket(racket);
    };
    ShoeFactory.prototype.createShoe = function (shoe) {
        return new Shoe(shoe);
    };
    return ShoeFactory;
}());
exports.ShoeFactory = ShoeFactory;
