// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	Plugin     *plugin
	Step       *step
	StepLog    *stepLog
	StepPlugin *stepPlugin
	Task       *task
	TaskLog    *taskLog
	TaskStep   *taskStep
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Plugin = &Q.Plugin
	Step = &Q.Step
	StepLog = &Q.StepLog
	StepPlugin = &Q.StepPlugin
	Task = &Q.Task
	TaskLog = &Q.TaskLog
	TaskStep = &Q.TaskStep
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Plugin:     newPlugin(db, opts...),
		Step:       newStep(db, opts...),
		StepLog:    newStepLog(db, opts...),
		StepPlugin: newStepPlugin(db, opts...),
		Task:       newTask(db, opts...),
		TaskLog:    newTaskLog(db, opts...),
		TaskStep:   newTaskStep(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Plugin     plugin
	Step       step
	StepLog    stepLog
	StepPlugin stepPlugin
	Task       task
	TaskLog    taskLog
	TaskStep   taskStep
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Plugin:     q.Plugin.clone(db),
		Step:       q.Step.clone(db),
		StepLog:    q.StepLog.clone(db),
		StepPlugin: q.StepPlugin.clone(db),
		Task:       q.Task.clone(db),
		TaskLog:    q.TaskLog.clone(db),
		TaskStep:   q.TaskStep.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Plugin:     q.Plugin.replaceDB(db),
		Step:       q.Step.replaceDB(db),
		StepLog:    q.StepLog.replaceDB(db),
		StepPlugin: q.StepPlugin.replaceDB(db),
		Task:       q.Task.replaceDB(db),
		TaskLog:    q.TaskLog.replaceDB(db),
		TaskStep:   q.TaskStep.replaceDB(db),
	}
}

type queryCtx struct {
	Plugin     IPluginDo
	Step       IStepDo
	StepLog    IStepLogDo
	StepPlugin IStepPluginDo
	Task       ITaskDo
	TaskLog    ITaskLogDo
	TaskStep   ITaskStepDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Plugin:     q.Plugin.WithContext(ctx),
		Step:       q.Step.WithContext(ctx),
		StepLog:    q.StepLog.WithContext(ctx),
		StepPlugin: q.StepPlugin.WithContext(ctx),
		Task:       q.Task.WithContext(ctx),
		TaskLog:    q.TaskLog.WithContext(ctx),
		TaskStep:   q.TaskStep.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
