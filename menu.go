//The menu package provides a simple analasys of a navigation structure
//This aims to take strings formatted similar to
//name:dest
//name2
//{
//  inner:menu
//}
//In this name2 holds an inner menu with 1 element, namely "inner"
package menu

import (
	"errors"
	"strconv"
	"strings"
)

type MenuEntry struct {
	Name     string
	Dest     string
	Children []*MenuEntry
}

//NewMenu Creates a new menu object from string array (lines)
func NewMenu(ar []string) (*MenuEntry, error) {
	res := MenuEntry{"TOP", "", nil}
	chids, _, err := newMenu(ar, 0)
	res.Children = chids
	return &res, err
}

//newMenu uses the array of strings to create a menu struct, p is the current array position- use 0 for beginning
func newMenu(ar []string, p int) ([]*MenuEntry, int, error) {

	res := make([]*MenuEntry, 0)
	var curr *MenuEntry

	for i := p; i < len(ar); i++ {
		s := strings.Trim(ar[i], "\t \r")
		if s == "{" {
			if i == p {
				return res, i, errors.New("No parent for Line - " + strconv.Itoa(i))
			}
			chids, ni, err := newMenu(ar, i+1)
			curr.Children = chids
			if err != nil {
				return res, ni, err
			}
			i = ni
		} else if s == "}" {
			return res, i, nil
		} else {
			a := strings.Split(s, ":")
			b := a[0]
			if len(a) > 1 {
				b = a[1]
			}
			curr = &MenuEntry{a[0], b, nil}
			res = append(res, curr)

		}
	}
	return res, len(ar), nil
}

func (self *MenuEntry) String() string {
	res := self.Name + "--" + self.Dest + "("
	for i := 0; i < len(self.Children); i++ {
		res += self.Children[i].String()
		if i+1 < len(self.Children) {
			res += ","
		}
	}
	res += ")"
	return res
}
