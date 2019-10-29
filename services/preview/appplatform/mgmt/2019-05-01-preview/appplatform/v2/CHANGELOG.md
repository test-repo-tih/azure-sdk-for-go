## Breaking Changes

### Removed Constants

1. DeploymentResourceProvisioningState.DeploymentResourceProvisioningStateProcessing
1. DeploymentResourceStatus.DeploymentResourceStatusProcessing

## Signature Changes

### Funcs

1. AppsClient.CreateOrUpdate
	- Returns
		- From: AppResource,error
		- To: AppsCreateOrUpdateFuture,error
1. AppsClient.CreateOrUpdateSender
	- Returns
		- From: *http.Response,error
		- To: AppsCreateOrUpdateFuture,error
1. AppsClient.Update
	- Returns
		- From: AppResource,error
		- To: AppsUpdateFuture,error
1. AppsClient.UpdateSender
	- Returns
		- From: *http.Response,error
		- To: AppsUpdateFuture,error

## New Content

### New Constants

1. AppResourceProvisioningState.Creating
1. AppResourceProvisioningState.Updating
1. DeploymentResourceProvisioningState.DeploymentResourceProvisioningStateUpdating

### New Funcs

1. *AppsCreateOrUpdateFuture.Result(AppsClient) (AppResource,error)
1. *AppsUpdateFuture.Result(AppsClient) (AppResource,error)

## Struct Changes

### New Structs

1. AppsCreateOrUpdateFuture
1. AppsUpdateFuture
