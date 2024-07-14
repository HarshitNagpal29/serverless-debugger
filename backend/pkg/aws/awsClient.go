package aws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/xray"
)

type AWSlambda struct {
	svc  *lambda.Lambda
	cwl  *cloudwatchlogs.CloudWatchLogs
	xray *xray.XRay
}

// This is a function that creates and returns a new AWSlambda instance.
func NewLambdaClient(region, accessKey, secretAccessKey string) *AWSlambda {
	//creating AWS credentials, NewStaticCredentials is used to create a new static credentials value provider.
	cred := credentials.NewStaticCredentials(accessKey, secretAccessKey, "")
	//creating a new session with the given region and credentials.Must is used to panic if the session is not created.
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: cred,
	}))
	return &AWSlambda{
		svc:  lambda.New(sess),
		cwl:  cloudwatchlogs.New(sess),
		xray: xray.New(sess),
	}
}

func (c *AWSlambda) ListFunctions() ([]*lambda.FunctionConfiguration, error) {
	result, err := c.svc.ListFunctions(nil)
	if err != nil {
		return nil, err
	}
	return result.Functions, nil
}

func (c *AWSlambda) InvokeFunction(functionName string) error {
	_, err := c.svc.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(functionName),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSlambda) UpdateFunctionCode(functionName, zipFile string) error {
	//This line updates the code of the specified Lambda function.
	// The zipFile string is converted to a byte slice, which represents the new code.
	_, err := c.svc.UpdateFunctionCode(&lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ZipFile:      []byte(zipFile),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *AWSlambda) GetFunctionLogs(functionName string, startTime, endTime time.Time) ([]*cloudwatchlogs.FilteredLogEvent, error) {
	logGroupName := "/aws/lambda/" + functionName
	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: aws.String(logGroupName),
		StartTime:    aws.Int64(startTime.UnixNano() / int64(time.Millisecond)), // Convert to milliseconds
		EndTime:      aws.Int64(endTime.UnixNano() / int64(time.Millisecond)),
	}
	resp, err := c.cwl.FilterLogEvents(params)
	if err != nil {
		return nil, err
	}
	return resp.Events, nil
}

func (c *AWSlambda) AddBreakPoint(functionName, fileName string, lineNumber int) error {
	_, err := c.xray.PutTraceSegments(&xray.PutTraceSegmentsInput{
		TraceSegmentDocuments: []*string{
			aws.String(fmt.Sprintf(`{
			"name": %s,
			"type": "subsegment",
			"start_time": %d,
			"end_time": %d,
			"annotations":{
			"BreakpointType": "Line",
			"FileName": %s,
			"LineNumber": %d
			}
		}`, functionName, time.Now().Unix(), time.Now().Unix(), fileName, lineNumber)),
		},
	})
	return err
}

func (c *AWSlambda) RemoveBreakPoint(functionName, fileName string, lineNumber int) error {
	// In a real-world scenario, you might need to query and update X-Ray traces to remove breakpoints.
	// Since X-Ray does not directly support breakpoint removal, this function serves as a placeholder.
	// Implement the necessary logic based on your application's needs.

	// Example: Remove a breakpoint annotation from X-Ray traces.
	_, err := c.xray.PutTraceSegments(&xray.PutTraceSegmentsInput{
		TraceSegmentDocuments: []*string{
			aws.String(fmt.Sprintf(`{
			"name": "%s",
			"type": "subsegment",
			"start_time": %d,
			"end_time": %d,
			"annotations":{
			"BreakpointType": "Line",
			"FileName": "%s",
			"LineNumber": %d
			}
			}`, functionName, time.Now().Unix(), time.Now().Unix(), fileName, lineNumber)),
		},
	})
	return err
}
