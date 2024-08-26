package serializer

import (
	"bytes"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hungtcs/clickhouse-sql-parser/grammar"
)

type SerializerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *SerializerVisitor) VisitQueryStmt(ctx *grammar.QueryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitQuery(ctx *grammar.QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCtes(ctx *grammar.CtesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitNamedQuery(ctx *grammar.NamedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnAliases(ctx *grammar.ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableStmt(ctx *grammar.AlterTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseAddColumn(ctx *grammar.AlterTableClauseAddColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseAddIndex(ctx *grammar.AlterTableClauseAddIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseAddProjection(ctx *grammar.AlterTableClauseAddProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseAttach(ctx *grammar.AlterTableClauseAttachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseClearColumn(ctx *grammar.AlterTableClauseClearColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseClearIndex(ctx *grammar.AlterTableClauseClearIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseClearProjection(ctx *grammar.AlterTableClauseClearProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseComment(ctx *grammar.AlterTableClauseCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDelete(ctx *grammar.AlterTableClauseDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDetach(ctx *grammar.AlterTableClauseDetachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDropColumn(ctx *grammar.AlterTableClauseDropColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDropIndex(ctx *grammar.AlterTableClauseDropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDropProjection(ctx *grammar.AlterTableClauseDropProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseDropPartition(ctx *grammar.AlterTableClauseDropPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseFreezePartition(ctx *grammar.AlterTableClauseFreezePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseMaterializeIndex(ctx *grammar.AlterTableClauseMaterializeIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseMaterializeProjection(ctx *grammar.AlterTableClauseMaterializeProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModifyCodec(ctx *grammar.AlterTableClauseModifyCodecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModifyComment(ctx *grammar.AlterTableClauseModifyCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModifyRemove(ctx *grammar.AlterTableClauseModifyRemoveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModify(ctx *grammar.AlterTableClauseModifyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModifyOrderBy(ctx *grammar.AlterTableClauseModifyOrderByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseModifyTTL(ctx *grammar.AlterTableClauseModifyTTLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseMovePartition(ctx *grammar.AlterTableClauseMovePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseRemoveTTL(ctx *grammar.AlterTableClauseRemoveTTLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseRename(ctx *grammar.AlterTableClauseRenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseReplace(ctx *grammar.AlterTableClauseReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlterTableClauseUpdate(ctx *grammar.AlterTableClauseUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAssignmentExprList(ctx *grammar.AssignmentExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAssignmentExpr(ctx *grammar.AssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableColumnPropertyType(ctx *grammar.TableColumnPropertyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitPartitionClause(ctx *grammar.PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAttachDictionaryStmt(ctx *grammar.AttachDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCheckStmt(ctx *grammar.CheckStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateDatabaseStmt(ctx *grammar.CreateDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateDictionaryStmt(ctx *grammar.CreateDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateLiveViewStmt(ctx *grammar.CreateLiveViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateMaterializedViewStmt(ctx *grammar.CreateMaterializedViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateTableStmt(ctx *grammar.CreateTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCreateViewStmt(ctx *grammar.CreateViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionarySchemaClause(ctx *grammar.DictionarySchemaClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionaryAttrDfnt(ctx *grammar.DictionaryAttrDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionaryEngineClause(ctx *grammar.DictionaryEngineClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionaryPrimaryKeyClause(ctx *grammar.DictionaryPrimaryKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionaryArgExpr(ctx *grammar.DictionaryArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSourceClause(ctx *grammar.SourceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLifetimeClause(ctx *grammar.LifetimeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLayoutClause(ctx *grammar.LayoutClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitRangeClause(ctx *grammar.RangeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDictionarySettingsClause(ctx *grammar.DictionarySettingsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitClusterClause(ctx *grammar.ClusterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitUuidClause(ctx *grammar.UuidClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDestinationClause(ctx *grammar.DestinationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSubqueryClause(ctx *grammar.SubqueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSchemaDescriptionClause(ctx *grammar.SchemaDescriptionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSchemaAsTableClause(ctx *grammar.SchemaAsTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSchemaAsFunctionClause(ctx *grammar.SchemaAsFunctionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitEngineClause(ctx *grammar.EngineClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitPartitionByClause(ctx *grammar.PartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitPrimaryKeyClause(ctx *grammar.PrimaryKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSampleByClause(ctx *grammar.SampleByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTtlClause(ctx *grammar.TtlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitEngineExpr(ctx *grammar.EngineExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableElementExprColumn(ctx *grammar.TableElementExprColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableElementExprConstraint(ctx *grammar.TableElementExprConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableElementExprIndex(ctx *grammar.TableElementExprIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableElementExprProjection(ctx *grammar.TableElementExprProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableColumnDfnt(ctx *grammar.TableColumnDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableColumnPropertyExpr(ctx *grammar.TableColumnPropertyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableIndexDfnt(ctx *grammar.TableIndexDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableProjectionDfnt(ctx *grammar.TableProjectionDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCodecExpr(ctx *grammar.CodecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitCodecArgExpr(ctx *grammar.CodecArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTtlExpr(ctx *grammar.TtlExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDescribeStmt(ctx *grammar.DescribeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDropDatabaseStmt(ctx *grammar.DropDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDropTableStmt(ctx *grammar.DropTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitExistsDatabaseStmt(ctx *grammar.ExistsDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitExistsTableStmt(ctx *grammar.ExistsTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitExplainASTStmt(ctx *grammar.ExplainASTStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitExplainSyntaxStmt(ctx *grammar.ExplainSyntaxStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitInsertStmt(ctx *grammar.InsertStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnsClause(ctx *grammar.ColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDataClauseFormat(ctx *grammar.DataClauseFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDataClauseValues(ctx *grammar.DataClauseValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDataClauseSelect(ctx *grammar.DataClauseSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAssignmentValues(ctx *grammar.AssignmentValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAssignmentValue(ctx *grammar.AssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitKillMutationStmt(ctx *grammar.KillMutationStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitOptimizeStmt(ctx *grammar.OptimizeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitRenameStmt(ctx *grammar.RenameStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitProjectionSelectStmt(ctx *grammar.ProjectionSelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSelectUnionStmt(ctx *grammar.SelectUnionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSelectStmtWithParens(ctx *grammar.SelectStmtWithParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSelectStmt(ctx *grammar.SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWithClause(ctx *grammar.WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTopClause(ctx *grammar.TopClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitFromClause(ctx *grammar.FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitArrayJoinClause(ctx *grammar.ArrayJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWindowClause(ctx *grammar.WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitPrewhereClause(ctx *grammar.PrewhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWhereClause(ctx *grammar.WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitGroupByClause(ctx *grammar.GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitHavingClause(ctx *grammar.HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitOrderByClause(ctx *grammar.OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitProjectionOrderByClause(ctx *grammar.ProjectionOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLimitByClause(ctx *grammar.LimitByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLimitClause(ctx *grammar.LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSettingsClause(ctx *grammar.SettingsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinExprOp(ctx *grammar.JoinExprOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinExprTable(ctx *grammar.JoinExprTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinExprParens(ctx *grammar.JoinExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinExprCrossOp(ctx *grammar.JoinExprCrossOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinOpInner(ctx *grammar.JoinOpInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinOpLeftRight(ctx *grammar.JoinOpLeftRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinOpFull(ctx *grammar.JoinOpFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinOpCross(ctx *grammar.JoinOpCrossContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitJoinConstraintClause(ctx *grammar.JoinConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSampleClause(ctx *grammar.SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLimitExpr(ctx *grammar.LimitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitOrderExprList(ctx *grammar.OrderExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitOrderExpr(ctx *grammar.OrderExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitRatioExpr(ctx *grammar.RatioExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSettingExprList(ctx *grammar.SettingExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSettingExpr(ctx *grammar.SettingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWindowExpr(ctx *grammar.WindowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWinPartitionByClause(ctx *grammar.WinPartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWinOrderByClause(ctx *grammar.WinOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWinFrameClause(ctx *grammar.WinFrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitFrameStart(ctx *grammar.FrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitFrameBetween(ctx *grammar.FrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWinFrameBound(ctx *grammar.WinFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSetStmt(ctx *grammar.SetStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowCreateDatabaseStmt(ctx *grammar.ShowCreateDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowCreateDictionaryStmt(ctx *grammar.ShowCreateDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowCreateTableStmt(ctx *grammar.ShowCreateTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowDatabasesStmt(ctx *grammar.ShowDatabasesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowDictionariesStmt(ctx *grammar.ShowDictionariesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitShowTablesStmt(ctx *grammar.ShowTablesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitSystemStmt(ctx *grammar.SystemStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTruncateStmt(ctx *grammar.TruncateStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitUseStmt(ctx *grammar.UseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitWatchStmt(ctx *grammar.WatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnTypeExprSimple(ctx *grammar.ColumnTypeExprSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnTypeExprNested(ctx *grammar.ColumnTypeExprNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnTypeExprEnum(ctx *grammar.ColumnTypeExprEnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnTypeExprComplex(ctx *grammar.ColumnTypeExprComplexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnTypeExprParam(ctx *grammar.ColumnTypeExprParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprList(ctx *grammar.ColumnExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnsExprAsterisk(ctx *grammar.ColumnsExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnsExprSubquery(ctx *grammar.ColumnsExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnsExprColumn(ctx *grammar.ColumnsExprColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprTernaryOp(ctx *grammar.ColumnExprTernaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprAlias(ctx *grammar.ColumnExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprExtract(ctx *grammar.ColumnExprExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprNegate(ctx *grammar.ColumnExprNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprSubquery(ctx *grammar.ColumnExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprLiteral(ctx *grammar.ColumnExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprArray(ctx *grammar.ColumnExprArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprSubstring(ctx *grammar.ColumnExprSubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprCast(ctx *grammar.ColumnExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprOr(ctx *grammar.ColumnExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprPrecedence1(ctx *grammar.ColumnExprPrecedence1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprPrecedence2(ctx *grammar.ColumnExprPrecedence2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprPrecedence3(ctx *grammar.ColumnExprPrecedence3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprInterval(ctx *grammar.ColumnExprIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprIsNull(ctx *grammar.ColumnExprIsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprWinFunctionTarget(ctx *grammar.ColumnExprWinFunctionTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprTrim(ctx *grammar.ColumnExprTrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprTuple(ctx *grammar.ColumnExprTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprArrayAccess(ctx *grammar.ColumnExprArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprBetween(ctx *grammar.ColumnExprBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprParens(ctx *grammar.ColumnExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprTimestamp(ctx *grammar.ColumnExprTimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprAnd(ctx *grammar.ColumnExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprTupleAccess(ctx *grammar.ColumnExprTupleAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprCase(ctx *grammar.ColumnExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprDate(ctx *grammar.ColumnExprDateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprNot(ctx *grammar.ColumnExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprWinFunction(ctx *grammar.ColumnExprWinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprIdentifier(ctx *grammar.ColumnExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprFunction(ctx *grammar.ColumnExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnExprAsterisk(ctx *grammar.ColumnExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnArgList(ctx *grammar.ColumnArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnArgExpr(ctx *grammar.ColumnArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnLambdaExpr(ctx *grammar.ColumnLambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitColumnIdentifier(ctx *grammar.ColumnIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitNestedIdentifier(ctx *grammar.NestedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableExprIdentifier(ctx *grammar.TableExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableExprSubquery(ctx *grammar.TableExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableExprAlias(ctx *grammar.TableExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableExprFunction(ctx *grammar.TableExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableFunctionExpr(ctx *grammar.TableFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableIdentifier(ctx *grammar.TableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableArgList(ctx *grammar.TableArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitTableArgExpr(ctx *grammar.TableArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitDatabaseIdentifier(ctx *grammar.DatabaseIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitFloatingLiteral(ctx *grammar.FloatingLiteralContext) interface{} {
	return ctx.GetText()
}

func (v *SerializerVisitor) VisitNumberLiteral(ctx *grammar.NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitLiteral(ctx *grammar.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitInterval(ctx *grammar.IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitKeyword(ctx *grammar.KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitKeywordForAlias(ctx *grammar.KeywordForAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitAlias(ctx *grammar.AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitIdentifier(ctx *grammar.IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitIdentifierOrNull(ctx *grammar.IdentifierOrNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitEnumValue(ctx *grammar.EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SerializerVisitor) VisitChildren(ctx antlr.RuleNode) any {
	if terminalNode, ok := ctx.(antlr.TerminalNode); ok {
		return v.VisitTerminal(terminalNode)
	}

	result := bytes.NewBuffer(nil)
	for idx, child := range ctx.GetChildren() {
		if idx > 0 && !IsDelimiterToken(child.(antlr.ParseTree)) {
			result.WriteRune(' ')
		}
		literal := child.(antlr.ParseTree).Accept(v).(string)
		result.WriteString(literal)
	}

	return result.String()
}

func (v *SerializerVisitor) VisitTerminal(ctx antlr.TerminalNode) any {
	if ctx.GetSymbol().GetTokenType() == grammar.ClickHouseParserEOF {
		return ""
	}
	return ctx.GetText()
}

func NewSerializerVisitor() *SerializerVisitor {
	return &SerializerVisitor{}
}

var (
	_ grammar.ClickHouseParserVisitor = (*SerializerVisitor)(nil)
)
