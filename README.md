# el-goog
the ethical file sync and backup system

## installation:
1. install go1.9 for your system via https://golang.org/doc/install
    by default, the instalation uses `~/go/` as the default workspace
    where you will store code. If you want to change this, you'll also
    need to set the $GOPATH env variable:

      `$ export GOPATH=$HOME/uni/compethics/go`

2. `$ export PATH=$PATH:$(go env GOPATH)/bin`
3. `$ go get -u github.com/golang/dep/cmd/dep`
4. ```
   $ cd $GOPATH/src/
   $ mkdir bmon
   ```

   then:

   `$ git clone git@github.com:bmon/el-goog.git`

   OR

   `$ git clone https://github.com/bmon/el-goog.git`

   (do the second one if you haven't set up your ssh key with github)

   lastly do

   `git checkout develop`

5. `$ cd el-goog`
   `$ dep ensure`
   this will install any extra (golang) libraries we've added to the code

6. install nvm: https://github.com/creationix/nvm

6. last, run `go build`, then
   `$ ./el-goog`
   and go to http://localhost:8000 in your browser



## making changes:

to make changes to the assingment:

```
$ git pull
$ dep ensure
```

_make your changes to the code_

```
$ git add thefile/i_changed.go theother/file_i_added.jsx
$ git commit -m "add profile delete button
\
\ create new react template file
\ add new profile delete route"
$ git push
```
