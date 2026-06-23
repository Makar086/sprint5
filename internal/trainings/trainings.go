package trainings
import(
	"time"
	"errors"
	"strconv"
	"strings"
		
)
type Training struct {
	// TODO: добавить поля
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
	
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slice := strings.Split(datastring, ",")
	if len(slice)==3{
	step, err := strconv.Atoi(slice[0])
	if err != nil {
			return err
			}
	if step<=0{
		return errors.New("ошибка 0 шагов")
		}
	t.Steps=step
	t.TrainingType=slice[1]
	time, err := time.ParseDuration(slice[1])
		if err != nil {
			
			return err
		}
	if time<=0{
		return errors.New("ошибка 0 время")
		}
	t.Duration =time
	return nil
	}
	return errors.New("ошибка")
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distan:=spentenergy.Distance(t.Steps, t.personaldata.Personal.height)
	speed:=spentenergy.MeanSpeed(t.Steps, t.personaldata.Personal.height, t.Duration)




//f weight<=0||height <=0{
//	return "", errors.New("Ошибка")
//	}
//	steps, vid, time, err:=parseTraining(data)
//if err != nil {
			
//			return  "", err
//		}
	switch t.TrainingType{
	case "Бег":
	calore, err:=spentenergy.RunningSpentCalories(t.Steps, t.personaldata.Personal.Weight, t.personaldata.Personal.Height, t.Duration)
if err != nil {
			
			
			return  "", err
		}
str:="Тип тренировки: Бег\nДлительность: "+t.Duration.String()+" ч.\nДистанция: "+strconv.Itoa(distan)+" км.\nСкорость: "+strconv.Itoa(speed)+" км/ч\nСожгли калорий: "+strconv.Itoa(calore)
return str, nil

	case "Ходьба":
	calore, err:=spentenergy.RunningSpentCalories(t.Steps, t.personaldata.Personal.Weight, t.personaldata.Personal.Height, t.Duration)
if err != nil {
			
			
			return  "", err
		}
str:="Тип тренировки: Ходьба\nДлительность: "+t.Duration.String()+" ч.\nДистанция: "+strconv.Itoa(distan)+" км.\nСкорость: "+strconv.Itoa(speed)+" км/ч\nСожгли калорий: "+strconv.Itoa(calore)
return str, nil
//	calore, err:=WalkingSpentCalories(steps, weight, height, time)
//if err != nil {
//			
//			return  "", err
//		}
//	distan:=distance(steps, height)
//	speed:=meanSpeed(steps, height, time)
	
//	minutes := time.Minutes()
//	time3:=minutes/minInH
//	time1:=strconv.FormatFloat(time3, 'f', 2, 64)
//	distan1:=strconv.FormatFloat(distan, 'f', 2, 64)
//	speed1:=strconv.FormatFloat(speed, 'f', 2, 64)
//	calore1:=strconv.FormatFloat(calore, 'f', 2, 64)
//str:= "Тип тренировки: Ходьба\nДлительность: "+time1+" ч.\nДистанция: "+distan1+" км.\nСкорость: "+speed1+" км/ч\nСожгли калорий: "+calore1+"\n"
//return str, err
	default:
return "", errors.New("неизвестный тип тренировки")
	}
}
