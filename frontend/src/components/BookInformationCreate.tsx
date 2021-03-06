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

//import MenuItem from '@material-ui/core/MenuItem';

import TextField from "@material-ui/core/TextField";

import Select from '@material-ui/core/Select';

import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";

import DateFnsUtils from "@date-io/date-fns";

import { BookInformationInterface } from "../models/IBookInformation";

import { BookOrderInterface } from "../models/IBookOrder";

import { BookLocationInterface } from "../models/IBookLocation";

import { BookCategoryInterface } from "../models/IBookCategory";

import {MembersInterface} from "../models/IMember";


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

 const [bookCategorys, setBookCategorys] = React.useState<BookCategoryInterface[]>([]);
 const [bookLocations, setBookLocations] = React.useState<BookLocationInterface[]>([]);
 const [bookOrders, setBookOrders] = React.useState<BookOrderInterface[]>([]);
 const [bookInformation, setBookInformation] = useState<Partial<BookInformationInterface>>({});
 const [members,setMembers] = React.useState<MembersInterface>()
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const [errorMessage, setErrorMessage] = useState("");

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

 const getBookCategorys = async () => {
  fetch(`${apiUrl}/book_categories`, requestOptions)
    .then((response) => response.json())
      .then((res) => {
        if (res.data) {
            setBookCategorys(res.data);
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

const geMembers = async () => {
  let uid = localStorage.getItem("uid");
  fetch(`${apiUrl}/member/${uid}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      bookInformation.MemberID = res.data.ID
      if (res.data) {
        setMembers(res.data);
      } else {
        console.log("else");
      }
    });
};

    useEffect(() => {
        getBookCategorys();
        getBookLocations();
        getBookOrder();
        geMembers();
    }, []);


 const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
};

 function submit() {
   let data = {
    BookCategoryID: convertType(bookInformation.BookCategoryID),
    BookLocationID: convertType(bookInformation.BookLocationID),
    BookOrderID: convertType(bookInformation.BookOrderID),
    MemberID: convertType(bookInformation.MemberID),
    Date: selectedDate,
    CallNumber: bookInformation.CallNumber ?? "",
    YearPublication: typeof bookInformation.YearPublication === "string" ? parseInt(bookInformation.YearPublication) : 0,

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
        console.log("???????????????????????????")
         setSuccess(true);
         setErrorMessage("")
       } else {
        console.log("????????????????????????????????????")
         setError(true);
         setErrorMessage(res.error)
       }
     }); 
     
 }

 return (

   <Container className={classes.container} maxWidth="md">
     <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="success">
         ??????????????????????????????????????????????????????
       </Alert>
     </Snackbar>
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
         ???????????????????????????????????????????????????????????????: {errorMessage}
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
             ???????????????????????????????????????
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={4} className={classes.root}>
        <Grid item xs={12}>
          <p>?????????????????????????????????</p>
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
                  ???????????????????????????????????????????????????????????????
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
          <p>??????????????????????????????????????????</p>
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
                  ????????????????????????????????????????????????????????????????????????
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
          <p>?????????????????????????????????????????????</p>
            <FormControl fullWidth variant="outlined">
                <Select
                native
                value={bookInformation.BookCategoryID}
                onChange={handleChange}
                inputProps={{
                  name: "BookCategoryID",
                }}
              >
              <option aria-label="None" value="">
                  ???????????????????????????????????????????????????????????????????????????
                </option>
                {bookCategorys.map((item: BookCategoryInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Category}
                  </option>
                ))}
              </Select>
            </FormControl>
        </Grid>
        <Grid item xs={6}>
           <p>?????????????????????????????????????????????</p>
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
             <p>??????????????????????????????</p>
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
             <p>????????????????????????????????????</p>
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
          <p>???????????????????????????</p>
            <FormControl fullWidth variant="outlined">
                <Select
                native
                value={bookInformation.MemberID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "MemberID",
                }}
              >
                <option value={members?.ID} key={members?.ID}>
                    {members?.Name}
                </option>
              </Select>
            </FormControl>
        </Grid>

        <Grid item xs={12}>
           <Button component={RouterLink} 
           to="/book_informations" 
           variant="contained" 
           size="large">
             ????????????
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             size="large"
             color="primary"
           >
             ????????????????????????????????????
           </Button>
         </Grid>
        </Grid>
     </Paper>
   </Container>
 );
}
export default BookInformationCreate;