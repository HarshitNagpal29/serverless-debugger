package handlers

import (
	"net/http"

	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/aws"
	//"github.com/HarshitNagpal29/severless-debugger/backend/pkg/azure"
	"github.com/HarshitNagpal29/severless-debugger/backend/pkg/gcp"
	"github.com/gin-gonic/gin"
)

type AWSLambdaHandler struct {
	service1 *aws.AWSlambda
}

// type AzureFunctionHandler struct {
// 	service2 *azure.AzureClient
// }

type GCFunctionHandler struct {
	service3 *gcp.GCFunction
}

func NewAWSLambdaHandler(service *aws.AWSlambda) *AWSLambdaHandler {
	return &AWSLambdaHandler{
		service1: service,
	}
}

// func NewAzureFunctionHandler(service *azure.AzureClient) *AzureFunctionHandler {
// 	return &AzureFunctionHandler{
// 		service2: service,
// 	}
// }

func NewGCFunctionHandler(service *gcp.GCFunction) *GCFunctionHandler {
	return &GCFunctionHandler{
		service3: service,
	}
}

func (h *AWSLambdaHandler) ListFunctions(c *gin.Context) {
	functions, err := h.service1.ListFunctions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"functions": functions,
	})
}

func (h *AWSLambdaHandler) InvokeFunction(c *gin.Context) {
	functionName := c.Param("functionName")
	err := h.service1.InvokeFunction(functionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Function invoked successfully",
	})
}

func (h *AWSLambdaHandler) UpdateFunctionCode(c *gin.Context) {
	functionName := c.Param("functionName")
	zipFile := c.Param("zipFile")
	err := h.service1.UpdateFunctionCode(functionName, zipFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Function code updated successfully",
	})
}

// func (h *AzureFunctionHandler) ListFunctions(c *gin.Context) {
// 	functions, err := h.service2.ListFunctions()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"functions": functions,
// 	})
// }

// func (h *AzureFunctionHandler) InvokeFunction(c *gin.Context) {
// 	functionName := c.Param("functionName")
// 	err := h.service2.InvokeFunction(functionName)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Function invoked successfully",
// 	})
// }

// func (h *AzureFunctionHandler) UpdateFunction(c *gin.Context) {
// 	functionName := c.Param("functionName")
// 	zipFilePath := c.Param("zipFilePath")
// 	err := h.service2.CreateOrUpdate(functionName, zipFilePath)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Function code updated successfully",
// 	})
// }

func (h *GCFunctionHandler) ListFunctions(c *gin.Context) {
	projectID := c.Param("projectID")
	region := c.Param("region")

	if projectID == "" || region == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "projectID and region are required",
		})
		return
	}

	functions, err := h.service3.ListFunctions(projectID, region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"functions": functions,
	})
}

func (h *GCFunctionHandler) InvokeFunction(c *gin.Context) {
	functionName := c.Param("functionName")

	if functionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "functionName is required",
		})
		return
	}

	err := h.service3.InvokeFunction(functionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Function invoked successfully",
	})
}

func (h *GCFunctionHandler) UpdateFunctionCode(c *gin.Context) {
	functionName := c.Param("functionName")
	sourceArchiveURL := c.Param("sourceArchiveURL")

	if functionName == "" || sourceArchiveURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "functionName and sourceArchiveURL are required",
		})
		return
	}

	err := h.service3.UpdateFunctionCode(functionName, sourceArchiveURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Function code updated successfully",
	})
}
