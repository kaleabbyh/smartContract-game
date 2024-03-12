package repositories

import (
	"cicada/web-service-gin/models"
	"errors"

	"gorm.io/gorm"
)

type ContractRequest struct {
	ContractHoF  *models.EscrowHoF
	ContractHire *models.EscrowHire
}

func GetContractByContractAddress(db *gorm.DB, ContractAddress string) (*string, error) {

	var contract models.EscrowHoF
	if err := db.First(&contract, "contract_address = ?", ContractAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contract doesn't exist on record")
		}
		return nil, err
	}

	return &contract.UserId, nil
}

func CheckUserId(GetUserIdByContractAddress *ContractRequest, chosenTable string) (*string, error) {
	var UserIdComp string

	switch chosenTable {
	case "hof":
		UserIdComp = GetUserIdByContractAddress.ContractHoF.UserId
		return &UserIdComp, nil
	case "hire":
		UserIdComp = GetUserIdByContractAddress.ContractHire.UserId
		return &UserIdComp, nil
	default:
		return nil, errors.New("unknown table selected")
	}

}

func GetUserIdByContractAddress(db *gorm.DB, ContractAddress string, chosenTable string) (*ContractRequest, error) {

	var contractHoF models.EscrowHoF
	var contractHire models.EscrowHire

	switch chosenTable {
	case "hof":
		if err := db.First(&contractHoF, "contract_address = ?", ContractAddress).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("contract doesn't exist on record")
			}
			return nil, err
		}
		contractData := ContractRequest{
			ContractHoF:  &contractHoF,
			ContractHire: &contractHire,
		}
		return &contractData, nil
	case "hire":
		if err := db.First(&contractHire, "contract_address = ?", ContractAddress).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("contract doesn't exist on record")
			}
			return nil, err
		}
		contractData := ContractRequest{
			ContractHoF:  &contractHoF,
			ContractHire: &contractHire,
		}
		return &contractData, nil
	default:
		return nil, errors.New("chosenTable doesnot exist on record")
	}

}

func CreateContract(db *gorm.DB, contract *models.EscrowHoF) (*models.EscrowHoF, error) {

	if err := db.Create(&contract).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("duplicated key")
		}
		return nil, err
	}
	return contract, nil
}
