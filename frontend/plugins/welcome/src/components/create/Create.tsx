import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  //Container,
  //Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  //Avatar,
  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

import { EntRoomage } from '../../api/models/EntRoomage';
import { EntRoom } from '../../api/models/EntRoom';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntUser } from '../../api/models/EntUser';
import { Alert } from '@material-ui/lab';




import { Link as RouterLink } from 'react-router-dom';
/*
const HeaderCustom = {
  minHeight: '50px',
};
*/
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },


  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },

}));

interface Rent_Type {
  /**
   * 
   * @type {string}
   * @memberof ControllersRent
   */
  cidUser?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersRent
   */
  insurance?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersRent
   */
  rentAge?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersRent
   */
  rentDate?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersRent
   */
  room?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersRent
   */
  roomage?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersRent
   */
  user?: number;
}

const NewRent: FC<{}> = () => {
  const classes = useStyles();
  const profile = { givenName: 'สัญญาเช่า' };
  const http = new DefaultApi();

  const [Rent, setRent] = React.useState<Partial<Rent_Type>>({});

  const [Roomage, setRoomage] =     React.useState<EntRoomage[]>([]);
 const [Room, setRoom] =            React.useState<EntRoom[]>([]);
 const [Insurance, setInsurance] =  React.useState<EntInsurance[]>([]);
 const [User, setUser] =            React.useState<EntUser[]>([]);
 const [status, setStatus] = React.useState(false);
 const [alert, setAlert] = React.useState(true);


  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUser(res);
  };

  const getRoomage = async () => {
    const res = await http.listRoomage({ limit: 10, offset: 0 });
    setRoomage(res);
  };

  const getRoom = async () => {
    const res = await http.listRoom({ limit: 10, offset: 0 });
    setRoom(res);
  };

  const getInsurance = async () => {
    const res = await http.listInsurance({ limit: 10, offset: 0 });
    setInsurance(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getUsers();
    getRoomage();
    getRoom();
    getInsurance();
  }, []);

  // set data to object Rent
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewRent;
    const { value } = event.target;
    setRent({ ...Rent, [name]: value });
  };


 


  const CreateRent = async () => {
   
    const res: any = await http.createRent({ 
      rent:Rent
    
    
    });
    setStatus(true);
  
    
    if (res.id != ''){
      setAlert(true);
    } else {
      setAlert(false);
    }
  
  };
  


 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ลงทะเบียน ${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ><Timer /></Header>
     <Content>
       <ContentHeader title="เช่าห้องพักอาศัย">
         
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 สร้างสัญญาสำเร็จ
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

        
       <div className={classes.root}>

         
         <form noValidate autoComplete="off">
         
         <FormControl variant="filled" className={classes.formControl}>
          <TextField
               name="rentDate"
               label="RentDate"
               
               type="date"
               size="medium"
               
               value={Rent.rentDate}
               onChange={handleChange}

               InputLabelProps={{
                shrink: true,
              }}

             />

          </FormControl>
            


         </form>
       </div>
       
       <div className={classes.root}>

         
         <form noValidate autoComplete="off">

         
         
         <FormControl variant="filled" className={classes.formControl}>
          <TextField
               name="rentAge"
               label="RentAge"
               variant="outlined"
               type="string"
               size="medium"
               
               value={Rent.rentAge}
               onChange={handleChange}
             />

          </FormControl>
          
          <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Roomage</InputLabel>
                <Select
                  name="roomage"
                  value={Rent.roomage}
                  onChange={handleChange}
                >
                  {Roomage.map((item:any) => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.text}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>


        </form>
       </div>
        
       <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Room</InputLabel>
                <Select
                  name="room"
                  value={Rent.room}
                  onChange={handleChange}
                >
                  {Room.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
         
 
 
        
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Insurance</InputLabel>
                <Select
                  name="insurance"
                  value={Rent.insurance}
                  onChange={handleChange}
                >
                  {Insurance.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.insurance}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>

           </form>
       </div>











        <ContentHeader title="ข้อมูลผู้เช่า"/>
         




        <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผู้อยู่อาศัย</InputLabel>
                <Select
                 name="user"
                  value={Rent.user}
                  onChange={handleChange}
                >
                  {User.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.userEmail}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
        
         
         

        


        </form>
        
       </div>

       <div className={classes.root}>
         <form noValidate autoComplete="off">
         <FormControl variant="filled" className={classes.formControl}>
          <TextField
               name="cidUser"
               label="CidUser"
               variant="outlined"
               type="string"
               size="medium"
               
               value={Rent.cidUser}
               onChange={handleChange}
             />

          </FormControl>
             
        </form>
        
        </div>










       <div className={classes.root}>
         <form noValidate autoComplete="off">
           
         
           <div className={classes.margin}>
             <Button
               onClick={() => {
                CreateRent();
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
              
         <Link component={RouterLink} to="">
           <Button variant="contained" color="primary">
           กลับสู่หน้าหลัก
           </Button>
         </Link>


             

           </div>
         </form>
       </div>
  
     </Content>
   </Page>
 );
};
export default NewRent;