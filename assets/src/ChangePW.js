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
import axios from 'axios';
import qs from 'qs';
import Cookie from 'js-cookie';


var HashRouter = require('react-router-dom').HashRouter
var Route = require('react-router-dom').Route
var Link = require('react-router-dom').Link

// css to be applied to elements
const styles = {
  button: {
    textAlign: 'center',
    margin: 12
  }
};

export default class ChangePW extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      open: false,
      oldPassword: "",
      NewPassword: "",
    }

    // Bind methods
    this.handleOpen = this.handleOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handleOldPassword = this.handleOldPassword.bind(this);
    this.handleNewPassword = this.handleNewPassword.bind(this);
    this.sendForm = this.sendForm.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };

  handleOldPassword (event) {
    var key = "oldPassword"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }

  handleNewPassword (event) {
    var key = "NewPassword"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }

  // need to change this stuff for a new folder
  sendForm() {
    window.userID = Cookie.get("user_id");
    console.log(userID);
    axios.put(
        ('/users/'+userID), qs.stringify({
          password: this.state.password,
        })
    ).then(function(response) {
      window.location = "/#/profile";
      location.reload()
    }).catch(function (error) {
      alert(error.response.data)
    });

    // TODO instead have user-friendly response and maintain close button
    this.setState({open: false});
  };

  render() {
    const actions = [
      <FlatButton
        label="Cancel"
        onClick={this.handleClose}
      />,
      <FlatButton
        label="Save"
        primary={true}
        onClick={this.sendForm}
      />,
    ];

    return (
      <div>
        <FlatButton label="Change Password" onClick={this.handleOpen} primary={true} />
        <Dialog
          title="Change Password" subtitle="Enter your old and new passwords."
          actions={actions}
          modal={true}
          open={this.state.open}
        >
          <TextField ref='Old password'
             name='Old password'
             required={true}
             floatingLabelText="Old password"
             value={this.state.OldPassword}
             onChange={this.handleOldPassword}
             type="password">
            </TextField>
          <br />
          <TextField ref='New password'
             name='password'
             required={true}
             floatingLabelText="New password"
             value={this.state.NewPassword}
             onChange={this.handleNewPassword}
             type="password">
            </TextField>
            <br/>
        </Dialog>
      </div>
    );
  }
}
