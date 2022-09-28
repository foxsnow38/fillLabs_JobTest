const reverseFnuc =  (input: number) => {
    let  array: any = []

    const findRecrusive = (input: number) => {
        if (input > 1) array.push(input)
        const number = Math.floor(input / 2) as number
        if (number > 1) findRecrusive(number)
    }
    // I CAN MAKE LIKE 9/1 9/2 9/4 BUT THIS SOLVE STYLE IS NOT PRETTY FIT TO RECRUSIVE FUNC IDEA

    findRecrusive(input)
   array.reverse().forEach((item:number) => console.log(item));
}

reverseFnuc(9)

export default reverseFnuc