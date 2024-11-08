package domain

import "github.com/Panthaweekan/EngRoomBookingAPI/internal/core/model"


type AccountService interface {
	GetByCMUITAccount(cmuitAccount string) (*model.Account, error)
	Save(account model.Account) error
	Delete(cmuitAccount string) error
}
