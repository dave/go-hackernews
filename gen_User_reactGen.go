// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type UserElem struct {
	react.Element
}

func buildUser(cd react.ComponentDef) react.Component {
	return UserDef{ComponentDef: cd}
}

func buildUserElem(props UserProps, children ...react.Element) *UserElem {
	return &UserElem{
		Element: react.CreateElement(buildUser, props, children...),
	}
}

func (u UserDef) RendersElement() react.Element {
	return u.Render()
}

// IsProps is an auto-generated definition so that UserProps implements
// the myitcv.io/react.Props interface.
func (u UserProps) IsProps() {}

// Props is an auto-generated proxy to the current props of User
func (u UserDef) Props() UserProps {
	uprops := u.ComponentDef.Props()
	return uprops.(UserProps)
}

func (u UserProps) EqualsIntf(val react.Props) bool {
	return u == val.(UserProps)
}

var _ react.Props = UserProps{}