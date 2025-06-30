package service

import (
	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/repository"
)

func GetLogsByUserService(userID uint) ([]dto.LogWithExerciseResponse, error) {
	db := config.GetDB()
	repo := repository.NewLogGetRepository(db)

	logs, err := repo.GetLogsByUserID(userID)
	if err != nil {
		return nil, ErrInternal
	}

	var response []dto.LogWithExerciseResponse
	for _, log := range logs {
		response = append(response, dto.LogWithExerciseResponse{
			ID:         log.ID,
			ExerciseID: log.ExerciseID,
			UserID:     log.UserID,
			Weight:     log.Weight,
			RepCount:   log.RepCount,
			SetCount:   log.SetCount,
			Exercise: dto.ExerciseItem{
				ID:          log.Exercise.ID,
				Name:        log.Exercise.Name,
				Description: log.Exercise.Description,
			},
		})
	}

	return response, nil
}
