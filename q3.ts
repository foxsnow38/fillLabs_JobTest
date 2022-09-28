

const findRecrusive= (input:string[]) =>{
const inputObject:any={} // make object
for (let index = 0; index < input.length; index++) {
    if(!inputObject[`${input[index]}`])inputObject[`${input[index]}`]=0 //if object keys input[index] is null than change zero
    inputObject[`${input[index]}`]+=1 // than object keys +1
}


return Object.entries(inputObject).sort((a,b):any=> +b[1]- +a[1])[0][0] 
  // than I convert array to object than I sort.
  // than I max value take.  
 // here the a and b problem solve// https://stackoverflow.com/a/68876658/18456713
}
console.log(findRecrusive(["apple","pie","apple","red","red","red"]))

export default findRecrusive