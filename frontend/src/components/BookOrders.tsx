import React, { useEffect,useState } from "react";
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
import { BookOrderInterface } from "../models/IBookOrder";
 
import moment from 'moment';
 
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   container: {marginTop: theme.spacing(3)},
   table: { minWidth: 1200},
   tableSpace: {marginTop: 40},
})
);
 
function BookOrders() {
 const classes = useStyles();
 const [bookorders, setBookOrders] = useState<BookOrderInterface[]>([]);
 
 const getBookOrders = async () => {
   const apiUrl = "http://localhost:8080";
   const requestOptions = {
     method: "GET",
     headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`, 
      "Content-Type": "application/json" },
   };
 
   fetch(`${apiUrl}/book_orders`, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
         setBookOrders(res.data);
       } else {
         console.log("else");
       }
     });
 };
 
 useEffect(() => {
   getBookOrders();
 }, []);
 
 return (
   <div>
     <Container className={classes.container}>
       <Box display="flex">
         <Box flexGrow={1}>
           <Typography
             component="h2"
             variant="h4"
             color="primary"
             gutterBottom
           >
             BOOKORDERS
           </Typography>
         </Box>
         <Box>
           <Button
  component={RouterLink}
  to="/bookordercreate"
  variant="contained"
  color="primary"
>
  สร้าง BookOrder
</Button>
</Box>
</Box>
<TableContainer component={Paper} className={classes.tableSpace}>
<Table width="10%" >
<TableHead>
  <TableRow>
    <TableCell align="center" width="2%">
      ลำดับ
    </TableCell>
    <TableCell align="center" width="17%">
      ชื่อเรื่องหนังสือ
    </TableCell>
    <TableCell align="center" width="10%">
      ผู้แต่ง
    </TableCell>
    <TableCell align="center" width="15%">
      ประเภทหนังสือ
    </TableCell>
    <TableCell align="center" width="15%">
      บริษัท
    </TableCell>
    <TableCell align="center" width="10%">
      จำนวน
    </TableCell>
    <TableCell align="center" width="10%">
      ราคา
    </TableCell>
    <TableCell align="center" width="10%">
      สถานะ
    </TableCell>
    <TableCell align="center" width="10%">
      วันที่
    </TableCell>
  </TableRow>
</TableHead>
<TableBody>
  {bookorders.map((item: BookOrderInterface) => (
    <TableRow key={item.ID}>
      <TableCell align="right">{item.ID}</TableCell>
      <TableCell align="left" >{item.BookTitle}</TableCell>
      <TableCell align="left">{item.Author}</TableCell>
      <TableCell align="center">{item.BookType.Type}</TableCell>
      <TableCell align="left">{item.Company.NameThai}</TableCell>
      <TableCell align="center">{item.OrderAmount}</TableCell>
      <TableCell align="center">{item.Price}</TableCell>
      <TableCell align="center">{item.OrderStatus.Status}</TableCell>
      <TableCell align="center">{moment(item.OrderDate).format("DD/MM/YYYY")}</TableCell>
    </TableRow>
   ))}
   </TableBody>
 </Table>
</TableContainer>
</Container>
</div>
);
}

export default BookOrders;

