
import './App.css';

import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Button } from '@mui/material';
import TextField from '@mui/material/TextField';
import { useRestContext } from './data/restContext';

function App() {
  
const [users,setUsers] = React.useState(null)
  const {getUser} = useRestContext()
  React.useEffect(()=>{
    setUsers([{id:"1",user:"ali"},
    {id:"2",user:"veli"},
    {id:"3",user:"ayse"}])
getUser()
  },[])

  return (
    <div className="App">
     Hello

    {users !=null &&    <TableContainer component={Paper}>
      <div style={{width:`100%`}}>
        <div style={{marginLeft:`20px`}}>
        <TextField id="standard-basic" label="Standard" variant="standard" />
        </div>
     
      </div>
      <Table sx={{ minWidth: 700 }} aria-label="spanning table">
        <TableHead>
         
          <TableRow>
            <TableCell >ID</TableCell>
            <TableCell >Item</TableCell>
            <TableCell >Edit</TableCell>
            <TableCell >Delete</TableCell>
            
          </TableRow>
        </TableHead>
        <TableBody>
        
       {users.map(item => (   <TableRow key={item.id}>
            <TableCell >{item.id}</TableCell>
            <TableCell >{item.user}</TableCell>
            <TableCell >
            <Button variant="outlined" color='primary'
            size='sm'>Edit</Button></TableCell>
            <TableCell ><Button variant="outlined" color='warning'
            size='sm'>Delete</Button></TableCell>
          </TableRow>))} 
        
       
        </TableBody>
      </Table>
    </TableContainer>}
    </div>
  );
}

export default App;
