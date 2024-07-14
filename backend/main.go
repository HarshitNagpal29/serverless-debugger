package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HarshitNagpal29/severless-debugger/backend/handlers/debuggerhandlers"
	"github.com/joho/godotenv"

	"github.com/HarshitNagpal29/severless-debugger/backend/handlers"
	"github.com/HarshitNagpal29/severless-debugger/backend/handlers/log_handlers"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/aws"

	//"github.com/HarshitNagpal29/severless-debugger/pkg/azure"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/gcp"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// For now it is not handling any kind of azure functions

	fmt.Println("AWS_REGION", os.Getenv("AWS_REGION"))
	fmt.Println("AWS_ACCESS_KEY", os.Getenv("AWS_ACCESS_KEY"))
	fmt.Println("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	awslambdaService := aws.NewLambdaClient(os.Getenv("AWS_REGION"), os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	awslambdaHandler := handlers.NewAWSLambdaHandler(awslambdaService)
	awslogHandler := log_handlers.AWSLogHandler(awslambdaService)
	awsDebuggingHandler := debuggerhandlers.AWSDebuggerHandler(awslambdaService)

	// azureService := azure.NewAzureClient(os.Getenv("AZURE_REGION"), os.Getenv("AZURE_ACCESS_KEY"), os.Getenv("AZURE_SECRET_ACCESS_KEY"))
	// azureHandler := handlers.NewAzureFunctionHandler(azureService)

	gcpService, _ := gcp.NewGCFunctionClient(os.Getenv("Project_ID"), os.Getenv("Credentials_file"))
	gcpHandler := handlers.NewGCFunctionHandler(gcpService)
	gcplogHandler := log_handlers.GCPLogHandler(gcpService)
	gcpDebuggingHandler := debuggerhandlers.CreateGCDebuggerHandler(gcpService)

	router.GET("/aws/functions", awslambdaHandler.ListFunctions)
	router.POST("/aws/invoke/:functionName", awslambdaHandler.InvokeFunction)
	router.POST("/aws/update/:functionName", awslambdaHandler.UpdateFunctionCode)
	router.GET("/aws/logs/:functionName", awslogHandler.GetAWSLambdaLogs)
	router.POST("/aws/debugger/addBreakPoint/:functionName", awsDebuggingHandler.AddBreakPointForAWS)
	// router.GET("/azure/functions", azureHandler.ListFunctions)
	// router.POST("/azure/invoke/:functionName", azureHandler.InvokeFunction)(not handling any azure functions currently)
	// router.POST("/azure/update/:functionName", azureHandler.UpdateFunction)
	router.GET("/gcp/functions", gcpHandler.ListFunctions)
	router.POST("/gcp/invoke/:functionName", gcpHandler.InvokeFunction)
	router.POST("/gcp/update/:functionName", gcpHandler.UpdateFunctionCode)
	router.GET("/gcp/logs/:functionName", gcplogHandler.GetGCPFunctionLogs)
	router.POST("/gcp/debugger/addBreakPoint/:functionName", gcpDebuggingHandler.AddBreakPointForGCP)

	router.Run(":" + port)

}
