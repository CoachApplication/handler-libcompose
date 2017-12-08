package libcompose

import (
	"testing"
	"context"
	"os"
	"path"
	"time"

	libcompose_config "github.com/docker/libcompose/config"
	libcompose_project "github.com/docker/libcompose/project"
	libcompose_project_options "github.com/docker/libcompose/project/options"
)

func testOrchestrate_Project(t *testing.T, ctx context.Context) (libcompose_project.APIProject, error) {
	wd, _ := os.Getwd()
	testPath := path.Join(wd, "test")

	name := "coach_test"
	composePath := path.Join(testPath, "docker-compose.yml")
	parseOptions := libcompose_config.ParseOptions{}

	settings := NewSettings(name, composePath, ctx, parseOptions)
	project, err := settings.LibComposeProject()

	return project, err
}

/**
 * Test the Up Operation
 */
func TestOrchestrate_Up(t *testing.T) {
	dur, _ := time.ParseDuration("60s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	upOptions := libcompose_project_options.Up{}
	upOptions.ForceRecreate = true
	upOptions.ForceBuild = true

	if project, err := testOrchestrate_Project(t, ctx); err != nil {
		t.Error("Failed to create a libcompose project", err.Error())
	} else if err := project.Up(ctx, upOptions); err != nil {
		t.Error("Failed to UP the libcompose project", err.Error())
	}
}

/**
 * Test the Down Operation
 */
func TestOrchestrate_Down(t *testing.T) {
	dur, _ := time.ParseDuration("60s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	downOptions := libcompose_project_options.Down{}
	downOptions.RemoveVolume = true
	downOptions.RemoveOrphans = true

	if project, err := testOrchestrate_Project(t, ctx); err != nil {
		t.Error("Failed to create a libcompose project", err.Error())
	} else if err := project.Down(ctx, downOptions); err != nil {
		t.Error("Failed to DOWN the libcompose project", err.Error())
	}
}
