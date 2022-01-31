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
import { BorrowDetailInterface } from "../models/IBorrowDetail";
import { format } from 'date-fns'
import { id } from "date-fns/locale";

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

function Borrows() {
  const classes = useStyles();
  const [borrows, setBorrows] = useState<BorrowDetailInterface[]>([]);


  const getBorrows = async () => {
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/borrow/member/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setBorrows(res.data);
        } else {
          console.log("else");
        }
      });
  };




  useEffect(() => {
    getBorrows();

  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการยืมหนังสือ
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                {/* <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell> */}
                <TableCell align="center" width="20%">
                  ชื่อหนังสือ
                </TableCell>
                <TableCell align="center" width="20%">
                  ระยะเวลาที่จะยืม
                </TableCell>
                <TableCell align="center" width="20%">
                  ยืมวันที่
                </TableCell>
                <TableCell align="center" width="20%">
                  สถานะ
                </TableCell>
               
              </TableRow>
            </TableHead>
            <TableBody>
              {borrows.map((item: BorrowDetailInterface) => (
                <TableRow key={item.ID}>
                  {/* <TableCell align="center">{item.ID}</TableCell> */}
                  <TableCell align="center">{item.Info.BookOrder.BookTitle}</TableCell>
                  <TableCell align="center">{item.BorrowDuration}</TableCell>
                  <TableCell align="center">{format((new Date(item.DateToBorrow)), 'dd MMMM yyyy')}</TableCell>
                  <TableCell align="center">{item.Status.Name}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Borrows;