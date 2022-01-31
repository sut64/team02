import React from "react";

import { useEffect, useState } from "react";

import { Link as RouterLink } from "react-router-dom";

import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";

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

import TextField from "@material-ui/core/TextField";

import Select from '@material-ui/core/Select';

import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";

import DateFnsUtils from "@date-io/date-fns";

import { BookInformationInterface } from "../models/IBookInformation";

import { BookOrderInterface } from "../models/IBookOrder";

import { BookLocationInterface } from "../models/IBookLocation";

import { BookTypeInterface } from "../models/IBookType";


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


function BookInformationCreate() {
 const classes = useStyles();
 const [selectedDate, setSelectedDate] = React.useState<Date | null>(
  new Date()
);

 const [bookTypes, setBookTypes] = React.useState<BookTypeInterface[]>([]);
 const [bookLocations, setBookLocations] = React.useState<BookLocationInterface[]>([]);
 const [bookOrders, setBookOrders] = React.useState<BookOrderInterface[]>([]);
 const [bookInformation, setBookInformation] = useState<Partial<BookInformationInterface>>({});
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);

 const apiUrl = "http://localhost:8080";
 const requestOptions = {
  method: "GET",
  headers: { 
    Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
   },
};

 const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };

 const handleDateChange = (date: Date | null) => {
  console.log(date);
  setSelectedDate(date);
};
const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof bookInformation;
    setBookInformation({ 
      ...bookInformation, 
      [name]: event.target.value, 
     });
  };
 
  const handleInputChange = (
   event: React.ChangeEvent<{ id?: string; value: any }>
 ) => {
   const id = event.target.id as keyof typeof bookInformation;
   const { value } = event.target;
   setBookInformation({ ...bookInformation, [id]: value });
 };

 const getBookTypes = async () => {
  fetch(`${apiUrl}/book_types`, requestOptions)
    .then((response) => response.json())
      .then((res) => {
        if (res.data) {
            setBookTypes(res.data);
        } else {
          console.log("else");
        }
      });
  };

const getBookLocations = async () => {
  fetch(`${apiUrl}/book_locations`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setBookLocations(res.data);
      } else {
        console.log("else");
      }
    });
};

const getBookOrder = async () => {
  fetch(`${apiUrl}/book_orders`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setBookOrders(res.data);
      } else {
        console.log("else");
      }
    });
};

    useEffect(() => {
        getBookTypes();
        getBookLocations();
        getBookOrder();
    }, []);


 const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
};

 function submit() {
   let data = {
    BookTypeID: convertType(bookInformation.BookTypeID),
    BookLocationID: convertType(bookInformation.BookLocationID),
    BookOrderID: convertType(bookInformation.BookOrderID),
    Date: selectedDate,
    CallNumber: bookInformation.CallNumber ?? "",
    YearPublication: typeof bookInformation.YearPublication === "string" ? parseInt(bookInformation.YearPublication) : 0,
    //YearPublication: bookInformation.YearPublication ?? "",

   };

   console.log(data)
   
   const requestOptionsPost = {
     method: "POST",
     headers: {  
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json", },
     body: JSON.stringify(data),
   };

   fetch(`${apiUrl}/book_informations`, requestOptionsPost)
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
             variant="h4"
             color="primary"
             gutterBottom
           >
             ข้อมูลหนังสือ
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={4} className={classes.root}>
        <Grid item xs={12}>
          <p>ชื่อหนังสือ</p>
            <FormControl fullWidth variant="outlined">
              <Select
                native
                value={bookInformation.BookOrderID}
                onChange={handleChange}
                inputProps={{
                  name: "BookOrderID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกชื่อหนังสือ
                </option>
                {bookOrders.map((item: BookOrderInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.BookTitle}
                  </option>
                ))}
              </Select>
            </FormControl>
        </Grid>
        <Grid item xs={12}>
          <p>สถานที่จัดเก็บ</p>
            <FormControl fullWidth variant="outlined">
                <Select
                native
                value={bookInformation.BookLocationID}
                onChange={handleChange}
                inputProps={{
                  name: "BookLocationID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกสถานที่จัดเก็บ
                </option>
                {bookLocations.map((item: BookLocationInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Location}
                </option>
              ))}
              </Select>
            </FormControl>
        </Grid>
        <Grid item xs={12}>
          <p>ประเภทหนังสือ</p>
            <FormControl fullWidth variant="outlined">
                <Select
                native
                value={bookInformation.BookTypeID}
                onChange={handleChange}
                inputProps={{
                  name: "BookTypeID",
                }}
              >
              <option aria-label="None" value="">
                  กรุณาเลือกประเภทหนังสือ
                </option>
                {bookTypes.map((item: BookTypeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Type}
                  </option>
                ))}
              </Select>
            </FormControl>
        </Grid>
        <Grid item xs={6}>
           <p>เลขเรียกหนังสือ</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="CallNumber"
               variant="outlined"
               type="string"
               size="medium"
               value={bookInformation.CallNumber || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>ปีที่พิมพ์</p>
             <TextField
               id="YearPublication"
               variant="outlined"
               type="uint"
               size="medium"
               value={bookInformation.YearPublication || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
        <Grid item xs={12}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่อัพเดท</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 minDate = {new Date("2018-01-01")}
                 format="yyyy-MM-dd"
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
           <Button component={RouterLink} 
           to="/book_informations" 
           variant="contained" 
           size="large">
             กลับ
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             size="large"
             color="primary"
           >
             บันทึกข้อมูล
           </Button>
         </Grid>
        </Grid>
     </Paper>
   </Container>
 );
}
export default BookInformationCreate;