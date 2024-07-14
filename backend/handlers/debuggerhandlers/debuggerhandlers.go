package debuggerhandlers

import (
	"net/http"
	"strconv"

	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/aws"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/gcp"
	"github.com/gin-gonic/gin"
)

type AWSdebuggerHandler struct {
	awsClient *aws.AWSlambda
}

type GCdebuggerHandler struct {
	gcpClient *gcp.GCFunction
}

func AWSDebuggerHandler(awsClient *aws.AWSlambda) *AWSdebuggerHandler {
	return &AWSdebuggerHandler{
		awsClient: awsClient,
	}
}

func CreateGCDebuggerHandler(gcpClient *gcp.GCFunction) *GCdebuggerHandler {
	return &GCdebuggerHandler{
		gcpClient: gcpClient,
	}
}

func (h *AWSdebuggerHandler) AddBreakPointForAWS(c *gin.Context) {
	functionName := c.Param("functionName")
	fileName := c.Query("fileName")
	lineNumber, err := strconv.Atoi(c.Query("lineNumber"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lineNumber"})
		return
	}

	err = h.awsClient.AddBreakPoint(functionName, fileName, lineNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Breakpoint added successfully"})
}

func (h *AWSdebuggerHandler) RemoveBreakPointForAWS(c *gin.Context) {
	functionName := c.Param("functionName")
	fileName := c.Query("fileName")
	lineNumber, err := strconv.Atoi(c.Query("lineNumber"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lineNumber"})
		return
	}

	err = h.awsClient.RemoveBreakPoint(functionName, fileName, lineNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Breakpoint removed successfully"})
}

func (h *GCdebuggerHandler) AddBreakPointForGCP(c *gin.Context) {
	functionName := c.Param("functionName")
	fileName := c.Query("fileName")
	lineNumber, err := strconv.Atoi(c.Query("lineNumber"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lineNumber"})
		return
	}

	err = h.gcpClient.AddBreakPoint(functionName, fileName, lineNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Breakpoint added successfully"})
}

func (h *GCdebuggerHandler) RemoveBreakPointForGCP(c *gin.Context) {
	functionName := c.Param("functionName")
	fileName := c.Query("fileName")
	lineNumber, err := strconv.Atoi(c.Query("lineNumber"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lineNumber"})
		return
	}

	err = h.gcpClient.RemoveBreakPoint(functionName, fileName, lineNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Breakpoint removed successfully"})
}
