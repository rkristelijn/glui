package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGitLabClient struct {
	mock.Mock
}

func (m *MockGitLabClient) GetPipelines(projectID string) ([]Pipeline, error) {
	args := m.Called(projectID)
	return args.Get(0).([]Pipeline), args.Error(1)
}

func TestPipelineService_ListPipelines(t *testing.T) {
	tests := []struct {
		name           string
		projectID      string
		mockPipelines  []Pipeline
		mockError      error
		expectedResult []Pipeline
		expectedError  string
	}{
		{
			name:      "successful pipeline fetch",
			projectID: "123",
			mockPipelines: []Pipeline{
				{ID: 1, Status: "success", Ref: "main"},
				{ID: 2, Status: "failed", Ref: "develop"},
			},
			expectedResult: []Pipeline{
				{ID: 1, Status: "success", Ref: "main"},
				{ID: 2, Status: "failed", Ref: "develop"},
			},
		},
		{
			name:          "API error propagation",
			projectID:     "456",
			mockError:     errors.New("API error: 404 not found"),
			expectedError: "failed to fetch pipelines: API error: 404 not found",
		},
		{
			name:           "empty pipeline list",
			projectID:      "789",
			mockPipelines:  []Pipeline{},
			expectedResult: []Pipeline{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockGitLabClient)
			mockClient.On("GetPipelines", tt.projectID).Return(tt.mockPipelines, tt.mockError)

			service := NewPipelineService(mockClient)
			result, err := service.ListPipelines(tt.projectID)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			mockClient.AssertExpectations(t)
		})
	}
}

func TestPipelineService_ListPipelines_Caching(t *testing.T) {
	mockClient := new(MockGitLabClient)
	pipelines := []Pipeline{{ID: 1, Status: "success", Ref: "main"}}

	// Should only call API once due to caching
	mockClient.On("GetPipelines", "123").Return(pipelines, nil).Once()

	service := NewPipelineService(mockClient)

	// First call
	result1, err1 := service.ListPipelines("123")
	assert.NoError(t, err1)
	assert.Equal(t, pipelines, result1)

	// Second call should use cache
	result2, err2 := service.ListPipelines("123")
	assert.NoError(t, err2)
	assert.Equal(t, pipelines, result2)

	mockClient.AssertExpectations(t)
}
