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
import { BookingRoomInterface } from "../models/IBookingRoom";
import { RoomAndTimeInterface } from "../models/IRoomAndTime";
import { RoomObjectiveInterface } from "../models/IRoomObjective";
import { RoomTypeInterface } from "../models/IRoomType";
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
 
function BookingRoomCreate() {
 const classes = useStyles();
 const [selectedDate, setSelectedDate] = React.useState<Date | null>(new Date());
 const [bookingroom, setBookingroom] = useState<Partial<BookingRoomInterface>>({});
 const [roomandtimes, setRoomAndTimes] = useState<RoomAndTimeInterface[]>([]);
 const [roomtypes, setRoomTypes] = useState<RoomTypeInterface[]>([]);
 const [roomobjectives, setRoomObjectives] = useState<RoomObjectiveInterface[]>([]);
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
   const id = event.target.id as keyof typeof bookingroom;
   const { value } = event.target;
   setBookingroom({ ...bookingroom, [id]: value });
 };

 const handleChange = (
  event: React.ChangeEvent<{ name?: string; value: unknown }>
) => {
  const name = event.target.name as keyof typeof bookingroom;
  setBookingroom({ 
    ...bookingroom, 
    [name]: event.target.value, 
   });
};

 const getRoomAndTimes = async () => {
  fetch(`${apiUrl}/roomandtimes`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setRoomAndTimes(res.data);
    } else {
      console.log("else");
    }
  })
};

const getRoomTypes = async () => {
  fetch(`${apiUrl}/roomtypes`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setRoomTypes(res.data);
    } else {
      console.log("else");
    }
  })
};

const getRoomObjectives = async () => {
  fetch(`${apiUrl}/roomobjectives`, requestOptions)
  .then((response) => response.json())
  .then((res) => {
    if (res.data) {
      setRoomObjectives(res.data);
    } else {
      console.log("else");
    }
  })
};

useEffect(() => {
  getRoomAndTimes();
  getRoomTypes();
  getRoomObjectives();
}, []);

const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
};
 
 function submit() {
   let data = {
     PhoneBooker: bookingroom.PhoneBooker ?? "",
     QuantityMember: typeof bookingroom.QuantityMember === "string" ? parseInt(bookingroom.QuantityMember) : 0,
     BookingRoomAt: selectedDate, 
     RoomAndTimeID: convertType(bookingroom.RoomAndTimeID),
     RoomTypeID: convertType(bookingroom.RoomTypeID),
     RoomObjectiveID: convertType(bookingroom.RoomObjectiveID),
   };
 

   const requestOptionsPost = {
    method: "POST",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
     },
    body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/bookingrooms`, requestOptionsPost)
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
            สร้าง การจองห้องติว
          </Typography>
        </Box>
      </Box>
      <Divider />


       <Grid container spacing={3} className={classes.root}>


       <Grid item xs={6}>
             <p>ห้องและเวลา</p>
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookingroom.RoomAndTimeID}
                  onChange={handleChange}
                  inputProps={{
                    name: "RoomAndTimeID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกห้องและเวลา
                  </option>
                  {roomandtimes.map((item: RoomAndTimeInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={6}>
             <p>ประเภทของห้อง</p>
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookingroom.RoomTypeID}
                  onChange={handleChange}
                  inputProps={{
                    name: "RoomTypeID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกประเภทของห้อง
                  </option>
                  {roomtypes.map((item: RoomTypeInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={6}>
             <p>จุดประสงค์การใช้ห้อง</p>
              <FormControl fullWidth variant="outlined">
                <Select
                  native
                  value={bookingroom.RoomObjectiveID}
                  onChange={handleChange}
                  inputProps={{
                    name: "RoomObjectiveID",
                  }}
                  >
                  <option aria-label="None" value="">
                    กรุณาเลือกจุดประสงค์การใช้ห้อง
                  </option>
                  {roomobjectives.map((item: RoomObjectiveInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>



         <Grid item xs={8}>
           <p>เบอร์โทรศัพท์ผู้จอง</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="PhoneBooker"
               variant="outlined"
               type="string"
               size="medium"
               value={bookingroom.PhoneBooker || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

         

         <Grid item xs={8}>
           <FormControl fullWidth variant="outlined">
             <p>จำนวน</p>
             <TextField
               id="QuantityMember"
               variant="outlined"
               type="number"
               size="medium"
               value={bookingroom.QuantityMember || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>


         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 id="BookingRoomAt"
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
           <Button component={RouterLink} to="/booking_rooms" variant="contained">
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
 
export default BookingRoomCreate;