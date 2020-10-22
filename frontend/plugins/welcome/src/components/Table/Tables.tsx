import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntRent } from '../../api/models/EntRent';
/*
import { Link as RouterLink } from 'react-router-dom';

import Button from '@material-ui/core/Button';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
 } from '@backstage/core';
 */
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 


export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Rents, setRents] = useState<EntRent[]>(Array);
 //const [User, setUser] = useState(Array);



 const [loading, setLoading] = useState(true);
 
 const getRents = async () => {
    const res = await api.listRent({ limit: 10, offset: 0 });
  
  setLoading(false);
  setRents(res);
};

console.log(Rents)


 useEffect(() => {
   
   getRents();



   
 }, [loading]);
 
 
 const deleteRentss = async (id: number) => {
  const res = await api.deleteRent({ id: id });
  setLoading(true);
};

 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">ชื่อผู้พักอาศัย</TableCell>
           <TableCell align="center">เลขห้อง</TableCell>
           <TableCell align="center">ระยะเวลาในสัญญา</TableCell>
           <TableCell align="center">วันทำสัญญา</TableCell>
           <TableCell align="center">Manage</TableCell>
           
         </TableRow>
       </TableHead>
       <TableBody>
       {Rents === undefined 
          ? null
          : Rents.map((item :any)=> (
           <TableRow>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.rentUser.nAME}</TableCell>
             <TableCell align="center">{item.edges.rentRoom.id}</TableCell>
             
             <TableCell align="center">{item.rentAge} {item.edges.rentRoomage.text}</TableCell>
             <TableCell align="center">{item.rentDate}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteRentss(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>

               
               

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
