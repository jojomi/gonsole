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

func (ctrl *ContainerControl) ChildrenDeep() []Control {
	controls := make([]Control, 0)
	for _, control := range ctrl.children {
		container, ok := control.(*ContainerControl)
		if ok {
			children := container.ChildrenDeep()
			for _, child := range children {
				controls = append(controls, child)
			}
		} else {
			controls = append(controls, control)
		}
	}
	return controls
}
