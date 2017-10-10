# whether or not you actually run this, its a nice reminder
# THIS KILLS THE DATABASE
rm elgoog.db
goose -dir migrations/ sqlite3 ./elgoog.db up
dep ensure
go build
npm install
./node_modules/.bin/webpack --config webpack.config.js
./el-goog
