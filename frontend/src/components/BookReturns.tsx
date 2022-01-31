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
import { BookReturnsInterface } from "../models/IBookReturn";
import { format } from 'date-fns'

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

function BookReturns() {
    const classes = useStyles();
    const [book_returns, setBookReturns] = useState<BookReturnsInterface[]>([]);
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };

    const getBookReturns = async () => {
        fetch(`${apiUrl}/book_returns`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setBookReturns(res.data);
                } else {
                    console.log("else");
                }
            });
    };
    console.log(book_returns);
    useEffect(() => {
        getBookReturns();
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
                            ข้อมูลการคืนหนังสือ
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/book_return/create"
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
                                <TableCell align="center" width="10%">
                                    ลำดับ
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    หนังสือ
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    สถานบริการ
                                </TableCell>
                                {/* <TableCell align="center" width="20%">
                                    จำนวนจุดที่ชำรุด
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    เบอร์โทรศัพท์
                                </TableCell> */}
                                <TableCell align="center" width="30%">
                                    วันที่และเวลา
                                </TableCell>
                                <TableCell align="center" width="20%">
                                    สถานะ
                                </TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {book_returns.map((item: BookReturnsInterface) => (
                                <TableRow >
                                    <TableCell align="center">{item.ID}</TableCell>
                                    <TableCell align="center">{item.BorrowDetail.Info.BookOrder.BookTitle}</TableCell>
                                    <TableCell align="center">{item.SevicePlace.Name}</TableCell>
                                    {/* <TableCell align="center">{item.Damage}</TableCell>
                                    <TableCell align="center">{item.Tel}</TableCell> */}
                                    <TableCell align="center">{format((new Date(item.DateReturn)), 'dd MMMM yyyy')}</TableCell>
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

export default BookReturns;