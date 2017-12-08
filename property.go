package libcompose

import (
	// libCompose_options "github.com/docker/libcompose/project/options"

	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_property "github.com/CoachApplication/base/property"
)

const (
	/**
	 * Configs used in all operations to build the libCompose project
	 */

	// config for a project name (not really used)
	PROPERTY_ID_PROJECTNAME = "compose.projectname"
	// config for a project yml files (not really used)
	PROPERTY_ID_COMPOSEFILES = "compose.composefiles"

	// Input/Output objects
	PROPERTY_ID_OUTPUT = "compose.output"
	PROPERTY_ID_ERROR  = "compose.error"

	/**
	 * General Properties for most operations
	 */

	// config for an orchestration context limiter
	PROPERTY_ID_CONTEXT = "compose.context"

	/**
	 * Operation specific contexts
	 */

	// Individual possible libcompose properties
	PROPERTY_ID_FORCEREMOVE      = "compose.forceremove"
	PROPERTY_ID_NOCACHE          = "compose.nocache"
	PROPERTY_ID_PULL             = "compose.pull"
	PROPERTY_ID_DETACH           = "compose.detach"
	PROPERTY_ID_NORECREATE       = "compose.norecreate"
	PROPERTY_ID_NOBUILD          = "compose.nobuild"
	PROPERTY_ID_FORCERECREATE    = "compose.forcerecreate"
	PROPERTY_ID_FORCEREBUILD     = "compose.forcerebuild"
	PROPERTY_ID_REMOVEVOLUMES    = "compose.removevolumes"
	PROPERTY_ID_REMOVEORPHANS    = "compose.removeorphans"
	PROPERTY_ID_REMOVEIMAGETYPES = "compose.removeimagetypes"
	PROPERTY_ID_REMOVERUNNING    = "compose.removerunning"
	PROPERTY_ID_TIMEOUT          = "compose.timeout"
)

/**
 * Properties which the libCompose handler uses
 */

// Project Name Property for a docker.libCompose project
type ProjectnameProperty struct {
	base_property.StringProperty
}

// Id for the Property
func (name *ProjectnameProperty) Id() string {
	return PROPERTY_ID_PROJECTNAME
}

// Label for the Property
func (name *ProjectnameProperty) Label() string {
	return "Project name"
}

// Description for the Property
func (name *ProjectnameProperty) Description() string {
	return "Compose project name, which is used in container, volume and network naming."
}

// Is the Property internal only
func (name *ProjectnameProperty) Usage() api.Usage {
	return base.InternalPropertyUsage{}.Usage()
}

// YAML file list Property for a docker.libCompose project
type ComposefilesProperty struct {
	base_property.StringSliceProperty
}

// Id for the Property
func (files *ComposefilesProperty) Id() string {
	return PROPERTY_ID_COMPOSEFILES
}

// Label for the Property
func (files *ComposefilesProperty) Label() string {
	return "docker-compose yml file list"
}

// Description for the Property
func (files *ComposefilesProperty) Description() string {
	return "An ordered list of docker-compose yml files, which are passed to libcompose."
}

// Is the Property internal only
func (files *ComposefilesProperty) Usage() api.Usage {
	return base.InternalOperationUsage{}.Usage()
}

// A libcompose Property for net context limiting
type ContextProperty struct {
	base_property.ContextProperty
}

// Id for the Property
func (contextConf *ContextProperty) Id() string {
	return PROPERTY_ID_CONTEXT
}

// Label for the Property
func (contextConf *ContextProperty) Label() string {
	return "context limiter"
}

// Description for the Property
func (contextConf *ContextProperty) Description() string {
	return "A context for controling execution."
}

// Is the Property internal only
func (contextConf *ContextProperty) Usage() api.Usage {
	return base.InternalPropertyUsage{}
}

// Output handler Property for a docker.libCompose project
type OutputProperty struct {
	base_property.WriterProperty
}

// Id for the Property
func (output *OutputProperty) Id() string {
	return PROPERTY_ID_OUTPUT
}

// Label for the Property
func (output *OutputProperty) Label() string {
	return "Output writer"
}

// Description for the Property
func (output *OutputProperty) Description() string {
	return "Output io.Writer which will receive compose output from containers."
}

// Is the Property internal only
func (output *OutputProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// Error handler Property for a docker.libCompose project
type ErrorProperty struct {
	base_property.WriterProperty
}

// Id for the Property
func (err *ErrorProperty) Id() string {
	return PROPERTY_ID_ERROR
}

// Label for the Property
func (err *ErrorProperty) Label() string {
	return "Error writer"
}

// Description for the Property
func (err *ErrorProperty) Description() string {
	return "Error io.Writer which will receive compose output from containers."
}

// Is the Property internal only
func (err *ErrorProperty) Usage() api.Usage {
	return base.InternalPropertyUsage{}.Usage()
}

/**
 * These Properties are wrappers for the various libCompose options
 * structs in https://github.com/docker/libcompose/blob/master/project/options/types.go
 */

// BUILD : Property for a docker.libCompose project to indicate that a build should ignore cached image layers
type NoCacheProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (nocache *NoCacheProperty) Id() string {
	return PROPERTY_ID_NOCACHE
}

// Label for the Property
func (nocache *NoCacheProperty) Label() string {
	return "nocache"
}

// Description for the Property
func (nocache *NoCacheProperty) Description() string {
	return "When capturing building, ignore cached docker layers?"
}

// Is the Property internal only
func (nocache *NoCacheProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// Property for a docker.libCompose project to indicate that a process remove ... something
type ForceRemoveProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (forceremove *ForceRemoveProperty) Id() string {
	return PROPERTY_ID_FORCEREMOVE
}

// Label for the Property
func (forceremove *ForceRemoveProperty) Label() string {
	return "Force remove"
}

// Description for the Property
func (forceremove *ForceRemoveProperty) Description() string {
	return "When building, force remove .... something?"
}

// Is the Property internal only
func (forceremove *ForceRemoveProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// Property for a docker.libCompose project to indicate that a process hsould stay attached and follow
type PullProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (pull *PullProperty) Id() string {
	return PROPERTY_ID_PULL
}

// Label for the Property
func (pull *PullProperty) Label() string {
	return "Pull"
}

// Description for the Property
func (pull *PullProperty) Description() string {
	return "When building, pull all images before using them?"
}

// Is the Property internal only
func (pull *PullProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// Property for a docker.libCompose project to indicate that a process hsould stay attached and follow
type DetachProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (detach *DetachProperty) Id() string {
	return PROPERTY_ID_DETACH
}

// Label for the Property
func (detach *DetachProperty) Label() string {
	return "Detach"
}

// Description for the Property
func (detach *DetachProperty) Description() string {
	return "When capturing output, detach from the output?"
}

// Is the Property internal only
func (detach *DetachProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// UP : Property for a docker.libCompose project to indicate that a process should not create missing containers
type NoRecreateProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (norecreate *NoRecreateProperty) Id() string {
	return PROPERTY_ID_NORECREATE
}

// Label for the Property
func (norecreate *NoRecreateProperty) Label() string {
	return "Create"
}

// Description for the Property
func (norecreate *NoRecreateProperty) Description() string {
	return "When starting a container, create it first, if it is missing?"
}

// Is the Property internal only
func (norecreate *NoRecreateProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// UP|RECREATE : Property for a docker.libCompose project to indicate that a process should build containers even if they are found
type ForceRecreateProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (forcerecreate *ForceRecreateProperty) Id() string {
	return PROPERTY_ID_FORCERECREATE
}

// Label for the Property
func (forcerecreate *ForceRecreateProperty) Label() string {
	return "Force Recreate"
}

// Description for the Property
func (forcerecreate *ForceRecreateProperty) Description() string {
	return "Force recreating containers, even if they exist already?"
}

// Is the Property internal only
func (forcerecreate *ForceRecreateProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// UP|CREATE : Property for a docker.libCompose project to indicate that a process should not build any containers
type NoBuildProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (dontbuild *NoBuildProperty) Id() string {
	return PROPERTY_ID_NOBUILD
}

// Label for the Property
func (dontbuild *NoBuildProperty) Label() string {
	return "Don't Build"
}

// Description for the Property
func (dontbuild *NoBuildProperty) Description() string {
	return "Don't build any missing images?"
}

// Is the Property internal only
func (dontbuild *NoBuildProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// UP|CREATE : Property for a docker.libCompose project to indicate that a process should force rebuilding images
type ForceRebuildProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (forcerebuild *ForceRebuildProperty) Id() string {
	return PROPERTY_ID_FORCEREBUILD
}

// Label for the Property
func (forcerebuild *ForceRebuildProperty) Label() string {
	return "Force rebuild"
}

// Description for the Property
func (forcerebuild *ForceRebuildProperty) Description() string {
	return "Force rebuilding any images, even if they exist already?"
}

// Is the Property internal only
func (forcerebuild *ForceRebuildProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// DOWN|DELETE : Property for a docker.libCompose project to indicate that a process should remove any volumes
type RemoveVolumesProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (removevolumes *RemoveVolumesProperty) Id() string {
	return PROPERTY_ID_REMOVEVOLUMES
}

// Label for the Property
func (removevolumes *RemoveVolumesProperty) Label() string {
	return "Remove volumes"
}

// Description for the Property
func (removevolumes *RemoveVolumesProperty) Description() string {
	return "When removing containers, remove any volumes?"
}

// Is the Property internal only
func (removevolumes *RemoveVolumesProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// DOWN : Property for a docker.libCompose project to indicate that a process should remove any orphan containers
type RemoveOrphansProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (removeorphans *RemoveOrphansProperty) Id() string {
	return PROPERTY_ID_REMOVEORPHANS
}

// Label for the Property
func (removeorphans *RemoveOrphansProperty) Label() string {
	return "Remove orphans"
}

// Description for the Property
func (removeorphans *RemoveOrphansProperty) Description() string {
	return "When removing containers, remove any orphans?"
}

// Is the Property internal only
func (removeorphans *RemoveOrphansProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// DOWN : Property for a docker.libCompose project to indicate that a process should remove images of a certain type (local, all)
type RemoveImageTypeProperty struct {
	base_property.StringProperty
}

// Id for the Property
func (removeimagetypes *RemoveImageTypeProperty) Id() string {
	return PROPERTY_ID_REMOVEIMAGETYPES
}

// Label for the Property
func (removeimagetypes *RemoveImageTypeProperty) Label() string {
	return "Remove image types"
}

// Description for the Property
func (removeimagetypes *RemoveImageTypeProperty) Description() string {
	return "When removing containers, remove either 'none' local' or 'all' images?"
}

// Is the Property internal only
func (removeimagetypes *RemoveImageTypeProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// DELETE : Property for a docker.libCompose project to indicate that a process should delete running containers
type RemoveRunningProperty struct {
	base_property.BoolProperty
}

// Id for the Property
func (removerunning *RemoveRunningProperty) Id() string {
	return PROPERTY_ID_REMOVERUNNING
}

// Label for the Property
func (removerunning *RemoveRunningProperty) Label() string {
	return "Remove running"
}

// Description for the Property
func (removerunning *RemoveRunningProperty) Description() string {
	return "When removing containers, remove running containers?"
}

// Is the Property internal only
func (removerunning *RemoveRunningProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// STOP : Property for a docker.libCompose project to indicate that how many seconds a process should run for before timing out
type TimeoutProperty struct {
	base_property.IntProperty
}

// Id for the Property
func (timeout *TimeoutProperty) Id() string {
	return PROPERTY_ID_TIMEOUT
}

// Label for the Property
func (timeout *TimeoutProperty) Label() string {
	return "Timeout"
}

// Description for the Property
func (timeout *TimeoutProperty) Description() string {
	return "Timeout in seconds before an operation should force"
}

// Is the Property internal only
func (timeout *TimeoutProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}