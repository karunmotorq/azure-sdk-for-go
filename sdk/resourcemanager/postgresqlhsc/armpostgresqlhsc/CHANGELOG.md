# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ServerGroupsClient.BeginStop` parameter(s) have been changed from `(context.Context, string, string, *ServerGroupsBeginStopOptions)` to `(context.Context, string, string, *ServerGroupsClientBeginStopOptions)`
- Function `*ServerGroupsClient.BeginStop` return value(s) have been changed from `(ServerGroupsStopPollerResponse, error)` to `(ServerGroupsClientStopPollerResponse, error)`
- Function `*RolesClient.ListByServerGroup` parameter(s) have been changed from `(context.Context, string, string, *RolesListByServerGroupOptions)` to `(context.Context, string, string, *RolesClientListByServerGroupOptions)`
- Function `*RolesClient.ListByServerGroup` return value(s) have been changed from `(RolesListByServerGroupResponse, error)` to `(RolesClientListByServerGroupResponse, error)`
- Function `*FirewallRulesClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *FirewallRulesBeginDeleteOptions)` to `(context.Context, string, string, string, *FirewallRulesClientBeginDeleteOptions)`
- Function `*FirewallRulesClient.BeginDelete` return value(s) have been changed from `(FirewallRulesDeletePollerResponse, error)` to `(FirewallRulesClientDeletePollerResponse, error)`
- Function `*ServerGroupsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, ServerGroup, *ServerGroupsBeginCreateOrUpdateOptions)` to `(context.Context, string, string, ServerGroup, *ServerGroupsClientBeginCreateOrUpdateOptions)`
- Function `*ServerGroupsClient.BeginCreateOrUpdate` return value(s) have been changed from `(ServerGroupsCreateOrUpdatePollerResponse, error)` to `(ServerGroupsClientCreateOrUpdatePollerResponse, error)`
- Function `*RolesClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *RolesBeginDeleteOptions)` to `(context.Context, string, string, string, *RolesClientBeginDeleteOptions)`
- Function `*RolesClient.BeginDelete` return value(s) have been changed from `(RolesDeletePollerResponse, error)` to `(RolesClientDeletePollerResponse, error)`
- Function `*ServerGroupsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, ServerGroupForUpdate, *ServerGroupsBeginUpdateOptions)` to `(context.Context, string, string, ServerGroupForUpdate, *ServerGroupsClientBeginUpdateOptions)`
- Function `*ServerGroupsClient.BeginUpdate` return value(s) have been changed from `(ServerGroupsUpdatePollerResponse, error)` to `(ServerGroupsClientUpdatePollerResponse, error)`
- Function `*ConfigurationsClient.ListByServerGroup` parameter(s) have been changed from `(string, string, *ConfigurationsListByServerGroupOptions)` to `(string, string, *ConfigurationsClientListByServerGroupOptions)`
- Function `*ConfigurationsClient.ListByServerGroup` return value(s) have been changed from `(*ConfigurationsListByServerGroupPager)` to `(*ConfigurationsClientListByServerGroupPager)`
- Function `*ServerGroupsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *ServerGroupsGetOptions)` to `(context.Context, string, string, *ServerGroupsClientGetOptions)`
- Function `*ServerGroupsClient.Get` return value(s) have been changed from `(ServerGroupsGetResponse, error)` to `(ServerGroupsClientGetResponse, error)`
- Function `*RolesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, Role, *RolesBeginCreateOptions)` to `(context.Context, string, string, string, Role, *RolesClientBeginCreateOptions)`
- Function `*RolesClient.BeginCreate` return value(s) have been changed from `(RolesCreatePollerResponse, error)` to `(RolesClientCreatePollerResponse, error)`
- Function `*ConfigurationsClient.ListByServer` parameter(s) have been changed from `(string, string, string, *ConfigurationsListByServerOptions)` to `(string, string, string, *ConfigurationsClientListByServerOptions)`
- Function `*ConfigurationsClient.ListByServer` return value(s) have been changed from `(*ConfigurationsListByServerPager)` to `(*ConfigurationsClientListByServerPager)`
- Function `*FirewallRulesClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, FirewallRule, *FirewallRulesBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, FirewallRule, *FirewallRulesClientBeginCreateOrUpdateOptions)`
- Function `*FirewallRulesClient.BeginCreateOrUpdate` return value(s) have been changed from `(FirewallRulesCreateOrUpdatePollerResponse, error)` to `(FirewallRulesClientCreateOrUpdatePollerResponse, error)`
- Function `*ServerGroupsClient.BeginRestart` parameter(s) have been changed from `(context.Context, string, string, *ServerGroupsBeginRestartOptions)` to `(context.Context, string, string, *ServerGroupsClientBeginRestartOptions)`
- Function `*ServerGroupsClient.BeginRestart` return value(s) have been changed from `(ServerGroupsRestartPollerResponse, error)` to `(ServerGroupsClientRestartPollerResponse, error)`
- Function `*FirewallRulesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *FirewallRulesGetOptions)` to `(context.Context, string, string, string, *FirewallRulesClientGetOptions)`
- Function `*FirewallRulesClient.Get` return value(s) have been changed from `(FirewallRulesGetResponse, error)` to `(FirewallRulesClientGetResponse, error)`
- Function `*ServerGroupsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *ServerGroupsBeginDeleteOptions)` to `(context.Context, string, string, *ServerGroupsClientBeginDeleteOptions)`
- Function `*ServerGroupsClient.BeginDelete` return value(s) have been changed from `(ServerGroupsDeletePollerResponse, error)` to `(ServerGroupsClientDeletePollerResponse, error)`
- Function `*ConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ConfigurationsGetOptions)` to `(context.Context, string, string, string, *ConfigurationsClientGetOptions)`
- Function `*ConfigurationsClient.Get` return value(s) have been changed from `(ConfigurationsGetResponse, error)` to `(ConfigurationsClientGetResponse, error)`
- Function `*ServerGroupsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *ServerGroupsListByResourceGroupOptions)` to `(string, *ServerGroupsClientListByResourceGroupOptions)`
- Function `*ServerGroupsClient.ListByResourceGroup` return value(s) have been changed from `(*ServerGroupsListByResourceGroupPager)` to `(*ServerGroupsClientListByResourceGroupPager)`
- Function `*FirewallRulesClient.ListByServerGroup` parameter(s) have been changed from `(context.Context, string, string, *FirewallRulesListByServerGroupOptions)` to `(context.Context, string, string, *FirewallRulesClientListByServerGroupOptions)`
- Function `*FirewallRulesClient.ListByServerGroup` return value(s) have been changed from `(FirewallRulesListByServerGroupResponse, error)` to `(FirewallRulesClientListByServerGroupResponse, error)`
- Function `*ConfigurationsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, ServerGroupConfiguration, *ConfigurationsBeginUpdateOptions)` to `(context.Context, string, string, string, ServerGroupConfiguration, *ConfigurationsClientBeginUpdateOptions)`
- Function `*ConfigurationsClient.BeginUpdate` return value(s) have been changed from `(ConfigurationsUpdatePollerResponse, error)` to `(ConfigurationsClientUpdatePollerResponse, error)`
- Function `*ServersClient.ListByServerGroup` parameter(s) have been changed from `(context.Context, string, string, *ServersListByServerGroupOptions)` to `(context.Context, string, string, *ServersClientListByServerGroupOptions)`
- Function `*ServersClient.ListByServerGroup` return value(s) have been changed from `(ServersListByServerGroupResponse, error)` to `(ServersClientListByServerGroupResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*ServersClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ServersGetOptions)` to `(context.Context, string, string, string, *ServersClientGetOptions)`
- Function `*ServersClient.Get` return value(s) have been changed from `(ServersGetResponse, error)` to `(ServersClientGetResponse, error)`
- Function `*ServerGroupsClient.CheckNameAvailability` parameter(s) have been changed from `(context.Context, NameAvailabilityRequest, *ServerGroupsCheckNameAvailabilityOptions)` to `(context.Context, NameAvailabilityRequest, *ServerGroupsClientCheckNameAvailabilityOptions)`
- Function `*ServerGroupsClient.CheckNameAvailability` return value(s) have been changed from `(ServerGroupsCheckNameAvailabilityResponse, error)` to `(ServerGroupsClientCheckNameAvailabilityResponse, error)`
- Function `*ServerGroupsClient.BeginStart` parameter(s) have been changed from `(context.Context, string, string, *ServerGroupsBeginStartOptions)` to `(context.Context, string, string, *ServerGroupsClientBeginStartOptions)`
- Function `*ServerGroupsClient.BeginStart` return value(s) have been changed from `(ServerGroupsStartPollerResponse, error)` to `(ServerGroupsClientStartPollerResponse, error)`
- Function `*ServerGroupsClient.List` parameter(s) have been changed from `(*ServerGroupsListOptions)` to `(*ServerGroupsClientListOptions)`
- Function `*ServerGroupsClient.List` return value(s) have been changed from `(*ServerGroupsListPager)` to `(*ServerGroupsClientListPager)`
- Function `*ServerGroupsListPager.PageResponse` has been removed
- Function `*FirewallRulesCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*ServerGroupsDeletePoller.Done` has been removed
- Function `*ServerGroupsStopPoller.Done` has been removed
- Function `*FirewallRulesDeletePoller.Done` has been removed
- Function `*RolesDeletePoller.Done` has been removed
- Function `*ServerGroupsUpdatePoller.ResumeToken` has been removed
- Function `FirewallRulesCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsRestartPollerResponse.Resume` has been removed
- Function `*ConfigurationsListByServerPager.PageResponse` has been removed
- Function `*ConfigurationsListByServerGroupPager.PageResponse` has been removed
- Function `*ConfigurationsUpdatePoller.ResumeToken` has been removed
- Function `*ConfigurationsUpdatePoller.Done` has been removed
- Function `*ServerGroupsDeletePoller.ResumeToken` has been removed
- Function `ServerConfiguration.MarshalJSON` has been removed
- Function `*ServerGroupsCreateOrUpdatePoller.Done` has been removed
- Function `ServerGroupsDeletePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*ConfigurationsListByServerGroupPager.Err` has been removed
- Function `*ServerGroupsCreateOrUpdatePoller.Poll` has been removed
- Function `*ServerGroupsRestartPoller.ResumeToken` has been removed
- Function `CloudError.Error` has been removed
- Function `*ConfigurationsListByServerPager.Err` has been removed
- Function `*RolesCreatePoller.ResumeToken` has been removed
- Function `*RolesCreatePoller.Poll` has been removed
- Function `ServerGroupsStopPollerResponse.PollUntilDone` has been removed
- Function `*FirewallRulesDeletePoller.Poll` has been removed
- Function `*RolesDeletePollerResponse.Resume` has been removed
- Function `*ServerGroupsDeletePollerResponse.Resume` has been removed
- Function `*FirewallRulesCreateOrUpdatePoller.Poll` has been removed
- Function `ConfigurationsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsListPager.Err` has been removed
- Function `*ServerGroupsUpdatePoller.FinalResponse` has been removed
- Function `*ConfigurationsListByServerGroupPager.NextPage` has been removed
- Function `*ServerGroupsUpdatePollerResponse.Resume` has been removed
- Function `ServerGroupsStartPollerResponse.PollUntilDone` has been removed
- Function `*FirewallRulesCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*RolesCreatePoller.Done` has been removed
- Function `*ServerGroupsRestartPoller.Poll` has been removed
- Function `*ConfigurationsListByServerPager.NextPage` has been removed
- Function `*ServerGroupsRestartPoller.Done` has been removed
- Function `*ServerGroupsStopPoller.FinalResponse` has been removed
- Function `*ServerGroupsStartPollerResponse.Resume` has been removed
- Function `*ServerGroupsListByResourceGroupPager.PageResponse` has been removed
- Function `RolesCreatePollerResponse.PollUntilDone` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `*ServerGroupsUpdatePoller.Poll` has been removed
- Function `FirewallRulesDeletePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsListByResourceGroupPager.NextPage` has been removed
- Function `ServerGroupServerProperties.MarshalJSON` has been removed
- Function `ServerGroupsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ConfigurationsUpdatePollerResponse.Resume` has been removed
- Function `ServerGroupsRestartPollerResponse.PollUntilDone` has been removed
- Function `ServerProperties.MarshalJSON` has been removed
- Function `*RolesCreatePollerResponse.Resume` has been removed
- Function `*RolesDeletePoller.ResumeToken` has been removed
- Function `*RolesDeletePoller.FinalResponse` has been removed
- Function `*ServerGroupsRestartPoller.FinalResponse` has been removed
- Function `*ServerGroupsDeletePoller.Poll` has been removed
- Function `*ServerGroupsListByResourceGroupPager.Err` has been removed
- Function `*FirewallRulesDeletePoller.ResumeToken` has been removed
- Function `*FirewallRulesDeletePoller.FinalResponse` has been removed
- Function `Role.MarshalJSON` has been removed
- Function `*ServerGroupsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*ServerGroupsStopPollerResponse.Resume` has been removed
- Function `*FirewallRulesCreateOrUpdatePoller.Done` has been removed
- Function `*ServerGroupsStartPoller.FinalResponse` has been removed
- Function `RolesDeletePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsDeletePoller.FinalResponse` has been removed
- Function `*RolesDeletePoller.Poll` has been removed
- Function `ServerGroupsUpdatePollerResponse.PollUntilDone` has been removed
- Function `*ServerGroupsUpdatePoller.Done` has been removed
- Function `*ConfigurationsUpdatePoller.FinalResponse` has been removed
- Function `*ServerGroupsStartPoller.Poll` has been removed
- Function `*ServerGroupsStartPoller.Done` has been removed
- Function `*ServerGroupsListPager.NextPage` has been removed
- Function `*RolesCreatePoller.FinalResponse` has been removed
- Function `*ServerGroupsStopPoller.Poll` has been removed
- Function `*ServerGroupsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*ServerGroupsStartPoller.ResumeToken` has been removed
- Function `*ServerGroupsStopPoller.ResumeToken` has been removed
- Function `FirewallRule.MarshalJSON` has been removed
- Function `*FirewallRulesCreateOrUpdatePollerResponse.Resume` has been removed
- Function `ServerGroupServer.MarshalJSON` has been removed
- Function `*ConfigurationsUpdatePoller.Poll` has been removed
- Function `*FirewallRulesDeletePollerResponse.Resume` has been removed
- Struct `ConfigurationsBeginUpdateOptions` has been removed
- Struct `ConfigurationsGetOptions` has been removed
- Struct `ConfigurationsGetResponse` has been removed
- Struct `ConfigurationsGetResult` has been removed
- Struct `ConfigurationsListByServerGroupOptions` has been removed
- Struct `ConfigurationsListByServerGroupPager` has been removed
- Struct `ConfigurationsListByServerGroupResponse` has been removed
- Struct `ConfigurationsListByServerGroupResult` has been removed
- Struct `ConfigurationsListByServerOptions` has been removed
- Struct `ConfigurationsListByServerPager` has been removed
- Struct `ConfigurationsListByServerResponse` has been removed
- Struct `ConfigurationsListByServerResult` has been removed
- Struct `ConfigurationsUpdatePoller` has been removed
- Struct `ConfigurationsUpdatePollerResponse` has been removed
- Struct `ConfigurationsUpdateResponse` has been removed
- Struct `ConfigurationsUpdateResult` has been removed
- Struct `FirewallRulesBeginCreateOrUpdateOptions` has been removed
- Struct `FirewallRulesBeginDeleteOptions` has been removed
- Struct `FirewallRulesCreateOrUpdatePoller` has been removed
- Struct `FirewallRulesCreateOrUpdatePollerResponse` has been removed
- Struct `FirewallRulesCreateOrUpdateResponse` has been removed
- Struct `FirewallRulesCreateOrUpdateResult` has been removed
- Struct `FirewallRulesDeletePoller` has been removed
- Struct `FirewallRulesDeletePollerResponse` has been removed
- Struct `FirewallRulesDeleteResponse` has been removed
- Struct `FirewallRulesGetOptions` has been removed
- Struct `FirewallRulesGetResponse` has been removed
- Struct `FirewallRulesGetResult` has been removed
- Struct `FirewallRulesListByServerGroupOptions` has been removed
- Struct `FirewallRulesListByServerGroupResponse` has been removed
- Struct `FirewallRulesListByServerGroupResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `RolesBeginCreateOptions` has been removed
- Struct `RolesBeginDeleteOptions` has been removed
- Struct `RolesCreatePoller` has been removed
- Struct `RolesCreatePollerResponse` has been removed
- Struct `RolesCreateResponse` has been removed
- Struct `RolesCreateResult` has been removed
- Struct `RolesDeletePoller` has been removed
- Struct `RolesDeletePollerResponse` has been removed
- Struct `RolesDeleteResponse` has been removed
- Struct `RolesListByServerGroupOptions` has been removed
- Struct `RolesListByServerGroupResponse` has been removed
- Struct `RolesListByServerGroupResult` has been removed
- Struct `ServerGroupsBeginCreateOrUpdateOptions` has been removed
- Struct `ServerGroupsBeginDeleteOptions` has been removed
- Struct `ServerGroupsBeginRestartOptions` has been removed
- Struct `ServerGroupsBeginStartOptions` has been removed
- Struct `ServerGroupsBeginStopOptions` has been removed
- Struct `ServerGroupsBeginUpdateOptions` has been removed
- Struct `ServerGroupsCheckNameAvailabilityOptions` has been removed
- Struct `ServerGroupsCheckNameAvailabilityResponse` has been removed
- Struct `ServerGroupsCheckNameAvailabilityResult` has been removed
- Struct `ServerGroupsCreateOrUpdatePoller` has been removed
- Struct `ServerGroupsCreateOrUpdatePollerResponse` has been removed
- Struct `ServerGroupsCreateOrUpdateResponse` has been removed
- Struct `ServerGroupsCreateOrUpdateResult` has been removed
- Struct `ServerGroupsDeletePoller` has been removed
- Struct `ServerGroupsDeletePollerResponse` has been removed
- Struct `ServerGroupsDeleteResponse` has been removed
- Struct `ServerGroupsGetOptions` has been removed
- Struct `ServerGroupsGetResponse` has been removed
- Struct `ServerGroupsGetResult` has been removed
- Struct `ServerGroupsListByResourceGroupOptions` has been removed
- Struct `ServerGroupsListByResourceGroupPager` has been removed
- Struct `ServerGroupsListByResourceGroupResponse` has been removed
- Struct `ServerGroupsListByResourceGroupResult` has been removed
- Struct `ServerGroupsListOptions` has been removed
- Struct `ServerGroupsListPager` has been removed
- Struct `ServerGroupsListResponse` has been removed
- Struct `ServerGroupsListResult` has been removed
- Struct `ServerGroupsRestartPoller` has been removed
- Struct `ServerGroupsRestartPollerResponse` has been removed
- Struct `ServerGroupsRestartResponse` has been removed
- Struct `ServerGroupsStartPoller` has been removed
- Struct `ServerGroupsStartPollerResponse` has been removed
- Struct `ServerGroupsStartResponse` has been removed
- Struct `ServerGroupsStopPoller` has been removed
- Struct `ServerGroupsStopPollerResponse` has been removed
- Struct `ServerGroupsStopResponse` has been removed
- Struct `ServerGroupsUpdatePoller` has been removed
- Struct `ServerGroupsUpdatePollerResponse` has been removed
- Struct `ServerGroupsUpdateResponse` has been removed
- Struct `ServerGroupsUpdateResult` has been removed
- Struct `ServersGetOptions` has been removed
- Struct `ServersGetResponse` has been removed
- Struct `ServersGetResult` has been removed
- Struct `ServersListByServerGroupOptions` has been removed
- Struct `ServersListByServerGroupResponse` has been removed
- Struct `ServersListByServerGroupResult` has been removed
- Field `TrackedResource` of struct `ServerGroup` has been removed
- Field `ProxyResource` of struct `Role` has been removed
- Field `Resource` of struct `ProxyResource` has been removed
- Field `ProxyResource` of struct `ServerGroupConfiguration` has been removed
- Field `ProxyResource` of struct `ServerConfiguration` has been removed
- Field `Resource` of struct `TrackedResource` has been removed
- Field `ServerProperties` of struct `ServerGroupServerProperties` has been removed
- Field `ProxyResource` of struct `ServerGroupServer` has been removed
- Field `ProxyResource` of struct `FirewallRule` has been removed
- Field `InnerError` of struct `CloudError` has been removed
- Field `ServerProperties` of struct `ServerRoleGroup` has been removed

### Features Added

- New function `*ServerGroupsClientRestartPollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*RolesClientCreatePoller.Done() bool`
- New function `ConfigurationsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ConfigurationsClientUpdateResponse, error)`
- New function `*ServerGroupsClientDeletePoller.Done() bool`
- New function `*ServerGroupsClientStopPoller.FinalResponse(context.Context) (ServerGroupsClientStopResponse, error)`
- New function `*ServerGroupsClientCreateOrUpdatePollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*ServerGroupsClientStopPoller.ResumeToken() (string, error)`
- New function `*ConfigurationsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*ConfigurationsClientUpdatePoller.Done() bool`
- New function `*ServerGroupsClientListPager.Err() error`
- New function `*RolesClientCreatePollerResponse.Resume(context.Context, *RolesClient, string) error`
- New function `ServerGroupsClientRestartPollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientRestartResponse, error)`
- New function `*ServerGroupsClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*ServerGroupsClientDeletePoller.ResumeToken() (string, error)`
- New function `*FirewallRulesClientDeletePoller.Done() bool`
- New function `*ServerGroupsClientStartPoller.Poll(context.Context) (*http.Response, error)`
- New function `*ServerGroupsClientStopPoller.Poll(context.Context) (*http.Response, error)`
- New function `*ServerGroupsClientStartPoller.FinalResponse(context.Context) (ServerGroupsClientStartResponse, error)`
- New function `*FirewallRulesClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ConfigurationsClientUpdatePollerResponse.Resume(context.Context, *ConfigurationsClient, string) error`
- New function `*FirewallRulesClientCreateOrUpdatePoller.Done() bool`
- New function `*ServerGroupsClientDeletePollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*ConfigurationsClientListByServerPager.NextPage(context.Context) bool`
- New function `*ServerGroupsClientListByResourceGroupPager.Err() error`
- New function `*ServerGroupsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*RolesClientCreatePoller.ResumeToken() (string, error)`
- New function `*RolesClientCreatePoller.FinalResponse(context.Context) (RolesClientCreateResponse, error)`
- New function `*ServerGroupsClientListByResourceGroupPager.PageResponse() ServerGroupsClientListByResourceGroupResponse`
- New function `*ServerGroupsClientRestartPoller.Done() bool`
- New function `*ServerGroupsClientCreateOrUpdatePoller.FinalResponse(context.Context) (ServerGroupsClientCreateOrUpdateResponse, error)`
- New function `*RolesClientDeletePoller.ResumeToken() (string, error)`
- New function `*ServerGroupsClientStopPoller.Done() bool`
- New function `*ServerGroupsClientUpdatePoller.ResumeToken() (string, error)`
- New function `*RolesClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FirewallRulesClientCreateOrUpdatePoller.FinalResponse(context.Context) (FirewallRulesClientCreateOrUpdateResponse, error)`
- New function `*FirewallRulesClientDeletePoller.ResumeToken() (string, error)`
- New function `FirewallRulesClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (FirewallRulesClientDeleteResponse, error)`
- New function `*ServerGroupsClientStartPollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*ServerGroupsClientRestartPoller.ResumeToken() (string, error)`
- New function `*ServerGroupsClientStopPollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*ConfigurationsClientUpdatePoller.FinalResponse(context.Context) (ConfigurationsClientUpdateResponse, error)`
- New function `*ServerGroupsClientListPager.PageResponse() ServerGroupsClientListResponse`
- New function `*ServerGroupsClientDeletePoller.FinalResponse(context.Context) (ServerGroupsClientDeleteResponse, error)`
- New function `*FirewallRulesClientDeletePollerResponse.Resume(context.Context, *FirewallRulesClient, string) error`
- New function `*ServerGroupsClientUpdatePollerResponse.Resume(context.Context, *ServerGroupsClient, string) error`
- New function `*ServerGroupsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*RolesClientDeletePoller.Done() bool`
- New function `*ServerGroupsClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*RolesClientDeletePollerResponse.Resume(context.Context, *RolesClient, string) error`
- New function `*ServerGroupsClientUpdatePoller.FinalResponse(context.Context) (ServerGroupsClientUpdateResponse, error)`
- New function `*FirewallRulesClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ServerGroupsClientStartPoller.Done() bool`
- New function `*ServerGroupsClientRestartPoller.FinalResponse(context.Context) (ServerGroupsClientRestartResponse, error)`
- New function `*ServerGroupsClientUpdatePoller.Done() bool`
- New function `*ServerGroupsClientRestartPoller.Poll(context.Context) (*http.Response, error)`
- New function `*ServerGroupsClientStartPoller.ResumeToken() (string, error)`
- New function `*ConfigurationsClientListByServerGroupPager.Err() error`
- New function `*FirewallRulesClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `RolesClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (RolesClientDeleteResponse, error)`
- New function `*ConfigurationsClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*FirewallRulesClientDeletePoller.FinalResponse(context.Context) (FirewallRulesClientDeleteResponse, error)`
- New function `*ServerGroupsClientCreateOrUpdatePoller.Done() bool`
- New function `ServerGroupsClientStartPollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientStartResponse, error)`
- New function `RolesClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (RolesClientCreateResponse, error)`
- New function `*RolesClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*ConfigurationsClientListByServerPager.PageResponse() ConfigurationsClientListByServerResponse`
- New function `*ConfigurationsClientListByServerGroupPager.NextPage(context.Context) bool`
- New function `*ConfigurationsClientListByServerGroupPager.PageResponse() ConfigurationsClientListByServerGroupResponse`
- New function `ServerGroupsClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientUpdateResponse, error)`
- New function `ServerGroupsClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientCreateOrUpdateResponse, error)`
- New function `*ServerGroupsClientListPager.NextPage(context.Context) bool`
- New function `*RolesClientDeletePoller.FinalResponse(context.Context) (RolesClientDeleteResponse, error)`
- New function `*ConfigurationsClientListByServerPager.Err() error`
- New function `*ServerGroupsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*FirewallRulesClientCreateOrUpdatePollerResponse.Resume(context.Context, *FirewallRulesClient, string) error`
- New function `ServerGroupsClientStopPollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientStopResponse, error)`
- New function `ServerGroupsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (ServerGroupsClientDeleteResponse, error)`
- New function `FirewallRulesClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (FirewallRulesClientCreateOrUpdateResponse, error)`
- New struct `ConfigurationsClientBeginUpdateOptions`
- New struct `ConfigurationsClientGetOptions`
- New struct `ConfigurationsClientGetResponse`
- New struct `ConfigurationsClientGetResult`
- New struct `ConfigurationsClientListByServerGroupOptions`
- New struct `ConfigurationsClientListByServerGroupPager`
- New struct `ConfigurationsClientListByServerGroupResponse`
- New struct `ConfigurationsClientListByServerGroupResult`
- New struct `ConfigurationsClientListByServerOptions`
- New struct `ConfigurationsClientListByServerPager`
- New struct `ConfigurationsClientListByServerResponse`
- New struct `ConfigurationsClientListByServerResult`
- New struct `ConfigurationsClientUpdatePoller`
- New struct `ConfigurationsClientUpdatePollerResponse`
- New struct `ConfigurationsClientUpdateResponse`
- New struct `ConfigurationsClientUpdateResult`
- New struct `FirewallRulesClientBeginCreateOrUpdateOptions`
- New struct `FirewallRulesClientBeginDeleteOptions`
- New struct `FirewallRulesClientCreateOrUpdatePoller`
- New struct `FirewallRulesClientCreateOrUpdatePollerResponse`
- New struct `FirewallRulesClientCreateOrUpdateResponse`
- New struct `FirewallRulesClientCreateOrUpdateResult`
- New struct `FirewallRulesClientDeletePoller`
- New struct `FirewallRulesClientDeletePollerResponse`
- New struct `FirewallRulesClientDeleteResponse`
- New struct `FirewallRulesClientGetOptions`
- New struct `FirewallRulesClientGetResponse`
- New struct `FirewallRulesClientGetResult`
- New struct `FirewallRulesClientListByServerGroupOptions`
- New struct `FirewallRulesClientListByServerGroupResponse`
- New struct `FirewallRulesClientListByServerGroupResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `RolesClientBeginCreateOptions`
- New struct `RolesClientBeginDeleteOptions`
- New struct `RolesClientCreatePoller`
- New struct `RolesClientCreatePollerResponse`
- New struct `RolesClientCreateResponse`
- New struct `RolesClientCreateResult`
- New struct `RolesClientDeletePoller`
- New struct `RolesClientDeletePollerResponse`
- New struct `RolesClientDeleteResponse`
- New struct `RolesClientListByServerGroupOptions`
- New struct `RolesClientListByServerGroupResponse`
- New struct `RolesClientListByServerGroupResult`
- New struct `ServerGroupsClientBeginCreateOrUpdateOptions`
- New struct `ServerGroupsClientBeginDeleteOptions`
- New struct `ServerGroupsClientBeginRestartOptions`
- New struct `ServerGroupsClientBeginStartOptions`
- New struct `ServerGroupsClientBeginStopOptions`
- New struct `ServerGroupsClientBeginUpdateOptions`
- New struct `ServerGroupsClientCheckNameAvailabilityOptions`
- New struct `ServerGroupsClientCheckNameAvailabilityResponse`
- New struct `ServerGroupsClientCheckNameAvailabilityResult`
- New struct `ServerGroupsClientCreateOrUpdatePoller`
- New struct `ServerGroupsClientCreateOrUpdatePollerResponse`
- New struct `ServerGroupsClientCreateOrUpdateResponse`
- New struct `ServerGroupsClientCreateOrUpdateResult`
- New struct `ServerGroupsClientDeletePoller`
- New struct `ServerGroupsClientDeletePollerResponse`
- New struct `ServerGroupsClientDeleteResponse`
- New struct `ServerGroupsClientGetOptions`
- New struct `ServerGroupsClientGetResponse`
- New struct `ServerGroupsClientGetResult`
- New struct `ServerGroupsClientListByResourceGroupOptions`
- New struct `ServerGroupsClientListByResourceGroupPager`
- New struct `ServerGroupsClientListByResourceGroupResponse`
- New struct `ServerGroupsClientListByResourceGroupResult`
- New struct `ServerGroupsClientListOptions`
- New struct `ServerGroupsClientListPager`
- New struct `ServerGroupsClientListResponse`
- New struct `ServerGroupsClientListResult`
- New struct `ServerGroupsClientRestartPoller`
- New struct `ServerGroupsClientRestartPollerResponse`
- New struct `ServerGroupsClientRestartResponse`
- New struct `ServerGroupsClientStartPoller`
- New struct `ServerGroupsClientStartPollerResponse`
- New struct `ServerGroupsClientStartResponse`
- New struct `ServerGroupsClientStopPoller`
- New struct `ServerGroupsClientStopPollerResponse`
- New struct `ServerGroupsClientStopResponse`
- New struct `ServerGroupsClientUpdatePoller`
- New struct `ServerGroupsClientUpdatePollerResponse`
- New struct `ServerGroupsClientUpdateResponse`
- New struct `ServerGroupsClientUpdateResult`
- New struct `ServersClientGetOptions`
- New struct `ServersClientGetResponse`
- New struct `ServersClientGetResult`
- New struct `ServersClientListByServerGroupOptions`
- New struct `ServersClientListByServerGroupResponse`
- New struct `ServersClientListByServerGroupResult`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `ID` in struct `ServerGroupServer`
- New field `Name` in struct `ServerGroupServer`
- New field `Type` in struct `ServerGroupServer`
- New field `ID` in struct `ServerConfiguration`
- New field `Name` in struct `ServerConfiguration`
- New field `Type` in struct `ServerConfiguration`
- New field `Tags` in struct `ServerGroup`
- New field `ID` in struct `ServerGroup`
- New field `Name` in struct `ServerGroup`
- New field `Type` in struct `ServerGroup`
- New field `Location` in struct `ServerGroup`
- New field `ID` in struct `ProxyResource`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `Type` in struct `FirewallRule`
- New field `ID` in struct `FirewallRule`
- New field `Name` in struct `FirewallRule`
- New field `EnableHa` in struct `ServerGroupServerProperties`
- New field `StorageQuotaInMb` in struct `ServerGroupServerProperties`
- New field `VCores` in struct `ServerGroupServerProperties`
- New field `ServerEdition` in struct `ServerGroupServerProperties`
- New field `EnablePublicIP` in struct `ServerGroupServerProperties`
- New field `Type` in struct `Role`
- New field `ID` in struct `Role`
- New field `Name` in struct `Role`
- New field `EnableHa` in struct `ServerRoleGroup`
- New field `ServerEdition` in struct `ServerRoleGroup`
- New field `StorageQuotaInMb` in struct `ServerRoleGroup`
- New field `VCores` in struct `ServerRoleGroup`
- New field `EnablePublicIP` in struct `ServerRoleGroup`
- New field `Error` in struct `CloudError`
- New field `Name` in struct `ServerGroupConfiguration`
- New field `Type` in struct `ServerGroupConfiguration`
- New field `ID` in struct `ServerGroupConfiguration`


## 0.1.0 (2021-12-22)

- Init release.