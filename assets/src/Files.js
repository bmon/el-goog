import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import ActionHome from 'material-ui/svg-icons/action/home';
import DeleteButton from 'material-ui/svg-icons/action/delete';
import FlatButton from 'material-ui/FlatButton';
import {Card, CardActions, CardHeader, CardMedia, CardTitle, CardText} from 'material-ui/Card';
import Divider from 'material-ui/Divider';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import FineUploaderTraditional from 'fine-uploader-wrappers'

import Gallery from 'react-fine-uploader'
import { Component } from 'react'
import {List, ListItem} from 'material-ui/List';
import Subheader from 'material-ui/Subheader';
import Avatar from 'material-ui/Avatar';
import FileFolder from 'material-ui/svg-icons/file/folder';
import ActionAssignment from 'material-ui/svg-icons/action/assignment';
import {blue500, yellow600} from 'material-ui/styles/colors';
import EditorInsertChart from 'material-ui/svg-icons/editor/insert-chart';
import Cookie from 'js-cookie';

import IconMenu from 'material-ui/IconMenu';
import FontIcon from 'material-ui/FontIcon';
import NavigationExpandMoreIcon from 'material-ui/svg-icons/navigation/expand-more';
import MenuItem from 'material-ui/MenuItem';
import DropDownMenu from 'material-ui/DropDownMenu';
import {Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle} from 'material-ui/Toolbar';

import axios from "axios";

import Register from './Register';
import LoginPU from './LoginPU';
import LogoutPU from './LogoutPU';
import Header from './Header';


import '../dist/gallery.css'

// currently unused
function handleTouchTap() {
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
    margin: 55,
  },
  mediumIcon: {
    width: 35,
    height: 35,
  },
  fileList: {
    textAlign: 'left'
  }
};


class UploadComponent extends React.Component {
    render() {
        window.folderID = Cookie.get("root_id")
        var uploader = new FineUploaderTraditional ({
            options: {
                chunking: {
                    // so max 1,2gb files for now
                    enabled: true,
                    mandatory: true,
                    // partSize: 2000000,

                },
                deleteFile: {
                    enabled: false,
                    //endpoint: '/upload'
                },
                request: {
                    endpoint: '/folders/'+folderID+'/files',
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

        return (
            <Gallery uploader={ uploader } />
        )
    }
}
//
const Files = () => (
  <div style={styles.body}>
  <Header />

  <Card style={styles.container}>
    <CardTitle title="All Files" />
    <br/>
    <UploadComponent/>
    <ObjectList/>
  </Card>
  </div>
);

class ObjectList extends Component {
  constructor(props){
    super(props);
    this.state = {
      files: [],
      folders: [],
      path: [],
      parent_id: -1,
    }
    var _this = this;
    axios.get("/folders/"+folderID)
    .then(function(result) {
      _this.setState({
        files: result.data.child_files,
        folders: result.data.child_folders,
        parent_id: result.data.parent_id,
        path: result.data.path.slice(0,-1).split("/").slice(2),
      });
    })
    this.gotoParent = this.gotoParent.bind(this)
  }

  updateLoc(id) {
    Cookie.set("root_id", id)
    location.reload()
  }
  downloadFile(id) {
    var link = document.createElement("a");
    link.href = "/files/"+id;
    link.click();
  }
  gotoParent() {
    if (this.state.parent_id > 0) {
      Cookie.set("root_id", this.state.parent_id)
      location.reload()
    }
  }

  render() {
    const _this = this
    const renderFiles = this.state.files.map(function(item, i) {
      return (
        <ListItem
        leftAvatar={<Avatar icon={<Avatar icon={<EditorInsertChart />} backgroundColor={yellow600} />} />}
        onClick={function (id) {_this.downloadFile(item.id)}}
        rightIcon={<DeleteButton />}
        primaryText={item.name}
        secondaryText={item.size + " bytes"}
        />
      )
    });
    const renderFolders = this.state.folders.map(function(item, i) {
      return (
        <ListItem
        leftAvatar={<Avatar icon={<FileFolder />} />}
        onClick={function (id) {_this.updateLoc(item.id)}}
        rightIcon={<DeleteButton />}
        primaryText={item.name}
        secondaryText={item.modified}
        />
      )
    });
    const renderPath = this.state.path.map(function(item, i) {
      var parts = item.split('.')
      var id = parts.pop()
      var name = parts.join()

      return (
        <RaisedButton
        style={styles.button}
        label={name}
        onClick={function() {_this.updateLoc(item.id)}}
        />
      )
    });
    return (
      <div>
        {renderPath}
        <div style={styles.fileContainer}>
        <Toolbar>
        <ToolbarGroup firstChild={true}>
          <RaisedButton style={styles.button} label="New Folder" />
        </ToolbarGroup>
        <ToolbarGroup>
          <ToolbarSeparator />

          <IconMenu
            iconButtonElement={
              <IconButton touch={true}>
                <NavigationExpandMoreIcon />
              </IconButton>
            }
          >
            <MenuItem primaryText="Size" />
            <MenuItem primaryText="Name" />
          </IconMenu>

        </ToolbarGroup>
      </Toolbar>



        <List style={styles.fileList}>
          <Subheader inset={false}>Files</Subheader>
          {renderFiles}
        </List>
        </div>
      </div>
    )
  }
}

class Fil extends Component {
  constructor(props){
    super(props);
    this.state = {
      items: []
    }
  }

  componentDidMount() {
    var _this = this;
    axios.get("/folders/"+folderID)
    .then(function(result) {
      _this.setState({
      items: result.data.items
      });
    })
  }

 /* componentWillUnmount() {
    this.serverRequest.abort();
  },
*/
  render() {
    const renderItems = this.state.items.map(function(item, i) {
      return <li key={i}>{item.title}</li>
    });
    return (
      <div>
        {renderItems}
        /* Render stuff here */
      </div>
    )
  }
}


export default Files;
