package errors

import "errors"

var NoConfig = errors.New("couldn't find configuration in the database")
var NoRoles = errors.New("couldn't find role command configuration in the database")
var CreatedConfig = errors.New("couldn't find configuration in the database. Created default one")
var NoModule = errors.New("couldn't find registered module")
var MessageNotSent = errors.New("couldn't replace message, since it doesn't exist")
var NoWelcomer = errors.New("couldn't find welcomer:main, welcomer:text:after or welcomer:text:before in the database")