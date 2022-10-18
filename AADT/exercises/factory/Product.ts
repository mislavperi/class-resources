interface AbstractFactory {
    createBall(ball: BallProduct): BallProduct
    createShoe(shoe: ShoeProduct): ShoeProduct
    createRacket(racket: RacketProduct): RacketProduct
}

interface BallProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number
}

interface ShoeProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number
}

interface RacketProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number
}

class Ball implements BallProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number

    constructor(ball) {
        this.name = ball.name
        this.price = ball.price
        this.discount = ball.discount
        this.discountedPrice = ball.discountedPrice
    }
}

class Racket implements RacketProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number

    constructor(racket) {
        this.name = racket.name
        this.price = racket.price
        this.discount = racket.discount
        this.discountedPrice = racket.discountedPrice
    }
}


class Shoe implements ShoeProduct {
    name: string
    price: string
    discount: string
    discountedPrice: number

    constructor(shoe) {
        this.name = shoe.name
        this.price = shoe.price
        this.discount = shoe.discount
        this.discountedPrice = shoe.discountedPrice
    }
}


class BallFactory implements AbstractFactory {
    public createBall(ball: Ball): Ball {
        return new Ball(ball)
    }
    public createRacket(racket: Racket): Racket {
        return new Racket(racket)
    }
    public createShoe(shoe: Shoe): Shoe {
        return new Shoe(shoe)
    }
}



class ShoeFactory implements AbstractFactory {
    public createBall(ball: Ball): Ball {
        return new Ball(ball)
    }
    public createRacket(racket: Racket): Racket {
        return new Racket(racket)
    }
    public createShoe(shoe: Shoe): Shoe {
        return new Shoe(shoe)
    }
}


export { ShoeFactory, BallFactory }