import * as fs from "fs";
import { BallFactory, ShoeFactory } from "./Product";

interface Item {
    name: string
    price: string
    discount: string
    discountPrice: number
}

function startServer(): Item {
    console.log("started")
    const str = fs.readFileSync("products.txt", "utf8")
    const splitted = str.split(";")
    switch (splitted[0]) {
        case "ball": {
            const factory = new BallFactory()
            const item = factory.createBall({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            })
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            }
        }
        case "shoe": {
            const factory = new ShoeFactory()
            const item = factory.createShoe({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            })
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            }
        }
        case "racket": {
            const factory = new ShoeFactory()
            const item = factory.createRacket({
                name: splitted[0],
                price: splitted[1],
                discount: splitted[2],
                discountedPrice: parseInt(splitted[1]) * (parseInt(splitted[2]) / 100)
            })
            return {
                name: item.name,
                price: item.price,
                discount: item.discount,
                discountPrice: item.discountedPrice
            }
        }
        default: {
            return {
                name: "",
                price: "item.price",
                discount: "",
                discountPrice: 0
            }
        }
    }
}

export default startServer