package docker

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io"
)

func (ctp _containerProvider) streamContainerStdOut(
	containerId string,
	writeCloser io.WriteCloser,
) error {

	readCloser, err := ctp.dockerClient.ContainerLogs(
		context.TODO(),
		containerId,
		types.ContainerLogsOptions{
			Follow:     true,
			ShowStdout: true,
		},
	)
	if nil != err {
		return err
	}

	_, err = io.Copy(writeCloser, readCloser)
	readCloser.Close()
	return err
}
