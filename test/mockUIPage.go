package test

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
)

type MockUIPage struct {
	name          string
	parent        components.UIComponent
	componentType components.NodeType
}

func NewMockUIComponent(name string, componentType components.NodeType, parent components.UIComponent) MockUIPage {
	return MockUIPage{name: name, componentType: componentType, parent: parent}
}

func (m MockUIPage) Render(req models.RequestWrapper) *components.RenderResponse {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) Type() components.NodeType {
	return m.componentType
}

func (m MockUIPage) Name() string {
	return m.name
}

func (m MockUIPage) UpdateSpec() *models.TreeSpec {
	return &models.TreeSpec{
		Name: m.name,
	}
}

func (m MockUIPage) GetSpec() *models.TreeSpec {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) GetChildren() []components.UIComponent {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) FindChild(name string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) FindChildByType(name string, componentType string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) GetParent() components.UIComponent {
	return m.parent
}

func (m MockUIPage) Id() string {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) FindChildById(id string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) SetParent(parent components.UIComponent) {
	m.parent = parent
}

func (m MockUIPage) AddChild(child components.UIComponent) error {
	//TODO implement me
	panic("implement me")
}

func (m MockUIPage) RemoveChild(child components.UIComponent) error {
	//TODO implement me
	panic("implement me")
}
