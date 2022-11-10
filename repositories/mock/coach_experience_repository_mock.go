package mock

import (
	"Mini-Project_Coaching-Clinic/models"

	"github.com/stretchr/testify/mock"
)

type MockCoachExperienceRepository struct {
	mock.Mock
}

func (_m *MockCoachExperienceRepository) FindByID(id string) (models.CoachExperience, error) {
	args := _m.Called(id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (_m *MockCoachExperienceRepository) Create(coachExperience models.CoachExperience) (models.CoachExperience, error) {
	args := _m.Called(coachExperience)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (_m *MockCoachExperienceRepository) Update(coachExperience models.CoachExperience, id string) (models.CoachExperience, error) {
	args := _m.Called(coachExperience, id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}

func (_m *MockCoachExperienceRepository) Delete(id string) (models.CoachExperience, error) {
	args := _m.Called(id)
	return args.Get(0).(models.CoachExperience), args.Error(1)
}
