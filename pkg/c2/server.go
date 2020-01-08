package c2

import (
	"context"
	"fmt"

	"github.com/kcarretto/paragon/ent"
	"github.com/kcarretto/paragon/graphql/models"
)

//go:generate protoc -I=./proto -I=${GOPATH}/pkg/mod/github.com/gogo/googleapis@v1.3.0/ -I=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/ --gogoslick_out=plugins=grpc,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:. transport.proto
//go:generate mockgen -destination=mocks/io.gen.go -package=mocks io Writer,WriteCloser
//go:generate mockgen -destination=mocks/teamserver.gen.go -package=mocks github.com/kcarretto/paragon/pkg/c2 Teamserver

// Teamserver provides client methods used to interact with a teamserver.
type Teamserver interface {
	ClaimTasks(context.Context, models.ClaimTaskRequest) ([]*ent.Task, error)
}

// Server manages communication with agents. Upon check-in, the server will claim and respond with
// any tasks available for the agent.
type Server struct {
	Teamserver
}

// HandleAgent is a transport-agnostic method for handling agent communications.
func (srv Server) HandleAgent(ctx context.Context, msg AgentMessage) (*ServerMessage, error) {
	// TODO: Report agent results

	// Determine target criteria based on reported agent metadata
	target, err := resolveTarget(msg.Metadata)
	if err != nil {
		return nil, err
	}

	// Claim tasks for the agent
	tasks, err := srv.ClaimTasks(ctx, target)
	if err != nil {
		return nil, fmt.Errorf("failed to claim tasks from teamserver: %w", err)
	}

	return &ServerMessage{
		Tasks: convertTasks(tasks),
	}, nil
}

// convertTasks coerces an array of ent.Task to an array of c2.Task
func convertTasks(models []*ent.Task) (tasks []*Task) {
	for _, task := range models {
		if task == nil {
			continue
		}

		tasks = append(tasks, &Task{
			Id:      int32(task.ID),
			Content: task.Content,
		})
	}

	return tasks
}

// resolveTarget determines parameters necessary to claim tasks based on received agent metadata.
func resolveTarget(metadata *AgentMetadata) (models.ClaimTaskRequest, error) {
	// Ensure Agent provided metadata to resolve a target
	if metadata == nil {
		return models.ClaimTaskRequest{}, fmt.Errorf("agent provided no valid metadata for target resolution")
	}

	// Use Machine UUID if available
	if uuid := metadata.MachineUUID; uuid != "" {
		return models.ClaimTaskRequest{
			MachineUUID: &uuid,
		}, nil
	}

	// Otherwise combine Primary IP and Primary MAC addresses
	return models.ClaimTaskRequest{
		PrimaryIP:  &metadata.PrimaryIP,
		PrimaryMac: &metadata.PrimaryMAC,
	}, nil
}