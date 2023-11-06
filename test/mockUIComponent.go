package test

import (
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pkg/components"
)

type MockUIComponent struct {
	name          string
	parent        components.UIComponent
	componentType components.NodeType
}

func NewMockUIComponent(name string, componentType components.NodeType, parent components.UIComponent) MockUIComponent {
	return MockUIComponent{name: name, componentType: componentType, parent: parent}
}

func (m MockUIComponent) Render(req models.RequestWrapper) *components.RenderResponse {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) Type() components.NodeType {
	return m.componentType
}

func (m MockUIComponent) Name() string {
	return m.name
}

func (m MockUIComponent) UpdateSpec() *models.TreeSpec {
	return &models.TreeSpec{
		Name: m.name,
	}
}

func (m MockUIComponent) GetSpec() *models.TreeSpec {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) GetChildren() []components.UIComponent {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) FindChild(name string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) FindChildByType(name string, componentType string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) GetParent() components.UIComponent {
	return m.parent
}

func (m MockUIComponent) Id() string {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) FindChildById(id string) (components.UIComponent, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) SetParent(parent components.UIComponent) {
	m.parent = parent
}

func (m MockUIComponent) AddChild(child components.UIComponent) error {
	//TODO implement me
	panic("implement me")
}

func (m MockUIComponent) KillChild(child components.UIComponent) error {
	//TODO implement me
	panic("implement me")
}
