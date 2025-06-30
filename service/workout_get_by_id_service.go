package service

import (
	"errors"
	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/repository"

	"fmt"

	"gorm.io/gorm"
)

func GetWorkoutDetailService(workoutID, userID uint) (*dto.WorkoutDetailResponse, error) {
	db := config.GetDB()
	repo := repository.NewWorkoutGetByIDRepository(db)

	workout, err := repo.GetWorkoutByID(workoutID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	fmt.Printf("WorkoutID: %d, Workout.UserID: %d, Token.UserID: %d\n", workoutID, workout.UserID, userID)

	if workout.UserID != userID {
		return nil, ErrForbidden
	}

	var exercises []dto.ExerciseItem
	for _, e := range workout.Exercises {
		exercises = append(exercises, dto.ExerciseItem{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
		})
	}

	return &dto.WorkoutDetailResponse{
		ID:          workout.ID,
		Name:        workout.Name,
		Description: workout.Description,
		Exercises:   exercises,
	}, nil
}
