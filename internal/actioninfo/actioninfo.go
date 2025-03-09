package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {

		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации: %v\n", err)
			continue
		}

		fmt.Println(info)
	}
}
