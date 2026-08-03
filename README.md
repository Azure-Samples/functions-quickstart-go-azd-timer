<!--
---
name: Azure Functions Go Timer Trigger using Azure Developer CLI
description: This repository contains an Azure Functions timer trigger quickstart written in Go and deployed to Azure Functions Flex Consumption using the Azure Developer CLI.
page_type: sample
languages:
- azdeveloper
- go
products:
- azure
- azure-functions
- entra-id
urlFragment: functions-quickstart-go-azd-timer
---
-->

# Azure Functions Go Timer Trigger using Azure Developer CLI

This template repository contains a timer trigger reference sample for Azure Functions written in Go and deployed with the Azure Developer CLI (`azd`). The sample uses managed identity and, by default, a virtual network for secure access to the function app's storage account.

The `scheduledTask` function runs automatically every 30 minutes and logs schedule, invocation, and retry details.

## Prerequisites

- [Go 1.24 or later](https://go.dev/dl/)
- [Azurite](https://learn.microsoft.com/azure/storage/common/storage-use-azurite)
- [Azure Functions Core Tools](https://learn.microsoft.com/azure/azure-functions/functions-run-local?pivots=programming-language-go#install-the-azure-functions-core-tools) 4.12.0 or later
- [Azure Developer CLI](https://learn.microsoft.com/azure/developer/azure-developer-cli/install-azd)
- To run and debug with Visual Studio Code:
  - [Visual Studio Code](https://code.visualstudio.com/)
  - [Azure Functions extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurefunctions)
  - [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Initialize the local project

Clone the repository and enter its directory:

```shell
git clone https://github.com/Azure-Samples/functions-quickstart-go-azd-timer.git
cd functions-quickstart-go-azd-timer
```

## Run locally

Copy the safe local settings template:

```powershell
Copy-Item local.settings.example.json local.settings.json
```

Start Azurite in a separate terminal:

```shell
azurite
```

Then start the Functions host:

```shell
func start
```

The host builds the Go project and invokes `scheduledTask` at minute 0 and minute 30 of each hour. The schedule uses the six-field NCrontab expression `0 */30 * * * *`.

For an immediate local invocation, temporarily change `timerSchedule` in [`main.go`](./main.go) to a more frequent development schedule. Restore the 30-minute schedule before deploying.

Press Ctrl+C to stop the Functions host.

## Source code

The timer function is registered directly in [`main.go`](./main.go):

```go
const timerSchedule = "0 */30 * * * *"

app.Timer(
    "scheduledTask",
    timerHandler,
    sdk.WithSchedule(timerSchedule),
)
```

The native Go worker indexes the app from code, so no `function.json` file is required.

## Deploy to Azure

Provision the Azure resources and deploy the function app:

```shell
azd up
```

The sample deploys to Azure Functions Flex Consumption with managed identity, Application Insights, and a virtual network by default. To deploy without the virtual network:

```shell
azd env set VNET_ENABLED false
azd up
```

## Redeploy

Run `azd up` again to provision infrastructure changes and deploy the latest code.

## Clean up resources

Delete the function app and related resources when they are no longer needed:

```shell
azd down
```
