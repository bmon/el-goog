import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import ActionHome from 'material-ui/svg-icons/action/home';
import FlatButton from 'material-ui/FlatButton';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Divider from 'material-ui/Divider';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import FineUploaderTraditional from 'fine-uploader-wrappers'
import Gallery from 'react-fine-uploader'
import { Component } from 'react'
import {List, ListItem} from 'material-ui/List';
import ActionInfo from 'material-ui/svg-icons/action/info';
import Subheader from 'material-ui/Subheader';
import Avatar from 'material-ui/Avatar';
import FileFolder from 'material-ui/svg-icons/file/folder';
import ActionAssignment from 'material-ui/svg-icons/action/assignment';
import EditorInsertChart from 'material-ui/svg-icons/editor/insert-chart';
import {Tabs, Tab} from 'material-ui/Tabs';
import FontIcon from 'material-ui/FontIcon';
import Cookie from 'js-cookie';
import axios from "axios";

import SettingIcon from 'material-ui/svg-icons/action/settings';

import Register from './Register';
import LogoutPU from './LogoutPU';
import Header from './Header';



// currently unused
function handleTouchTap() {
  alert('onClick triggered on the title component');
}

// css to be applied to elements
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
  container: {
    margin: 50,
    textAlign: 'center'
  },
  fileContainer: {
    margin: 90,
  },
  mediumIcon: {
    width: 35,
    height: 35,
  },
  tabs: {
    fontSize: 16,
    paddingTop: 16,
    marginBottom: 12,
    fontWeight: 400,
  }
};

class UserDetails extends Component {
  constructor (props) {
    super(props);
    this.state = {
      name: "",
      email: "",
      password: "",
    }
    this.handleName = this.handleName.bind(this);
    this.handleEmail = this.handleEmail.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.render = this.render.bind(this);

    var _this = this;
    window.userID = Cookie.get("user_id");
    console.log(userID);
    axios.get("/users/"+userID)
    .then(function(result) {
     // console.log(result.data.username);
      _this.setState({
        name: result.data.username,
        email: result.data.email,
        password: result.data.password,
    });
        console.log(this.state.name);

    })
  }
  handleName (event) {
    var key = "email"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }

  handleEmail (event) {
    var key = "email"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }
  handlePassword (event) {
    var key = "password"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }


  render(){
    return(
      <div style={styles.body}>
      <Header />
    
      <Card style={styles.container}>
        <div>
          <CardTitle title="Personal account"/>
              <Tabs>
                <Tab style={styles.tabs} label="General" icon={<IconButton><SettingIcon /></IconButton>} >
    
                  <div>
                    <h2 style={styles.tabs}>Tab One</h2>
                    <p>
                      This is an example tab.
                    </p>
                    <p>
                      You can put any sort of HTML or react component in here. It even keeps the component state!
                    </p>
                  </div>
                </Tab>
                <Tab icon={<FontIcon className="material-icons">phone</FontIcon>} label="Security" >
                  <div>
                    <h2 style={styles.tabs}>Tab Two</h2>
                    <p>
                      This is another example tab.
                    </p>
                  </div>
                </Tab>
              </Tabs>


          
        </div>
      </Card>

      </div>



    )
  }
}

const Account = () => (
  <div style={styles.body}>
  <Header />
  
  <Card style={styles.container}>
    <div>
      <CardTitle title="Personal account"/>
      <CardText>Name</CardText>
      <FlatButton style={styles.button} label="Edit Name" primary={true}/>
      <Divider />  
      <CardText>Primary Email</CardText>
      <CardActions>
        <FlatButton style={styles.button} label="Edit Email" primary={true}/>
      </CardActions>
      <Divider />        
      <CardText>Delete Account</CardText>
      <CardActions>
        <FlatButton style={styles.button} label="Delete Account" primary={true}/>
      </CardActions>
      <Divider /> 
    </div>
  </Card>

  </div>
);

export default UserDetails;
