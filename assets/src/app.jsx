var React = require('react')
var ReactDOM = require('react-dom')

var HashRouter = require('react-router-dom').HashRouter
var Route = require('react-router-dom').Route
var Link = require('react-router-dom').Link

var axios = require('axios').axios
var filesize = require('file-size');
//import {Router, Route, hashHistory } from 'react-router'

// importing cause requiring a class definition doesnt seem to work
import { Component } from 'react'
import FineUploaderTraditional from 'fine-uploader-wrappers'
import Gallery from 'react-fine-uploader'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import Register from './Register';
import HomePage from './Homepage';
import Login from './Login';
import FilesPage from './Files';
import ProfilePage from './Profile';

//import FineUploaderTraditional from 'fine-uploader-wrappers'
//import Gallery from 'react-fine-uploader'

var homeStyle = {
  padding: 10,
  margin: 10,
  backgroundColor: "#ffffff",
  color: "#333",
  display: "inline-block",
  fontFamily: "monospace",
  fontSize: 32,
  textAlign: "center"
}

var contentStyle = {
  padding: 10,
  margin: 10,
  backgroundColor: "#E5E1E0",
  color: "#333",
  display: "inline-block",
  fontFamily: "monospace",
  fontSize: 28,
  textAlign: "center"
}

var buttonStyle = {
  padding: 10,
  height: 20,
  // width: 60,
  margin: 10,
  backgroundColor: "#4CAF50",
  color: "#ffffff",
  display: "inline-block",
  fontFamily: "monospace",
  fontSize: 10,
  textAlign: "center",
  borderRadius: 8
}

var Content = React.createClass({
  render() {
      return (
        <div style={contentStyle}>
          <h2>How el-Goog works</h2>
          <p>[insert all of the pictures about how el-goog works]</p>
          <p>[insert all of the pictures about how el-goog works]</p>
        </div>
      )
    }
})

var Home = React.createClass({
  render() {
    return (
      <div style={homeStyle}>
        <h1>El-Goog</h1>
        <ul className="header">
          <div><li style={buttonStyle}><Link to="/login" activeClassName="active">Login</Link></li></div>
          <div><li style={buttonStyle}><Link to="/signup" activeClassName="active">SignUp</Link></li></div>
        </ul>
        <div className="content">
        <Content />
        </div>
      </div>
    )
  }
})

//==============================================================================
// React Material UI New Homepage

const NewHome = () => (
  <MuiThemeProvider>
    <HomePage />
  </MuiThemeProvider>
);

const NewRegister = () => (
  <MuiThemeProvider>
    <Register />
  </MuiThemeProvider>
);

const SignIn = () => (
  <MuiThemeProvider>
    <Login />
  </MuiThemeProvider>
);

const Files = () => (
  <MuiThemeProvider>
    <FilesPage />
  </MuiThemeProvider>
);

const UserProfile = () => (
  <MuiThemeProvider>
    <ProfilePage />
  </MuiThemeProvider>
);

function requireAuth() {
  if (!loggedIn()) {
    alert('You are not authorised')
  }
  alert('did the thing')
};

//ReactDOM.render(<Home/>, document.getElementById('react-app'));

ReactDOM.render((
  <HashRouter>
    <div>
      <Route exact path="/" component={NewHome}/>
    <Route path="/login" component={SignIn}/>
    <Route path="/signup" component={NewRegister}/>
    <Route path="/files" component={Files} onEnter={requireAuth}/>
    <Route path="/profile" component={UserProfile} onEnter={requireAuth}/>
    </div>
    </HashRouter>
  ), document.getElementById('react-app'))
