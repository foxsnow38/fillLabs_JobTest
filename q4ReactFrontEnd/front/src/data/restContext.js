import axios from "axios";
import { createContext } from "react";

const restContext = createContext(null)


const restContextProvider= ({children}) => {





return (<restContext.Provider value={null}>
    {children}
</restContext.Provider>)

    
}