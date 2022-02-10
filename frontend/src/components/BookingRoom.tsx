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
import { format } from 'date-fns'
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

  useEffect(() => {
    getEvent();
  }, []);

  

  return (
    <div>
      <Container className={classes.container} maxWidth = "lg" >
        <Box display="flex">
          <Box flexGrow={1}>
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

              </TableRow>
            </TableHead>
            
            <TableBody>
              {Evet.map((item: BookingRoomInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.RoomAndTime.Name}</TableCell>
                  <TableCell align="center">{item.RoomType.Name}</TableCell>
                  <TableCell align="center">{item.RoomObjective.Name}</TableCell>
                  <TableCell align="center">{item.PhoneBooker}</TableCell>
                  <TableCell align="center">{item.QuantityMember}</TableCell>
                  {/* //THH:mm */}
                  <TableCell align="center">{moment(item.BookingRoomAt).format("YYYY-MM-DD")}</TableCell> 
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