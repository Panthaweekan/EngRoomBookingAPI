package dto

import "github.com/Panthaweekan/EngRoomBookingAPI/internal/core/model"


type Account struct {
	UserData    model.Account  `json:"userData"`
	StudentData *model.Student `json:"studentData"`
}
