# ap(Apport)

## Usage(Only tested in windows operation system for now)

Make self defined shortcut command with "ap"

(I was rushed to get these `ap` functions and also I'm a Go newbie, so this is a very very immature project.)

like:
ap t [alias] eg:`ap t p12 js` means locate to the directory with alias ["p12","js"](It open another cmd window, and it's just a workarround)
ap [shortcut] [arguments...]`ap open .` means excute the application which shortcut "open" is pointed to.

## Install
go get github.com/glight2000/ap

## Config
Excute `ap init` will create a configuration file in GOBIN directory, and make sure the enviroment variable GOBIN is exist before that.

Then write customized configuration to ap.cfg in GOBIN directory.

The config is a json file.

The `to` array defines path targets' alias which are used follow with `ap t`. Alias is a array, so we can write it like `["ide","vs"]` or `["tool","vs"]`

The `customize`(bad name...) defines applications' shortcut which are used follow with. `argument_filter` is aregular used to seprate actions with same application, it only checks the first argument after the `ap [shortcut]`. `isArgumentsInherit` means if the arguments after `ap [shortcut]` should be passed to the application.
