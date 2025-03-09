package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid string format: %s", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("counter conversion error: %v", err)
	}
	t.Steps = steps

	if parts[1] != "Бег" && parts[1] != "Ходьба" {
		return fmt.Errorf("invalid training: %v", err)
	}
	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("incorrect duration format: %v", err)
	}

	t.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", fmt.Errorf("incorrect duration format")
	}

	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("invalid training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("failed to calculate calories: %w", err)
	}

	info := fmt.Sprintf("Тип тренировки:%s \nДлительность: %.2f ч. \nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал.",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories)

	return info, nil

}
