/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package compose

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/compose-spec/compose-go/types"
	moby "github.com/docker/docker/api/types"
	containerType "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/errdefs"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	compose "github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/mocks"
)

func TestDown(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	api := mocks.NewMockAPIClient(mockCtrl)
	cli := mocks.NewMockCli(mockCtrl)
	tested := composeService{
		dockerCli: cli,
	}
	cli.EXPECT().Client().Return(api).AnyTimes()

	api.EXPECT().ContainerList(gomock.Any(), projectFilterListOpt(false)).Return(
		[]moby.Container{
			testContainer("service1", "123", false),
			testContainer("service2", "456", false),
			testContainer("service2", "789", false),
			testContainer("service_orphan", "321", true),
		}, nil)
	api.EXPECT().VolumeList(gomock.Any(), filters.NewArgs(projectFilter(strings.ToLower(testProject)))).
		Return(volume.VolumeListOKBody{}, nil)

	// network names are not guaranteed to be unique, ensure Compose handles
	// cleanup properly if duplicates are inadvertently created
	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{Filters: filters.NewArgs(projectFilter(strings.ToLower(testProject)))}).
		Return([]moby.NetworkResource{
			{ID: "abc123", Name: "myProject_default"},
			{ID: "def456", Name: "myProject_default"},
		}, nil)

	stopOptions := containerType.StopOptions{}
	api.EXPECT().ContainerStop(gomock.Any(), "123", stopOptions).Return(nil)
	api.EXPECT().ContainerStop(gomock.Any(), "456", stopOptions).Return(nil)
	api.EXPECT().ContainerStop(gomock.Any(), "789", stopOptions).Return(nil)

	api.EXPECT().ContainerRemove(gomock.Any(), "123", moby.ContainerRemoveOptions{Force: true}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "456", moby.ContainerRemoveOptions{Force: true}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "789", moby.ContainerRemoveOptions{Force: true}).Return(nil)

	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{
		Filters: filters.NewArgs(filters.Arg("name", "myProject_default")),
	}).Return([]moby.NetworkResource{
		{ID: "abc123", Name: "myProject_default"},
		{ID: "def456", Name: "myProject_default"},
	}, nil)
	api.EXPECT().NetworkRemove(gomock.Any(), "abc123").Return(nil)
	api.EXPECT().NetworkRemove(gomock.Any(), "def456").Return(nil)

	err := tested.Down(context.Background(), strings.ToLower(testProject), compose.DownOptions{})
	assert.NilError(t, err)
}

func TestDownRemoveOrphans(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	api := mocks.NewMockAPIClient(mockCtrl)
	cli := mocks.NewMockCli(mockCtrl)
	tested := composeService{
		dockerCli: cli,
	}
	cli.EXPECT().Client().Return(api).AnyTimes()

	api.EXPECT().ContainerList(gomock.Any(), projectFilterListOpt(true)).Return(
		[]moby.Container{
			testContainer("service1", "123", false),
			testContainer("service2", "789", false),
			testContainer("service_orphan", "321", true),
		}, nil)
	api.EXPECT().VolumeList(gomock.Any(), filters.NewArgs(projectFilter(strings.ToLower(testProject)))).
		Return(volume.VolumeListOKBody{}, nil)
	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{Filters: filters.NewArgs(projectFilter(strings.ToLower(testProject)))}).
		Return([]moby.NetworkResource{{Name: "myProject_default"}}, nil)

	stopOptions := containerType.StopOptions{}
	api.EXPECT().ContainerStop(gomock.Any(), "123", stopOptions).Return(nil)
	api.EXPECT().ContainerStop(gomock.Any(), "789", stopOptions).Return(nil)
	api.EXPECT().ContainerStop(gomock.Any(), "321", stopOptions).Return(nil)

	api.EXPECT().ContainerRemove(gomock.Any(), "123", moby.ContainerRemoveOptions{Force: true}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "789", moby.ContainerRemoveOptions{Force: true}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "321", moby.ContainerRemoveOptions{Force: true}).Return(nil)

	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{
		Filters: filters.NewArgs(filters.Arg("name", "myProject_default")),
	}).Return([]moby.NetworkResource{{ID: "abc123", Name: "myProject_default"}}, nil)
	api.EXPECT().NetworkRemove(gomock.Any(), "abc123").Return(nil)

	err := tested.Down(context.Background(), strings.ToLower(testProject), compose.DownOptions{RemoveOrphans: true})
	assert.NilError(t, err)
}

func TestDownRemoveVolumes(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	api := mocks.NewMockAPIClient(mockCtrl)
	cli := mocks.NewMockCli(mockCtrl)
	tested := composeService{
		dockerCli: cli,
	}
	cli.EXPECT().Client().Return(api).AnyTimes()

	api.EXPECT().ContainerList(gomock.Any(), projectFilterListOpt(false)).Return(
		[]moby.Container{testContainer("service1", "123", false)}, nil)
	api.EXPECT().VolumeList(gomock.Any(), filters.NewArgs(projectFilter(strings.ToLower(testProject)))).
		Return(volume.VolumeListOKBody{
			Volumes: []*moby.Volume{{Name: "myProject_volume"}},
		}, nil)
	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{Filters: filters.NewArgs(projectFilter(strings.ToLower(testProject)))}).
		Return(nil, nil)

	api.EXPECT().ContainerStop(gomock.Any(), "123", containerType.StopOptions{}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "123", moby.ContainerRemoveOptions{Force: true, RemoveVolumes: true}).Return(nil)

	api.EXPECT().VolumeRemove(gomock.Any(), "myProject_volume", true).Return(nil)

	err := tested.Down(context.Background(), strings.ToLower(testProject), compose.DownOptions{Volumes: true})
	assert.NilError(t, err)
}

func TestDownRemoveImages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	opts := compose.DownOptions{
		Project: &types.Project{
			Name: strings.ToLower(testProject),
			Services: types.Services{
				{Name: "local-anonymous"},
				{Name: "local-named", Image: "local-named-image"},
				{Name: "remote", Image: "remote-image"},
				{Name: "remote-tagged", Image: "registry.example.com/remote-image-tagged:v1.0"},
				{Name: "no-images-anonymous"},
				{Name: "no-images-named", Image: "missing-named-image"},
			},
		},
	}

	api := mocks.NewMockAPIClient(mockCtrl)
	cli := mocks.NewMockCli(mockCtrl)
	tested := composeService{
		dockerCli: cli,
	}
	cli.EXPECT().Client().Return(api).AnyTimes()

	api.EXPECT().ContainerList(gomock.Any(), projectFilterListOpt(false)).
		Return([]moby.Container{
			testContainer("service1", "123", false),
		}, nil).
		AnyTimes()

	api.EXPECT().ImageList(gomock.Any(), moby.ImageListOptions{
		Filters: filters.NewArgs(
			projectFilter(strings.ToLower(testProject)),
			filters.Arg("dangling", "false"),
		),
	}).Return([]moby.ImageSummary{
		{
			Labels:   types.Labels{compose.ServiceLabel: "local-anonymous"},
			RepoTags: []string{"testproject-local-anonymous:latest"},
		},
		{
			Labels:   types.Labels{compose.ServiceLabel: "local-named"},
			RepoTags: []string{"local-named-image:latest"},
		},
	}, nil).AnyTimes()

	imagesToBeInspected := map[string]bool{
		"testproject-local-anonymous":     true,
		"local-named-image":               true,
		"remote-image":                    true,
		"testproject-no-images-anonymous": false,
		"missing-named-image":             false,
	}
	for img, exists := range imagesToBeInspected {
		var resp moby.ImageInspect
		var err error
		if exists {
			resp.RepoTags = []string{img}
		} else {
			err = errdefs.NotFound(fmt.Errorf("test specified that image %q should not exist", img))
		}

		api.EXPECT().ImageInspectWithRaw(gomock.Any(), img).
			Return(resp, nil, err).
			AnyTimes()
	}

	api.EXPECT().ImageInspectWithRaw(gomock.Any(), "registry.example.com/remote-image-tagged:v1.0").
		Return(moby.ImageInspect{RepoTags: []string{"registry.example.com/remote-image-tagged:v1.0"}}, nil, nil).
		AnyTimes()

	localImagesToBeRemoved := []string{
		"testproject-local-anonymous:latest",
	}
	for _, img := range localImagesToBeRemoved {
		// test calls down --rmi=local then down --rmi=all, so local images
		// get "removed" 2x, while other images are only 1x
		api.EXPECT().ImageRemove(gomock.Any(), img, moby.ImageRemoveOptions{}).
			Return(nil, nil).
			Times(2)
	}

	t.Log("-> docker compose down --rmi=local")
	opts.Images = "local"
	err := tested.Down(context.Background(), strings.ToLower(testProject), opts)
	assert.NilError(t, err)

	otherImagesToBeRemoved := []string{
		"local-named-image:latest",
		"remote-image:latest",
		"registry.example.com/remote-image-tagged:v1.0",
	}
	for _, img := range otherImagesToBeRemoved {
		api.EXPECT().ImageRemove(gomock.Any(), img, moby.ImageRemoveOptions{}).
			Return(nil, nil).
			Times(1)
	}

	t.Log("-> docker compose down --rmi=all")
	opts.Images = "all"
	err = tested.Down(context.Background(), strings.ToLower(testProject), opts)
	assert.NilError(t, err)
}

func TestDownRemoveImages_NoLabel(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	api := mocks.NewMockAPIClient(mockCtrl)
	cli := mocks.NewMockCli(mockCtrl)
	tested := composeService{
		dockerCli: cli,
	}
	cli.EXPECT().Client().Return(api).AnyTimes()

	container := testContainer("service1", "123", false)

	api.EXPECT().ContainerList(gomock.Any(), projectFilterListOpt(false)).Return(
		[]moby.Container{container}, nil)

	api.EXPECT().VolumeList(gomock.Any(), filters.NewArgs(projectFilter(strings.ToLower(testProject)))).
		Return(volume.VolumeListOKBody{
			Volumes: []*moby.Volume{{Name: "myProject_volume"}},
		}, nil)
	api.EXPECT().NetworkList(gomock.Any(), moby.NetworkListOptions{Filters: filters.NewArgs(projectFilter(strings.ToLower(testProject)))}).
		Return(nil, nil)

	// ImageList returns no images for the project since they were unlabeled
	// (created by an older version of Compose)
	api.EXPECT().ImageList(gomock.Any(), moby.ImageListOptions{
		Filters: filters.NewArgs(
			projectFilter(strings.ToLower(testProject)),
			filters.Arg("dangling", "false"),
		),
	}).Return(nil, nil)

	api.EXPECT().ImageInspectWithRaw(gomock.Any(), "testproject-service1").
		Return(moby.ImageInspect{}, nil, nil)

	api.EXPECT().ContainerStop(gomock.Any(), "123", containerType.StopOptions{}).Return(nil)
	api.EXPECT().ContainerRemove(gomock.Any(), "123", moby.ContainerRemoveOptions{Force: true}).Return(nil)

	api.EXPECT().ImageRemove(gomock.Any(), "testproject-service1:latest", moby.ImageRemoveOptions{}).Return(nil, nil)

	err := tested.Down(context.Background(), strings.ToLower(testProject), compose.DownOptions{Images: "local"})
	assert.NilError(t, err)
}
