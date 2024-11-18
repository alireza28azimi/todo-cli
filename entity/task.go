package entity

type Task struct {
	ID         int
	Title      string
	DueData    string
	CategoryID int
	IsDone     bool
	UserID     int
}
