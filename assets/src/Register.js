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


export default class Login extends React.Component {
  constructor(props) {
    super(props);

    // Initial state
    this.state = {
      open: false,
    }

    // Bind methods
    this.handleOpen = this.handleOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  handleOpen (){
    this.setState({open: true});
  };

  handleClose () {
    this.setState({open: false});
  };

render() {
    const actions = [
      <FlatButton
        label="Cancel"
        primary={true}
        onClick={this.handleClose}
      />,
      <FlatButton
        label="Login"
        primary={true}
        onClick={this.handleClose}
      />,
    ];

    return (
      <div>
        <RaisedButton label="Sign Up" onClick={this.handleOpen} />
        <Dialog
          title="Sign Up"
          actions={actions}
          modal={true}
          open={this.state.open}
        >
          <TextField ref='fullname'
             name='fullname'
             required={true}
             hintText="Full Name"
             //errorText={"error"}
             floatingLabelText="Full Name"
             type="text">
          </TextField>
          <br/>
          <TextField ref='email'
             name='email'
             required={true}
             hintText="Email"
             //errorText={"error"}
             floatingLabelText="Email"
             type="text">
          </TextField>
          <br/>
          <TextField ref='password'
             name='password'
             required={true}
             hintText="Password"
             //errorText={"errorrr"}
             floatingLabelText="Password"
             type="text">
          </TextField>
          <br/>
        </Dialog>
      </div>
    );
  }
}