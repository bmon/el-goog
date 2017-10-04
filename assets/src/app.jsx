var React = require('react')
var ReactDOM = require('react-dom')

//var Router = require('react-router').Router
//var Route = require('react-router')

import {Router, Route, hashHistory } from 'react-router'

  	var homeStyle = {
        padding: 10,
        margin: 10,
        backgroundColor: "#ffffff",
        color: "#333",
        display: "inline-block",
        fontFamily: "monospace",
        fontSize: 32,
        textAlign: "center"
      };

     var contentStyle = {
        padding: 10,
        margin: 10,
        backgroundColor: "#E5E1E0",
        color: "#333",
        display: "inline-block",
        fontFamily: "monospace",
        fontSize: 28,
        textAlign: "center"
      };

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
  	  };

var Content = React.createClass({
  render: function() {
      return (
        <div style={contentStyle}>
          <h2>How el-Goog works</h2>
          <p>[insert all of the pictures about how el-goog works]</p>
  
          <p>[insert all of the pictures about how el-goog works]</p>
        </div>
      );
    }
});

var Home = React.createClass({
  render: function() {
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

//ReactDOM.render(<Home/>, document.getElementById('react-app'));

render((
  <Router history={hashHistory}>
    <Route path="/" component={Home}/>
 
  </Router>
  ), document.getElementById('react-app'))


