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
  },
  container: {
    margin: 50,
    textAlign: 'center'
  },
  fileContainer: {
    margin: 90,
  }
};

const uploader = new FineUploaderTraditional ({
    options: {
        chunking: {
            // so max 1,2gb files for now
            enabled: false
        },
        deleteFile: {
            enabled: false,
            //endpoint: '/upload'
        },
        request: {
            endpoint: '/files',
            // using default names:
            //
            // filenameParam: 'qqfilename',
            // inputName: 'qqfile',
            // totalFileSizeName: 'qqtotalfilesize',
            // uuidName: 'qquuid',
        },
        retry: {
            enableAuto: true
        }

    }
})

//console.log(uploader)

class UploadComponent extends React.Component {
    render() {
        return (
            <Gallery uploader={ uploader } />
        )
    }
}

const HomePage = () => (
  <div style={styles.body}>
  <AppBar
    title={<span style={styles.title}></span>}
    onTitleTouchTap={handleTouchTap}
    iconElementLeft={<IconButton href="./"><NavigationClose /></IconButton>}
    iconElementRight={<RaisedButton style={styles.button} label=""><LoginPU /></RaisedButton>
	}
  />

  <Card style={styles.container}>
    <br/>
    <br/>
    <UploadComponent/>
    <div style={styles.fileContainer}>
    <List style={styles.container}>
      <Subheader inset={false}>Folders</Subheader>
      <ListItem
        leftAvatar={<Avatar icon={<FileFolder />} />}
        rightIcon={<ActionInfo />}
        primaryText="SampleFolder"
        secondaryText="Jan 9, 2014"
      />
      <ListItem
        leftAvatar={<Avatar icon={<FileFolder />} />}
        rightIcon={<ActionInfo />}
        primaryText="SampleFolder"
        secondaryText="Jan 17, 2014"
      />
      <ListItem
        leftAvatar={<Avatar icon={<FileFolder />} />}
        rightIcon={<ActionInfo />}
        primaryText="SampleFolder"
        secondaryText="Jan 28, 2014"
      />
    </List>
    <Divider inset={true} />
    <List style={styles.container}>
      <Subheader inset={false}>Files</Subheader>
      <ListItem
        leftAvatar={<Avatar icon={<ActionAssignment />} backgroundColor={blue500} />}
        rightIcon={<ActionInfo />}
        primaryText="SampleFile"
        secondaryText="Jan 20, 2014"
      />
      <ListItem
        leftAvatar={<Avatar icon={<EditorInsertChart />} backgroundColor={yellow600} />}
        rightIcon={<ActionInfo />}
        primaryText="SampleFile"
        secondaryText="Jan 10, 2014"
      />
    </List>
    </div>
  </Card>

  </div>
);

export default HomePage;
