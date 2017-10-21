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
import {blue500, yellow600} from 'material-ui/styles/colors';
import EditorInsertChart from 'material-ui/svg-icons/editor/insert-chart';

import Register from './Register';
import LogoutPU from './LogoutPU';
import Header from './Header';

// currently unused
function handleTouchTap() {
  alert('onClick triggered on the title component');
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
    textAlign: 'center'
  },
  fileContainer: {
    margin: 90,
  },
  mediumIcon: {
    width: 35,
    height: 35,
  }
};

class UserDetails extends Component {
  constructor (props) {
    super(props);
    this.state = {
      name: "",
      email: "",
      password: "",
    }
    this.handleName = this.handleName.bind(this);
    this.handleEmail = this.handleEmail.bind(this);
    this.handlePassword = this.handlePassword.bind(this);
    this.sendForm = this.sendForm.bind(this);
  }
  handleName (event) {
    var key = "email"
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
    axios.post(
        '/users', qs.stringify({
            name: this.state.name,
            email: this.state.email,
            password: this.state.password,
        })
    ).then(function(response) {
        _this.setState({
        name: result.data.username,
        email: result.data.email,
        password: result.data.password,
      })
    }).catch(function (error) {
      alert(error.response.data)
    });
/*
    var _this = this;
    axios.get("/users/"+userID)
    .then(function(result) {
      _this.setState({
        name: result.data.username,
        email: result.data.email,
        password: result.data.password,
    }).catch(function (error) {
      alert(error.response.data)
      });
    })*/
  };
  

  render(){
    return(
      <div style={styles.body}>
      <Header />
      
      <Card style={styles.container}>
        <div>
          <CardTitle title="Personal account"/>
          <CardText>Name: {this.state.name}</CardText>
          <FlatButton style={styles.button} label="Edit Name" primary={true}/>
          <Divider />  
          <CardText>Primary Email: {this.state.email}</CardText>
          <CardActions>
            <FlatButton style={styles.button} label="Edit Email" primary={true}/>
          </CardActions>
          <Divider />        
          <CardText>Delete Account: {this.state.password}</CardText>
          <CardActions>
            <FlatButton style={styles.button} label="Delete Account" primary={true}/>
          </CardActions>
          <Divider /> 
        </div>
      </Card>

      </div>



    )
  }
}

const Account = () => (
  <div style={styles.body}>
  <Header />
  
  <Card style={styles.container}>
    <div>
      <CardTitle title="Personal account"/>
      <CardText>Name</CardText>
      <FlatButton style={styles.button} label="Edit Name" primary={true}/>
      <Divider />  
      <CardText>Primary Email</CardText>
      <CardActions>
        <FlatButton style={styles.button} label="Edit Email" primary={true}/>
      </CardActions>
      <Divider />        
      <CardText>Delete Account</CardText>
      <CardActions>
        <FlatButton style={styles.button} label="Delete Account" primary={true}/>
      </CardActions>
      <Divider /> 
    </div>
  </Card>

  </div>
);

export default UserDetails;
