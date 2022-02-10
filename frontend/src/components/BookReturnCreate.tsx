import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
    makeStyles,
    Theme,
    createStyles,
    alpha,
} from "@material-ui/core/styles";
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
import MenuItem from '@material-ui/core/MenuItem';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { ServicePlacesInterface } from "../models/IServicePlace";
import { StatusInterface } from "../models/IStatus";
import { BorrowDetailInterface } from "../models/IBorrowDetail";
import { BookReturnsInterface } from "../models/IBookReturn";
import { MembersInterface } from "../models/IMember";
import {
    MuiPickersUtilsProvider,
    KeyboardDatePicker,
} from "@material-ui/pickers";
import { TextField } from "@material-ui/core";
import DateFnsUtils from '@date-io/date-fns';
import { isTemplateExpression } from "typescript";

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



function BookReturnCreate() {
    const classes = useStyles();
    const [selectedDate, setSelectedDate] = useState<Date | null>(new
        Date());
    const [servicePlaces, setServicePlaces] = useState<ServicePlacesInterface[]>([]);
    const [status, setStatus] = useState<StatusInterface[]>([]);
    const [borrowDetails, setborrowDetails] = useState<BorrowDetailInterface[]>([]);
    const [members, setMembers] = useState<MembersInterface>();
    const [book_return, setBookReturn] =
        useState<Partial<BookReturnsInterface>>(
            {}
        );
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

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof book_return;
        const { value } = event.target;
        setBookReturn({ ...book_return, [id]: value });
    };

    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>
    ) => {
        const name = event.target.name as keyof typeof book_return;
        setBookReturn({
            ...book_return,
            [name]: event.target.value,
        });
    };

    const handleMember = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>
    ) => {
        const id = Number(event.target.value);
        console.log(id)
        const name = event.target.name as keyof typeof book_return;
        setBookReturn({
            ...book_return,
            [name]: event.target.value,
        });
    }

    const handleDateChange = (date: Date | null) => {
        console.log(date);
        setSelectedDate(date);
    };

    console.log("value", book_return)
    const getMember = async () => {
        let uid = localStorage.getItem("uid");
        fetch(`${apiUrl}/member/${uid}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                book_return.MemberID = res.data.ID
                if (res.data) {
                    console.log(res.data);
                    setMembers(res.data);
                    getBorrowDetails(res.data.ID);
                } else {
                    console.log("else");
                }
            });
    };

    const getSevicePlace = async () => {

        fetch(`${apiUrl}/places`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setServicePlaces(res.data);
                } else {
                    console.log("else");
                }
            });
    };

    const getBorrowDetails = async (id: number) => {

        fetch(`${apiUrl}/borrow/member/${id}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res.data);
                if (res.data) {
                    setborrowDetails(res.data);
                    console.log(res.data);
                } else {
                    console.log("else");
                }
            });
    };
    console.log("value data", borrowDetails);

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
        getSevicePlace();
        getStatus();
    }, []);
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };
    function submit() {
        let data = {
            MemberID: convertType(book_return.MemberID),
            BorrowDetailID: convertType(book_return.BorrowDetailID),
            ServicePlaceID: convertType(book_return.ServicePlaceID),
            Damage: convertType(book_return.Damage),
            Tel: book_return.Tel ?? "",
            DateReturn: selectedDate,
            StatusID: 2,
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
        fetch(`${apiUrl}/book_return`, requestOptionsPost)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("")
                } else {
                    setError(true);
                    setErrorMessage(res.error)
                }
            });
    }
    return (
        <Container className={classes.container} maxWidth="md">
            <Snackbar open={success} autoHideDuration={6000}
                onClose={handleClose}>
                <Alert onClose={handleClose} severity="success">
                    บันทึกข้อมูลสําเร็จ
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000}
                onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    บันทึกข้อมูลไม่สําเร็จ: {errorMessage}
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
                            RETURN BOOK
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} className={classes.root}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard">
                            <p>ชื่อผู้คืน</p>
                            <TextField
                                disabled
                                id="MemberID"
                                value={members?.Name}

                            />
                        </FormControl>

                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard">
                            <p>หนังสือ</p>
                            <Select
                                native
                                value={book_return.BorrowDetailID}
                                onChange={handleMember}
                                inputProps={{
                                    name: "BorrowDetailID",
                                }}
                            >
                                {borrowDetails.map((item: BorrowDetailInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.BookOrder.BookTitle}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard">
                            <p>สถานบริการ</p>
                            <Select
                                native
                                value={book_return.ServicePlaceID}
                                onChange={handleChange}
                                inputProps={{
                                    name: "ServicePlaceID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    เลือกสถานที่ที่ไปใช้บริการ
                                </option>
                                {servicePlaces.map((item: ServicePlacesInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.Name}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard" >
                            <p>จำนวนจุดที่ชำรุด</p>
                            <TextField
                                id="Damage"
                                variant="standard"
                                type="number"
                                size="medium"
                                placeholder=""
                                value={book_return.Damage || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard">
                            <p>เบอร์โทรศัพท์</p>
                            <TextField
                                required
                                id="Tel"
                                variant="standard"
                                type="string"
                                size="medium"
                                placeholder="(+66)"
                                value={book_return.Tel || ""}
                                onChange={handleInputChange}

                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="standard">
                            <p>วันที่คืนหนังสือ</p>
                            <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                <KeyboardDatePicker
                                    name="DateReturn"
                                    value={selectedDate}
                                    onChange={handleDateChange}
                                    label="กรุณาเลือกวันที่"
                                    size="medium"
                                    minDate={new Date("2018-01-01T00:00")}
                                    format="yyyy/MM/dd "
                                />
                            </MuiPickersUtilsProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl fullWidth variant="standard">
                            <p>สถานะ</p>

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

                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/book_return"
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
export default BookReturnCreate;