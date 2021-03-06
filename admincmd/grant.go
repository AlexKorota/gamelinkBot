package admincmd

import (
	"context"
	"gamelink-bot/command_list"
	"gamelink-bot/iface"
	"gamelink-bot/parser"
	"gamelink-bot/service"
	"strings"
)

type (
	//GrantFabric - struct for Grant fabric
	GrantFabric struct{}
	//GrantCommand - struct for grant command
	GrantCommand struct {
		userName string
		params   []string
		res      iface.Responder
	}
)

//init - func for register fabric in parser
func init() {
	parser.SharedParser().RegisterFabric(GrantFabric{})
}

//CommandName - return human readable command name
func (c GrantFabric) CommandName() string {
	return command_list.CommandGrants
}

//RequireAdmin - func for checking if admin permissions required
func (c GrantFabric) RequireAdmin() bool {
	return true
}

//Require - return array of needed permissions
func (c GrantFabric) Require() []string {
	return []string{command_list.CommandGrants}
}

//TryParse - func for parsing request
func (c GrantFabric) TryParse(req iface.RequesterResponder) (iface.Command, error) {
	var (
		command GrantCommand
		err     error
	)
	if command.userName, command.params, err = service.CompareParsePermissionCommand(req.Request(), "/"+command_list.CommandGrants); err != nil {
		if err == service.UnknownCommandError {
			return nil, nil
		}
		return nil, err
	}
	command.res = req
	return command, nil
}

//Execute - execute command
func (cc GrantCommand) Execute(ctx context.Context) {
	user, err := Executor().GrantPermissions(cc.userName, cc.params)
	if err != nil {
		cc.res.Respond(err.Error())
		return
	}
	cc.res.Respond("Success " + user.Name + " now has next permissions: " + strings.Join(user.Permissions, ", "))
}
