#!/usr/bin/bash

# Update go dependencies

# cd to the root of the project directory (the directory that contains the .git directory)
cd "$(git rev-parse --show-toplevel)" || exit 1

# save the current changes
git stash

# Update dependencies
go get -u -v github.com/gogf/gf/contrib/drivers/pgsql/v2@master
go get -u -v github.com/gogf/gf/v2@master
go get -u -v all

git add .
git commit -a -m "update deps"

# restore the current changes
git stash pop

# exit
cd - || exit 1
