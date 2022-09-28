const reverseFnuc =  (input: number) => {
    let  array: any = []

    const findRecrusive = (input: number) => {
        if (input > 1) array.push(input)
        const number = Math.floor(input / 2) as number
        if (number > 1) findRecrusive(number)
    }

    findRecrusive(input)
   array.reverse().forEach((item:number) => console.log(item));
}

reverseFnuc(9)

export default reverseFnuc