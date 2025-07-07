package dtos

type ScheduleAppointmentDto struct {
	BarberID   string `json:"barber_id" binding:"required"`
	CustomerID string `json:"customer_id" binding:"required"`
	StartAt    string `json:"start_at" binding:"required"`
}
