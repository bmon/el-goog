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

import Register from './Register';
import LoginPU from './LoginPU';

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
  }
};

const HomePage = () => (
  <div style={styles.body}>
  <AppBar
    title={<span style={styles.title}></span>}
    onTitleTouchTap={handleTouchTap}
    iconElementLeft={<IconButton href="./"><NavigationClose /></IconButton>}
    iconElementRight={<RaisedButton style={styles.button}><LoginPU /></RaisedButton>
	}
  />

  <Card>
    <CardMedia>
      <img src="logo.png" alt="el-goog logo" />
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
