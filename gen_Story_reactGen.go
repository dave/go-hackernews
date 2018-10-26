// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type StoryElem struct {
	react.Element
}

func buildStory(cd react.ComponentDef) react.Component {
	return StoryDef{ComponentDef: cd}
}

func buildStoryElem(props StoryProps, children ...react.Element) *StoryElem {
	return &StoryElem{
		Element: react.CreateElement(buildStory, props, children...),
	}
}

func (s StoryDef) RendersElement() react.Element {
	return s.Render()
}

// IsProps is an auto-generated definition so that StoryProps implements
// the myitcv.io/react.Props interface.
func (s StoryProps) IsProps() {}

// Props is an auto-generated proxy to the current props of Story
func (s StoryDef) Props() StoryProps {
	uprops := s.ComponentDef.Props()
	return uprops.(StoryProps)
}

func (s StoryProps) EqualsIntf(val react.Props) bool {
	return s.Equals(val.(StoryProps))
}

var _ react.Props = StoryProps{}
