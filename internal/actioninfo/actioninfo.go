package actioninfo
import(
	"log"
	
	
//"github.com/Yandex-Practicum/tracker/internal/personaldata"
//"github.com/Yandex-Practicum/tracker/internal/spentenergy"
"github.com/Yandex-Practicum/tracker/internal/daysteps"
//"github.com/Yandex-Practicum/tracker/internal/trainings"
		
)

type DataParser interface {
	// TODO: добавить методы

	Parse()  daysteps.Parse()
	
	ActionInfo() daysteps.ActionInfo
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
for _, v := range dataset{
er:= dp.Parse(v)
	if er!=nil{
	log.Printf("Ошибка парсинга для элемента", er)
			continue
	}
str, er:= dp.ActionInfo()
if er!=nil{
	log.Printf("Ошибка парсинга для элемента", er)
			}
return str
}
}
