package libcompose

import (
	"context"

	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"

	libcompose_project_options "github.com/docker/libcompose/project/options"
)

const (
	OPERATION_ID_ORCHESTRATE_DOWN = "orchestrate.down"
)

// OrchestrateDownOperation orchestrate down operation using libcompose
type OrchestrateDownOperation struct {
	OperationBase
}

// NewOrchestrateDownOperation constructor for OrchestrateDownOperation
func NewOrchestrateDownOperation(base OperationBase) *OrchestrateDownOperation {
	return &OrchestrateDownOperation{
		OperationBase: base,
	}
}

func (odo *OrchestrateDownOperation) Operation() api.Operation {
	return api.Operation(odo)
}

func (odo *OrchestrateDownOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_DOWN
}

func (odo *OrchestrateDownOperation) Ui() api.Ui {
	return base.NewUi(
		odo.Id(),
		"Orchestrate down",
		"Bring down the application composition",
		"",
	)
}

func (odo *OrchestrateDownOperation) Usage() api.Usage {
	return (&base.ExternalOperationUsage{}).Usage()
}

func (odo *OrchestrateDownOperation) Properties() api.Properties {
	props := base.NewProperties()

	return props.Properties()
}

func (odo *OrchestrateDownOperation) Validate(props api.Properties) api.Result {
	return base.MakeSuccessfulResult()
}

func (odo *OrchestrateDownOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func() {
		if APIProject, err := odo.APIProject(); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {

			var ctx context.Context
			var downOptions libcompose_project_options.Down

			// net context
			if settingsCtx := odo.Settings().Context; &settingsCtx != nil {
				ctx = settingsCtx
			} else {
				ctx = context.Background()
			}
			
			// up options
			downOptions = libcompose_project_options.Down{}
			if downOptionsProp, err := props.Get(PROPERTY_ID_REMOVEVOLUMES); err == nil {
				downOptions.RemoveVolume = downOptionsProp.Get().(bool)
			}
			if downOptionsProp, err := props.Get(PROPERTY_ID_REMOVEIMAGETYPES); err == nil {
				downOptions.RemoveImages = libcompose_project_options.ImageType(downOptionsProp.Get().(string))
			}
			if downOptionsProp, err := props.Get(PROPERTY_ID_REMOVEORPHANS); err == nil {
				downOptions.RemoveOrphans = downOptionsProp.Get().(bool)
			}

			if res.Success() {
				if err := APIProject.Down(ctx, downOptions); err != nil {
					res.MarkFailed()
					res.AddError(err)
				} else {
					res.MarkSucceeded()
				}
			}

		}

		res.MarkFinished()
	}()

	return res.Result()
}