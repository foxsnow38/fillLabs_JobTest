

const findRecrusive= (input:number) =>{
if(input >=1) console.log(input) 
const number = Math.floor(input/2)  as number
if(number >1) findRecrusive(number)



}
findRecrusive(9)

export default findRecrusive