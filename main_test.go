package main

import (
	"context"
	"testing"

	"github.com/azure/azure-functions-golang-worker/sdk"
	"github.com/azure/azure-functions-golang-worker/sdk/bindings"
)

func TestTimerScheduleRunsEveryThirtyMinutes(t *testing.T) {
	if timerSchedule != "0 */30 * * * *" {
		t.Fatalf("timerSchedule = %q, want %q", timerSchedule, "0 */30 * * * *")
	}
}

func TestTimerHandlerCompletes(t *testing.T) {
	if err := timerHandler(context.Background(), bindings.TimerInfo{}); err != nil {
		t.Fatalf("timerHandler() error = %v", err)
	}
}

func TestFunctionAppRegistersTimerMetadata(t *testing.T) {
	app := newFunctionApp()

	var registeredCount int
	app.GetRegisteredFunctions().Range(func(_, value any) bool {
		registeredCount++

		function := value.(*sdk.RegisteredFunction)
		if function.FuncName != "scheduledTask" {
			t.Errorf("FuncName = %q, want %q", function.FuncName, "scheduledTask")
		}
		if function.TriggerType != string(bindings.TimerTriggerType) {
			t.Errorf("TriggerType = %q, want %q", function.TriggerType, bindings.TimerTriggerType)
		}

		trigger := function.TriggerBinding()
		if trigger == nil || trigger.TimerBinding == nil {
			t.Fatal("timer trigger binding was not generated")
		}
		if trigger.TimerBinding.Schedule != timerSchedule {
			t.Errorf("Schedule = %q, want %q", trigger.TimerBinding.Schedule, timerSchedule)
		}

		return true
	})

	if registeredCount != 1 {
		t.Fatalf("registered functions = %d, want 1", registeredCount)
	}
}
