package entities

type Task struct {
	Id_project         int
	IdTask             int
	Task               string
	DesignatedEmployee string
	Hours              int
	HoursSpent         int
	StatusTask         string
	TaskDescription    string
}
