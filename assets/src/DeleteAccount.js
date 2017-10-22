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

export default class DeleteAccount extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      open: false,
      password: "",
    }

    // Bind methods
    this.handleOpen = this.handleOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.sendForm = this.sendForm.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };

  handlePassword (event) {
    var key = "oldPassword"
    var val = event.target.value
    var rel = {}
    rel[key] = val
    this.setState( rel );
  }

  // need to change this stuff for a new folder
  sendForm() {
    window.userID = Cookie.get("user_id");
    console.log(userID);
    axios.post(
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
        label="Delete"
        primary={true}
        onClick={this.sendForm}
        href = "/#/"
      />,
    ];

    return (
      <div>
        <FlatButton label="Change Password" onClick={this.handleOpen} primary={true} />
        <Dialog
          title="Are you sure you want to delete your el-goog account?" 
          actions={actions}
          modal={true}
          open={this.state.open}
        >
            <br/>
        </Dialog>
      </div>
    );
  }
}
