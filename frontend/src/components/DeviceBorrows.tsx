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
import { DeviceBorrowInterface } from "../models/IDeviceBorrow";
import { format } from 'date-fns';

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

function DeviceBorrows() {
  const classes = useStyles();
  const [deviceborrows, setDeviceBorrows] = useState<DeviceBorrowInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getDeviceBorrows = async () => {
    fetch(`${apiUrl}/deviceborrows`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setDeviceBorrows(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getDeviceBorrows();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              style={{ color: '#13c2c2' }}
              gutterBottom
            >
              ข้อมูลการยืมอุปกรณ์
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/deviceborrow/create"
              variant="contained"
              style={{ backgroundColor: '#13c2c2', color: '#ffffff' }}
            >
              ยืมอุปกรณ์
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="10%">
                  รหัสยืม
                </TableCell>
                <TableCell align="center" width="20%">
                  ชื่อผู้ยืม
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่ออุปกรณ์
                </TableCell>
                <TableCell align="center" width="15%">
                  ประเภทอุปกรณ์
                </TableCell>
                <TableCell align="center" width="10%">
                  จำนวน
                </TableCell>
                <TableCell align="center" width="20%">
                  วันที่ยืม
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {deviceborrows.map((item: DeviceBorrowInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.BorrowCode}</TableCell>
                  <TableCell align="center">{item.Member.Name}</TableCell> 
                  <TableCell align="center">{item.DeviceName}</TableCell>
                  <TableCell align="center">{item.DeviceType.Type}</TableCell>
                  <TableCell align="center">{item.Amount}</TableCell>    
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

export default DeviceBorrows;