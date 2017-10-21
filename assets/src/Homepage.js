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
import { Link } from 'react-router-dom';

import Register from './Register';
import LoginPU from './LoginPU';
import LogoutPU from './LogoutPU';
// currently unused
function handleTouchTap() {
  alert('onClick triggered on the title component');
}

// css to be applied to elements
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

class Login extends Component {
  render() {
    return (
      <FlatButton style={styles.button}><LoginPU /></FlatButton>
    );
  }
}
//
const Logged = (props) => (
  <IconMenu {...props}
    iconButtonElement={<IconButton iconStyle={styles.mediumIcon}><AccountIcon /></IconButton>}
    targetOrigin={{horizontal: 'right', vertical: 'top'}}
    anchorOrigin={{horizontal: 'right', vertical: 'top'}}>
    <MenuItem primaryText="My Account"  href="/#/profile" />
    <MenuItem primaryText="My Files" href="/#/files" />
    <MenuItem disabled href="/#/"><LogoutPU /></MenuItem>
  </IconMenu>

);
  //
class AppBarr extends Component {
  constructor(props) {
    super(props);
    this.state = {
      logged: true,
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
          iconElementLeft={<IconButton iconStyle={styles.mediumIcon} href="./"><ActionHome /></IconButton>}
          iconElementRight={
            this.state.logged ? <Logged /> : <Login />
          }
        />
      </div>
    );
  }
}


//
const HomePage = () => (
  <div style={styles.body}>
  <AppBarr />

  <Card>
    <CardMedia>
      <img src="/assets/logo.png" alt="el-goog logo" />
    </CardMedia>
    <CardTitle style={styles.title} title="Ethical File Sync and Backup" />
    <CardActions>
      <RaisedButton style={styles.button}><Register /></RaisedButton>
    </CardActions>

    <Divider />
    <br/>
    <CardTitle titleStyle={styles.content} title="The elegant way to store your files and access them anytime, anywhere" />
    <br/>
    <CardTitle titleStyle={styles.content} title="[insert visual instructions on how to use el-goog]" />
    <br/>
    <CardTitle titleStyle={styles.content} title="product plan: FREE. GETCHA FILES ON 'ERE" />

  </Card>

  </div>
);

export default HomePage;
