// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type HeaderElem struct {
	react.Element
}

func buildHeader(cd react.ComponentDef) react.Component {
	return HeaderDef{ComponentDef: cd}
}

func buildHeaderElem(props HeaderProps, children ...react.Element) *HeaderElem {
	return &HeaderElem{
		Element: react.CreateElement(buildHeader, props, children...),
	}
}

func (h HeaderDef) RendersElement() react.Element {
	return h.Render()
}

// IsProps is an auto-generated definition so that HeaderProps implements
// the myitcv.io/react.Props interface.
func (h HeaderProps) IsProps() {}

// Props is an auto-generated proxy to the current props of Header
func (h HeaderDef) Props() HeaderProps {
	uprops := h.ComponentDef.Props()
	return uprops.(HeaderProps)
}

func (h HeaderProps) EqualsIntf(val react.Props) bool {
	return h == val.(HeaderProps)
}

var _ react.Props = HeaderProps{}
