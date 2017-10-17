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
    this.sendForm = this.sendForm.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };
  sendForm() {
    axios.get(
        '/logout', qs.stringify({})
    ).then(function(response) {
        // TODO proper form responses
        console.log(response)
    }).catch(function (error) {
        console.log(response)
    });

    // TODO instead have user-friendly response and maintain close button
    //this.setState({open: false});
  };

render() {
    const actions = [
      <FlatButton
        label="Cancel"
        primary={true}
        onClick={this.handleClose}
      />,
      <FlatButton
        label="Logout"
        primary={true}
        onClick={this.sendForm}
      />,
    ];

    return (
      <div>
        <RaisedButton label="Logout" onClick={this.handleOpen} />
        <Dialog
          title="Are you sure you want to logout?"
          actions={actions}
          modal={true}
          open={this.state.open}
        >
        </Dialog>
      </div>
    );
  }
}
