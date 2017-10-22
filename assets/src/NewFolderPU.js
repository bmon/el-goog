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

export default class Login extends React.Component {
  constructor(props) {
    super(props);

    // Initial state
    this.state = {
      open: false,
      email: "",
      password: "",
    }

    // Bind methods
    this.handleOpen = this.handleOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
    this.handleEmail = this.handleEmail.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.sendForm = this.sendForm.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };
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
  // need to change this stuff for a new folder
  sendForm() {
    axios.post(
        '/login', qs.stringify({
            email: this.state.email,
            password: this.state.password,
        })
    ).then(function(response) {
        // TODO proper form responses
      window.location = "/#/files";
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
        label="Make Folder"
        primary={true}
        onClick={this.sendForm}
      />,
    ];

    return (
      <div>
        <RaisedButton label="New Folder" onClick={this.handleOpen} />
        <Dialog
          title="Make a New Folder"
          actions={actions}
          modal={true}
          open={this.state.open}
        >
          <TextField ref='folderName'
             name='name'
             required={true}
           hintText="Folder Name"
               floatingLabelText="Folder Name"
               // need to change this stuff for the new folder name stuff
                value={this.state.email}
                onChange={this.handleEmail}
               type="text">
            </TextField>
            <br/>
        </Dialog>
      </div>
    );
  }
}
