// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_1-0]
	_ = x[Start21_1PLUS-1]
	_ = x[Start21_2-2]
	_ = x[ZonesTableForSecondaryTenants-3]
	_ = x[UseKeyEncodeForHashShardedIndexes-4]
	_ = x[DatabasePlacementPolicy-5]
	_ = x[GeneratedAsIdentity-6]
	_ = x[OnUpdateExpressions-7]
	_ = x[SpanConfigurationsTable-8]
	_ = x[BoundedStaleness-9]
	_ = x[DateAndIntervalStyle-10]
	_ = x[TenantUsageSingleConsumptionColumn-11]
	_ = x[SQLStatsTables-12]
	_ = x[SQLStatsCompactionScheduledJob-13]
	_ = x[V21_2-14]
	_ = x[Start22_1-15]
	_ = x[AvoidDrainingNames-16]
	_ = x[DrainingNamesMigration-17]
	_ = x[TraceIDDoesntImplyStructuredRecording-18]
	_ = x[AlterSystemTableStatisticsAddAvgSizeCol-19]
	_ = x[AlterSystemStmtDiagReqs-20]
	_ = x[MVCCAddSSTable-21]
	_ = x[InsertPublicSchemaNamespaceEntryOnRestore-22]
	_ = x[UnsplitRangesInAsyncGCJobs-23]
	_ = x[ValidateGrantOption-24]
	_ = x[PebbleFormatBlockPropertyCollector-25]
	_ = x[ProbeRequest-26]
	_ = x[SelectRPCsTakeTracingInfoInband-27]
	_ = x[PreSeedTenantSpanConfigs-28]
	_ = x[SeedTenantSpanConfigs-29]
	_ = x[PublicSchemasWithDescriptors-30]
}

const _Key_name = "V21_1Start21_1PLUSStart21_2ZonesTableForSecondaryTenantsUseKeyEncodeForHashShardedIndexesDatabasePlacementPolicyGeneratedAsIdentityOnUpdateExpressionsSpanConfigurationsTableBoundedStalenessDateAndIntervalStyleTenantUsageSingleConsumptionColumnSQLStatsTablesSQLStatsCompactionScheduledJobV21_2Start22_1AvoidDrainingNamesDrainingNamesMigrationTraceIDDoesntImplyStructuredRecordingAlterSystemTableStatisticsAddAvgSizeColAlterSystemStmtDiagReqsMVCCAddSSTableInsertPublicSchemaNamespaceEntryOnRestoreUnsplitRangesInAsyncGCJobsValidateGrantOptionPebbleFormatBlockPropertyCollectorProbeRequestSelectRPCsTakeTracingInfoInbandPreSeedTenantSpanConfigsSeedTenantSpanConfigsPublicSchemasWithDescriptors"

var _Key_index = [...]uint16{0, 5, 18, 27, 56, 89, 112, 131, 150, 173, 189, 209, 243, 257, 287, 292, 301, 319, 341, 378, 417, 440, 454, 495, 521, 540, 574, 586, 617, 641, 662, 690}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
