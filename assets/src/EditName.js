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

export default class EditName extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      open: false,
      Username: "",
      password: "",
      oldpassword: "",
    }

    // Bind methods
    this.handleOpen = this.handleOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handleName = this.handleName.bind(this);
    this.sendForm = this.sendForm.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };
  handleName (event) {
    var key = "Username"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }

  // need to change this stuff for a new folder
  sendForm() {
    window.userID = Cookie.get("user_id");
    console.log(userID);
    axios.put('/users/'+userID, {
          username: this.state.Username,
          newPassword: this.state.password,
          oldPassword: this.state.oldpassword,
        }
    ).then(function(response) {
   //   window.location = "/#/profile";
   //   location.reload()
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
        <FlatButton label="Edit Name" onClick={this.handleOpen} primary={true} />
        <Dialog
          title="Change Name"
          actions={actions}
          modal={true}
          open={this.state.open}
        >
          <TextField ref='New name'
             name='name'
             required={true}
             hintText="New name"
             floatingLabelText="New name"
             value={this.state.Username}
             onChange={this.handleName}
             type="text">
            </TextField>
            <br/>
        </Dialog>
      </div>
    );
  }
}
