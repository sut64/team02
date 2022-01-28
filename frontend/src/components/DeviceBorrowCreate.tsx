import React from "react";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles, } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import TextField from "@material-ui/core/TextField";
import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";
import DateFnsUtils from '@date-io/date-fns';

import { DeviceBorrowInterface } from "../models/IDeviceBorrow";

import { MembersInterface } from "../models/IMember";
import { DeviceListsInterface } from "../models/IDeviceList";
import { DeviceTypesInterface } from "../models/IDeviceType";

const Alert = (props: AlertProps) => {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function DeviceBorrowCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(new Date());
  const [members, setMembers] = useState<MembersInterface>();
  const [devicelists, setDeviceLists] = useState<DeviceListsInterface[]>([]);
  const [devicetypes, setDeviceTypes] = useState<DeviceTypesInterface[]>([]);
  const [deviceborrow, setDeviceBorrow] = useState<Partial<DeviceBorrowInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

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

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof deviceborrow;
    const { value } = event.target;
    setDeviceBorrow({
      ...deviceborrow,
      [id]: value });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof deviceborrow;
    setDeviceBorrow({
      ...deviceborrow,
      [name]: event.target.value,
    });
  };

  const getMembers = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/member/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
          deviceborrow.MemberID = res.data.ID
        if (res.data) {
          setMembers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDeviceLists = async () => {
    fetch(`${apiUrl}/devicelists`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDeviceLists(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDeviceTypes = async () => {
    fetch(`${apiUrl}/devicetypes`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDeviceTypes(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getMembers();
    getDeviceLists();
    getDeviceTypes();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
    
  };

  function submit() {
    let data = {
      MemberID: convertType(deviceborrow.MemberID),
      DeviceListID: convertType(deviceborrow.DeviceListID),
      DeviceTypeID: convertType(deviceborrow.DeviceTypeID),
      DeviceName: deviceborrow.DeviceName,
      BorrowCode: deviceborrow.BorrowCode,
      Amount: typeof deviceborrow.Amount === "string" ? parseInt(deviceborrow.Amount) : 0,
      Date: selectedDate,
    };

    const requestOptionsPost = {
      method: "POST",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/deviceborrows`, requestOptionsPost)
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
              บันทึกข้อมูลการยืมอุปกรณ์
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={1} className={classes.root}>

        <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>ผู้บันทึกการยืม</p>
              <Select
                native
                value={deviceborrow.MemberID}
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

          <Grid item xs={6}>
            <p>BorrowCode</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="BorrowCode"
                variant="outlined"
                type="ิstring"
                size="medium"
                placeholder="กรุณากรอกรหัสยืม"
                value={deviceborrow.BorrowCode || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>DeviceCode</p>
              <Select
                native
                value={deviceborrow.DeviceListID}
                onChange={handleChange}
                inputProps={{
                  name: "DeviceListID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกรหัสอุปกรณ์
                </option>
                {devicelists.map((item: DeviceListsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.DeviceCode}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่ออุปกรณ์</p>
              <TextField
                id="DeviceName"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="กรุณากรอกชื่ออุปกรณ์"
                value={deviceborrow.DeviceName || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทอุปกรณ์</p>
              <Select
                native
                value={deviceborrow.DeviceTypeID}
                onChange={handleChange}
                inputProps={{
                  name: "DeviceTypeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทอุปกรณ์
                </option>
                {devicetypes.map((item: DeviceTypesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Type}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
	          <FormControl fullWidth variant="outlined">
	            <p>จำนวน</p>
		         <TextField
                id="Amount"
                variant="outlined"
                type="number"
                size="medium"
                placeholder="กรุณากรอกจำนวน"
                value={deviceborrow.Amount}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่ยืม</p>
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
            <Button
              component={RouterLink}
              to="/deviceborrows"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
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

export default DeviceBorrowCreate;