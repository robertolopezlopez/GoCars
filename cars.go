package awesomeProject

import "fmt"

type (
	CarBrand string

	Model struct {
		Brand       CarBrand `json:"brand"`
		ModelName   string   `json:"modelName"`
		Version     int      `json:"-"`
		ReleaseYear int
	}

	Car struct {
		CarModel           Model
		ManufacturingYear  int
		EngineSerialNumber string
	}

	Defect struct {
		CarModel      Model
		AffectedYears []int
		Code          string
	}

	RemoteService interface {
		FindDefects([]Model) ([]Defect, error)
	}
)

var (
	Bmw  = CarBrand("BMW")
	Vw   = CarBrand("VW")
	Audi = CarBrand("Audi")
)

func NewCar() Model {
	return Model{Version: 1}
}

func (m Model) WithBrand(brand CarBrand) Model {
	m.Brand = brand
	return m
}

func (m Model) WithModelName(modelName string) Model {
	m.ModelName = modelName
	return m
}

func (m Model) String() string {
	return fmt.Sprintf("this is a %s model %s, version %d", m.Brand, m.ModelName, m.Version)
}

func (m Model) LogThis() string {
	return fmt.Sprintf(`%+v`, m)
}

func (m Model) IncrementVersion() Model {
	m.Version += 1
	return m
}

func FindDefects(service RemoteService, cars []Car) ([]Defect, error) {
	if len(cars) == 0 {
		return nil, fmt.Errorf("no defects for no cars")
	}

	models := []Model{}
	for _, coche := range cars {
		models = append(models, coche.CarModel)
	}

	defects, err := service.FindDefects(models)
	if err != nil {
		return nil, err
	}

	returnDefects := []Defect{}
	for _, car := range cars {
		for _, defect := range defects {
			for _, affectedYear := range defect.AffectedYears {
				if affectedYear == car.ManufacturingYear {
					returnDefects = append(returnDefects, defect)
					break
				}
			}
		}
	}

	return returnDefects, err
}
