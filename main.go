package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func randBool() bool {
	return rand.Intn(2) == 1
}

type vehicle struct {
	VehicleInfo vehicleInfo `json:"vehicleInfo"`
}

type policy struct {
	PhoneNumber      string `json:"phoneNumber"`
	PolicyNumber     string `json:"policyNumber"`
	Company          string `json:"company"`
	RenewalDate      string `json:"renewalDate"`
	DateLastModified int64  `json:"dateLastModified"`
}

type service struct {
	PhoneNumber      string `json:"phoneNumber"`
	Garage           string `json:"garage"`
	DueDate          string `json:"dueDate"`
	DateLastModified int64  `json:"dateLastModified"`
}

type experianInfo struct {
	VehicleReg          string `json:"vehicleReg,omitempty"`
	Make                string `json:"make,omitempty"`
	MakeModel           string `json:"makeModel,omitempty"`
	Model               string `json:"model,omitempty"`
	Color               string `json:"color,omitempty"`
	VIN                 string `json:"vin,omitempty"`
	DateFirstRegistered string `json:"dateFirstRegistered,omitempty"`
	YearOfManufacture   int    `json:"yearOfManufacture,omitempty"`
	Transmission        string `json:"transmission,omitempty"`
	EngineNumber        string `json:"engineNumber,omitempty"`
	EngineSize          string `json:"engineSize,omitempty"`
	Fuel                string `json:"fuel,omitempty"`
}

type vehicleInfo struct {
	RegistrationNumber string       `json:"registrationNumber"`
	TaxRenewalDate     string       `json:"taxRenewalDate"`
	DateLastModified   int64        `json:"dateLastModified"`
	Breakdown          policy       `json:"breakdown,omitempty"`
	CarInsurance       policy       `json:"carInsurance,omitempty"`
	CarService         service      `json:"carService,omitempty"`
	MOT                service      `json:"mot,omitempty"`
	ExperianInfo       experianInfo `json:"experianInfo"`
}

var incompleteVehicle vehicle = vehicle{vehicleInfo{
	RegistrationNumber: "INC DATA",
	TaxRenewalDate:     "2024-10-25",
	DateLastModified:   1657939200, // Thursday, July 18, 2024 1:00:00 AM PST
	// Breakdown:          policy{PhoneNumber: "01234567890", PolicyNumber: "BRK-12345", Company: "Breakdown Rescue", RenewalDate: "2025-03-14", DateLastModified: 1652451200},   // Wednesday, May 15, 2024 1:00:00 AM PST
	// CarInsurance:       policy{PhoneNumber: "09876543210", PolicyNumber: "INS-98765", Company: "Reliable Insurance", RenewalDate: "2025-01-01", DateLastModified: 1650777600}, // Tuesday, April 23, 2024 1:00:00 AM PST
	CarService: service{PhoneNumber: "01987654321", Garage: "Speedy Fix", DueDate: "2024-09-12", DateLastModified: 1657344000}, // Sunday, July 14, 2024 1:00:00 AM PST
	MOT:        service{PhoneNumber: "", DueDate: "2025-05-20", DateLastModified: 1649097600},                                  // Monday, April 1, 2024 1:00:00 AM PST
	ExperianInfo: experianInfo{
		VehicleReg:          "INC DATA",
		Make:                "Ford",
		MakeModel:           "Fiesta",
		Model:               "Zetec",
		Color:               "Blue",
		VIN:                 "WF05XXGBB2G123456",
		DateFirstRegistered: "2018-01-30",
		// YearOfManufacture:   2017,
		// Transmission:        "Automatic",
		// EngineNumber:        "123456A",
		// EngineSize:          "1.0L",
		// Fuel:                "Petrol",
	},
}}

var vehicles = []vehicleInfo{
	{
		RegistrationNumber: "AA12 ABC",
		TaxRenewalDate:     "2024-10-25",
		DateLastModified:   1657939200,                                                                                                                                            // Thursday, July 18, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "01234567890", PolicyNumber: "BRK-12345", Company: "Breakdown Rescue", RenewalDate: "2025-03-14", DateLastModified: 1652451200},   // Wednesday, May 15, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "09876543210", PolicyNumber: "INS-98765", Company: "Reliable Insurance", RenewalDate: "2025-01-01", DateLastModified: 1650777600}, // Tuesday, April 23, 2024 1:00:00 AM PST
		CarService:         service{PhoneNumber: "01987654321", Garage: "Speedy Fix", DueDate: "2024-09-12", DateLastModified: 1657344000},                                        // Sunday, July 14, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "", DueDate: "2025-05-20", DateLastModified: 1649097600},                                                                         // Monday, April 1, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "AA12 ABC",
			Make:                "Ford",
			MakeModel:           "Fiesta",
			Model:               "Zetec",
			Color:               "Blue",
			VIN:                 "WF05XXGBB2G123456",
			DateFirstRegistered: "2018-01-30",
			YearOfManufacture:   2017,
			Transmission:        "Automatic",
			EngineNumber:        "123456A",
			EngineSize:          "1.0L",
			Fuel:                "Petrol",
		},
	}, {
		RegistrationNumber: "MN34 DEF",
		TaxRenewalDate:     "2024-12-09",
		DateLastModified:   1657939200,                                                                                                                                         // Thursday, July 18, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "02345678910", PolicyNumber: "BRK-54321", Company: "Roadside Assist", RenewalDate: "2025-07-05", DateLastModified: 1644691200}, // Sunday, February 10, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "08765432109", PolicyNumber: "INS-34567", Company: "Budget Cover", RenewalDate: "2024-11-18", DateLastModified: 1654998400},    // Monday, June 10, 2024 1:00:00 AM PST
		CarService:         service{},                                                                                                                                          // No car service info
		MOT:                service{PhoneNumber: "08654321079", DueDate: "2025-03-01", DateLastModified: 1652451200},                                                           // Wednesday, May 15, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "MN34 DEF",
			Make:                "Volkswagen",
			MakeModel:           "Golf",
			Model:               "GTI",
			Color:               "Black",
			VIN:                 "WVWZZZAUHZH123456",
			DateFirstRegistered: "2015-06-12",
			YearOfManufacture:   2015,
			Transmission:        "Manual",
			EngineNumber:        "789012B",
			EngineSize:          "2.0L",
			Fuel:                "Petrol",
		},
	}, {
		RegistrationNumber: "TY56 UIO",
		TaxRenewalDate:     "2024-09-20",
		DateLastModified:   1657939200,                                                                                                                                         // Thursday, July 18, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "04321078965", PolicyNumber: "BRK-98765", Company: "Reliable Assist", RenewalDate: "2025-06-12", DateLastModified: 1648003200}, // Tuesday, March 26, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "03210789654", PolicyNumber: "INS-67890", Company: "DriveSafe", RenewalDate: "2023-12-25", DateLastModified: 1640390400},       // Saturday, December 23, 2023 1:00:00 AM PST  Needs renewal!
		CarService:         service{PhoneNumber: "05432109876", Garage: "Midas Touch Motors", DueDate: "2025-01-22", DateLastModified: 1649718400},                             // Tuesday, April 9, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "02107896543", DueDate: "2024-12-13", DateLastModified: 1651392000},                                                           // Friday, May 3, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "TY56 UIO",
			Make:                "BMW",
			MakeModel:           "X5",
			Model:               "M Sport",
			Color:               "Black",
			VIN:                 "WBAUZ9C50J123456",
			DateFirstRegistered: "2020-08-15",
			YearOfManufacture:   2020,
			Transmission:        "Automatic",
			EngineNumber:        "345678C",
			EngineSize:          "3.0L",
			Fuel:                "Petrol",
		},
	}}

var dbVehicles = []vehicleInfo{
	{
		RegistrationNumber: "BB34 BCD",
		TaxRenewalDate:     "2024-11-12",
		DateLastModified:   1658534400,                                                                                                                                      // Friday, July 21, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "01112223344", PolicyNumber: "BRK-23456", Company: "Quick Assist", RenewalDate: "2025-04-10", DateLastModified: 1653451200}, // Wednesday, May 22, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "09988776655", PolicyNumber: "INS-87654", Company: "Secure Cover", RenewalDate: "2025-02-15", DateLastModified: 1651777600}, // Friday, April 26, 2024 1:00:00 AM PST
		CarService:         service{PhoneNumber: "01987654322", Garage: "Fast Fix", DueDate: "2024-10-18", DateLastModified: 1658344000},                                    // Monday, July 19, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "", DueDate: "2025-06-25", DateLastModified: 1649197600},                                                                   // Tuesday, April 2, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "BB34 BCD",
			Make:                "Toyota",
			MakeModel:           "Corolla",
			Model:               "Icon",
			Color:               "Red",
			VIN:                 "JT15XXGBB3G654321",
			DateFirstRegistered: "2019-02-15",
			YearOfManufacture:   2018,
			Transmission:        "Manual",
			EngineNumber:        "654321B",
			EngineSize:          "1.8L",
			Fuel:                "Hybrid",
		},
	},

	{
		RegistrationNumber: "CC56 CDE",
		TaxRenewalDate:     "2025-01-30",
		DateLastModified:   1658937600,                                                                                                                                        // Tuesday, July 25, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "02223334455", PolicyNumber: "BRK-34567", Company: "Assist Now", RenewalDate: "2025-06-20", DateLastModified: 1654451200},     // Friday, June 21, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "08877665544", PolicyNumber: "INS-76543", Company: "Safe Insurance", RenewalDate: "2025-03-20", DateLastModified: 1652777600}, // Tuesday, April 30, 2024 1:00:00 AM PST
		CarService:         service{PhoneNumber: "01987654323", Garage: "Quick Service", DueDate: "2024-11-15", DateLastModified: 1659244000},                                 // Tuesday, July 30, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "", DueDate: "2025-07-30", DateLastModified: 1649297600},                                                                     // Friday, April 5, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "CC56 CDE",
			Make:                "Honda",
			MakeModel:           "Civic",
			Model:               "SE",
			Color:               "Black",
			VIN:                 "HN18XXGBB4G789012",
			DateFirstRegistered: "2020-03-10",
			YearOfManufacture:   2019,
			Transmission:        "Automatic",
			EngineNumber:        "789012C",
			EngineSize:          "2.0L",
			Fuel:                "Petrol",
		},
	},

	{
		RegistrationNumber: "DD78 DEF",
		TaxRenewalDate:     "2025-02-25",
		DateLastModified:   1659439200,                                                                                                                                         // Friday, July 28, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "03334445566", PolicyNumber: "BRK-45678", Company: "Rescue Service", RenewalDate: "2025-07-10", DateLastModified: 1655451200},  // Sunday, June 23, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "07766554433", PolicyNumber: "INS-65432", Company: "Trust Insurance", RenewalDate: "2025-04-25", DateLastModified: 1653777600}, // Thursday, May 2, 2024 1:00:00 AM PST
		CarService:         service{PhoneNumber: "01987654324", Garage: "Best Service", DueDate: "2024-12-20", DateLastModified: 1659245000},                                   // Friday, August 2, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "", DueDate: "2025-08-15", DateLastModified: 1649397600},                                                                      // Monday, April 8, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "DD78 DEF",
			Make:                "BMW",
			MakeModel:           "3 Series",
			Model:               "Sport",
			Color:               "White",
			VIN:                 "BM22XXGBB5G012345",
			DateFirstRegistered: "2021-04-05",
			YearOfManufacture:   2020,
			Transmission:        "Automatic",
			EngineNumber:        "012345D",
			EngineSize:          "2.5L",
			Fuel:                "Diesel",
		},
	},

	{
		RegistrationNumber: "EE90 EFG",
		TaxRenewalDate:     "2025-03-15",
		DateLastModified:   1659837600,                                                                                                                                    // Tuesday, July 31, 2024 1:00:00 AM PST
		Breakdown:          policy{PhoneNumber: "04445556677", PolicyNumber: "BRK-56789", Company: "Aid Plus", RenewalDate: "2025-08-15", DateLastModified: 1656451200},   // Monday, June 25, 2024 1:00:00 AM PST
		CarInsurance:       policy{PhoneNumber: "06655443322", PolicyNumber: "INS-54321", Company: "Sure Cover", RenewalDate: "2025-05-10", DateLastModified: 1654777600}, // Friday, May 5, 2024 1:00:00 AM PST
		CarService:         service{PhoneNumber: "01987654325", Garage: "Quick Fix", DueDate: "2024-12-25", DateLastModified: 1659344000},                                 // Monday, August 5, 2024 1:00:00 AM PST
		MOT:                service{PhoneNumber: "", DueDate: "2025-09-10", DateLastModified: 1649497600},                                                                 // Wednesday, April 10, 2024 1:00:00 AM PST
		ExperianInfo: experianInfo{
			VehicleReg:          "EE90 EFG",
			Make:                "Audi",
			MakeModel:           "A4",
			Model:               "Premium",
			Color:               "Grey",
			VIN:                 "AU33XXGBB6G123456",
			DateFirstRegistered: "2022-05-10",
			YearOfManufacture:   2021,
			Transmission:        "Automatic",
			EngineNumber:        "123456E",
			EngineSize:          "2.0L",
			Fuel:                "Diesel",
		},
	},
}

type findVehicleRequest struct {
	RegistrationNumber *string `json:"registrationNumber"`
	MembershipNumber   *string `json:"membershipNumber"`
	UserId             *string `json:"userId"`
}

type vehicleDetails struct {
	RegistrationNumber string   `json:"registrationNumber"`
	Breakdown          *policy  `json:"breakdown"`
	CarInsurance       *policy  `json:"carInsurance"`
	CarService         *service `json:"carService"`
	MOT                *service `json:"mot"`
}

type updateVehicleRequest struct {
	MembershipNumber string `json:"membershipNumber"`
	UserId           string `json:"userId"`
	Vehicles         []vehicleDetails
}

type deleteVehicleRequest struct {
	MembershipNumber   string `json:"membershipNumber"`
	UserId             string `json:"userId"`
	RegistrationNumber string `json:"registrationNumber"`
}

type searchVehicleRequest struct {
	MembershipNumber   string `json:"membershipNumber"`
	UserId             string `json:"userId"`
	RegistrationNumber string `json:"registrationNumber"`
}

func getVehicles(c *gin.Context) {
	var requestBody findVehicleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if requestBody.RegistrationNumber != nil {
		var vehicle vehicle
		if randBool() {
			vehicle = getVehicleByVRN(*requestBody.RegistrationNumber)
		} else {
			c.IndentedJSON(http.StatusPartialContent, incompleteVehicle)
		}
		c.IndentedJSON(http.StatusOK, vehicle)
	}
	if requestBody.MembershipNumber != nil {
		vehicles := getVehiclesByMembershipNumber(*requestBody.MembershipNumber)
		c.IndentedJSON(http.StatusOK, vehicles)
	}
	if requestBody.UserId != nil {
		vehicles := getVehiclesByUserId(*requestBody.UserId)
		c.IndentedJSON(http.StatusOK, vehicles)
	}
	c.AbortWithStatus(http.StatusBadRequest)
	// if requestBody.RegistrationNumber != nil {
	// 	vehicle := getVehicleByVRN(*requestBody.RegistrationNumber)
	// 	c.IndentedJSON(http.StatusOK, vehicle)
	// }
}
func updateVehicle(c *gin.Context) {
	var requestBody updateVehicleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, vehicleDetails := range requestBody.Vehicles {
		for idx, currentVehicle := range vehicles {
			if vehicleDetails.RegistrationNumber == currentVehicle.ExperianInfo.VehicleReg {
				vehicles[idx] = mapDetailsToVehicleInfo(currentVehicle, vehicleDetails)
			}
		}
	}

	c.IndentedJSON(http.StatusOK, vehicles)
}

func mapDetailsToVehicleInfo(vehicleInfo vehicleInfo, vehicleDetails vehicleDetails) vehicleInfo {
	if vehicleDetails.Breakdown != nil {
		vehicleInfo.Breakdown = *vehicleDetails.Breakdown
	}
	if vehicleDetails.CarInsurance != nil {
		vehicleInfo.CarInsurance = *vehicleDetails.CarInsurance
	}
	if vehicleDetails.CarService != nil {
		vehicleInfo.CarService = *vehicleDetails.CarService
	}
	if vehicleDetails.MOT != nil {
		vehicleInfo.MOT = *vehicleDetails.MOT
	}
	return vehicleInfo
}
func getVehiclesByMembershipNumber(membershipNumber string) []vehicle {
	print(membershipNumber)
	var lookupVehicles []vehicle
	for _, vehicleInfo := range vehicles[0:1] {
		lookupVehicles = append(lookupVehicles, vehicle{vehicleInfo})
	}
	return lookupVehicles
}

func getVehiclesByUserId(userId string) []vehicle {
	print(userId)
	var lookupVehicles []vehicle
	for _, vehicleInfo := range vehicles[1:2] {
		lookupVehicles = append(lookupVehicles, vehicle{vehicleInfo})
	}
	return lookupVehicles
}

func getVehicleByVRN(vrn string) vehicle {
	var lookupVehicleInfo vehicleInfo
	for _, vehicle := range vehicles {
		if vehicle.ExperianInfo.VehicleReg == vrn {
			lookupVehicleInfo = vehicle
			break
		}
	}
	return vehicle{lookupVehicleInfo}
}

func deleteVehicle(c *gin.Context) {
	var requestBody deleteVehicleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	didFindVehicle := false
	var newVehicles []vehicleInfo
	for _, vehicleInfo := range vehicles {
		if vehicleInfo.ExperianInfo.VehicleReg == requestBody.RegistrationNumber {
			didFindVehicle = true
		}
		if vehicleInfo.ExperianInfo.VehicleReg != requestBody.RegistrationNumber {
			newVehicles = append(newVehicles, vehicleInfo)
		}
	}
	if !didFindVehicle {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	vehicles = newVehicles
	if len(vehicles) == 0 {
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.IndentedJSON(http.StatusOK, newVehicles)
}

func createVehicle(c *gin.Context) {
	var requestBody searchVehicleRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for _, vehicleInfo := range dbVehicles {
		if vehicleInfo.ExperianInfo.VehicleReg == requestBody.RegistrationNumber {
			if vehicleAlreadyExists(vehicleInfo) {
				c.IndentedJSON(http.StatusConflict, vehicles)
				return
			}
			vehicles = append(vehicles, vehicleInfo)
			c.IndentedJSON(http.StatusOK, vehicle{vehicleInfo})
			return
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func vehicleAlreadyExists(vehicle vehicleInfo) bool {
	for _, vehicleInfo := range vehicles {
		if vehicle.ExperianInfo.VehicleReg == vehicleInfo.ExperianInfo.VehicleReg {
			return true
		}
	}
	return false
}

func main() {
	router := gin.Default()
	router.GET("/sync/vehicles", getVehicles)
	router.GET("/vehicle", getVehicles)
	router.POST("/add/vehicle", createVehicle)
	router.POST("/sync/vehicle", updateVehicle)
	router.DELETE("/vehicle", deleteVehicle)

	router.Run("localhost:8080")
}
