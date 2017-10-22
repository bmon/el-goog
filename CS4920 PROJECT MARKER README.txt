Welcome to el-goog!

DESCRIPTION:
el-goog allows you to create, login, edit and delete an account
with your account, you own a file directory and can create and delete folders
within each folder, you can upload and delete files!

TO RUN THE PROJECT:
just run ./el-goog 
then check out localhost:5000 in a web browser

TO REBUILD THE PROJECT:
~~~~~ to save submission space, 
~~~~~ nodejs modules and golang dependencies were not packaged in the zip
youll need to fetch nodejs modules and configure golang.
ensure you have an internet connection and a unix (preferably linux) environment
follow the readme.md for golang setup, "dep" installation, nodejs install
then simply run ./rebuild.sh
rebuild.sh does all the necessary steps after your golang environment is
configured

TECH:
golang backend files are in the root directory to package main
web assets are in assets/static, and webpack also copies them to assets/dist
javascript jsx files are in assets/src
database migrations are in migrations, each one adds to the final schema

