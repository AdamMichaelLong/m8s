package server

import (
	"bufio"
	"fmt"
	"io"

	pb "github.com/previousnext/m8s/pb"
	"github.com/previousnext/m8s/server/k8s/utils"
)

// Step is used for defining a single "command" step
func (srv Server) Step(in *pb.StepRequest, stream pb.M8S_StepServer) error {
	if in.Credentials.Token != srv.Token {
		return fmt.Errorf("token is incorrect")
	}

	if in.Name == "" {
		return fmt.Errorf("pod name was not provided")
	}

	if in.Container == "" {
		return fmt.Errorf("exec container was not provided")
	}

	if in.Command == "" {
		return fmt.Errorf("exec container was not provided")
	}

	// This is what we will use to communicate back to the CLI client.
	r, w := io.Pipe()

	go func(reader io.Reader, stream pb.M8S_StepServer) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			err := stream.Send(&pb.StepResponse{
				Message: scanner.Text(),
			})
			if err != nil {
				fmt.Println("failed to send response:", err)
			}
		}
	}(r, stream)

	err := stream.Send(&pb.StepResponse{
		Message: fmt.Sprintf("Running command: %s", in.Command),
	})
	if err != nil {
		return err
	}

	err = utils.PodExec(srv.client, srv.config, w, srv.Namespace, in.Name, in.Container, in.Command)
	if err != nil {
		return fmt.Errorf("command failed: %s", err)
	}

	return nil
}
