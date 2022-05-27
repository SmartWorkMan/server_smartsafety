package safety

type RouterGroup struct {
	SafetyFactoryRouter
	InspectorRouter
	AreaRouter
	ItemRouter
	TaskRouter
}
