package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("некорректный формат данных о тренировке")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("некорректное количество шагов")
	}
	trainingType := parts[1]
	if trainingType == "" {
		return fmt.Errorf("не указан тип тренировки")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return fmt.Errorf("некорректная длительность тренировки")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Weight <= 0 || t.Height <= 0 || t.Duration <= 0 {
		return "", fmt.Errorf("некорректные данные тренировки")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
