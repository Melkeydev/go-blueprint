package database

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
	"strings"
	"testing"
)

const (
	port = nat.Port("19042/tcp")
)

func mustStartScyllaDBContainer() (testcontainers.Container, error) {

	// Define the container
	image := "scylladb/scylla:6.2"
	exposedPorts := []string{"9042/tcp", "19042/tcp"}
	commands := []string{
		"--smp=2",
		"--memory=1G",
		"--developer-mode=1",
		"--overprovisioned=1",
	}

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{},
		Image:          image,
		ExposedPorts:   exposedPorts,
		Cmd:            commands,
		WaitingFor: wait.ForAll(
			wait.ForLog(".*initialization completed.").AsRegexp(),
			wait.ForListeningPort(port),
			wait.ForExec([]string{"cqlsh", "-e", "SELECT bootstrapped FROM system.local"}).WithResponseMatcher(func(body io.Reader) bool {
				data, _ := io.ReadAll(body)
				return strings.Contains(string(data), "COMPLETED")
			}),
		),
	}

	// Start the container
	scyllaDBContainer, err := testcontainers.GenericContainer(
		context.Background(), testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		return nil, err
	}

	mappedPort, err := scyllaDBContainer.MappedPort(context.Background(), "19042/tcp")
	if err != nil {
		return nil, err
	}

	hosts = fmt.Sprintf("localhost:%v", mappedPort.Port())

	return scyllaDBContainer, nil
}

func TestMain(m *testing.M) {

	container, err := mustStartScyllaDBContainer()
	if err != nil {
		log.Fatalf("could not start scylla container: %v", err)
	}

	m.Run()

	err = container.Terminate(context.Background())
	if err != nil {
		return
	}

}

func TestNew(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New() returned nil")
	}

	err := srv.Close()
	if err != nil {
		t.Fatalf("expected Close() to return nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}

	if stats["scylla_cluster_nodes_up"] != "1" {
		t.Fatalf("expected nodes up  '1', got %s", stats["scylla_cluster_nodes_up"])
	}

	if stats["scylla_cluster_nodes_down"] != "0" {
		t.Fatalf("expected nodes down '0', got %s", stats["scylla_cluster_nodes_down"])
	}

	if stats["scylla_current_datacenter"] != "datacenter1" {
		t.Fatalf("expected connected dc 'datacenter', got %s", stats["scylla_current_datacenter"])
	}

	err := srv.Close()
	if err != nil {
		t.Fatalf("expected Close() to return nil")
	}
}

func TestClose(t *testing.T) {
	srv := New()

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
