# XSwitch XCC Proto Buffer 协议参考文档

<a name="top"></a>
<a name="user-content-top"></a>

这是[XCC API文档](https://docs.xswitch.cn/xcc-api/)的协议参考，使用[Google Protocol Buffers](https://protobuf.dev/)描述。

本文档只是对具体协议数据格式及类型的参考说明，详细的字段说明和用法请参考[XCC API列表](https://docs.xswitch.cn/xcc-api/api/)，原始的`.proto`文件可以在[proto](../proto)相关目录中找到。

## 目录

- [cman.proto](#cman.proto)
  - [BootstrapRequest](#cman.BootstrapRequest)
  - [BootstrapRequestParamsData](#cman.BootstrapRequestParamsData)
  - [BootstrapRequestParamsDataLiveArray](#cman.BootstrapRequestParamsDataLiveArray)
  - [BootstrapRequestParamsDataLiveArrayObj](#cman.BootstrapRequestParamsDataLiveArrayObj)
  - [BootstrapResponse](#cman.BootstrapResponse)
  - [BootstrapResponseParamsData](#cman.BootstrapResponseParamsData)
  - [BootstrapResponseParamsDataConfInfo](#cman.BootstrapResponseParamsDataConfInfo)
  - [BootstrapResponseParamsDataConfInfoGroupMembers](#cman.BootstrapResponseParamsDataConfInfoGroupMembers)
  - [ChangeLeaderRequest](#cman.ChangeLeaderRequest)
  - [ChangeLeaderResponse](#cman.ChangeLeaderResponse)
  - [ConferenceInfoRequest](#cman.ConferenceInfoRequest)
  - [ConferenceInfoResponse](#cman.ConferenceInfoResponse)
  - [ConferenceInfoResponseConference](#cman.ConferenceInfoResponseConference)
  - [ConferenceInfoResponseConferenceMembers](#cman.ConferenceInfoResponseConferenceMembers)
  - [ConferenceInfoResponseConferenceMembersStatus](#cman.ConferenceInfoResponseConferenceMembersStatus)
  - [ConferenceInfoResponseConferenceMembersStatusAudio](#cman.ConferenceInfoResponseConferenceMembersStatusAudio)
  - [ConferenceInfoResponseConferenceMembersStatusVideo](#cman.ConferenceInfoResponseConferenceMembersStatusVideo)
  - [ConferenceInfoResponseConferenceNodes](#cman.ConferenceInfoResponseConferenceNodes)
  - [ConferenceInfoResponseConferenceNodesGroupMembers](#cman.ConferenceInfoResponseConferenceNodesGroupMembers)
  - [ConferenceInfoResponseConferenceNodesNodeMembers](#cman.ConferenceInfoResponseConferenceNodesNodeMembers)
  - [ConferenceInfoResponseFlags](#cman.ConferenceInfoResponseFlags)
  - [ConferenceInfoResponseMembers](#cman.ConferenceInfoResponseMembers)
  - [ConferenceInfoResponseVariables](#cman.ConferenceInfoResponseVariables)
  - [ConferenceObject](#cman.ConferenceObject)
  - [ConferenceObjectNodes](#cman.ConferenceObjectNodes)
  - [ConferenceObjectNodesGroupMembers](#cman.ConferenceObjectNodesGroupMembers)
  - [ConferenceObjectNodesNodeMembers](#cman.ConferenceObjectNodesNodeMembers)
  - [EmptyMessage](#cman.EmptyMessage)
  - [GetCManInfoRequest](#cman.GetCManInfoRequest)
  - [GetCManInfoResponse](#cman.GetCManInfoResponse)
  - [GetConferenceListRequest](#cman.GetConferenceListRequest)
  - [GetConferenceListResponse](#cman.GetConferenceListResponse)
  - [confControlRequest](#cman.confControlRequest)
  - [confControlResponse](#cman.confControlResponse)
  - [getJWTRequest](#cman.getJWTRequest)
  - [getJWTRequestClaims](#cman.getJWTRequestClaims)
  - [getJWTResponse](#cman.getJWTResponse)
  - [getNodeForConferenceRequest](#cman.getNodeForConferenceRequest)
  - [getNodeForConferenceResponse](#cman.getNodeForConferenceResponse)
  - [loopSomeRequest](#cman.loopSomeRequest)
  - [loopSomeRequestMembers](#cman.loopSomeRequestMembers)
  - [loopSomeResponse](#cman.loopSomeResponse)
  - [pushMainRequest](#cman.pushMainRequest)
  - [pushMainResponse](#cman.pushMainResponse)
  - [pushMainSubCanvasRequest](#cman.pushMainSubCanvasRequest)
  - [pushMainSubCanvasRequestMembers](#cman.pushMainSubCanvasRequestMembers)
  - [pushMainSubCanvasResponse](#cman.pushMainSubCanvasResponse)
  - [pushSubCanvasRequest](#cman.pushSubCanvasRequest)
  - [pushSubCanvasRequestMembers](#cman.pushSubCanvasRequestMembers)
  - [pushSubCanvasResponse](#cman.pushSubCanvasResponse)
  - [pushWholeRequest](#cman.pushWholeRequest)
  - [pushWholeRequestMembers](#cman.pushWholeRequestMembers)
  - [pushWholeResponse](#cman.pushWholeResponse)



  - [cMan](#cman.cMan)


- [Scalar Value Types](#scalar-value-types)



<a name="user-content-cman.proto"/>
<a name="cman.proto"/>
<p align="right"><a href="#top">Top</a></p>

## cman.proto



<a name="user-content-cman.BootstrapRequest"/>
<a name="cman.BootstrapRequest"/>

### BootstrapRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_channel | [string](#string) |  |  |
| data | [BootstrapRequestParamsData](#cman.BootstrapRequestParamsData) |  |  |
| sessid | [string](#string) |  |  |






<a name="user-content-cman.BootstrapRequestParamsData"/>
<a name="cman.BootstrapRequestParamsData"/>

### BootstrapRequestParamsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_array | [BootstrapRequestParamsDataLiveArray](#cman.BootstrapRequestParamsDataLiveArray) |  |  |






<a name="user-content-cman.BootstrapRequestParamsDataLiveArray"/>
<a name="cman.BootstrapRequestParamsDataLiveArray"/>

### BootstrapRequestParamsDataLiveArray



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| context | [string](#string) |  |  |
| name | [string](#string) |  |  |
| obj | [BootstrapRequestParamsDataLiveArrayObj](#cman.BootstrapRequestParamsDataLiveArrayObj) |  |  |






<a name="user-content-cman.BootstrapRequestParamsDataLiveArrayObj"/>
<a name="cman.BootstrapRequestParamsDataLiveArrayObj"/>

### BootstrapRequestParamsDataLiveArrayObj







<a name="user-content-cman.BootstrapResponse"/>
<a name="cman.BootstrapResponse"/>

### BootstrapResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_channel | [string](#string) |  |  |
| sessid | [string](#string) |  |  |
| data | [BootstrapResponseParamsData](#cman.BootstrapResponseParamsData) |  |  |






<a name="user-content-cman.BootstrapResponseParamsData"/>
<a name="cman.BootstrapResponseParamsData"/>

### BootstrapResponseParamsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  |  |
| name | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| wire_serno | [int32](#int32) |  |  |
| conf_info | [BootstrapResponseParamsDataConfInfo](#cman.BootstrapResponseParamsDataConfInfo) |  |  |






<a name="user-content-cman.BootstrapResponseParamsDataConfInfo"/>
<a name="cman.BootstrapResponseParamsDataConfInfo"/>

### BootstrapResponseParamsDataConfInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| size | [int32](#int32) |  |  |
| group_members | [BootstrapResponseParamsDataConfInfoGroupMembers](#cman.BootstrapResponseParamsDataConfInfoGroupMembers) |  |  |






<a name="user-content-cman.BootstrapResponseParamsDataConfInfoGroupMembers"/>
<a name="cman.BootstrapResponseParamsDataConfInfoGroupMembers"/>

### BootstrapResponseParamsDataConfInfoGroupMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no_group | [int32](#int32) |  |  |






<a name="user-content-cman.ChangeLeaderRequest"/>
<a name="cman.ChangeLeaderRequest"/>

### ChangeLeaderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | the instance UUID |
| mode | [string](#string) |  |  |






<a name="user-content-cman.ChangeLeaderResponse"/>
<a name="cman.ChangeLeaderResponse"/>

### ChangeLeaderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | the instance UUID |
| is_leader | [bool](#bool) |  |  |
| mode | [string](#string) |  | ACTIVE or STANDBY in HA,LEADER, FOLLOWER, or CANDIDATE in RAFT |






<a name="user-content-cman.ConferenceInfoRequest"/>
<a name="cman.ConferenceInfoRequest"/>

### ConferenceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| showMembers | [bool](#bool) |  |  |
| memberFilters | [map<string, string>](#map-string-string) |  |  |






<a name="user-content-cman.ConferenceInfoResponse"/>
<a name="cman.ConferenceInfoResponse"/>

### ConferenceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| conference | [ConferenceInfoResponseConference](#cman.ConferenceInfoResponseConference) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConference"/>
<a name="cman.ConferenceInfoResponseConference"/>

### ConferenceInfoResponseConference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| nodes | [ConferenceInfoResponseConferenceNodes](#cman.ConferenceInfoResponseConferenceNodes) | repeated |  |
| size | [int32](#int32) |  |  |
| member_count | [int32](#int32) |  |  |
| states | [string](#string) |  |  |
| members | [ConferenceInfoResponseConferenceMembers](#cman.ConferenceInfoResponseConferenceMembers) | repeated |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceMembers"/>
<a name="cman.ConferenceInfoResponseConferenceMembers"/>

### ConferenceInfoResponseConferenceMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberID | [string](#string) |  |  |
| cidNumber | [string](#string) |  |  |
| cidName | [string](#string) |  |  |
| codec | [string](#string) |  |  |
| status | [ConferenceInfoResponseConferenceMembersStatus](#cman.ConferenceInfoResponseConferenceMembersStatus) |  |  |
| email | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| active | [bool](#bool) |  |  |
| uuid | [string](#string) |  |  |
| groupID | [string](#string) |  |  |
| memberType | [string](#string) |  |  |
| nodeType | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| canvasID | [int32](#int32) |  |  |
| watchingCanvasID | [int32](#int32) |  |  |
| roleID | [string](#string) |  |  |
| layerID | [int32](#int32) |  |  |
| reservationID | [string](#string) |  |  |
| channelName | [string](#string) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceMembersStatus"/>
<a name="cman.ConferenceInfoResponseConferenceMembersStatus"/>

### ConferenceInfoResponseConferenceMembersStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| audio | [ConferenceInfoResponseConferenceMembersStatusAudio](#cman.ConferenceInfoResponseConferenceMembersStatusAudio) |  |  |
| video | [ConferenceInfoResponseConferenceMembersStatusVideo](#cman.ConferenceInfoResponseConferenceMembersStatusVideo) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceMembersStatusAudio"/>
<a name="cman.ConferenceInfoResponseConferenceMembersStatusAudio"/>

### ConferenceInfoResponseConferenceMembersStatusAudio



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| talking | [bool](#bool) |  |  |
| deaf | [bool](#bool) |  |  |
| muted | [bool](#bool) |  |  |
| onHold | [bool](#bool) |  |  |
| energyScore | [int32](#int32) |  |  |
| floor | [bool](#bool) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceMembersStatusVideo"/>
<a name="cman.ConferenceInfoResponseConferenceMembersStatusVideo"/>

### ConferenceInfoResponseConferenceMembersStatusVideo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| visible | [bool](#bool) |  |  |
| noRecover | [bool](#bool) |  |  |
| avatarPresented | [bool](#bool) |  |  |
| mediaFlow | [bool](#bool) |  |  |
| muted | [bool](#bool) |  |  |
| floor | [bool](#bool) |  |  |
| reservationID | [string](#string) |  |  |
| roleID | [string](#string) |  |  |
| videoLayerID | [int32](#int32) |  |  |
| canvasID | [int32](#int32) |  |  |
| watchingCanvasID | [int32](#int32) |  |  |
| order | [int32](#int32) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceNodes"/>
<a name="cman.ConferenceInfoResponseConferenceNodes"/>

### ConferenceInfoResponseConferenceNodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeUUID | [string](#string) |  |  |
| confMemberCount | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| GroupMembers | [ConferenceInfoResponseConferenceNodesGroupMembers](#cman.ConferenceInfoResponseConferenceNodesGroupMembers) | repeated |  |
| confCreatTime | [string](#string) |  |  |
| confUUID | [string](#string) |  |  |
| nodeType | [string](#string) |  |  |
| localIP | [string](#string) |  |  |
| nodeMembers | [ConferenceInfoResponseConferenceNodesNodeMembers](#cman.ConferenceInfoResponseConferenceNodesNodeMembers) |  |  |
| moderatorMemberID | [string](#string) |  |  |
| destNumber | [string](#string) |  |  |
| action | [string](#string) |  |  |
| state | [string](#string) |  |  |
| lastUpdated | [string](#string) |  |  |
| port | [int32](#int32) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceNodesGroupMembers"/>
<a name="cman.ConferenceInfoResponseConferenceNodesGroupMembers"/>

### ConferenceInfoResponseConferenceNodesGroupMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  |  |
| count | [int32](#int32) |  |  |






<a name="user-content-cman.ConferenceInfoResponseConferenceNodesNodeMembers"/>
<a name="cman.ConferenceInfoResponseConferenceNodesNodeMembers"/>

### ConferenceInfoResponseConferenceNodesNodeMembers







<a name="user-content-cman.ConferenceInfoResponseFlags"/>
<a name="cman.ConferenceInfoResponseFlags"/>

### ConferenceInfoResponseFlags



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| can_hear | [bool](#bool) |  |  |
| can_see | [bool](#bool) |  |  |
| can_speak | [bool](#bool) |  |  |
| hold | [bool](#bool) |  |  |
| mute_detect | [bool](#bool) |  |  |
| talking | [bool](#bool) |  |  |
| has_video | [bool](#bool) |  |  |
| video_bridge | [bool](#bool) |  |  |
| has_floor | [bool](#bool) |  |  |
| is_moderator | [bool](#bool) |  |  |
| end_conference | [bool](#bool) |  |  |






<a name="user-content-cman.ConferenceInfoResponseMembers"/>
<a name="cman.ConferenceInfoResponseMembers"/>

### ConferenceInfoResponseMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [int32](#int32) |  |  |
| uuid | [string](#string) |  |  |
| caller_id_name | [string](#string) |  |  |
| caller_id_number | [string](#string) |  |  |
| join_time | [int32](#int32) |  |  |
| last_talking | [int32](#int32) |  |  |
| energy | [int32](#int32) |  |  |
| volume_in | [int32](#int32) |  |  |
| volume_out | [int32](#int32) |  |  |
| output_volume | [int32](#int32) |  |  |
| input_volume | [int32](#int32) |  |  |
| flags | [ConferenceInfoResponseFlags](#cman.ConferenceInfoResponseFlags) |  |  |






<a name="user-content-cman.ConferenceInfoResponseVariables"/>
<a name="cman.ConferenceInfoResponseVariables"/>

### ConferenceInfoResponseVariables







<a name="user-content-cman.ConferenceObject"/>
<a name="cman.ConferenceObject"/>

### ConferenceObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| nodes | [ConferenceObjectNodes](#cman.ConferenceObjectNodes) | repeated |  |
| size | [int32](#int32) |  |  |
| memberCount | [string](#string) |  |  |
| state | [string](#string) |  |  |






<a name="user-content-cman.ConferenceObjectNodes"/>
<a name="cman.ConferenceObjectNodes"/>

### ConferenceObjectNodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeUUID | [string](#string) |  |  |
| confMemberCount | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| groupMembers | [ConferenceObjectNodesGroupMembers](#cman.ConferenceObjectNodesGroupMembers) | repeated |  |
| confCreatedTime | [string](#string) |  |  |
| confUUID | [string](#string) |  |  |
| localIP | [string](#string) |  |  |
| nodeMembers | [ConferenceObjectNodesNodeMembers](#cman.ConferenceObjectNodesNodeMembers) | repeated |  |
| moderatorMemberID | [string](#string) |  |  |
| destNumber | [string](#string) |  |  |
| action | [string](#string) |  |  |
| state | [string](#string) |  |  |
| lastUpdated | [string](#string) |  |  |






<a name="user-content-cman.ConferenceObjectNodesGroupMembers"/>
<a name="cman.ConferenceObjectNodesGroupMembers"/>

### ConferenceObjectNodesGroupMembers







<a name="user-content-cman.ConferenceObjectNodesNodeMembers"/>
<a name="cman.ConferenceObjectNodesNodeMembers"/>

### ConferenceObjectNodesNodeMembers







<a name="user-content-cman.EmptyMessage"/>
<a name="cman.EmptyMessage"/>

### EmptyMessage







<a name="user-content-cman.GetCManInfoRequest"/>
<a name="cman.GetCManInfoRequest"/>

### GetCManInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reply | [string](#string) |  | the reply address |






<a name="user-content-cman.GetCManInfoResponse"/>
<a name="cman.GetCManInfoResponse"/>

### GetCManInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | the instance UUID |
| is_leader | [bool](#bool) |  |  |
| mode | [string](#string) |  | ACTIVE or STANDBY in HA,LEADER, FOLLOWER, or CANDIDATE in RAFT |






<a name="user-content-cman.GetConferenceListRequest"/>
<a name="cman.GetConferenceListRequest"/>

### GetConferenceListRequest







<a name="user-content-cman.GetConferenceListResponse"/>
<a name="cman.GetConferenceListResponse"/>

### GetConferenceListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| conferences | [ConferenceObject](#cman.ConferenceObject) | repeated |  |






<a name="user-content-cman.confControlRequest"/>
<a name="cman.confControlRequest"/>

### confControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conferenceName | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| memberID | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="user-content-cman.confControlResponse"/>
<a name="cman.confControlResponse"/>

### confControlResponse







<a name="user-content-cman.getJWTRequest"/>
<a name="cman.getJWTRequest"/>

### getJWTRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [string](#string) |  |  |
| claims | [getJWTRequestClaims](#cman.getJWTRequestClaims) |  |  |






<a name="user-content-cman.getJWTRequestClaims"/>
<a name="cman.getJWTRequestClaims"/>

### getJWTRequestClaims



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |
| expires | [int32](#int32) |  |  |
| username | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| login | [string](#string) |  |  |
| session_uuid | [string](#string) |  |  |






<a name="user-content-cman.getJWTResponse"/>
<a name="cman.getJWTResponse"/>

### getJWTResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="user-content-cman.getNodeForConferenceRequest"/>
<a name="cman.getNodeForConferenceRequest"/>

### getNodeForConferenceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="user-content-cman.getNodeForConferenceResponse"/>
<a name="cman.getNodeForConferenceResponse"/>

### getNodeForConferenceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeUUID | [string](#string) |  |  |






<a name="user-content-cman.loopSomeRequest"/>
<a name="cman.loopSomeRequest"/>

### loopSomeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conference_name | [string](#string) |  |  |
| interval | [string](#string) |  |  |
| layout | [string](#string) |  |  |
| members | [loopSomeRequestMembers](#cman.loopSomeRequestMembers) | repeated |  |






<a name="user-content-cman.loopSomeRequestMembers"/>
<a name="cman.loopSomeRequestMembers"/>

### loopSomeRequestMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberID | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |






<a name="user-content-cman.loopSomeResponse"/>
<a name="cman.loopSomeResponse"/>

### loopSomeResponse







<a name="user-content-cman.pushMainRequest"/>
<a name="cman.pushMainRequest"/>

### pushMainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conference_name | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| stream | [int32](#int32) |  |  |






<a name="user-content-cman.pushMainResponse"/>
<a name="cman.pushMainResponse"/>

### pushMainResponse







<a name="user-content-cman.pushMainSubCanvasRequest"/>
<a name="cman.pushMainSubCanvasRequest"/>

### pushMainSubCanvasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conference_name | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| layout | [string](#string) |  |  |
| members | [pushMainSubCanvasRequestMembers](#cman.pushMainSubCanvasRequestMembers) | repeated |  |






<a name="user-content-cman.pushMainSubCanvasRequestMembers"/>
<a name="cman.pushMainSubCanvasRequestMembers"/>

### pushMainSubCanvasRequestMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberID | [string](#string) |  |  |
| position | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |






<a name="user-content-cman.pushMainSubCanvasResponse"/>
<a name="cman.pushMainSubCanvasResponse"/>

### pushMainSubCanvasResponse







<a name="user-content-cman.pushSubCanvasRequest"/>
<a name="cman.pushSubCanvasRequest"/>

### pushSubCanvasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conference_name | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| layout | [string](#string) |  |  |
| members | [pushSubCanvasRequestMembers](#cman.pushSubCanvasRequestMembers) | repeated |  |






<a name="user-content-cman.pushSubCanvasRequestMembers"/>
<a name="cman.pushSubCanvasRequestMembers"/>

### pushSubCanvasRequestMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberID | [string](#string) |  |  |
| position | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |






<a name="user-content-cman.pushSubCanvasResponse"/>
<a name="cman.pushSubCanvasResponse"/>

### pushSubCanvasResponse







<a name="user-content-cman.pushWholeRequest"/>
<a name="cman.pushWholeRequest"/>

### pushWholeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  |  |
| conference_name | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |
| layout | [string](#string) |  |  |
| members | [pushWholeRequestMembers](#cman.pushWholeRequestMembers) | repeated |  |






<a name="user-content-cman.pushWholeRequestMembers"/>
<a name="cman.pushWholeRequestMembers"/>

### pushWholeRequestMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberID | [string](#string) |  |  |
| position | [string](#string) |  |  |
| nodeUUID | [string](#string) |  |  |






<a name="user-content-cman.pushWholeResponse"/>
<a name="cman.pushWholeResponse"/>

### pushWholeResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="user-content-cman.cMan"/>
<a name="cman.cMan"/>

### cMan


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| getConferenceList | [GetConferenceListRequest](#cman.GetConferenceListRequest) | [GetConferenceListResponse](#cman.GetConferenceListRequest) | 获取会议列表 |
| conferenceInfo | [ConferenceInfoRequest](#cman.ConferenceInfoRequest) | [ConferenceInfoResponse](#cman.ConferenceInfoRequest) | 获取会议信息 |
| bootstrap | [BootstrapRequest](#cman.BootstrapRequest) | [BootstrapResponse](#cman.BootstrapRequest) | bootstrap 客户端侧冷启动时发送，刷新当前会议状态和成员列表 |
| getNodeForConference | [getNodeForConferenceRequest](#cman.getNodeForConferenceRequest) | [getNodeForConferenceResponse](#cman.getNodeForConferenceRequest) | getNodeForConference 小会参会者入会前询问小会所在节点 |
| confControl | [confControlRequest](#cman.confControlRequest) | [confControlResponse](#cman.confControlRequest) | confControl 通用的会控接口 |
| pushMain | [pushMainRequest](#cman.pushMainRequest) | [pushMainResponse](#cman.pushMainRequest) | 推送主会场。主会场推送给全体参会人员 |
| pushSubCanvas | [pushSubCanvasRequest](#cman.pushSubCanvasRequest) | [pushSubCanvasResponse](#cman.pushSubCanvasRequest) | 推送合屏。指定一部分参会者合成一屏推送给全体参会人员，当前版本API，这些参会者必须在同一个节点上。 |
| pushMainSubCanvas | [pushMainSubCanvasRequest](#cman.pushMainSubCanvasRequest) | [pushMainSubCanvasResponse](#cman.pushMainSubCanvasRequest) | 主会场单看。指定一部分参会者合成一屏单独推送给主会场，当前版本API，这些参会者必须在同一个节点上；除了主会场外的其他参会人员看主会场。 |
| pushWhole | [pushWholeRequest](#cman.pushWholeRequest) | [pushWholeResponse](#cman.pushWholeRequest) | 推送整体。指定一部分参会者跟主会场一起推送给所有人，当前版本API，这些参会者必须在同一个节点上。 |
| loopSome | [loopSomeRequest](#cman.loopSomeRequest) | [loopSomeResponse](#cman.loopSomeRequest) | 轮播指定的参会者，当前版本API，这些参会者必须在同一个节点上。 |
| getJWT | [getJWTRequest](#cman.getJWTRequest) | [getJWTResponse](#cman.getJWTRequest) | 获取JWT Token |
| getCManInfo | [GetCManInfoRequest](#cman.GetCManInfoRequest) | [GetCManInfoResponse](#cman.GetCManInfoRequest) | 获取cMan实例列表 |
| changeLeader | [ChangeLeaderRequest](#cman.ChangeLeaderRequest) | [ChangeLeaderResponse](#cman.ChangeLeaderRequest) |  |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="user-content-double" /><a name="double" /> double |  | double | double | float |
| <a name="user-content-float" /><a name="float" /> float |  | float | float | float |
| <a name="user-content-int32" /><a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="user-content-int64" /><a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="user-content-uint32" /><a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="user-content-uint64" /><a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="user-content-sint32" /><a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="user-content-sint64" /><a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="user-content-fixed32" /><a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="user-content-fixed64" /><a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="user-content-sfixed32" /><a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="user-content-sfixed64" /><a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="user-content-bool" /><a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="user-content-string" /><a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="user-content-bytes" /><a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |


<a name="user-content-map-string-string" />

map&lt;[string](#string), [string](#string)&gt;
