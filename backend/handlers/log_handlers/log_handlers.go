package log_handlers

import (
	"net/http"
	"time"

	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/aws"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/gcp"
	"github.com/gin-gonic/gin"
)

type awsLogHandler struct {
	awsClient *aws.AWSlambda
}

type gcpLogHandler struct {
	gcpClient *gcp.GCFunction
}

func AWSLogHandler(awsClient *aws.AWSlambda) *awsLogHandler {
	return &awsLogHandler{
		awsClient: awsClient,
	}
}
func GCPLogHandler(gcpClient *gcp.GCFunction) *gcpLogHandler {
	return &gcpLogHandler{
		gcpClient: gcpClient,
	}
}

func (h *awsLogHandler) GetAWSLambdaLogs(c *gin.Context) {
	functionName := c.Param("functionName")

	startTimestr := c.Query("startTime")
	endTimestr := c.Query("endTime")

	startTime, err := time.Parse(time.RFC3339, startTimestr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid startTime"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, endTimestr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid endTime"})
		return
	}

	logs, err := h.awsClient.GetFunctionLogs(functionName, startTime, endTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func (h *gcpLogHandler) GetGCPFunctionLogs(c *gin.Context) {
	functionName := c.Param("functionName")
	region := c.Param("region")
	startTimeStr := c.Query("startTime")
	endTimeStr := c.Query("endTime")

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startTime"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endTime"})
		return
	}

	logs, err := h.gcpClient.GetFunctionLogs(functionName, region, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"logs": logs})

}
