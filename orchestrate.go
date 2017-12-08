package libcompose

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
)

/**
 * Orchestration is handled via Docker using the libCompose library
 */

func (b *Builder) orchestrateOperations() api.Operations {
	baseOp := NewOperationBase(b.settings)

	ops := base.NewOperations()

	// include the Up operation
	ops.Add(NewOrchestrateUpOperation(*baseOp))

	return ops.Operations()
}
