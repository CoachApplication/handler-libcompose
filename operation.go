package libcompose

import libcompose_project "github.com/docker/libcompose/project"

// OperationBse base operations class which contains and can provide libcompose settings
type OperationBase struct {
	settings Settings
}

func NewOperationBase(settings Settings) *OperationBase {
	return &OperationBase{
		settings: settings,
	}
}

// Settings get the settings for the project
func (ob *OperationBase) Settings() Settings {
	return ob.settings
}

// APIProject from the libcompose library
func (ob *OperationBase) APIProject() (libcompose_project.APIProject, error) {
	project, err := ob.settings.LibComposeProject()
	return project, err
}
