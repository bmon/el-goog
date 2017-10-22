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
import {cyan400} from 'material-ui/styles/colors';
import axios from "axios";

import SettingIcon from 'material-ui/svg-icons/action/settings';
import SecurityIcon from 'material-ui/svg-icons/hardware/security';

import Register from './Register';
import LogoutPU from './LogoutPU';
import Header from './Header';
import EditName from './EditName';
import ChangePW from './ChangePW';
import DeleteAccount from './DeleteAccount';

// currently unused
function handleTouchTap() {
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
    textAlign: 'center',
  },
  containerCard: {
    margin: 50,
    textAlign: 'center',
    boxShadow: 'none'
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
    fontWeight: 400,
    height: 80,
    backgroundColor: cyan400,
  }
};

class UserDetails extends Component {
  constructor (props) {
    super(props);
    this.state = {
      Username: "",
      email: "",
    }

    this.handleName = this.handleName.bind(this);
    this.handleEmail = this.handleEmail.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.render = this.render.bind(this);

    var _this = this;
    window.userID = Cookie.get("user_id");
    axios.get("/users/"+userID)
    .then(function(result) {
      _this.setState({
        Username: result.data.Username,
        email: result.data.email,
      });
    })
  }

  handleName (event) {
    var key = "Username"
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

  sendForm() {
    window.userID = Cookie.get("user_id");
    axios.post(
      '/users/'+userID, qs.stringify({
        Username: this.state.Username,
    }))
  }

  render(){
    return(
      <div style={styles.body}>
      <Header />
    
      <Card style={styles.container}>
        <div>
          <CardTitle title="Personal account"/>
              <Tabs>
                <Tab style={styles.tabs} label="General" icon={<IconButton ><SettingIcon /></IconButton>} >
                  <div>
                    <Card style={styles.containerCard}>
                    <div>
                      <CardTitle>NAME: {this.state.Username}</CardTitle>
                      <EditName />
                      <br />
                      <Divider />  
                      <br />
                      <CardTitle>EMAIL: {this.state.email}</CardTitle>
                      <CardTitle subtitle="E-mail cannot be changed after account is created."/>
                    </div>
                    </Card>
                  </div>

                </Tab>
                <Tab style={styles.tabs} label="Security" icon={<IconButton><SecurityIcon /></IconButton>} >
                  <div>
                    <Card style={styles.containerCard}>
                    <div>     
                      <CardTitle>Password</CardTitle> 
                      <CardTitle subtitle="Set a unique password to protect your personal account."/>
                      <ChangePW />
                      <br />
                      <Divider />
                      <br /> 
                      <CardTitle>Delete Account</CardTitle>
                      <CardTitle subtitle="If you delete your account, you data will be gone forever."/>
                      <DeleteAccount />
                      <Divider /> 
                    </div>
                    </Card>
                  </div>
                </Tab>
              </Tabs>          
        </div>
      </Card>

      </div>



    )
  }
}


export default UserDetails;
