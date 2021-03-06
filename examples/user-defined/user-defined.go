// User-defined gorbac example
package main

import (
	"fmt"

	"github.com/mikespook/gorbac"
)

type MyRole struct {
	gorbac.Role
	Label       string
	Description string
}

func loadByName(name string) (label, description string) {
	// loading data from storages or somewhere
	return name, "This is the description for " + name
}

func NewMyRole(rbac *gorbac.Rbac, name string) gorbac.Role {
	// loading extra properties by `name`.
	label, desc := loadByName(name)
	role := &MyRole{
		Role:        gorbac.NewBaseRole(rbac, name),
		Label:       label,
		Description: desc,
	}
	return role
}

func main() {
	rbac := gorbac.NewWithFactory(NewMyRole)
	rbac.Add("role-1", []string{"a", "b", "c"}, nil)
	role := rbac.Get("role-1")
	if myRole, ok := role.(*MyRole); ok {
		fmt.Printf("Name:\t%s\nLabel:\t%s\nDesc:\t%s\n",
			myRole.Name(), myRole.Label, myRole.Description)
	}
}
