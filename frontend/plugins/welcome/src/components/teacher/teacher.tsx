import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
   } from '@backstage/core';
import FormControl from '@material-ui/core/FormControl';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);
    
export default function BasicTextFields() {
  const classes = useStyles();
  const profile = { givenName: 'to System Analysis and Design 63' };

  return (
    <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
     <div className={classes.root}>    
       <ContentHeader title="อาจารย์ Pantip"> </ContentHeader>
       <form className={classes.root} noValidate autoComplete="off">
            
            <FormControl fullWidth className={classes.margin}variant="outlined">
                <TextField id="outlined-basic" label="รหัสวิชา" variant="outlined" />
            </FormControl>
            <FormControl fullWidth className={classes.margin}variant="outlined">
                <TextField id="outlined-basic" label="ห้องเรียน" variant="outlined" />
            </FormControl>
            <FormControl fullWidth className={classes.margin}variant="outlined">
                <TextField id="outlined-basic" label="จำนวนคน" variant="outlined" />
            </FormControl>
            <FormControl fullWidth className={classes.margin}variant="outlined">
                <TextField id="outlined-basic" label="วันและเวลา" variant="outlined" />
            </FormControl>

        <div className={classes.margin}>
             <Button
               onClick={() => {
                
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
             >
               Back
             </Button>
           </div>
           </form>
       </div>
     </Content>
   </Page>
  );
}