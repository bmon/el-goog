#!/bin/bash -x
# whether or not you actually run this, its a nice reminder
# THIS KILLS THE DATABASE
dep ensure
rm elgoog.db
goose -dir migrations/ sqlite3 ./elgoog.db up
go build -i
npm install
./node_modules/.bin/webpack --config webpack.config.js
go build
echo "build finished, running el-goog"
./el-goog
