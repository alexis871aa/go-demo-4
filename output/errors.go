package output

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintError(value any) { // interface{}
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")
	// АЛЬТЕРНАТИВА
	//switch t := value.(type) {
	//case string:
	//	color.Red(t)
	//case int:
	//	color.Red("Код ошибки: %d", t)
	//case error:
	//	color.Red(t.Error())
	//default:
	//	color.Red("Неизвестный тип ошибки")
	//}
}

// мы не можем добавить в union types интерфейсы
func sum[T int | float32 | float64 | int16 | int32 | string](a, b T) T {
	//intValue, ok := value.(int) // так не работает

	// type switch не работает с дженериком
	//switch d := a.(type) {
	//}

	// мы не можем вернуть просто 1, так в union типе есть string, но мы можем вернуть сумму T
	return a + b
}

func example[T int | string](a, b T) T {
	// лайфхак, как обойти ограничение дженериков
	switch d := any(a).(type) {
	case string:
		fmt.Println(d)
	}
	return a + b
}

//type List[T any] struct {
//	elements []T
//}
//
//func (l *List[T]) addElement() {
//
//}
