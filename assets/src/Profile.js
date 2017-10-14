import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import FlatButton from 'material-ui/FlatButton';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Divider from 'material-ui/Divider';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import {List, ListItem} from 'material-ui/List';
import Subheader from 'material-ui/Subheader';
import Paper from 'material-ui/Paper';


function handleTouchTap() {
  alert('onClick triggered on the title component');
}

const styles = {
  title: {
    cursor: 'pointer',
    fontSize: 20
  },
  button: {
    textAlign: 'center',
    margin: 12
  },
  content: {
    textAlign: 'center',
    fontSize: 15
  },
  body: {
    textAlign: 'center'
  },
  root: {
    display: 'flex',
    flexWrap: 'wrap',
  }
};

const Profile = () => (
  <div style={styles.body}>
  <AppBar
    title={<span style={styles.title}></span>}
    onTitleTouchTap={handleTouchTap}
    iconElementLeft={<IconButton href="./"><NavigationClose /></IconButton>}
 //   iconElementRight={<RaisedButton style={styles.button} label="Login" />}
  />

  <Card>
    <CardTitle style={styles.title} title="Personal Account" />
      <div style={styles.root}>      
          <List>
            <Subheader>My Account</Subheader>
            <ListItem
              primaryText="General"
              secondaryText="Edit and view my information"
            />
            <ListItem
              primaryText="Files"
              secondaryText="My files and folders"
            />
            <ListItem
              primaryText="Plan"
              secondaryText="Personal el-goog space"
            />
            <ListItem
              primaryText="Security"
              secondaryText="Change password"
            />
          </List>
         
        <Card>
            <CardTitle
              title="General"
              subtitle="Edit and view my information"
            />
              <CardText>Name</CardText>
            <CardActions>
              <FlatButton label="Edit" primary={true}/>
            </CardActions>
            <Divider />  
            <CardText>Personal email</CardText>
            <CardActions>
              <FlatButton label="Edit" primary={true}/>
            </CardActions>
            <Divider />        
            <CardText>Delete account</CardText>
            <CardActions>
              <FlatButton label="Delete account" primary={true}/>
            </CardActions>
            <Divider /> 

             

        </Card>    
       </div>
  </Card>
  </div>
);



export default Profile;