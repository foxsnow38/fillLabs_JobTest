import axios from "axios";

import { createContext, useContext } from "react";

const restContext = createContext(null)
const getUser = async  () =>{
    
 const axi = await axios.get("http://localhost:8081/user")

     console.log(axi)
}


const RestContextProvider= ({children}) => {

const values ={
    getUser
}



return (<restContext.Provider value={values}>
    {children}
</restContext.Provider>)

    
}

const useRestContext = () => useContext(restContext)

export {RestContextProvider ,useRestContext}