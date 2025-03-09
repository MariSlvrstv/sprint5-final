package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid string format: %s", datastring)
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("counter conversion error: %v", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("incorrect duration format: %v", err)
	}

	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("incorrect duration format")
	}

	if ds.Steps < 0 {
		return "", fmt.Errorf("invalid steps count: %d", ds.Steps)
	}

	distance := spentenergy.Distance(ds.Steps)

	var calories float64
	var err error
	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("failed to calculate calories: %w", err)
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f.\nВы сожгли: %.2f ккал.",
		ds.Steps, distance, calories)
	return info, nil
}
