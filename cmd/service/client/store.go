package client

import (
	"database/sql"
	"fmt"
	"github.com/roman-haidarov/go-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetClientByIin(iin string) (*types.Client, error) {
	rows, err := s.db.Query("select * from clients where identification_no = ?", iin)
	if err != nil {
		return nil, err
	}

	c := new(types.Client)
	for rows.Next() {
		c, err = scanRowIntoClient(rows)
		if err != nil {
			return nil, err
		}
	}

	if c.ID == 0 {
		return nil, fmt.Errorf("client not found")
	}

	return c, nil
}

func scanRowIntoClient(rows *sql.Rows) (*types.Client, error) {
	client := new(types.Client)
	err := rows.Scan(clientPointers(client))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func clientPointers(client *types.Client) []interface{} {
	return []interface{}{
		&client.ID,
		&client.Name,
		&client.PrivateFirstName,
		&client.PrivateLastName,
		&client.IdentificationNo,
		&client.Phone,
		&client.Email,
		&client.ClientType,
		&client.CreatedAt,
		&client.UpdatedAt,
		&client.Sex,
		&client.PermitNotifications,
		&client.Locale,
		&client.Address,
		&client.AddressDeclared,
		&client.Phone2,
		&client.BankID,
		&client.BankAccount,
		&client.Income,
		&client.OfficialIncome,
		&client.BusinessIncome,
		&client.Debts,
		&client.DebtsAmount,
		&client.DebtsAmountSecondary,
		&client.Region,
		&client.RegionDeclared,
		&client.AgreementSigner,
		&client.Address2,
		&client.Comment,
		&client.IP,
		&client.AdditionalIncome,
		&client.HouseholdExpenses,
		&client.MonthlyLiabilities,
		&client.MonthlyLiabilitiesSecondary,
		&client.PhotoUID,
		&client.AdditionalIncomeFoundation,
		&client.SaisInformation,
		&client.DateOfBirth,
		&client.VatNo,
		&client.SodraRequestTime,
		&client.GrtRequestTime,
		&client.XroadRrRequestTime,
		&client.PhoneNotes,
		&client.Phone2Notes,
		&client.AverageIncome,
		&client.Workplace,
		&client.WorkExperience,
		&client.DriversLicenceNo,
		&client.DriverName,
		&client.DriverIdentificationNo,
		&client.DriverPhone,
		&client.DriverEmail,
		&client.DriverAddress,
		&client.MissingDriversLicenceReason,
		&client.SaisInformationFull,
		&client.RemoteClientID,
		&client.SpouseID,
		&client.SpouseStatus,
		&client.SpouseHash,
		&client.SpouseVerificationType,
		&client.SpouseVerificationManual,
		&client.SpouseStatusForDti,
		&client.WorkplaceAddress,
		&client.WorkplacePhone,
		&client.IDCardDateOfIssuance,
		&client.IDCardDateOfExpire,
		&client.IDCardIssuedBy,
		&client.SocialCardID,
		&client.SecondName,
		&client.AddressMunicipality,
		&client.AddressCity,
		&client.AddressBuilding,
		&client.AddressApartment,
		&client.AddressPostalCode,
		&client.AddressCountry,
		&client.AddressDeclaredCountry,
		&client.AddressDeclaredMunicipality,
		&client.AddressDeclaredCity,
		&client.AddressDeclaredBuilding,
		&client.AddressDeclaredApartment,
		&client.AddressDeclaredPostalCode,
		&client.APIAccessToken,
		&client.PhoneConfirmed,
		&client.EmailConfirmed,
		&client.MiscData,
		&client.UnconfirmedEmail,
		&client.UnconfirmedPhone,
		&client.Phone3,
		&client.Phone4,
		&client.Phone3Notes,
		&client.Phone4Notes,
		&client.ProtectedByGdpr,
		&client.IDCardNumber,
		&client.PassportNumber,
		&client.PassportDateOfIssuance,
		&client.PassportDateOfExpire,
		&client.PassportIssuedBy,
		&client.RubieID,
		&client.RiskLevel,
		&client.Charity,
		&client.PassportSeries,
		&client.IsCitizen,
		&client.PlaceOfBirth,
		&client.BankMfo,
		&client.PassportType,
		&client.PassportIssueRegion,
		&client.PassportIssueMunicipality,
		&client.StateRegistrationNumber,
		&client.StateRegistrationDate,
		&client.OrganizationOkpo,
		&client.OrganizationDirectorName,
		&client.OrganizationFormOfOwnership,
		&client.OrganizationIndustryTopSection,
		&client.OrganizationIndustrySection,
		&client.OrganizationIndustryGroup,
		&client.OrganizationIndustryClass,
		&client.OrganizationIndustryName,
		&client.RentObject,
		&client.Phone2Confirmed,
		&client.Phone3Confirmed,
		&client.Phone4Confirmed,
		&client.AddressDeclaredConfirmed,
		&client.PropertyConfirmed,
		&client.VehicleConfirmed,
		&client.UsesOnlinePaymentSystem,
		&client.PhoneRegistrationDate,
		&client.WorkSubdivision,
		&client.WorkPosition,
		&client.MarriageCertificate,
		&client.DateOfReceiptMarriageCertificate,
		&client.RegistryOfficeName,
		&client.SpouseIdentificationNumber,
		&client.ManualAge,
		&client.WorkplaceIdentificationNo,
		&client.WorkplaceEmail,
		&client.WorkplaceLegalAddress,
		&client.WorkplaceEmploymentDate,
		&client.AverageDebitCardOutcome,
		&client.RemainingBalance,
		&client.AverageDepositAccountFunding,
		&client.AverageWithdrawnFromDepositAccount,
		&client.MonthlyBankCardFunding,
		&client.Kdn,
		&client.PkbIncome,
		&client.TotalMonthlyPaymentFromPkb,
		&client.CreatedAtInTbm,
	}
}

func (s *Store) GetClientById(id int) (*types.Client, error) {
	return nil, nil
}

func (s *Store) CreateClient(client types.Client) error {
	return nil
}
