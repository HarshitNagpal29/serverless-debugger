package gcp

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	functions "cloud.google.com/go/functions/apiv1"
	functionspb "cloud.google.com/go/functions/apiv1/functionspb"
	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

type GCFunction struct {
	client          *functions.CloudFunctionsClient
	loggingClient   *logging.Client
	firestoreClient *firestore.Client
	projectID       string
}

func NewGCFunctionClient(projectID, credentialsFile string) (*GCFunction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	opts := []option.ClientOption{
		//we are using the credentials file to authenticate the client.
		option.WithCredentialsFile(credentialsFile),
	}
	client, err := functions.NewCloudFunctionsClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	loggingClient, err := logging.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}
	firestoreClient, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}

	return &GCFunction{
		client:          client,
		loggingClient:   loggingClient,
		firestoreClient: firestoreClient,
		projectID:       projectID,
	}, nil

}

func (g *GCFunction) ListFunctions(projectID, region string) ([]*functionspb.CloudFunction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	req := &functionspb.ListFunctionsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, region),
	}
	//This line calls the ListFunctions method on the g.client object.
	it := g.client.ListFunctions(ctx, req)

	var function_s []*functionspb.CloudFunction
	for {
		function, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		function_s = append(function_s, function)
	}
	return function_s, nil
}

func (g *GCFunction) InvokeFunction(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	req := &functionspb.CallFunctionRequest{
		Name: name,
	}
	_, err := g.client.CallFunction(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (g *GCFunction) UpdateFunctionCode(name, sourceArchiveURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	req := &functionspb.UpdateFunctionRequest{
		Function: &functionspb.CloudFunction{
			Name: name,
			SourceCode: &functionspb.CloudFunction_SourceArchiveUrl{
				SourceArchiveUrl: sourceArchiveURL,
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			//This line specifies the field that we want to update.
			Paths: []string{"source.source_archive_url"}, //This indicates that only the source_archive_url field of the Cloud Function object should be updated with the provided value.
		},
	}
	_, err := g.client.UpdateFunction(ctx, req)
	return err

}

func (g *GCFunction) GetFunctionLogs(functionName, region string, startTime, endTime time.Time) ([]*logging.Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	adminClient, err := logadmin.NewClient(ctx, g.projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create logadmin client: %v", err)
	}
	defer adminClient.Close()

	filter := fmt.Sprintf(`resource.type="cloud_function" AND resource.labels.function_name="%s" AND resource.labels.region="%s" AND timestamp >= "%s" AND timestamp <= "%s"`,
		functionName, region, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))

	iter := adminClient.Entries(ctx, logadmin.Filter(filter))

	var entries []*logging.Entry

	for {
		entry, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (g *GCFunction) AddBreakPoint(functionName, fileName string, lineNumber int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, _, err := g.firestoreClient.Collection("Breakpoints").Add(ctx, map[string]interface{}{
		"FunctionName": functionName,
		"FileName":     fileName,
		"LineNumber":   lineNumber,
	})
	return err
}

func (g *GCFunction) RemoveBreakPoint(functionName, fileName string, lineNumber int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	iter := g.firestoreClient.Collection("Breakpoints").Where("FunctionName", "==", functionName).Where("FileName", "==", fileName).Where("LineNumber", "==", lineNumber).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		_, err = g.firestoreClient.Collection("Breakpoints").Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			return err
		}

	}
	return nil
}
