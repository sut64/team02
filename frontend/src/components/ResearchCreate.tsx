import React from "react";
import { useEffect } from 'react';
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
import { TypeResearchInterface } from "../models/ITypeResearch";
import { AuthorNameInterface } from "../models/IAuthorName";
import { InstitutionNameInterface } from "../models/IInstitutionName";
import { ResearchInterface } from "../models/IResearch";
import {MembersInterface} from "../models/IMember";
import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
 
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
 
function ResearchCreate() {
 const classes = useStyles();
 const [selectedDate, setSelectedDate] = React.useState<Date | null>(
   new Date()
 );
 const [members,setMembers] = React.useState<MembersInterface>()
 const [typeresearch, setTypeResearch] = React.useState<TypeResearchInterface[]>([]);
 const [authorname, setAuthorName] = React.useState<AuthorNameInterface[]>([]);
 const [institutionname, setInstitutionName] = React.useState<InstitutionNameInterface[]>([]);
 const [research, setResearch] = React.useState<Partial<ResearchInterface>>({});

 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);

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
 
 const handleInputChange = (
   event: React.ChangeEvent<{ id?: string; value: any }>
 ) => {
   const id = event.target.id as keyof typeof research;
   const { value } = event.target;
   setResearch({ ...research, [id]: value });
 };

 const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
  const name = event.target.name as keyof typeof research;
  console.log(event.target.value);
  setResearch({...research, [name]: event.target.value,});
  
 };
 
 const getTypeResearch = async () => {
  fetch(`${apiUrl}/typeresearches`, requestOptions)
    .then((response) => response.json())
      .then((res) => {
        if (res.data) {
            console.log(res.data);
            setTypeResearch(res.data);
        } else {
          console.log("else");
        }
      });
  };

const getAuthorName = async () => {
  fetch(`${apiUrl}/authornames`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setAuthorName(res.data);
      } else {
        console.log("else");
      }
    });
};

const getInstitutionName = async () => {
  fetch(`${apiUrl}/institutionnames`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setInstitutionName(res.data);
      } else {
        console.log("else");
      }
    });
};

const getMembers = async () => {
  let uid = localStorage.getItem("uid");
  fetch(`${apiUrl}/member/${uid}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      research.MemberID = res.data.ID
      if (res.data) {
        setMembers(res.data);
      } else {
        console.log("else");
      }
    });
};

    useEffect(() => {
        getTypeResearch();
        getAuthorName();
        getInstitutionName();
        getMembers();
    }, []);



 const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
 };

 function submit() {
   let data = {
     MemberID: convertType(research.MemberID),
     NameResearch: research.NameResearch,
     YearOfPublication: convertType(research.YearOfPublication),
     TypeResearchID: convertType(research.TypeResearchID),
     AuthorNameID: convertType(research.AuthorNameID),
     InstitutionNameID: convertType(research.InstitutionNameID),
     RecordingDate: selectedDate,
      
   };
   console.log(data)
 
   const requestOptionsPost = {
    method: "POST",
    headers: {  
     Authorization: `Bearer ${localStorage.getItem("token")}`,
     "Content-Type": "application/json", },
    body: JSON.stringify(data),
   };
 
   fetch(`${apiUrl}/researches`, requestOptionsPost)
     .then((response) => response.json())
     .then((res) => {
       if (res.data) {
         setSuccess(true);
       } else {
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
             เพิ่มงานวิจัย
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={3} className={classes.root}>
         <Grid item xs={6}>
           <p>ชื่องานวิจัย</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="NameResearch"
               variant="outlined"
               type="string"
               size="medium"
               value={research.NameResearch || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>ปีที่ตีพิมพ์</p>
             <TextField
               id="YearOfPublication"
               variant="outlined"
               type="string"
               size="medium"
               value={research.YearOfPublication || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>
         <Grid item xs={12}>
           <FormControl fullWidth variant="outlined">
           <p>กรุณาเลือกประเภท</p>
              <Select
                value={research.TypeResearchID}
                onChange={handleChange}
                inputProps={{name: "TypeResearchID"}}
              >
                {typeresearch.map((item: TypeResearchInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Value}
                  </MenuItem>
                ))}
              </Select>
           </FormControl>
         </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
            <p>กรุณาเลือกชื่อผู้แต่ง</p>
              <Select
                value={research.AuthorNameID}
                onChange={handleChange}
                inputProps={{name: "AuthorNameID"}}
              >
                  {authorname.map((item: AuthorNameInterface) => (
                    <MenuItem value={item.ID} key={item.ID}>
                      {item.AuthorName}
                    </MenuItem>
                  ))}
              </Select>
           </FormControl>
         </Grid>
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
            <p>กรุณาเลือกสถาบัน</p>
              <Select
                value={research.InstitutionNameID}
                onChange={handleChange}
                inputProps={{name: "InstitutionNameID"}}
              >
                  {institutionname.map((item: InstitutionNameInterface) => (
                    <MenuItem value={item.ID} key={item.ID}>
                      {item.InstitutionName}
                    </MenuItem>
                  ))}
              </Select>
           </FormControl>
          </Grid>
          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่บันทึก</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 id="RecordingDate"
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
          <Grid item xs={6}>
          <Grid item xs={12}>
            <p>ผู้บันทึก</p>
              <FormControl fullWidth variant="outlined">
                  <Select
                    native
                    value={research.MemberID}
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
           <Button component={RouterLink} to="/" variant="contained">
             กลับ
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
           >
             บันทึกข้อมูล
           </Button>
          </Grid>
       </Grid>
     </Paper>
   </Container>
 );
}
 
export default ResearchCreate;