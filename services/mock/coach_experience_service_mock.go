package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachExperienceService struct {
	mock.Mock
}

func (m *MockCoachExperienceService) FindByID(id string) (models.CoachExperience, error) {
	args := m.Called(id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (m *MockCoachExperienceService) Create(coachExperience models.CoachExperience) (models.CoachExperience, error) {
	args := m.Called(coachExperience)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (m *MockCoachExperienceService) Update(coachExperience models.CoachExperience, id string) (models.CoachExperience, error) {
	args := m.Called(coachExperience, id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (m *MockCoachExperienceService) Delete(id string) (models.CoachExperience, error) {
	args := m.Called(id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}
