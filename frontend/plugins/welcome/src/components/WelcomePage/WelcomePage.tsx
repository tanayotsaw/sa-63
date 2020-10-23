import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntTeacher } from '../../api/models/EntTeacher'; // import interface User
import { EntSection } from '../../api/models/EntSection'; // import interface User
import { EntSubject } from '../../api';
import { Link as RouterLink } from 'react-router-dom';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface lesson_plan {
  GroupssID:   Number
	TeacherssID: Number
	SubjectssID: Number
	Room:        String
  // create_by: number;
}

const WelcomePage: FC<{}> = () => {
 const classes = useStyles();
 const http = new DefaultApi();
 
 const [lessonplans, setLessonplans] = React.useState<Partial<lesson_plan>>({});
 const [teachers, setTeachers] = React.useState<EntTeacher[]>([]);
 const [secs, setSecs] = React.useState<EntSection[]>([]);
 const [subjects, setSubjects] = React.useState<EntSubject[]>([]);

 const getTeachers = async () => {
  const res = await http.listTeacher({ limit: 10, offset: 0 });
  setTeachers(res);
};

const getSecs = async () => {
  const res = await http.listSection({ limit: 10, offset: 0 });
  setSecs(res);
};

const getSubjects = async () => {
  const res = await http.listSubject({ limit: 10, offset: 0 });
  setSubjects(res);
};



useEffect(() => {
  getTeachers();
  getSecs();
  getSubjects();
}, []);


interface lessonplanss {
  GroupssID : number;
  TeacherssID: number;
  SubjectssID: number;
  Room: string;
  // create_by: number;
}

const [lessonplanss, setlessonplans] = React.useState<Partial<lessonplanss>>({});

const handleChange = (
  event: React.ChangeEvent<{ name?: string; value: unknown }>,
) => {
  const name = event.target.name as keyof typeof WelcomePage;
  const { value } = event.target;
  setlessonplans({ ...lessonplanss, [name]: value });
  console.log(lessonplanss);
};

function save() {
  const apiUrl = 'http://localhost:8080/api/v1/lessonplans';
  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(lessonplanss),
};

  console.log(lessonplanss); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

  fetch(apiUrl, requestOptions)
    .then(response => response.json())
    .then(data => {
      console.log(data);
    });
}

 return (
  <Page theme={pageTheme.home}>
  
  <Header style={HeaderCustom} title={`Lesson Plan`}>
    
  </Header>
  
  <Content>
  <Grid item xs={3}></Grid>
        <Grid item xs={9}>
          <Button
            style={{ width: 200, height: 40, marginLeft: 100 }}
            variant="contained"
            color="primary"
            component={RouterLink}
            to="/Table"
        >
            แผนการสอน
          </Button>
        </Grid>

      <Grid item xs={3}></Grid>
        <Grid item xs={9}>
          <Button
            style={{ width: 200, height: 40, marginLeft: 1300 }}
            variant="contained"
            color="primary"
            component={RouterLink}
            to="/"
        >
            Log Out
          </Button>
        </Grid>
  
    <Container maxWidth="sm">
      <Grid container spacing={3}>
        <Grid item xs={12}></Grid>
        <Grid item xs={3}>
          <div className={classes.paper}>วิชา</div>
        </Grid>
        <Grid item xs={9}>
          <FormControl variant="outlined" className={classes.formControl}>
            <InputLabel></InputLabel>
            <Select
                  name="SubjectssID"
                  value={lessonplanss.SubjectssID || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                {subjects.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.subjectName}
                      </MenuItem>
                    );
                  })}
                  </Select>
          </FormControl>
        </Grid>

        <Grid item xs={3}>
          <div className={classes.paper}>เซค</div>
        </Grid>
        <Grid item xs={9}>
          <FormControl variant="outlined" className={classes.formControl}>
            <InputLabel></InputLabel>
            <Select
                  name="GroupssID"
                  value={lessonplanss.GroupssID|| ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                {secs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.group}
                      </MenuItem>
                    );
                  })}
              </Select>
          </FormControl>
        </Grid>

        <Grid item xs={3}>
          <div className={classes.paper}>ห้อง</div>
        </Grid>
        <Grid item xs={9}>
          <FormControl variant="outlined" className={classes.formControl}>
            <InputLabel></InputLabel>
            <Select
                  name="Room"
                  value={lessonplanss.Room|| ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                {secs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.room}>
                        {item.room}
                      </MenuItem>
                    );
                  })}
              </Select>
          </FormControl>
        </Grid>

        <Grid item xs={3}>
          <div className={classes.paper}>สมาชิกระบบ</div>
        </Grid>
        <Grid item xs={9}>
          <FormControl variant="outlined" className={classes.formControl}>
          {teachers.map(item => {
                    return (
                      <InputLabel>
                        
                      </InputLabel>
                    );
                  })}
            <Select
                   value={lessonplanss.TeacherssID || ''} // (undefined || '') = ''
                   onChange={handleChange}
                  name="TeacherssID"
                >
                  {teachers.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.teacherName}
                      </MenuItem>
                    );
                  })}
              </Select>
             
          </FormControl>
        </Grid>

        

        <Grid item xs={3}></Grid>
        <Grid item xs={9}>
          <Button
            variant="contained"
            color="primary"
            size="large"
            startIcon={<SaveIcon />}
            onClick= {save}
          >
            บันทึกแผนการสอน
          </Button>
        </Grid>
      </Grid>
    </Container>
  </Content>
</Page>
);
};
export default WelcomePage;
 
