package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/projects/types"
)

// add a default project to a subscription, add the subscription key as
func (k Keeper) CreateDefaultProject(ctx sdk.Context, subscriptionAddress string) error {
	return k.CreateProject(ctx, subscriptionAddress, types.DEFAULT_PROJECT_NAME, subscriptionAddress, true)
}

// add a new project to the subscription
func (k Keeper) CreateProject(ctx sdk.Context, subscriptionAddress string, projectName string, adminAddress string, enable bool) error {
	project := types.CreateProject(subscriptionAddress, projectName)
	var emptyProject types.Project

	blockHeight := uint64(ctx.BlockHeight())
	found := k.projectsFS.FindEntry(ctx, project.Index, blockHeight, &emptyProject)
	// the project with the same name already exists if no error has returned
	if found {
		return utils.LavaError(ctx, ctx.Logger(), "CreateEmptyProject_already_exist", map[string]string{"subscription": subscriptionAddress}, "project already exist for the current subscription with the same name")
	}

	if subscriptionAddress != adminAddress {
		project.AppendKey(types.ProjectKey{Key: adminAddress, Types: []types.ProjectKey_KEY_TYPE{types.ProjectKey_ADMIN}})
	}

	err := k.RegisterDeveloperKey(ctx, adminAddress, project.Index, blockHeight)
	if err != nil {
		return err
	}

	project.Enabled = enable
	return k.projectsFS.AppendEntry(ctx, project.Index, blockHeight, &project)
}

func (k Keeper) RegisterDeveloperKey(ctx sdk.Context, developerKey string, projectIndex string, blockHeight uint64) error {
	var projectID types.ProtoString
	found := k.developerKeysFS.FindEntry(ctx, developerKey, blockHeight, &projectID)
	// a developer key with this address is not registered, add it to the developer keys list
	if !found {
		projectID.String_ = projectIndex
		err := k.developerKeysFS.AppendEntry(ctx, developerKey, blockHeight, &projectID)
		if err != nil {
			return err
		}
	}

	return nil
}

// snapshot project, create a snapshot of a project and reset the cu
func (k Keeper) SnapshotProject(ctx sdk.Context, projectID string) error {
	var project types.Project
	found := k.projectsFS.FindEntry(ctx, projectID, uint64(ctx.BlockHeight()), &project)
	if !found {
		return utils.LavaError(ctx, ctx.Logger(), "SnapshotProject_project_not_found", map[string]string{"projectID": projectID}, "snapshot of project failed, project does not exist")
	}

	project.UsedCu = 0

	return k.projectsFS.AppendEntry(ctx, project.Index, uint64(ctx.BlockHeight()), &project)
}

func (k Keeper) DeleteProject(ctx sdk.Context, projectID string) error {
	var project types.Project
	found := k.projectsFS.FindEntry(ctx, projectID, uint64(ctx.BlockHeight()), &project)
	if !found {
		return utils.LavaError(ctx, ctx.Logger(), "DeleteProject_project_not_found", map[string]string{"projectID": projectID}, "project to delete was not found")
	}

	project.Enabled = false
	// TODO: delete all developer keys from the fixation

	return k.projectsFS.AppendEntry(ctx, project.Index, uint64(ctx.BlockHeight()), &project)
}
