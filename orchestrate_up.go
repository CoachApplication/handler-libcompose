package libcompose

import (
	"context"

	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"

	libcompose_project_options "github.com/docker/libcompose/project/options"
)

const (
	OPERATION_ID_ORCHESTRATE_UP = "orchestrate.up"
)

// OrchestrateUpOperation orchestrate up operation using libcompose
type OrchestrateUpOperation struct {
	OperationBase
}

// NewOrchestrateUpOperation constructor for OrchestrateUpOperation
func NewOrchestrateUpOperation(base OperationBase) *OrchestrateUpOperation {
	return &OrchestrateUpOperation{
		OperationBase: base,
	}
}

func (ouo *OrchestrateUpOperation) Operation() api.Operation {
	return api.Operation(ouo)
}

func (ouo *OrchestrateUpOperation) Id() string {
	return OPERATION_ID_ORCHESTRATE_UP
}

func (ouo *OrchestrateUpOperation) Ui() api.Ui {
	return base.NewUi(
		ouo.Id(),
		"Orchestrate up",
		"Bring up the application composition",
		"",
	)
}

func (ouo *OrchestrateUpOperation) Usage() api.Usage {
	return (&base.ExternalOperationUsage{}).Usage()
}

func (ouo *OrchestrateUpOperation) Properties() api.Properties {
	props := base.NewProperties()

	return props.Properties()
}

func (ouo *OrchestrateUpOperation) Validate(props api.Properties) api.Result {
	return base.MakeSuccessfulResult()
}

func (ouo *OrchestrateUpOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func() {
		if APIProject, err := ouo.APIProject(); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {

			var ctx context.Context
			var upOptions libcompose_project_options.Up

			// net context
			if settingsCtx := ouo.Settings().Context; &settingsCtx != nil {
				ctx = settingsCtx
			} else {
				ctx = context.Background()
			}

			// up options
			upOptions = libcompose_project_options.Up{}
			if upOptionsProp, err := props.Get(PROPERTY_ID_NORECREATE); err == nil {
				upOptions.NoRecreate = upOptionsProp.Get().(bool)
			}
			if upOptionsProp, err := props.Get(PROPERTY_ID_FORCERECREATE); err == nil {
				upOptions.ForceRecreate = upOptionsProp.Get().(bool)
			}
			if upOptionsProp, err := props.Get(PROPERTY_ID_NOBUILD); err == nil {
				upOptions.NoBuild = upOptionsProp.Get().(bool)
			}
			if upOptionsProp, err := props.Get(PROPERTY_ID_FORCEREBUILD); err == nil {
				upOptions.ForceBuild = upOptionsProp.Get().(bool)
			}

			if res.Success() {
				if err := APIProject.Up(ctx, upOptions); err != nil {
					res.AddError(err)
					res.MarkFailed()
				} else {
					res.MarkSucceeded()
				}
			}

		}

		res.MarkFinished()
	}()

	return res.Result()
}