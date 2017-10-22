import React from 'react';
import {Component} from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import ActionHome from 'material-ui/svg-icons/action/home';
import FlatButton from 'material-ui/FlatButton';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Divider from 'material-ui/Divider';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import AccountIcon from 'material-ui/svg-icons/action/account-circle';
import Cookie from 'js-cookie';
import {white, black} from 'material-ui/styles/colors';

import Register from './Register';
import LoginPU from './LoginPU';
import LogoutPU from './LogoutPU';


function handleTouchTap() {
}

const styles = {
  title: {
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
  mediumIcon: {
    width: 35,
    height: 35,
  }
};

const Logged = (props) => (

  <IconMenu {...props}
    iconButtonElement={
      <FlatButton label='user' icon={<AccountIcon />}></FlatButton>
    }

    targetOrigin={{horizontal: 'right', vertical: 'top'}}
    anchorOrigin={{horizontal: 'right', vertical: 'top'}}>
    <MenuItem primaryText="My Account"  href="/#/profile" />
    <MenuItem primaryText="My Files" href="/#/files" />
    <Divider />
    <LogoutPU />
  </IconMenu>

);

class AppBarHeader extends Component {
  constructor(props) {
    super(props);
    this.state = {
      logged: true,
      userName: "",
      userEmail: "",
    }
    this.handleLoggedIn = this.handleLoggedIn.bind(this);
    this.handleLoggedOut = this.handleLoggedOut.bind(this);
    this.checkLogged = this.checkLogged.bind(this);
  }

  handleLoggedIn () {
    this.setState({logged: true});
  };
  handleLoggedOut () {
    this.setState({logged: false});
  };

  checkLogged(){
    axios.post(
      '/login', qs.stringify({
        userEmail: this.state.email,
        userName: this.state.name,
    })
    ).then(function (response) {
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });
    if (userEmail == "") {
      this.handleLoggedOut;
    } else {
      this.handleLoggedIn;
    }
  }

  render() {
    return (
      <div>
        <AppBar
          title={<span style={styles.title}></span>}
          onTitleTouchTap={handleTouchTap}
          iconElementLeft={
            <img src="/assets/transparent_logo.gif" width="auto" height="auto" alt="el-goog logo" />
          }
          iconElementRight={<Logged />}
        />
      </div>
    );
  }
}


export default AppBarHeader;
