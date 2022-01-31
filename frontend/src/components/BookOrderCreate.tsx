import React from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { BookOrderInterface } from "../models/IBookOrder";
import { BookTypeInterface } from "../models/IBookType";
import {CompaniesInterface} from "../models/ICompany";
import {OrderStatusesInterface} from "../models/IOrderStatus";
import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { useEffect, useState } from "react";
import { Select } from "@material-ui/core";
 
function Alert(props: AlertProps) {
 return <MuiAlert elevation={6} variant="filled" {...props} />;
}
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {flexGrow: 1},
   container: {marginTop: theme.spacing(2)},
   paper: {padding: theme.spacing(2),color: theme.palette.text.secondary},
})
);
 
function BookOrderCreate() {
 const classes = useStyles();
 const [selectedDate, setSelectedDate] = React.useState<Date | null>(new Date());
 const [bookorder, setBookOrder] = useState<Partial<BookOrderInterface>>({});
 const [booktypes, setBookTypes] = useState<BookTypeInterface[]>([]);
 const [companies, setCompanies] = useState<CompaniesInterface[]>([]);
 const [orderrstatuses, setOrderStatuses] = useState<OrderStatusesInterface[]>([]);
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);

 const apiUrl = "http://localhost:8080";
 const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" },
 };
 
 const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };
 
 const handleDateChange = (date: Date | null) => {
   setSelectedDate(date);
 };
 
 const handleInputChange = (
   event: React.ChangeEvent<{ id?: string; value: any }>
 ) => {
   const id = event.target.id as keyof typeof bookorder;
   const { value } = event.target;
   setBookOrder({ ...bookorder, [id]: value });
 };

 const handleChange = (
  event: React.ChangeEvent<{ name?: string; value: unknown }>
) => {
  const name = event.target.name as keyof typeof bookorder;
  setBookOrder({ 
    ...bookorder, 
    [name]: event.target.value, 
   });
};

 const getBookType = async () => {
  fetch(`${apiUrl}/book_types`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setBookTypes(res.data);
    } else {
      console.log("else");
    }
  })
};

const getCompany = async () => {
  fetch(`${apiUrl}/companies`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setCompanies(res.data);
    } else {
      console.log("else");
    }
  })
};

const getOrderStatus = async () => {
  fetch(`${apiUrl}/order_statuses`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setOrderStatuses(res.data);
    } else {
      console.log("else");
    }
  })
};

useEffect(() => {
  getBookType();
  getCompany();
  getOrderStatus();
}, []);

const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
};
 
 function submit() {
   let data = {
     BookTitle: bookorder.BookTitle ?? "",
     Author: bookorder.Author ?? "",
     OrderAmount: typeof bookorder.OrderAmount === "string" ? parseInt(bookorder.OrderAmount) : 0,
     Price: typeof bookorder.Price === "string" ? parseFloat(bookorder.Price) : 0,
     OrderDate: selectedDate, 
     BookTypeID: convertType(bookorder.BookTypeID),
     CompanyID: convertType(bookorder.CompanyID),
     OrderStatusID: convertType(bookorder.OrderStatusID),
   };
 

   const requestOptionsPost = {
    method: "POST",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
     },
    body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/book_orders`, requestOptionsPost)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        console.log("บันทึกได้")
        setSuccess(true);
      } else {
        console.log("บันทึกไม่ได้")
        setError(true);
      }
    });
}

return (
  <Container className={classes.container} maxWidth="md">
    <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
      <Alert onClose={handleClose} severity="success">
        บันทึกข้อมูลสำเร็จ
      </Alert>
    </Snackbar>
    <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
      <Alert onClose={handleClose} severity="error">
        บันทึกข้อมูลไม่สำเร็จ
      </Alert>
    </Snackbar>
    <Paper className={classes.paper}>
      <Box display="flex">
        <Box flexGrow={1}>
          <Typography
            component="h2"
            variant="h6"
            color="primary"
            gutterBottom
          >
            สร้าง BookOrder
          </Typography>
        </Box>
      </Box>
      <Divider />
       <Grid container spacing={3} className={classes.root}>
         <Grid item xs={8}>
           <p>ชื่อเรื่องหนังสือ</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="BookTitle"
               variant="outlined"
               type="string"
               size="medium"
               value={bookorder.BookTitle || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={8}>
           <FormControl fullWidth variant="outlined">
             <p>ผู้แต่ง</p>
             <TextField
               id="Author"
               variant="outlined"
               type="string"
               size="medium"
               value={bookorder.Author || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
             <p>ประเภทหนังสือ</p>
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookorder.BookTypeID}
                  onChange={handleChange}
                  inputProps={{
                    name: "BookTypeID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกประเภทหนังสือ
                  </option>
                  {booktypes.map((item: BookTypeInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Type}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
             <p>บริษัท</p>
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookorder.CompanyID}
                  onChange={handleChange}
                  inputProps={{
                    name: "CompanyID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกบริษัทที่ต้องการสั่งซื้อ
                  </option>
                  {companies.map((item: CompaniesInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.NameThai}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>จำนวน</p>
             <TextField
               id="OrderAmount"
               variant="outlined"
               type="string"
               size="medium"
               value={bookorder.OrderAmount || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
         <FormControl fullWidth variant="outlined">
             <p>ราคา</p>
             <TextField
               id="Price"
               variant="outlined"
               type="float"
               size="medium"
               value={bookorder.Price || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
             <p>สถานะ</p>
            
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookorder.OrderStatusID}
                  onChange={handleChange}
                  inputProps={{
                    name: "OrderStatusID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกสถานะการสั่งซื้อ
                  </option>
                  {orderrstatuses.map((item: OrderStatusesInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Status}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 id="OrderDate"
                 format="dd-MM-yyyy"
                 value={selectedDate}
                 onChange={handleDateChange}
                 KeyboardButtonProps={{
                   "aria-label": "change date",
                 }}
               />
             </MuiPickersUtilsProvider>
           </FormControl>
         </Grid>
         <Grid item xs={12}>
           <Button component={RouterLink} to="/" variant="contained">
             กลับ
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
             >
             บันทึก
           </Button>
         </Grid>
       </Grid>
     </Paper>
   </Container>
 );
}
 
export default BookOrderCreate;
