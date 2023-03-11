package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/projects/types"
)

func (k Keeper) GetProjectForBlock(ctx sdk.Context, projectID string, blockHeight uint64) (types.Project, error) {
	var project types.Project

	found := k.projectsFS.FindEntry(ctx, projectID, blockHeight, &project)
	if !found {
		return project, utils.LavaError(ctx, ctx.Logger(), "GetProjectForBlock_not_found", map[string]string{"project": projectID, "blockHeight": strconv.FormatUint(blockHeight, 10)}, "project not found")
	}

	return project, nil
}

func (k Keeper) GetProjectIDForDeveloper(ctx sdk.Context, developerKey string, blockHeight uint64) (string, error) {
	var projectIDstring types.ProtoString
	found := k.developerKeysFS.FindEntry(ctx, developerKey, blockHeight, &projectIDstring)
	if !found {
		return "", utils.LavaError(ctx, ctx.Logger(), "GetProjectIDForDeveloper_invalid_key", map[string]string{"developer": developerKey}, "the requesting key is not registered to a project")
	}

	return projectIDstring.String_, nil
}

func (k Keeper) GetProjectForDeveloper(ctx sdk.Context, developerKey string, blockHeight uint64) (types.Project, error) {
	var project types.Project
	projectID, err := k.GetProjectIDForDeveloper(ctx, developerKey, blockHeight)
	if err != nil {
		return project, err
	}

	found := k.projectsFS.FindEntry(ctx, projectID, blockHeight, &project)
	if !found {
		return project, utils.LavaError(ctx, ctx.Logger(), "GetProjectForDeveloper_project_not_found", map[string]string{"developer": developerKey, "project": projectID}, "the developers project was not found")
	}

	return project, nil
}

func (k Keeper) AddKeysToProject(ctx sdk.Context, projectID string, adminKey string, projectKeys []types.ProjectKey) error {
	var project types.Project
	found := k.projectsFS.FindEntry(ctx, projectID, uint64(ctx.BlockHeight()), &project)
	if !found {
		return utils.LavaError(ctx, ctx.Logger(), "AddProjectKeys_project_not_found", map[string]string{"project": projectID}, "project id not found")
	}

	// check if the admin key is valid
	if !project.HasKeyType(adminKey, types.ProjectKey_ADMIN) && project.Subscription != adminKey {
		return utils.LavaError(ctx, ctx.Logger(), "AddProjectKeys_not_admin", map[string]string{"project": projectID}, "the requesting key is not admin key")
	}

	// check that those keys are unique for developers
	for _, projectKey := range projectKeys {
		err := k.RegisterDeveloperKey(ctx, projectKey.Key, project.Index, uint64(ctx.BlockHeight()))
		if err != nil {
			return err
		}

		project.AppendKey(projectKey)
	}

	return k.projectsFS.AppendEntry(ctx, projectID, uint64(ctx.BlockHeight()), &project)
}

func (k Keeper) ValidateDeveloperRequest(ctx sdk.Context, developerKey string, chainID string, apiName string, blockHeight uint64) (valid bool, policy types.Policy, err error) {
	project, err := k.GetProjectForDeveloper(ctx, developerKey, blockHeight)
	if err != nil {
		return false, types.Policy{}, err
	}

	if project.UsedCu >= project.Policy.TotalCuLimit {
		return false, project.Policy, nil
	}

	for _, chain := range project.Policy.ChainPolicies {
		if chain.ChainId == chainID {
			for _, api := range chain.Apis {
				if api == apiName {
					return true, project.Policy, nil
				}
			}
		}
	}

	return false, project.Policy, nil
}
