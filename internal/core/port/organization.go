package port

import "github.com/Panthaweekan/EngRoomBookingAPI/internal/core/model"

type OrganizationRepo interface {
	GetAll() ([]model.Organization, error)
	Create(organization *model.Organization) error
	Update(organization *model.Organization) error
	Delete(organizationID int) error
}
