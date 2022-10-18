import startServer from "./Server";

const timer = ms => new Promise(res => setTimeout(res, ms))

let i = 0
// while (i < 10) {
//     await timer(3000)
//     i++
// }

const data: Item = startServer()
console.log(`Name: ${data.name}`)
console.log(`Price: ${data.price}`)
console.log(`Discount: ${data.discount}`)
console.log(`Discounted Price: ${data.discountPrice}`)

interface Item {
    name: string
    price: string
    discount: string
    discountPrice: number
}