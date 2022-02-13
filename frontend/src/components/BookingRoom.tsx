import React from 'react'
import { useEffect, useState } from "react";
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
import { BookingRoomInterface } from "../models/IBookingRoom";
import { IconButton, Snackbar } from "@material-ui/core";
import DeleteIcon from '@material-ui/icons/Delete';
import { format } from 'date-fns'
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import moment from "moment";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function BookingRoom() {
  const classes = useStyles();
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [ErrorMessage, setErrorMessage] = React.useState("");
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {

  if (reason === "clickaway") {
    return;
  }
  setSuccess(false);
  setError(false);
};

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

  const [Evet, setEvent] = useState<BookingRoomInterface[]>([]);
  const apiUrl = "http://localhost:8080/bookingrooms";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
              "Content-Type": "application/json",
    },
  };

  const getEvent = async () => {
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setEvent(res.data);
        } else {
          console.log("else");
        }
      });
  };

  


  const DeleteBookingRoom = async (id: string | number | undefined) => {
    const apiUrl = "http://localhost:8080/bookingrooms";
    const requestOptions = {
      method: "DELETE",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",},
    };
  
    fetch(`${apiUrl}/${id}`, requestOptions)
    .then((response) => response.json())
    .then(
      (res) => {
        if (res.data) {
          setSuccess(true)
          setErrorMessage("")
        } 
        else { 
          setErrorMessage(res.error)
          setError(true)
        }  
        getEvent(); 
      }
    )
  }


  useEffect(() => {
    getEvent();
  }, []);

  

  return (
    <div>
      <Container className={classes.container} maxWidth = "lg" >
        <Box display="flex">
          <Box flexGrow={1}>

          <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
              <Alert onClose={handleClose} severity="success">
                ลบข้อมูลสำเร็จ
              </Alert>
            </Snackbar>

            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
              <Alert onClose={handleClose} severity="error">
                {ErrorMessage}
              </Alert>
            </Snackbar>
            <br/><br/> 


            <Typography
              component="h2"
              variant="h4"
              color="primary"
              gutterBottom
            >
              ข้อมูลการจองห้องติว
            </Typography>
            </Box>
         <Box>
           <Button
             component={RouterLink}
             to="/booking_roomscreate"
             
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
                  ผู้จอง
                </TableCell>
                <TableCell align="center" width="15%">
                  ห้องและเวลา
                </TableCell>
                <TableCell align="center" width="10%">
                  ประเภทของห้อง
                </TableCell>
                <TableCell align="center" width="10%">
                  จุดประสงค์การใช้ห้อง
                </TableCell>
                <TableCell align="center" width="5%">
                  เบอร์โทรศัพท์ผู้จอง
                </TableCell>
                <TableCell align="center" width="5%">
                  จำนวน
                </TableCell>
                <TableCell align="center" width="15%">
                  วันที่จอง
                </TableCell>
                <TableCell align="center" width="15%">
                  ลบ
                </TableCell>

              </TableRow>
            </TableHead>
            
            <TableBody>
              {Evet.map((item: BookingRoomInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Member.Name}</TableCell>
                  <TableCell align="center">{item.RoomAndTime.Name}</TableCell>
                  <TableCell align="center">{item.RoomType.Name}</TableCell>
                  <TableCell align="center">{item.RoomObjective.Name}</TableCell>
                  <TableCell align="center">{item.PhoneBooker}</TableCell>
                  <TableCell align="center">{item.QuantityMember}</TableCell>
                  {/* //THH:mm */}
                  <TableCell align="center">{moment(item.BookingRoomAt).format("YYYY-MM-DD")}</TableCell> 
                  <TableCell align="center"> 
                  
                  <IconButton color="secondary" onClick={() => DeleteBookingRoom(item.ID)} > <DeleteIcon /> </IconButton>

                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default BookingRoom;