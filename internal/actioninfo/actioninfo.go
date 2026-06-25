package actioninfo
import(
	"log"
	"fmt"
	
//"github.com/Yandex-Practicum/tracker/internal/personaldata"
//"github.com/Yandex-Practicum/tracker/internal/spentenergy"
//"github.com/Yandex-Practicum/tracker/internal/daysteps"
//"github.com/Yandex-Practicum/tracker/internal/trainings"
		
)

type DataParser interface {
	// TODO: добавить методы

	Parse(string) (error)  
	
	ActionInfo() (string, error)

}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
for q := range dataset{
er:= dp.Parse(dataset[q])
	if er !=nil{
	log.Println("Ошибка парсинга для элемента", er)
			continue
	}
str, er:= dp.ActionInfo()
if er !=nil{
	log.Println("Ошибка вывода строки", er)
	continue
			}
fmt.Println(str)
 
}
return
}
