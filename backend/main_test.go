package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/HarshitNagpal29/severless-debugger/backend/handlers"
	"github.com/HarshitNagpal29/severless-debugger/backend/handlers/debuggerhandlers"
	"github.com/HarshitNagpal29/severless-debugger/backend/handlers/log_handlers"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/aws"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/gcp"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	//Mocking service Implementations
	awslambdaService := aws.NewLambdaClient("us-east-1", "access_key", "secret_access_key")
	awslambdaHandler := handlers.NewAWSLambdaHandler(awslambdaService)
	awslogHandler := log_handlers.AWSLogHandler(awslambdaService)
	awsDebuggingHandler := debuggerhandlers.AWSDebuggerHandler(awslambdaService)

	gcpService, _ := gcp.NewGCFunctionClient("project_id", "credentials_file")
	gcpHandler := handlers.NewGCFunctionHandler(gcpService)
	gcplogHandler := log_handlers.GCPLogHandler(gcpService)
	gcpDebuggingHandler := debuggerhandlers.CreateGCDebuggerHandler(gcpService)

	router.GET("/aws/functions", awslambdaHandler.ListFunctions)
	router.POST("/aws/invoke/:functionName", awslambdaHandler.InvokeFunction)
	router.POST("/aws/update/:functionName", awslambdaHandler.UpdateFunctionCode)
	router.GET("/aws/logs/:functionName", awslogHandler.GetAWSLambdaLogs)
	router.POST("/aws/debugger/addBreakPoint/:functionName", awsDebuggingHandler.AddBreakPointForAWS)

	router.GET("/gcp/functions", gcpHandler.ListFunctions)
	router.POST("/gcp/invoke/:functionName", gcpHandler.InvokeFunction)
	router.POST("/gcp/update/:functionName", gcpHandler.UpdateFunctionCode)
	router.GET("/gcp/logs/:functionName", gcplogHandler.GetGCPFunctionLogs)
	router.POST("/gcp/debugger/addBreakPoint/:functionName", gcpDebuggingHandler.AddBreakPointForGCP)
}

func TestListAWSFunctions(t *testing.T) {
	req, _ := http.NewRequest("GET", "/aws/functions", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestListGCPFunctions(t *testing.T) {
	req, _ := http.NewRequest("GET", "/gcp/functions", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestInvokeAWSFunction(t *testing.T) {
	req, _ := http.NewRequest("POST", "/aws/invoke/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestInvokeGCPFunction(t *testing.T) {
	req, _ := http.NewRequest("POST", "/gcp/invoke/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestUpdateAWSFunction(t *testing.T) {
	req, _ := http.NewRequest("POST", "/aws/update/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestUpdateGCPFunction(t *testing.T) {
	req, _ := http.NewRequest("POST", "/gcp/update/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestGetAWSFunctionLogs(t *testing.T) {
	req, _ := http.NewRequest("GET", "/aws/logs/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestGetGCPFunctionLogs(t *testing.T) {
	req, _ := http.NewRequest("GET", "/gcp/logs/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestAddBreakPointForAWS(t *testing.T) {
	req, _ := http.NewRequest("POST", "/aws/debugger/addBreakPoint/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestAddBreakPointForGCP(t *testing.T) {
	req, _ := http.NewRequest("POST", "/gcp/debugger/addBreakPoint/testFunction", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}
