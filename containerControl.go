package gonsole

type ContainerControl struct {
	BasicControl

	children []Control
}

func (ctrl *ContainerControl) Repaint() {
	ctrl.BasicControl.Repaint()
	// repaint children
	for _, child := range ctrl.Children() {
		child.Repaint()
	}
}

func (ctrl *ContainerControl) AddControl(control Control) {
	if ctrl.children == nil {
		ctrl.children = make([]Control, 0)
	}
	control.SetParent(ctrl)

	ctrl.children = append(ctrl.children, control)
}

func (ctrl *ContainerControl) Children() []Control {
	return ctrl.children
}
