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
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import Cookie from 'js-cookie';
import {blueGrey100} from 'material-ui/styles/colors';
import {cyan400, yellow600} from 'material-ui/styles/colors';

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
  	fontSize: 17
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
    iconButtonElement={<IconButton><MoreVertIcon /></IconButton>}
    targetOrigin={{horizontal: 'right', vertical: 'top'}}
    anchorOrigin={{horizontal: 'right', vertical: 'top'}}>

    <MenuItem primaryText="My Account" />
    <MenuItem primaryText="My Files" />
    <MenuItem primaryText="Logout" />
  </IconMenu>
);
  //
class AppBarr extends Component {
  constructor(props) {
    super(props);
    this.state = {
      logged: false,
      userEmail: "",
    }
    this.handleLoggedIn = this.handleLoggedIn.bind(this);
    this.handleLoggedOut = this.handleLoggedOut.bind(this);
    this.checkLogged = this.checkLogged.bind(this);

      if(Cookie.get("session_id")) window.location = "/#/files";
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

          iconElementLeft={<IconButton> </IconButton>}
          iconElementRight={
            this.state.logged ? <Logged /> : <Login />
          }
          style={{backgroundColor: cyan400}}
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
    <CardTitle title="The elegant way to store your files and access them anytime, anywhere" />
    <CardTitle titleStyle={styles.content} title="• Upload one or more files at a time •" />
    <CardTitle titleStyle={styles.content} title="• Create folders to store your files in •" />
    <CardTitle titleStyle={styles.content} title="• Download your uploaded files •" />
    <CardTitle titleStyle={styles.content} title="• Delete your files and folders •" />

    <br/>

  </Card>

  </div>
);

export default HomePage;
