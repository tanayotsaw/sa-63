import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntLessonplan } from '../../api/models/EntLessonplan';
import {
    Grid,
  } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { Content, Header, Page, pageTheme } from '@backstage/core';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});

 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Lessonplan, setLessonplan] = useState<EntLessonplan[]>(Array);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getLessonplans = async () => {
     const res = await api.listLessonplan({ limit: 10, offset: 0 });
      setLoading(false);
      setLessonplan(res);
   };
   getLessonplans();
 }, [loading]);
 
 const deleteLessonplans = async (id: number) => {
   const res = await api.deleteLessonplan({ id: id });
   setLoading(true);
 };
 
 return (
    <Page theme={pageTheme.home}>
             <Header
                title="ตารางแผนการสอน"
                
             ></Header>
     <Content>
     <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
       <TableHead>
        <Grid item xs={3}></Grid>
        <Grid item xs={9}>
          <Button
            style={{ width: 300, height: 40, marginLeft: 130 }}
            variant="contained"
            color="primary"
            component={RouterLink}
            to="/WelcomePage"
          >
            วางแผนการสอน
          </Button>
        </Grid>
        
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Room</TableCell>
           <TableCell align="center">Section</TableCell>
           <TableCell align="center">Subject</TableCell>
           <TableCell align="center">Teacher</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {Lessonplan.map((item: any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.room}</TableCell>
             <TableCell align="center">{item.edges?.groupID?.group}</TableCell>
             <TableCell align="center">{item.edges?.courseID?.subjectName}</TableCell>
             <TableCell align="center">{item.edges?.professorID?.teacherName}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteLessonplans(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button></TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
   </Page>
 );
}
