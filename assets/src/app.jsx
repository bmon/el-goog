var React = require('react')
var ReactDOM = require('react-dom')

var HashRouter = require('react-router-dom').HashRouter
var Route = require('react-router-dom').Route
//var Link = require('react-router-dom').Link
//import {Router, Route, hashHistory } from 'react-router'

// importing cause requiring a class definition doesnt seem to work
import { Component } from 'react'
import FineUploaderTraditional from 'fine-uploader-wrappers'
import Gallery from 'react-fine-uploader'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Register from './Register';

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
        <div style={buttonStyle}>
          <h2>Login</h2>
         </div>
         <div style={buttonStyle}>
          <h2>Sign Up</h2>
         </div>
        <div className="content">
 		<Content/>
        </div>
      </div>
    )
  }
})

//==============================================================================
// React Material-UI
const MuiThemeExample = () => (
  <MuiThemeProvider>
    <Register />
  </MuiThemeProvider>
);


//==============================================================================
// React Fine Uploader Example
//


const uploader = new FineUploaderTraditional ({
    options: {
        chunking: {
            enabled: true
        },
        deleteFile: {
            enabled: true,
            endpoint: '/uploads'
        },
        request: {
            endpoint: '/uploads'
        },
        retry: {
            enableAuto: true
        }
    }
})

console.log(uploader)

class UploadComponent extends Component {
    render() {
        return (
            <Gallery uploader={ uploader } />
        )
    }
}

var UploadTest = new UploadComponent()


//==============================================================================
// David's sandbox
var DavidTest = React.createClass({
  render() {
    return (
      <div style={homeStyle}>
        <h1>Hey what are you doing here</h1>
        <Gallery uploader={ uploader  } />
      </div>
    )
  }
})


//ReactDOM.render(<Home/>, document.getElementById('react-app'));

ReactDOM.render((
  <HashRouter>
	<div>
		<Route exact path="/" component={Home}/>
		<Route path="/test" component={DavidTest}/>
		<Route path="/mui" component={MuiThemeExample}/> 
	</div>
  </HashRouter>
), document.getElementById('react-app'))


