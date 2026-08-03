package main

import (
	"context"
	"log/slog"

	"github.com/azure/azure-functions-golang-worker/sdk"
	"github.com/azure/azure-functions-golang-worker/sdk/bindings"
	"github.com/azure/azure-functions-golang-worker/worker"
)

const timerSchedule = "0 */30 * * * *"

func main() {
	worker.Start(newFunctionApp())
}

func newFunctionApp() *sdk.App {
	app := sdk.FunctionApp()
	app.Timer(
		"scheduledTask",
		timerHandler,
		sdk.WithSchedule(timerSchedule),
	)
	return app
}

func timerHandler(ctx context.Context, timer bindings.TimerInfo) error {
	ic, _ := sdk.FromContext(ctx)

	attrs := append([]any{
		"is_past_due", timer.IsPastDue,
		"last", timer.ScheduleStatus.Last,
		"next", timer.ScheduleStatus.Next,
	}, invocationAttrs(ic)...)

	slog.InfoContext(ctx, "timer trigger executed", attrs...)

	if ic != nil && ic.RetryContext.RetryCount > 0 {
		slog.WarnContext(ctx, "timer trigger running on retry",
			"retry_count", ic.RetryContext.RetryCount,
			"max_retry_count", ic.RetryContext.MaxRetryCount,
		)
	}

	return nil
}

func invocationAttrs(ic *sdk.InvocationContext) []any {
	if ic == nil {
		return nil
	}

	return []any{
		"invocation_id", ic.InvocationID,
		"function_id", ic.FunctionID,
		"function_name", ic.FunctionName,
		"trigger_type", ic.TriggerType,
		"trace_parent", ic.TraceContext.TraceParent,
	}
}
