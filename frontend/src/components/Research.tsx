import React, { useEffect } from "react";
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
import { ResearchInterface } from "../models/IResearch";
 
import moment from 'moment';
 
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   container: {marginTop: theme.spacing(2)},
   table: { minWidth: 650},
   tableSpace: {marginTop: 20},
 })
);
 
function Researches() {
 const classes = useStyles();
 const [research, setResearchs] = React.useState<ResearchInterface[]>([]);
 const apiUrl = "http://localhost:8080";
   const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    "Content-Type": "application/json", },
   };
 const getResearchs = async () => {
  fetch(`${apiUrl}/researches`, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
         setResearchs(res.data);
       } else {
         console.log("else");
       }
     });
 };
 
 useEffect(() => {
   getResearchs();
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
             งานวิจัย
           </Typography>
         </Box>
         <Box>
           <Button
             component={RouterLink}
             to="/researCreate/create"
             variant="contained"
             color="primary"
           >
             เพิ่มงานวิจัย
           </Button>
         </Box>
       </Box>
       <TableContainer component={Paper} className={classes.tableSpace}>
         <Table className={classes.table} aria-label="simple table">
           <TableHead>
             <TableRow>
               <TableCell align="center" width="25%">
                 ชื่องานวิจัย
               </TableCell>
               <TableCell align="center" width="5%">
                 ปีที่ตีพิมพ์
               </TableCell>
               <TableCell align="center" width="20%">
                 ประเภทงานวิจัย
               </TableCell>
               <TableCell align="center" width="15%">
                 ชื่อผู้แต่ง
               </TableCell>
               <TableCell align="center" width="25%">
                 ชื่อสถาบัน
               </TableCell>
               <TableCell align="center" width="10%">
                 วันที่บันทึก
               </TableCell>
             </TableRow>
           </TableHead>
           <TableBody>
             {research.map((researches: ResearchInterface) => (
               <TableRow key={researches.ID}>
                 <TableCell align="center">
                   {researches.NameResearch}
                 </TableCell>
                 <TableCell align="center" size="medium">
                   {researches.YearOfPublication}
                 </TableCell>
                 <TableCell align="center">{researches.TypeResearch.Value}</TableCell>
                 <TableCell align="center">{researches.AuthorName.AuthorName}</TableCell>
                 <TableCell align="center">{researches.InstitutionName.InstitutionName}</TableCell>
                 <TableCell align="center">{moment(researches.RecordingDate).format("DD/MM/YYYY")}</TableCell>
               </TableRow>
             ))}
           </TableBody>
         </Table>
       </TableContainer>
     </Container>
   </div>
 );
}
 
export default Researches;