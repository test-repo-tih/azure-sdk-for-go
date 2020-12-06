## Breaking Changes

### Removed Funcs

1. *DedicatedCloudNodeCreateOrUpdateFuture.Result(DedicatedCloudNodeClient) (DedicatedCloudNode,error)
1. *DedicatedCloudServiceDeleteFuture.Result(DedicatedCloudServiceClient) (autorest.Response,error)
1. *VirtualMachineCreateOrUpdateFuture.Result(VirtualMachineClient) (VirtualMachine,error)
1. *VirtualMachineDeleteFuture.Result(VirtualMachineClient) (autorest.Response,error)
1. *VirtualMachineStartFuture.Result(VirtualMachineClient) (autorest.Response,error)
1. *VirtualMachineStopFuture.Result(VirtualMachineClient) (autorest.Response,error)
1. *VirtualMachineUpdateFuture.Result(VirtualMachineClient) (VirtualMachine,error)
1. AvailableOperationsClient.List(context.Context) (AvailableOperationsListResponsePage,error)
1. AvailableOperationsClient.ListComplete(context.Context) (AvailableOperationsListResponseIterator,error)
1. AvailableOperationsClient.ListPreparer(context.Context) (*http.Request,error)
1. AvailableOperationsClient.ListResponder(*http.Response) (AvailableOperationsListResponse,error)
1. AvailableOperationsClient.ListSender(*http.Request) (*http.Response,error)
1. BaseClient.GetOperationResultByRegion(context.Context,string) (OperationResource,error)
1. BaseClient.GetOperationResultByRegionPreparer(context.Context,string) (*http.Request,error)
1. BaseClient.GetOperationResultByRegionResponder(*http.Response) (OperationResource,error)
1. BaseClient.GetOperationResultByRegionSender(*http.Request) (*http.Response,error)
1. BaseClient.GetPrivateCloud(context.Context,string) (PrivateCloud,error)
1. BaseClient.GetPrivateCloudPreparer(context.Context,string) (*http.Request,error)
1. BaseClient.GetPrivateCloudResponder(*http.Response) (PrivateCloud,error)
1. BaseClient.GetPrivateCloudSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodeClient.CreateOrUpdate(context.Context,string,string,DedicatedCloudNode) (DedicatedCloudNodeCreateOrUpdateFuture,error)
1. DedicatedCloudNodeClient.CreateOrUpdatePreparer(context.Context,string,string,DedicatedCloudNode) (*http.Request,error)
1. DedicatedCloudNodeClient.CreateOrUpdateResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodeClient.CreateOrUpdateSender(*http.Request) (DedicatedCloudNodeCreateOrUpdateFuture,error)
1. DedicatedCloudNodeClient.Delete(context.Context,string,string) (autorest.Response,error)
1. DedicatedCloudNodeClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudNodeClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DedicatedCloudNodeClient.DeleteSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodeClient.Get(context.Context,string,string) (DedicatedCloudNode,error)
1. DedicatedCloudNodeClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudNodeClient.GetResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodeClient.GetSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodeClient.ListByResourceGroup(context.Context,string,string,*int32,string) (DedicatedCloudNodeListResponsePage,error)
1. DedicatedCloudNodeClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (DedicatedCloudNodeListResponseIterator,error)
1. DedicatedCloudNodeClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. DedicatedCloudNodeClient.ListByResourceGroupResponder(*http.Response) (DedicatedCloudNodeListResponse,error)
1. DedicatedCloudNodeClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodeClient.ListBySubscription(context.Context,string,*int32,string) (DedicatedCloudNodeListResponsePage,error)
1. DedicatedCloudNodeClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (DedicatedCloudNodeListResponseIterator,error)
1. DedicatedCloudNodeClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. DedicatedCloudNodeClient.ListBySubscriptionResponder(*http.Response) (DedicatedCloudNodeListResponse,error)
1. DedicatedCloudNodeClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodeClient.Update(context.Context,string,string,PatchPayload) (DedicatedCloudNode,error)
1. DedicatedCloudNodeClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. DedicatedCloudNodeClient.UpdateResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodeClient.UpdateSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServiceClient.CreateOrUpdate(context.Context,string,string,DedicatedCloudService) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.CreateOrUpdatePreparer(context.Context,string,string,DedicatedCloudService) (*http.Request,error)
1. DedicatedCloudServiceClient.CreateOrUpdateResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServiceClient.Delete(context.Context,string,string) (DedicatedCloudServiceDeleteFuture,error)
1. DedicatedCloudServiceClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudServiceClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DedicatedCloudServiceClient.DeleteSender(*http.Request) (DedicatedCloudServiceDeleteFuture,error)
1. DedicatedCloudServiceClient.Get(context.Context,string,string) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudServiceClient.GetResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.GetSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServiceClient.ListByResourceGroup(context.Context,string,string,*int32,string) (DedicatedCloudServiceListResponsePage,error)
1. DedicatedCloudServiceClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (DedicatedCloudServiceListResponseIterator,error)
1. DedicatedCloudServiceClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. DedicatedCloudServiceClient.ListByResourceGroupResponder(*http.Response) (DedicatedCloudServiceListResponse,error)
1. DedicatedCloudServiceClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServiceClient.ListBySubscription(context.Context,string,*int32,string) (DedicatedCloudServiceListResponsePage,error)
1. DedicatedCloudServiceClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (DedicatedCloudServiceListResponseIterator,error)
1. DedicatedCloudServiceClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. DedicatedCloudServiceClient.ListBySubscriptionResponder(*http.Response) (DedicatedCloudServiceListResponse,error)
1. DedicatedCloudServiceClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServiceClient.Update(context.Context,string,string,PatchPayload) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. DedicatedCloudServiceClient.UpdateResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServiceClient.UpdateSender(*http.Request) (*http.Response,error)
1. NewAvailableOperationsClient(string,string,string) AvailableOperationsClient
1. NewAvailableOperationsClientWithBaseURI(string,string,string,string) AvailableOperationsClient
1. NewDedicatedCloudNodeClient(string,string,string) DedicatedCloudNodeClient
1. NewDedicatedCloudNodeClientWithBaseURI(string,string,string,string) DedicatedCloudNodeClient
1. NewDedicatedCloudServiceClient(string,string,string) DedicatedCloudServiceClient
1. NewDedicatedCloudServiceClientWithBaseURI(string,string,string,string) DedicatedCloudServiceClient
1. NewPrivateCloudByRegionClient(string,string,string) PrivateCloudByRegionClient
1. NewPrivateCloudByRegionClientWithBaseURI(string,string,string,string) PrivateCloudByRegionClient
1. NewResourcePoolByPCClient(string,string,string) ResourcePoolByPCClient
1. NewResourcePoolByPCClientWithBaseURI(string,string,string,string) ResourcePoolByPCClient
1. NewResourcePoolsByPCClient(string,string,string) ResourcePoolsByPCClient
1. NewResourcePoolsByPCClientWithBaseURI(string,string,string,string) ResourcePoolsByPCClient
1. NewSkusAvailabilityWithinRegionClient(string,string,string) SkusAvailabilityWithinRegionClient
1. NewSkusAvailabilityWithinRegionClientWithBaseURI(string,string,string,string) SkusAvailabilityWithinRegionClient
1. NewUsagesWithinRegionClient(string,string,string) UsagesWithinRegionClient
1. NewUsagesWithinRegionClientWithBaseURI(string,string,string,string) UsagesWithinRegionClient
1. NewVirtualMachineClient(string,string,string) VirtualMachineClient
1. NewVirtualMachineClientWithBaseURI(string,string,string,string) VirtualMachineClient
1. NewVirtualMachineTemplateByPCClient(string,string,string) VirtualMachineTemplateByPCClient
1. NewVirtualMachineTemplateByPCClientWithBaseURI(string,string,string,string) VirtualMachineTemplateByPCClient
1. NewVirtualMachineTemplatesByPCClient(string,string,string) VirtualMachineTemplatesByPCClient
1. NewVirtualMachineTemplatesByPCClientWithBaseURI(string,string,string,string) VirtualMachineTemplatesByPCClient
1. NewVirtualNetworkByPCClient(string,string,string) VirtualNetworkByPCClient
1. NewVirtualNetworkByPCClientWithBaseURI(string,string,string,string) VirtualNetworkByPCClient
1. NewVirtualNetworksByPCClient(string,string,string) VirtualNetworksByPCClient
1. NewVirtualNetworksByPCClientWithBaseURI(string,string,string,string) VirtualNetworksByPCClient
1. PrivateCloudByRegionClient.List(context.Context) (PrivateCloudListPage,error)
1. PrivateCloudByRegionClient.ListComplete(context.Context) (PrivateCloudListIterator,error)
1. PrivateCloudByRegionClient.ListPreparer(context.Context) (*http.Request,error)
1. PrivateCloudByRegionClient.ListResponder(*http.Response) (PrivateCloudList,error)
1. PrivateCloudByRegionClient.ListSender(*http.Request) (*http.Response,error)
1. ResourcePoolByPCClient.Get(context.Context,string,string) (ResourcePool,error)
1. ResourcePoolByPCClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ResourcePoolByPCClient.GetResponder(*http.Response) (ResourcePool,error)
1. ResourcePoolByPCClient.GetSender(*http.Request) (*http.Response,error)
1. ResourcePoolsByPCClient.List(context.Context,string) (ResourcePoolsListResponsePage,error)
1. ResourcePoolsByPCClient.ListComplete(context.Context,string) (ResourcePoolsListResponseIterator,error)
1. ResourcePoolsByPCClient.ListPreparer(context.Context,string) (*http.Request,error)
1. ResourcePoolsByPCClient.ListResponder(*http.Response) (ResourcePoolsListResponse,error)
1. ResourcePoolsByPCClient.ListSender(*http.Request) (*http.Response,error)
1. SkusAvailabilityWithinRegionClient.List(context.Context,string) (SkuAvailabilityListResponsePage,error)
1. SkusAvailabilityWithinRegionClient.ListComplete(context.Context,string) (SkuAvailabilityListResponseIterator,error)
1. SkusAvailabilityWithinRegionClient.ListPreparer(context.Context,string) (*http.Request,error)
1. SkusAvailabilityWithinRegionClient.ListResponder(*http.Response) (SkuAvailabilityListResponse,error)
1. SkusAvailabilityWithinRegionClient.ListSender(*http.Request) (*http.Response,error)
1. UsagesWithinRegionClient.List(context.Context,string) (UsageListResponsePage,error)
1. UsagesWithinRegionClient.ListComplete(context.Context,string) (UsageListResponseIterator,error)
1. UsagesWithinRegionClient.ListPreparer(context.Context,string) (*http.Request,error)
1. UsagesWithinRegionClient.ListResponder(*http.Response) (UsageListResponse,error)
1. UsagesWithinRegionClient.ListSender(*http.Request) (*http.Response,error)
1. VirtualMachineClient.CreateOrUpdate(context.Context,string,string,VirtualMachine) (VirtualMachineCreateOrUpdateFuture,error)
1. VirtualMachineClient.CreateOrUpdatePreparer(context.Context,string,string,VirtualMachine) (*http.Request,error)
1. VirtualMachineClient.CreateOrUpdateResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachineClient.CreateOrUpdateSender(*http.Request) (VirtualMachineCreateOrUpdateFuture,error)
1. VirtualMachineClient.Delete(context.Context,string,string) (VirtualMachineDeleteFuture,error)
1. VirtualMachineClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachineClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. VirtualMachineClient.DeleteSender(*http.Request) (VirtualMachineDeleteFuture,error)
1. VirtualMachineClient.Get(context.Context,string,string) (VirtualMachine,error)
1. VirtualMachineClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachineClient.GetResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachineClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualMachineClient.ListByResourceGroup(context.Context,string,string,*int32,string) (VirtualMachineListResponsePage,error)
1. VirtualMachineClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (VirtualMachineListResponseIterator,error)
1. VirtualMachineClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. VirtualMachineClient.ListByResourceGroupResponder(*http.Response) (VirtualMachineListResponse,error)
1. VirtualMachineClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. VirtualMachineClient.ListBySubscription(context.Context,string,*int32,string) (VirtualMachineListResponsePage,error)
1. VirtualMachineClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (VirtualMachineListResponseIterator,error)
1. VirtualMachineClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. VirtualMachineClient.ListBySubscriptionResponder(*http.Response) (VirtualMachineListResponse,error)
1. VirtualMachineClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. VirtualMachineClient.Start(context.Context,string,string) (VirtualMachineStartFuture,error)
1. VirtualMachineClient.StartPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachineClient.StartResponder(*http.Response) (autorest.Response,error)
1. VirtualMachineClient.StartSender(*http.Request) (VirtualMachineStartFuture,error)
1. VirtualMachineClient.Stop(context.Context,string,string,*VirtualMachineStopMode,StopMode) (VirtualMachineStopFuture,error)
1. VirtualMachineClient.StopPreparer(context.Context,string,string,*VirtualMachineStopMode,StopMode) (*http.Request,error)
1. VirtualMachineClient.StopResponder(*http.Response) (autorest.Response,error)
1. VirtualMachineClient.StopSender(*http.Request) (VirtualMachineStopFuture,error)
1. VirtualMachineClient.Update(context.Context,string,string,PatchPayload) (VirtualMachineUpdateFuture,error)
1. VirtualMachineClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. VirtualMachineClient.UpdateResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachineClient.UpdateSender(*http.Request) (VirtualMachineUpdateFuture,error)
1. VirtualMachineTemplateByPCClient.Get(context.Context,string,string) (VirtualMachineTemplate,error)
1. VirtualMachineTemplateByPCClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachineTemplateByPCClient.GetResponder(*http.Response) (VirtualMachineTemplate,error)
1. VirtualMachineTemplateByPCClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualMachineTemplatesByPCClient.List(context.Context,string,string) (VirtualMachineTemplateListResponsePage,error)
1. VirtualMachineTemplatesByPCClient.ListComplete(context.Context,string,string) (VirtualMachineTemplateListResponseIterator,error)
1. VirtualMachineTemplatesByPCClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachineTemplatesByPCClient.ListResponder(*http.Response) (VirtualMachineTemplateListResponse,error)
1. VirtualMachineTemplatesByPCClient.ListSender(*http.Request) (*http.Response,error)
1. VirtualNetworkByPCClient.Get(context.Context,string,string) (VirtualNetwork,error)
1. VirtualNetworkByPCClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualNetworkByPCClient.GetResponder(*http.Response) (VirtualNetwork,error)
1. VirtualNetworkByPCClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualNetworksByPCClient.List(context.Context,string,string) (VirtualNetworkListResponsePage,error)
1. VirtualNetworksByPCClient.ListComplete(context.Context,string,string) (VirtualNetworkListResponseIterator,error)
1. VirtualNetworksByPCClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualNetworksByPCClient.ListResponder(*http.Response) (VirtualNetworkListResponse,error)
1. VirtualNetworksByPCClient.ListSender(*http.Request) (*http.Response,error)

## Struct Changes

### Removed Structs

1. AvailableOperationsClient
1. DedicatedCloudNodeClient
1. DedicatedCloudNodeCreateOrUpdateFuture
1. DedicatedCloudServiceClient
1. DedicatedCloudServiceDeleteFuture
1. PrivateCloudByRegionClient
1. ResourcePoolByPCClient
1. ResourcePoolsByPCClient
1. SkusAvailabilityWithinRegionClient
1. UsagesWithinRegionClient
1. VirtualMachineClient
1. VirtualMachineCreateOrUpdateFuture
1. VirtualMachineDeleteFuture
1. VirtualMachineStartFuture
1. VirtualMachineStopFuture
1. VirtualMachineTemplateByPCClient
1. VirtualMachineTemplatesByPCClient
1. VirtualMachineUpdateFuture
1. VirtualNetworkByPCClient
1. VirtualNetworksByPCClient

### Removed Struct Fields

1. BaseClient.RegionID

## Signature Changes

### Funcs

1. New
	- Params
		- From: string,string,string
		- To: string,string
1. NewWithBaseURI
	- Params
		- From: string,string,string,string
		- To: string,string,string

## New Content

### New Funcs

1. *DedicatedCloudNodesCreateOrUpdateFuture.Result(DedicatedCloudNodesClient) (DedicatedCloudNode,error)
1. *DedicatedCloudServicesDeleteFuture.Result(DedicatedCloudServicesClient) (autorest.Response,error)
1. *VirtualMachinesCreateOrUpdateFuture.Result(VirtualMachinesClient) (VirtualMachine,error)
1. *VirtualMachinesDeleteFuture.Result(VirtualMachinesClient) (autorest.Response,error)
1. *VirtualMachinesStartFuture.Result(VirtualMachinesClient) (autorest.Response,error)
1. *VirtualMachinesStopFuture.Result(VirtualMachinesClient) (autorest.Response,error)
1. *VirtualMachinesUpdateFuture.Result(VirtualMachinesClient) (VirtualMachine,error)
1. DedicatedCloudNodesClient.CreateOrUpdate(context.Context,string,string,DedicatedCloudNode) (DedicatedCloudNodesCreateOrUpdateFuture,error)
1. DedicatedCloudNodesClient.CreateOrUpdatePreparer(context.Context,string,string,DedicatedCloudNode) (*http.Request,error)
1. DedicatedCloudNodesClient.CreateOrUpdateResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodesClient.CreateOrUpdateSender(*http.Request) (DedicatedCloudNodesCreateOrUpdateFuture,error)
1. DedicatedCloudNodesClient.Delete(context.Context,string,string) (autorest.Response,error)
1. DedicatedCloudNodesClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudNodesClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DedicatedCloudNodesClient.DeleteSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodesClient.Get(context.Context,string,string) (DedicatedCloudNode,error)
1. DedicatedCloudNodesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudNodesClient.GetResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodesClient.GetSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodesClient.ListByResourceGroup(context.Context,string,string,*int32,string) (DedicatedCloudNodeListResponsePage,error)
1. DedicatedCloudNodesClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (DedicatedCloudNodeListResponseIterator,error)
1. DedicatedCloudNodesClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. DedicatedCloudNodesClient.ListByResourceGroupResponder(*http.Response) (DedicatedCloudNodeListResponse,error)
1. DedicatedCloudNodesClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodesClient.ListBySubscription(context.Context,string,*int32,string) (DedicatedCloudNodeListResponsePage,error)
1. DedicatedCloudNodesClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (DedicatedCloudNodeListResponseIterator,error)
1. DedicatedCloudNodesClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. DedicatedCloudNodesClient.ListBySubscriptionResponder(*http.Response) (DedicatedCloudNodeListResponse,error)
1. DedicatedCloudNodesClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. DedicatedCloudNodesClient.Update(context.Context,string,string,PatchPayload) (DedicatedCloudNode,error)
1. DedicatedCloudNodesClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. DedicatedCloudNodesClient.UpdateResponder(*http.Response) (DedicatedCloudNode,error)
1. DedicatedCloudNodesClient.UpdateSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServicesClient.CreateOrUpdate(context.Context,string,string,DedicatedCloudService) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.CreateOrUpdatePreparer(context.Context,string,string,DedicatedCloudService) (*http.Request,error)
1. DedicatedCloudServicesClient.CreateOrUpdateResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServicesClient.Delete(context.Context,string,string) (DedicatedCloudServicesDeleteFuture,error)
1. DedicatedCloudServicesClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudServicesClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DedicatedCloudServicesClient.DeleteSender(*http.Request) (DedicatedCloudServicesDeleteFuture,error)
1. DedicatedCloudServicesClient.Get(context.Context,string,string) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. DedicatedCloudServicesClient.GetResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.GetSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServicesClient.ListByResourceGroup(context.Context,string,string,*int32,string) (DedicatedCloudServiceListResponsePage,error)
1. DedicatedCloudServicesClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (DedicatedCloudServiceListResponseIterator,error)
1. DedicatedCloudServicesClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. DedicatedCloudServicesClient.ListByResourceGroupResponder(*http.Response) (DedicatedCloudServiceListResponse,error)
1. DedicatedCloudServicesClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServicesClient.ListBySubscription(context.Context,string,*int32,string) (DedicatedCloudServiceListResponsePage,error)
1. DedicatedCloudServicesClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (DedicatedCloudServiceListResponseIterator,error)
1. DedicatedCloudServicesClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. DedicatedCloudServicesClient.ListBySubscriptionResponder(*http.Response) (DedicatedCloudServiceListResponse,error)
1. DedicatedCloudServicesClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. DedicatedCloudServicesClient.Update(context.Context,string,string,PatchPayload) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. DedicatedCloudServicesClient.UpdateResponder(*http.Response) (DedicatedCloudService,error)
1. DedicatedCloudServicesClient.UpdateSender(*http.Request) (*http.Response,error)
1. NewDedicatedCloudNodesClient(string,string) DedicatedCloudNodesClient
1. NewDedicatedCloudNodesClientWithBaseURI(string,string,string) DedicatedCloudNodesClient
1. NewDedicatedCloudServicesClient(string,string) DedicatedCloudServicesClient
1. NewDedicatedCloudServicesClientWithBaseURI(string,string,string) DedicatedCloudServicesClient
1. NewOperationsClient(string,string) OperationsClient
1. NewOperationsClientWithBaseURI(string,string,string) OperationsClient
1. NewPrivateCloudsClient(string,string) PrivateCloudsClient
1. NewPrivateCloudsClientWithBaseURI(string,string,string) PrivateCloudsClient
1. NewResourcePoolsClient(string,string) ResourcePoolsClient
1. NewResourcePoolsClientWithBaseURI(string,string,string) ResourcePoolsClient
1. NewSkusAvailabilityClient(string,string) SkusAvailabilityClient
1. NewSkusAvailabilityClientWithBaseURI(string,string,string) SkusAvailabilityClient
1. NewUsagesClient(string,string) UsagesClient
1. NewUsagesClientWithBaseURI(string,string,string) UsagesClient
1. NewVirtualMachineTemplatesClient(string,string) VirtualMachineTemplatesClient
1. NewVirtualMachineTemplatesClientWithBaseURI(string,string,string) VirtualMachineTemplatesClient
1. NewVirtualMachinesClient(string,string) VirtualMachinesClient
1. NewVirtualMachinesClientWithBaseURI(string,string,string) VirtualMachinesClient
1. NewVirtualNetworksClient(string,string) VirtualNetworksClient
1. NewVirtualNetworksClientWithBaseURI(string,string,string) VirtualNetworksClient
1. OperationsClient.Get(context.Context,string,string) (OperationResource,error)
1. OperationsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. OperationsClient.GetResponder(*http.Response) (OperationResource,error)
1. OperationsClient.GetSender(*http.Request) (*http.Response,error)
1. OperationsClient.List(context.Context) (AvailableOperationsListResponsePage,error)
1. OperationsClient.ListComplete(context.Context) (AvailableOperationsListResponseIterator,error)
1. OperationsClient.ListPreparer(context.Context) (*http.Request,error)
1. OperationsClient.ListResponder(*http.Response) (AvailableOperationsListResponse,error)
1. OperationsClient.ListSender(*http.Request) (*http.Response,error)
1. PrivateCloudsClient.Get(context.Context,string,string) (PrivateCloud,error)
1. PrivateCloudsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. PrivateCloudsClient.GetResponder(*http.Response) (PrivateCloud,error)
1. PrivateCloudsClient.GetSender(*http.Request) (*http.Response,error)
1. PrivateCloudsClient.List(context.Context,string) (PrivateCloudListPage,error)
1. PrivateCloudsClient.ListComplete(context.Context,string) (PrivateCloudListIterator,error)
1. PrivateCloudsClient.ListPreparer(context.Context,string) (*http.Request,error)
1. PrivateCloudsClient.ListResponder(*http.Response) (PrivateCloudList,error)
1. PrivateCloudsClient.ListSender(*http.Request) (*http.Response,error)
1. ResourcePoolsClient.Get(context.Context,string,string,string) (ResourcePool,error)
1. ResourcePoolsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ResourcePoolsClient.GetResponder(*http.Response) (ResourcePool,error)
1. ResourcePoolsClient.GetSender(*http.Request) (*http.Response,error)
1. ResourcePoolsClient.List(context.Context,string,string) (ResourcePoolsListResponsePage,error)
1. ResourcePoolsClient.ListComplete(context.Context,string,string) (ResourcePoolsListResponseIterator,error)
1. ResourcePoolsClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. ResourcePoolsClient.ListResponder(*http.Response) (ResourcePoolsListResponse,error)
1. ResourcePoolsClient.ListSender(*http.Request) (*http.Response,error)
1. SkusAvailabilityClient.List(context.Context,string,string) (SkuAvailabilityListResponsePage,error)
1. SkusAvailabilityClient.ListComplete(context.Context,string,string) (SkuAvailabilityListResponseIterator,error)
1. SkusAvailabilityClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. SkusAvailabilityClient.ListResponder(*http.Response) (SkuAvailabilityListResponse,error)
1. SkusAvailabilityClient.ListSender(*http.Request) (*http.Response,error)
1. UsagesClient.List(context.Context,string,string) (UsageListResponsePage,error)
1. UsagesClient.ListComplete(context.Context,string,string) (UsageListResponseIterator,error)
1. UsagesClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. UsagesClient.ListResponder(*http.Response) (UsageListResponse,error)
1. UsagesClient.ListSender(*http.Request) (*http.Response,error)
1. VirtualMachineTemplatesClient.Get(context.Context,string,string,string) (VirtualMachineTemplate,error)
1. VirtualMachineTemplatesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. VirtualMachineTemplatesClient.GetResponder(*http.Response) (VirtualMachineTemplate,error)
1. VirtualMachineTemplatesClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualMachineTemplatesClient.List(context.Context,string,string,string) (VirtualMachineTemplateListResponsePage,error)
1. VirtualMachineTemplatesClient.ListComplete(context.Context,string,string,string) (VirtualMachineTemplateListResponseIterator,error)
1. VirtualMachineTemplatesClient.ListPreparer(context.Context,string,string,string) (*http.Request,error)
1. VirtualMachineTemplatesClient.ListResponder(*http.Response) (VirtualMachineTemplateListResponse,error)
1. VirtualMachineTemplatesClient.ListSender(*http.Request) (*http.Response,error)
1. VirtualMachinesClient.CreateOrUpdate(context.Context,string,string,VirtualMachine) (VirtualMachinesCreateOrUpdateFuture,error)
1. VirtualMachinesClient.CreateOrUpdatePreparer(context.Context,string,string,VirtualMachine) (*http.Request,error)
1. VirtualMachinesClient.CreateOrUpdateResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachinesClient.CreateOrUpdateSender(*http.Request) (VirtualMachinesCreateOrUpdateFuture,error)
1. VirtualMachinesClient.Delete(context.Context,string,string) (VirtualMachinesDeleteFuture,error)
1. VirtualMachinesClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachinesClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. VirtualMachinesClient.DeleteSender(*http.Request) (VirtualMachinesDeleteFuture,error)
1. VirtualMachinesClient.Get(context.Context,string,string) (VirtualMachine,error)
1. VirtualMachinesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachinesClient.GetResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachinesClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualMachinesClient.ListByResourceGroup(context.Context,string,string,*int32,string) (VirtualMachineListResponsePage,error)
1. VirtualMachinesClient.ListByResourceGroupComplete(context.Context,string,string,*int32,string) (VirtualMachineListResponseIterator,error)
1. VirtualMachinesClient.ListByResourceGroupPreparer(context.Context,string,string,*int32,string) (*http.Request,error)
1. VirtualMachinesClient.ListByResourceGroupResponder(*http.Response) (VirtualMachineListResponse,error)
1. VirtualMachinesClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. VirtualMachinesClient.ListBySubscription(context.Context,string,*int32,string) (VirtualMachineListResponsePage,error)
1. VirtualMachinesClient.ListBySubscriptionComplete(context.Context,string,*int32,string) (VirtualMachineListResponseIterator,error)
1. VirtualMachinesClient.ListBySubscriptionPreparer(context.Context,string,*int32,string) (*http.Request,error)
1. VirtualMachinesClient.ListBySubscriptionResponder(*http.Response) (VirtualMachineListResponse,error)
1. VirtualMachinesClient.ListBySubscriptionSender(*http.Request) (*http.Response,error)
1. VirtualMachinesClient.Start(context.Context,string,string) (VirtualMachinesStartFuture,error)
1. VirtualMachinesClient.StartPreparer(context.Context,string,string) (*http.Request,error)
1. VirtualMachinesClient.StartResponder(*http.Response) (autorest.Response,error)
1. VirtualMachinesClient.StartSender(*http.Request) (VirtualMachinesStartFuture,error)
1. VirtualMachinesClient.Stop(context.Context,string,string,*VirtualMachineStopMode,StopMode) (VirtualMachinesStopFuture,error)
1. VirtualMachinesClient.StopPreparer(context.Context,string,string,*VirtualMachineStopMode,StopMode) (*http.Request,error)
1. VirtualMachinesClient.StopResponder(*http.Response) (autorest.Response,error)
1. VirtualMachinesClient.StopSender(*http.Request) (VirtualMachinesStopFuture,error)
1. VirtualMachinesClient.Update(context.Context,string,string,PatchPayload) (VirtualMachinesUpdateFuture,error)
1. VirtualMachinesClient.UpdatePreparer(context.Context,string,string,PatchPayload) (*http.Request,error)
1. VirtualMachinesClient.UpdateResponder(*http.Response) (VirtualMachine,error)
1. VirtualMachinesClient.UpdateSender(*http.Request) (VirtualMachinesUpdateFuture,error)
1. VirtualNetworksClient.Get(context.Context,string,string,string) (VirtualNetwork,error)
1. VirtualNetworksClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. VirtualNetworksClient.GetResponder(*http.Response) (VirtualNetwork,error)
1. VirtualNetworksClient.GetSender(*http.Request) (*http.Response,error)
1. VirtualNetworksClient.List(context.Context,string,string,string) (VirtualNetworkListResponsePage,error)
1. VirtualNetworksClient.ListComplete(context.Context,string,string,string) (VirtualNetworkListResponseIterator,error)
1. VirtualNetworksClient.ListPreparer(context.Context,string,string,string) (*http.Request,error)
1. VirtualNetworksClient.ListResponder(*http.Response) (VirtualNetworkListResponse,error)
1. VirtualNetworksClient.ListSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. DedicatedCloudNodesClient
1. DedicatedCloudNodesCreateOrUpdateFuture
1. DedicatedCloudServicesClient
1. DedicatedCloudServicesDeleteFuture
1. OperationsClient
1. PrivateCloudsClient
1. ResourcePoolsClient
1. SkusAvailabilityClient
1. UsagesClient
1. VirtualMachineTemplatesClient
1. VirtualMachinesClient
1. VirtualMachinesCreateOrUpdateFuture
1. VirtualMachinesDeleteFuture
1. VirtualMachinesStartFuture
1. VirtualMachinesStopFuture
1. VirtualMachinesUpdateFuture
1. VirtualNetworksClient
