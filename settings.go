package libcompose

import (
	"context"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/config"
)


// Settings for defining a local project
type Settings struct {
	// ProjectName string label for the project, which is used for machine names for things like containers, volumes etc
	ProjectName string
	// The path to the docker compose file
	ComposeFilePath string
	// Context for operation
	Context context.Context
	// ParseOptions for libcompose yml parsing (can be left as nil)
	ParseOptions config.ParseOptions
}

// Settings constructor from parametrized values
func NewSettings(name, compose string, ctx context.Context, parseOptions config.ParseOptions) Settings {
	return Settings{
		ProjectName: name,
		ComposeFilePath: compose,
		Context: ctx,
		ParseOptions: parseOptions,
	}
}

// libCompose APIProject from a settings object
func (s *Settings) LibComposeProject() (project.APIProject, error) {
	proj, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{s.ComposeFilePath},
			ProjectName:  s.ProjectName,
		},
	}, &s.ParseOptions)

	return proj, err
}
