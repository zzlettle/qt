#ifdef __cplusplus
extern "C" {
#endif

void* QGeoManeuver_NewQGeoManeuver();
void* QGeoManeuver_NewQGeoManeuver2(void* other);
int QGeoManeuver_Direction(void* ptr);
double QGeoManeuver_DistanceToNextInstruction(void* ptr);
char* QGeoManeuver_InstructionText(void* ptr);
int QGeoManeuver_IsValid(void* ptr);
void QGeoManeuver_SetDirection(void* ptr, int direction);
void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance);
void QGeoManeuver_SetInstructionText(void* ptr, char* instructionText);
void QGeoManeuver_SetPosition(void* ptr, void* position);
void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs);
void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate);
int QGeoManeuver_TimeToNextInstruction(void* ptr);
void QGeoManeuver_DestroyQGeoManeuver(void* ptr);
void* QGeoRoute_NewQGeoRoute();
void* QGeoRoute_NewQGeoRoute2(void* other);
double QGeoRoute_Distance(void* ptr);
char* QGeoRoute_RouteId(void* ptr);
void QGeoRoute_SetBounds(void* ptr, void* bounds);
void QGeoRoute_SetDistance(void* ptr, double distance);
void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment);
void QGeoRoute_SetRequest(void* ptr, void* request);
void QGeoRoute_SetRouteId(void* ptr, char* id);
void QGeoRoute_SetTravelMode(void* ptr, int mode);
void QGeoRoute_SetTravelTime(void* ptr, int secs);
int QGeoRoute_TravelMode(void* ptr);
int QGeoRoute_TravelTime(void* ptr);
void QGeoRoute_DestroyQGeoRoute(void* ptr);
void* QGeoRouteReply_NewQGeoRouteReply(int error, char* errorString, void* parent);
void QGeoRouteReply_Abort(void* ptr);
int QGeoRouteReply_Error(void* ptr);
char* QGeoRouteReply_ErrorString(void* ptr);
void QGeoRouteReply_ConnectFinished(void* ptr);
void QGeoRouteReply_DisconnectFinished(void* ptr);
int QGeoRouteReply_IsFinished(void* ptr);
void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr);
void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination);
void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other);
int QGeoRouteRequest_FeatureWeight(void* ptr, int featureType);
int QGeoRouteRequest_ManeuverDetail(void* ptr);
int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr);
int QGeoRouteRequest_RouteOptimization(void* ptr);
int QGeoRouteRequest_SegmentDetail(void* ptr);
void QGeoRouteRequest_SetFeatureWeight(void* ptr, int featureType, int featureWeight);
void QGeoRouteRequest_SetManeuverDetail(void* ptr, int maneuverDetail);
void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives);
void QGeoRouteRequest_SetRouteOptimization(void* ptr, int optimization);
void QGeoRouteRequest_SetSegmentDetail(void* ptr, int segmentDetail);
void QGeoRouteRequest_SetTravelModes(void* ptr, int travelModes);
int QGeoRouteRequest_TravelModes(void* ptr);
void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr);
void* QGeoRouteSegment_NewQGeoRouteSegment();
void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other);
double QGeoRouteSegment_Distance(void* ptr);
int QGeoRouteSegment_IsValid(void* ptr);
void QGeoRouteSegment_SetDistance(void* ptr, double distance);
void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver);
void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment);
void QGeoRouteSegment_SetTravelTime(void* ptr, int secs);
int QGeoRouteSegment_TravelTime(void* ptr);
void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr);
void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManager_ConnectError(void* ptr);
void QGeoRoutingManager_DisconnectError(void* ptr);
void QGeoRoutingManager_ConnectFinished(void* ptr);
void QGeoRoutingManager_DisconnectFinished(void* ptr);
char* QGeoRoutingManager_ManagerName(void* ptr);
int QGeoRoutingManager_ManagerVersion(void* ptr);
int QGeoRoutingManager_MeasurementSystem(void* ptr);
void QGeoRoutingManager_SetLocale(void* ptr, void* locale);
void QGeoRoutingManager_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManager_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManager_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManager_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManager_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManager_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManager_SupportedTravelModes(void* ptr);
void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position);
void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr);
void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request);
void QGeoRoutingManagerEngine_ConnectError(void* ptr);
void QGeoRoutingManagerEngine_DisconnectError(void* ptr);
void QGeoRoutingManagerEngine_ConnectFinished(void* ptr);
void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr);
char* QGeoRoutingManagerEngine_ManagerName(void* ptr);
int QGeoRoutingManagerEngine_ManagerVersion(void* ptr);
int QGeoRoutingManagerEngine_MeasurementSystem(void* ptr);
void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale);
void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, int system);
int QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr);
int QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr);
int QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr);
int QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr);
int QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr);
void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position);
void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr);
int QGeoServiceProvider_OnlineGeocodingFeature_Type();
int QGeoServiceProvider_OfflineGeocodingFeature_Type();
int QGeoServiceProvider_ReverseGeocodingFeature_Type();
int QGeoServiceProvider_LocalizedGeocodingFeature_Type();
int QGeoServiceProvider_AnyGeocodingFeatures_Type();
int QGeoServiceProvider_OnlineMappingFeature_Type();
int QGeoServiceProvider_OfflineMappingFeature_Type();
int QGeoServiceProvider_LocalizedMappingFeature_Type();
int QGeoServiceProvider_AnyMappingFeatures_Type();
int QGeoServiceProvider_OnlinePlacesFeature_Type();
int QGeoServiceProvider_OfflinePlacesFeature_Type();
int QGeoServiceProvider_SavePlaceFeature_Type();
int QGeoServiceProvider_RemovePlaceFeature_Type();
int QGeoServiceProvider_SaveCategoryFeature_Type();
int QGeoServiceProvider_RemoveCategoryFeature_Type();
int QGeoServiceProvider_PlaceRecommendationsFeature_Type();
int QGeoServiceProvider_SearchSuggestionsFeature_Type();
int QGeoServiceProvider_LocalizedPlacesFeature_Type();
int QGeoServiceProvider_NotificationsFeature_Type();
int QGeoServiceProvider_PlaceMatchingFeature_Type();
int QGeoServiceProvider_AnyPlacesFeatures_Type();
int QGeoServiceProvider_OnlineRoutingFeature_Type();
int QGeoServiceProvider_OfflineRoutingFeature_Type();
int QGeoServiceProvider_LocalizedRoutingFeature_Type();
int QGeoServiceProvider_RouteUpdatesFeature_Type();
int QGeoServiceProvider_AlternativeRoutesFeature_Type();
int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type();
int QGeoServiceProvider_AnyRoutingFeatures_Type();
char* QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders();
int QGeoServiceProvider_Error(void* ptr);
char* QGeoServiceProvider_ErrorString(void* ptr);
int QGeoServiceProvider_GeocodingFeatures(void* ptr);
void* QGeoServiceProvider_GeocodingManager(void* ptr);
int QGeoServiceProvider_MappingFeatures(void* ptr);
void* QGeoServiceProvider_PlaceManager(void* ptr);
int QGeoServiceProvider_PlacesFeatures(void* ptr);
int QGeoServiceProvider_RoutingFeatures(void* ptr);
void* QGeoServiceProvider_RoutingManager(void* ptr);
void QGeoServiceProvider_SetAllowExperimental(void* ptr, int allow);
void QGeoServiceProvider_SetLocale(void* ptr, void* locale);
void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr);
void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr);
char* QGeoServiceProviderFactory_ObjectNameAbs(void* ptr);
void QGeoServiceProviderFactory_SetObjectNameAbs(void* ptr, char* name);

#ifdef __cplusplus
}
#endif