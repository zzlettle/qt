#ifdef __cplusplus
extern "C" {
#endif

void* QSqlDatabase_NewQSqlDatabase();
void* QSqlDatabase_NewQSqlDatabase2(void* other);
void QSqlDatabase_Close(void* ptr);
int QSqlDatabase_Commit(void* ptr);
char* QSqlDatabase_ConnectOptions(void* ptr);
char* QSqlDatabase_ConnectionName(void* ptr);
char* QSqlDatabase_QSqlDatabase_ConnectionNames();
int QSqlDatabase_QSqlDatabase_Contains(char* connectionName);
char* QSqlDatabase_DatabaseName(void* ptr);
void* QSqlDatabase_Driver(void* ptr);
char* QSqlDatabase_DriverName(void* ptr);
char* QSqlDatabase_QSqlDatabase_Drivers();
char* QSqlDatabase_HostName(void* ptr);
int QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name);
int QSqlDatabase_IsOpen(void* ptr);
int QSqlDatabase_IsOpenError(void* ptr);
int QSqlDatabase_IsValid(void* ptr);
int QSqlDatabase_Open(void* ptr);
int QSqlDatabase_Open2(void* ptr, char* user, char* password);
char* QSqlDatabase_Password(void* ptr);
int QSqlDatabase_Port(void* ptr);
void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, void* creator);
void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName);
int QSqlDatabase_Rollback(void* ptr);
void QSqlDatabase_SetConnectOptions(void* ptr, char* options);
void QSqlDatabase_SetDatabaseName(void* ptr, char* name);
void QSqlDatabase_SetHostName(void* ptr, char* host);
void QSqlDatabase_SetPassword(void* ptr, char* password);
void QSqlDatabase_SetPort(void* ptr, int port);
void QSqlDatabase_SetUserName(void* ptr, char* name);
int QSqlDatabase_Transaction(void* ptr);
char* QSqlDatabase_UserName(void* ptr);
void QSqlDatabase_DestroyQSqlDatabase(void* ptr);
int QSqlDriver_BeginTransaction(void* ptr);
void QSqlDriver_Close(void* ptr);
int QSqlDriver_CommitTransaction(void* ptr);
void* QSqlDriver_CreateResult(void* ptr);
int QSqlDriver_DbmsType(void* ptr);
char* QSqlDriver_EscapeIdentifier(void* ptr, char* identifier, int ty);
char* QSqlDriver_FormatValue(void* ptr, void* field, int trimStrings);
void* QSqlDriver_Handle(void* ptr);
int QSqlDriver_HasFeature(void* ptr, int feature);
int QSqlDriver_IsIdentifierEscaped(void* ptr, char* identifier, int ty);
int QSqlDriver_IsOpen(void* ptr);
int QSqlDriver_IsOpenError(void* ptr);
void QSqlDriver_ConnectNotification(void* ptr);
void QSqlDriver_DisconnectNotification(void* ptr);
int QSqlDriver_Open(void* ptr, char* db, char* user, char* password, char* host, int port, char* options);
int QSqlDriver_RollbackTransaction(void* ptr);
char* QSqlDriver_SqlStatement(void* ptr, int ty, char* tableName, void* rec, int preparedStatement);
char* QSqlDriver_StripDelimiters(void* ptr, char* identifier, int ty);
int QSqlDriver_SubscribeToNotification(void* ptr, char* name);
char* QSqlDriver_SubscribedToNotifications(void* ptr);
int QSqlDriver_UnsubscribeFromNotification(void* ptr, char* name);
void QSqlDriver_DestroyQSqlDriver(void* ptr);
void* QSqlDriverCreatorBase_CreateObject(void* ptr);
void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(void* ptr);
char* QSqlDriverCreatorBase_ObjectNameAbs(void* ptr);
void QSqlDriverCreatorBase_SetObjectNameAbs(void* ptr, char* name);
void* QSqlDriverPlugin_Create(void* ptr, char* key);
void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr);
void* QSqlError_NewQSqlError3(void* other);
void* QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code);
char* QSqlError_DatabaseText(void* ptr);
char* QSqlError_DriverText(void* ptr);
int QSqlError_IsValid(void* ptr);
char* QSqlError_NativeErrorCode(void* ptr);
char* QSqlError_Text(void* ptr);
int QSqlError_Type(void* ptr);
void QSqlError_DestroyQSqlError(void* ptr);
void* QSqlField_NewQSqlField2(void* other);
void QSqlField_Clear(void* ptr);
void* QSqlField_DefaultValue(void* ptr);
int QSqlField_IsAutoValue(void* ptr);
int QSqlField_IsGenerated(void* ptr);
int QSqlField_IsNull(void* ptr);
int QSqlField_IsReadOnly(void* ptr);
int QSqlField_IsValid(void* ptr);
int QSqlField_Length(void* ptr);
char* QSqlField_Name(void* ptr);
int QSqlField_Precision(void* ptr);
int QSqlField_RequiredStatus(void* ptr);
void QSqlField_SetAutoValue(void* ptr, int autoVal);
void QSqlField_SetDefaultValue(void* ptr, void* value);
void QSqlField_SetGenerated(void* ptr, int gen);
void QSqlField_SetLength(void* ptr, int fieldLength);
void QSqlField_SetName(void* ptr, char* name);
void QSqlField_SetPrecision(void* ptr, int precision);
void QSqlField_SetReadOnly(void* ptr, int readOnly);
void QSqlField_SetRequired(void* ptr, int required);
void QSqlField_SetRequiredStatus(void* ptr, int required);
void QSqlField_SetValue(void* ptr, void* value);
void* QSqlField_Value(void* ptr);
void QSqlField_DestroyQSqlField(void* ptr);
void* QSqlIndex_NewQSqlIndex2(void* other);
void* QSqlIndex_NewQSqlIndex(char* cursorname, char* name);
void QSqlIndex_Append(void* ptr, void* field);
void QSqlIndex_Append2(void* ptr, void* field, int desc);
char* QSqlIndex_CursorName(void* ptr);
int QSqlIndex_IsDescending(void* ptr, int i);
char* QSqlIndex_Name(void* ptr);
void QSqlIndex_SetCursorName(void* ptr, char* cursorName);
void QSqlIndex_SetDescending(void* ptr, int i, int desc);
void QSqlIndex_SetName(void* ptr, char* name);
void QSqlIndex_DestroyQSqlIndex(void* ptr);
void* QSqlQuery_NewQSqlQuery3(void* db);
void* QSqlQuery_NewQSqlQuery(void* result);
void* QSqlQuery_NewQSqlQuery4(void* other);
void* QSqlQuery_NewQSqlQuery2(char* query, void* db);
int QSqlQuery_At(void* ptr);
void* QSqlQuery_BoundValue(void* ptr, char* placeholder);
void* QSqlQuery_BoundValue2(void* ptr, int pos);
void QSqlQuery_Clear(void* ptr);
void* QSqlQuery_Driver(void* ptr);
int QSqlQuery_Exec2(void* ptr);
int QSqlQuery_Exec(void* ptr, char* query);
int QSqlQuery_ExecBatch(void* ptr, int mode);
char* QSqlQuery_ExecutedQuery(void* ptr);
void QSqlQuery_Finish(void* ptr);
int QSqlQuery_First(void* ptr);
int QSqlQuery_IsActive(void* ptr);
int QSqlQuery_IsForwardOnly(void* ptr);
int QSqlQuery_IsNull2(void* ptr, char* name);
int QSqlQuery_IsNull(void* ptr, int field);
int QSqlQuery_IsSelect(void* ptr);
int QSqlQuery_IsValid(void* ptr);
int QSqlQuery_Last(void* ptr);
void* QSqlQuery_LastInsertId(void* ptr);
char* QSqlQuery_LastQuery(void* ptr);
int QSqlQuery_Next(void* ptr);
int QSqlQuery_NextResult(void* ptr);
int QSqlQuery_NumRowsAffected(void* ptr);
int QSqlQuery_Prepare(void* ptr, char* query);
int QSqlQuery_Previous(void* ptr);
void* QSqlQuery_Result(void* ptr);
int QSqlQuery_Seek(void* ptr, int index, int relative);
void QSqlQuery_SetForwardOnly(void* ptr, int forward);
int QSqlQuery_Size(void* ptr);
void* QSqlQuery_Value2(void* ptr, char* name);
void* QSqlQuery_Value(void* ptr, int index);
void QSqlQuery_DestroyQSqlQuery(void* ptr);
int QSqlQueryModel_RowCount(void* ptr, void* parent);
void* QSqlQueryModel_Data(void* ptr, void* item, int role);
int QSqlQueryModel_CanFetchMore(void* ptr, void* parent);
void QSqlQueryModel_Clear(void* ptr);
int QSqlQueryModel_ColumnCount(void* ptr, void* index);
void QSqlQueryModel_FetchMore(void* ptr, void* parent);
void* QSqlQueryModel_HeaderData(void* ptr, int section, int orientation, int role);
int QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent);
int QSqlQueryModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QSqlQueryModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role);
void QSqlQueryModel_SetQuery(void* ptr, void* query);
void QSqlQueryModel_SetQuery2(void* ptr, char* query, void* db);
void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr);
void* QSqlRecord_NewQSqlRecord();
void* QSqlRecord_NewQSqlRecord2(void* other);
void QSqlRecord_Append(void* ptr, void* field);
void QSqlRecord_Clear(void* ptr);
void QSqlRecord_ClearValues(void* ptr);
int QSqlRecord_Contains(void* ptr, char* name);
int QSqlRecord_Count(void* ptr);
char* QSqlRecord_FieldName(void* ptr, int index);
int QSqlRecord_IndexOf(void* ptr, char* name);
int QSqlRecord_IsEmpty(void* ptr);
int QSqlRecord_IsGenerated(void* ptr, char* name);
int QSqlRecord_IsGenerated2(void* ptr, int index);
int QSqlRecord_IsNull(void* ptr, char* name);
int QSqlRecord_IsNull2(void* ptr, int index);
void QSqlRecord_SetGenerated(void* ptr, char* name, int generated);
void QSqlRecord_SetGenerated2(void* ptr, int index, int generated);
void QSqlRecord_SetNull2(void* ptr, char* name);
void QSqlRecord_SetNull(void* ptr, int index);
void QSqlRecord_SetValue2(void* ptr, char* name, void* val);
void QSqlRecord_SetValue(void* ptr, int index, void* val);
void* QSqlRecord_Value2(void* ptr, char* name);
void* QSqlRecord_Value(void* ptr, int index);
void QSqlRecord_DestroyQSqlRecord(void* ptr);
void* QSqlRelation_NewQSqlRelation();
void* QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn);
char* QSqlRelation_DisplayColumn(void* ptr);
char* QSqlRelation_IndexColumn(void* ptr);
int QSqlRelation_IsValid(void* ptr);
char* QSqlRelation_TableName(void* ptr);
void QSqlRelationalTableModel_Clear(void* ptr);
void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role);
void* QSqlRelationalTableModel_RelationModel(void* ptr, int column);
int QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent);
void QSqlRelationalTableModel_RevertRow(void* ptr, int row);
int QSqlRelationalTableModel_Select(void* ptr);
int QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role);
void QSqlRelationalTableModel_SetJoinMode(void* ptr, int joinMode);
void QSqlRelationalTableModel_SetTable(void* ptr, char* table);
void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr);
void* QSqlResult_Handle(void* ptr);
void QSqlResult_DestroyQSqlResult(void* ptr);
char* QSqlResult_ObjectNameAbs(void* ptr);
void QSqlResult_SetObjectNameAbs(void* ptr, char* name);
void QSqlTableModel_ConnectBeforeDelete(void* ptr);
void QSqlTableModel_DisconnectBeforeDelete(void* ptr);
void QSqlTableModel_Clear(void* ptr);
void* QSqlTableModel_Data(void* ptr, void* index, int role);
int QSqlTableModel_EditStrategy(void* ptr);
int QSqlTableModel_FieldIndex(void* ptr, char* fieldName);
char* QSqlTableModel_Filter(void* ptr);
int QSqlTableModel_Flags(void* ptr, void* index);
void* QSqlTableModel_HeaderData(void* ptr, int section, int orientation, int role);
int QSqlTableModel_InsertRecord(void* ptr, int row, void* record);
int QSqlTableModel_InsertRows(void* ptr, int row, int count, void* parent);
int QSqlTableModel_IsDirty2(void* ptr);
int QSqlTableModel_IsDirty(void* ptr, void* index);
int QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent);
int QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent);
void QSqlTableModel_Revert(void* ptr);
void QSqlTableModel_RevertAll(void* ptr);
void QSqlTableModel_RevertRow(void* ptr, int row);
int QSqlTableModel_RowCount(void* ptr, void* parent);
int QSqlTableModel_Select(void* ptr);
int QSqlTableModel_SelectRow(void* ptr, int row);
int QSqlTableModel_SetData(void* ptr, void* index, void* value, int role);
void QSqlTableModel_SetEditStrategy(void* ptr, int strategy);
void QSqlTableModel_SetFilter(void* ptr, char* filter);
int QSqlTableModel_SetRecord(void* ptr, int row, void* values);
void QSqlTableModel_SetSort(void* ptr, int column, int order);
void QSqlTableModel_SetTable(void* ptr, char* tableName);
void QSqlTableModel_Sort(void* ptr, int column, int order);
int QSqlTableModel_Submit(void* ptr);
int QSqlTableModel_SubmitAll(void* ptr);
char* QSqlTableModel_TableName(void* ptr);
void QSqlTableModel_DestroyQSqlTableModel(void* ptr);

#ifdef __cplusplus
}
#endif