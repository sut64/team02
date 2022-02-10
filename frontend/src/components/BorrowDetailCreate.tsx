import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Box from "@material-ui/core/Box";
import { MembersInterface } from "../models/IMember";
import Snackbar from "@material-ui/core/Snackbar";
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import DateFnsUtils from '@date-io/date-fns';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
  DatePicker,
} from "@material-ui/pickers";
import { BookTypeInterface } from "../models/IBookType";
import { ServicePlacesInterface } from "../models/IServicePlace";
import { BorrowDetailInterface } from "../models/IBorrowDetail";
import { FormControl, ImageListItem } from "@material-ui/core";
import { StatusInterface } from "../models/IStatus";
import { BookOrderInterface } from "../models/IBookOrder";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    container: { marginTop: theme.spacing(2) },

    box: { marginTop: 20, marginBottom: 20 },

    table: { minWidth: 650 },

    tableSpace: { marginTop: 20 },

    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 200,
    },

  })

);



function CreateBorrowDetail() {

  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [members, setMembers] = useState<MembersInterface>();
  const [bookOrder, setBookOrder] = useState<BookOrderInterface[]>([]);
  const [status, setStatus] = useState<StatusInterface[]>([]);
  const [booktypes, setBookTypes] = useState<BookTypeInterface[]>([]);
  const [places, setPlaces] = useState<ServicePlacesInterface[]>([]);
  const [borrowDetail, setborrowDetail] = useState<Partial<BorrowDetailInterface>>({});


  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
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

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    console.log(event.target.value)
    const name = event.target.name as keyof typeof borrowDetail;
    setborrowDetail({
      ...borrowDetail,
      [name]: event.target.value,
    });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof borrowDetail;
    const { value } = event.target;
    setborrowDetail({ ...borrowDetail, [id]: value });
  };

  const handleType = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    getOrder(Number(event.target.value));
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getMember = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/member/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        borrowDetail.MemberID = res.data.ID
        if (res.data) {
          setMembers(res.data);
        } else {
          console.log("else");
        }
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
      });
  };

  const getOrder = async (id: number) => {
    fetch(`${apiUrl}/book_order/book_type/${id}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBookOrder(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getOrderList = async () => {
    fetch(`${apiUrl}/book_orders`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setBookOrder(res.data);
          console.log(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getPlaces = async () => {
    fetch(`${apiUrl}/places`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setPlaces(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getStatus = async () => {
    fetch(`${apiUrl}/statuses`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setStatus(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getMember();
    getPlaces();
    getBookType();
    getStatus();
    //getOrderList();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      MemberID: convertType(members?.ID),
      BookOrderID: convertType(borrowDetail.BookOrderID),
      ServicePlaceID: convertType(borrowDetail.ServicePlaceID),
      StatusID: 1,
      DateToBorrow: selectedDate,
      BorrowDuration: convertType(borrowDetail.BorrowDuration),
      Tel: borrowDetail.Tel ?? "",
    };
    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/borrows`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
          setErrorMessage(res.error)
        }
      });
  }

  return (

    <div>

      <Container className={classes.container} maxWidth="md">
        <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="success">
            บันทึกข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            บันทึกข้อมูลไม่สำเร็จ : {errorMessage}
          </Alert>
        </Snackbar>
        <Box display="flex" style={{ marginBottom: 55 }}>

          <Box flexGrow={1} >
          </Box>

          <Box display="flex" >
          </Box>
        </Box>

        <Box style={{ marginLeft: 25, marginTop: 5, marginBottom: 5 }} >
          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">ประเภท</Typography>
            </Grid>
            <Grid item xs={6}>
              <Select
                value={borrowDetail.BookOrder?.BookTypeID}
                onChange={handleType}
                style={{ width: 300 }}
                inputProps={{
                  name: "BookTypeID",
                }}
              >
                {booktypes.map((item: BookTypeInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Type}
                  </MenuItem>
                ))}
              </Select>
            </Grid>

            <Grid item xs={6}>
              <Typography align="center">ชื่อหนังสือ</Typography>
            </Grid>
            <Grid item xs={6}>
              <Select
                value={borrowDetail.BookOrderID}
                onChange={handleChange}
                style={{ width: 300 }}
                inputProps={{
                  name: "BookOrderID",
                }}
              >
                {bookOrder.map((item: BookOrderInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.BookTitle}
                  </MenuItem>
                ))}
              </Select>
            </Grid>
          </Grid>

          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">ข้อมูลผู้บันทึก</Typography>
            </Grid>
            <Grid item xs={6}>
              <Select
                native
                value={borrowDetail.MemberID}
                onChange={handleChange}
                style={{ width: 300 }}
                disabled
                inputProps={{
                  name: "MemberID",
                }}
              >
                <option aria-label="None" value="" >
                  {members?.Name}
                </option>

              </Select>
            </Grid>
          </Grid>


          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">สถานที่</Typography>
            </Grid>

            <Grid item xs={6}>
              <Select

                id="placeSelect"
                value={borrowDetail.ServicePlaceID}
                onChange={handleChange}
                style={{ width: 300 }}
                inputProps={{
                  name: "ServicePlaceID",
                }}
              >
                {places.map((item: ServicePlacesInterface) =>
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Name}
                  </MenuItem>
                )}
              </Select>
            </Grid>
          </Grid>



          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">วันที่จะยืม</Typography>
            </Grid>
            <Grid item xs={6}>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <DatePicker
                  name="DateToBorrow"
                  value={selectedDate}
                  onChange={handleDateChange}
                  style={{ width: 300 }}
                  minDate={new Date("2018-01-01T00:00")}
                  format="dd/MM/yyyy"
                />
              </MuiPickersUtilsProvider>
            </Grid>
          </Grid>

          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">เบอร์โทร</Typography>
            </Grid>
            <Grid item xs={6}>
              <TextField
                id="Tel"
                variant="standard"
                type="string"
                style={{ width: 300 }}
                value={borrowDetail.Tel || ""}
                onChange={handleInputChange}
              >
              </TextField>
            </Grid>
          </Grid>

          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">ระยะเวลาที่จะยืม</Typography>
            </Grid>
            <Grid item xs={6}>
              <TextField
                id="BorrowDuration"
                variant="standard"
                type="number"
                style={{ width: 300 }}
                value={borrowDetail.BorrowDuration}
                onChange={handleInputChange}
              >
              </TextField>
            </Grid>
          </Grid>

          <Grid container spacing={1}>
            <Grid item xs={6}>
              <Typography align="center">สถานะ</Typography>
            </Grid>
            <Grid item xs={6}>
              <Select
                id="status"
                value={1}
                onChange={handleChange}
                style={{ width: 300 }}
                disabled
                inputProps={{
                  name: "StatusID",
                }}
              >
                {status.map((item: StatusInterface) =>
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Name}
                  </MenuItem>
                )}
              </Select>
            </Grid>
          </Grid>



          <Box style={{ marginTop: 50 }} textAlign="center">
            <Button
              variant="contained"
              color="primary"
              size="large"
              onClick={submit}
              style={{ marginLeft: 100 }}
            >

              บันทึก
            </Button>

            <Button
              variant="contained"
              color="primary"
              size="large"
              component={RouterLink}
              to="/"
              style={{ marginLeft: 25 }}
            >

              ย้อนกลับ
            </Button>
          </Box>

        </Box>

      </Container>



    </div>

  );

}



export default CreateBorrowDetail;