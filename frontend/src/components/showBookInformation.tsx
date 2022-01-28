import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";

import Typography from "@material-ui/core/Typography";

import Button from "@material-ui/core/Button";

import Container from "@material-ui/core/Container";

import Paper from "@material-ui/core/Paper";

import Box from "@material-ui/core/Box";

import Table from "@material-ui/core/Table";

import TableBody from "@material-ui/core/TableBody";

import TableCell from "@material-ui/core/TableCell";

import TableContainer from "@material-ui/core/TableContainer";

import TableHead from "@material-ui/core/TableHead";

import TableRow from "@material-ui/core/TableRow";

import { BookInformationInterface } from "../models/IBookInformation";

import { format } from 'date-fns'


//import moment from 'moment';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   container: {marginTop: theme.spacing(2)},
   table: { minWidth: 650},
   tableSpace: {marginTop: 20},
 })
);

function BookInformations() {
 const classes = useStyles();
 const [bookInformations, setBookInformations] = React.useState<BookInformationInterface[]>([]);
 const apiUrl = "http://localhost:8080";
 const requestOptions = {
  method: "GET",
  headers: { 
    Authorization: `Bearer ${localStorage.getItem("token")}`,
  "Content-Type": "application/json", },
};

 const getStatusRepairs = async () => {
      fetch(`${apiUrl}/book_informations`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setBookInformations(res.data);
        } else {
          console.log("else");
       }
     });
 };
 useEffect(() => {
  getStatusRepairs();
 }, []);

 

 return (
   <div>
     <Container className={classes.container} maxWidth="md">
       <Box display="flex">
         <Box flexGrow={1}>
           <Typography
             component="h2"
             variant="h4"
             color="primary"
             gutterBottom
           >
             ข้อมูลหนังสือ
           </Typography>
         </Box>
         <Box>
           <Button
             component={RouterLink}
             to="/book_informations/create"
             variant="contained"
             color="primary"
           >
             เพิ่มข้อมูล
           </Button>
         </Box>
       </Box>
       <TableContainer component={Paper} className={classes.tableSpace}>
         <Table className={classes.table} aria-label="simple table">
           <TableHead>
             <TableRow>
               <TableCell align="center" width="5%">
                 ID
               </TableCell>
               <TableCell align="center" width="15%">
                  เลขเรียกหนังสือ
               </TableCell>
               <TableCell align="center" width="25%">
                 ชื่อเรื่อง
               </TableCell>
               <TableCell align="center" width="10%">
                 ประเภท
               </TableCell>
               <TableCell align="center" width="10%">
                 ที่จัดเก็บ
               </TableCell>
               <TableCell align="center" width="10%">
                 ปีที่พิมพ์
               </TableCell>
               <TableCell align="center" width="30%">
                 วันที่อัพเดท
              </TableCell>
             </TableRow>
           </TableHead>
           <TableBody>
             {bookInformations.map((item: BookInformationInterface) => (
               <TableRow key={item.ID}>
                 <TableCell align="center">{item.ID}</TableCell>
                 <TableCell align="center">{item.CallNumber}</TableCell>
                 <TableCell align="center">{item.BookOrder.BookTitle}</TableCell>
                 <TableCell align="center">{item.BookType.Type}</TableCell>
                 <TableCell align="center">{item.BookLocation.Location}</TableCell>
                 <TableCell align="center" size="medium">
                   {item.YearPublication}
                 </TableCell> 
                 <TableCell align="center">{format((new Date(item.Date)), 'dd MMMM yyyy')}</TableCell>
               </TableRow>
             ))}
           </TableBody>
         </Table>
       </TableContainer>
     </Container>
   </div>
 );

}

export default BookInformations;