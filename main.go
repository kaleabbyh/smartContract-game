// package main

// import "fmt"

// type AirtimeAdapter struct {
// 	Name  string
// 	Email string
// }

// type PrepaidPackageRequest struct {
// 	PhoneNumber string
// 	Amount      float64
// }

// type PrepaidPackageResponse struct {
// 	ConfirmationCode string
// 	Status           string
// }

// func (airtimeAdapter AirtimeAdapter) BuyAirtime(request PrepaidPackageRequest) (*PrepaidPackageResponse, error) {
// 	// Perform the necessary operations and return the response and error, if any

// 	// Simulate generating a confirmation code
// 	confirmationCode := "ABC123"

// 	response := &PrepaidPackageResponse{
// 		ConfirmationCode: confirmationCode,
// 		Status:           "Success",
// 	}

// 	return response, nil
// }

// func main() {
// 	// Create an instance of AirtimeAdapter
// 	adapter := AirtimeAdapter{
// 		Name:  "John Doe",
// 		Email: "johndoe@example.com",
// 	}

// 	// Create an instance of PrepaidPackageRequest
// 	request := PrepaidPackageRequest{
// 		PhoneNumber: "1234567890",
// 		Amount:      25.0,
// 	}

// 	// Call the BuyAirtime function on the AirtimeAdapter instance
// 	response, err := adapter.BuyAirtime(request)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Access the fields of the PrepaidPackageResponse struct
// 	fmt.Println("Confirmation Code:", response.ConfirmationCode)
// 	fmt.Println("Status:", response.Status)
// }

package main

import (
	"encoding/xml"
	"fmt"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
}

type Body struct {
	RMTFundTransfer RMTFundTransfer `xml:"http://temenos.com/CBEREMITANCE RMTFundtransfer"`
}

type RMTFundTransfer struct {
	WebRequestCommon WebRequestCommon
	OfsFunction      OfsFunction
	FundTransferType FUNDSTRANSFERCBEREMITANCEType `xml:"http://temenos.com/FUNDSTRANSFERCBEREMITANCE FUNDSTRANSFERCBEREMITANCEType"`
}

type WebRequestCommon struct {
	Company  string
	Password string
	UserName string
}

type OfsFunction struct {
}

type FUNDSTRANSFERCBEREMITANCEType struct {
	DEBITAMOUNT       string
	DEBITTHEIRREF     string
	CREDITTHEIRREF    string
	CREDITACCTNO      string
	CREDITCURRENCY    string
	CREDITAMOUNT      string
	CREDITVALUEDATE   string
	RemitterName      string
	BeneficiaryName   string
	BENCUST           string
	ORDCUST           string
}

func main() {
	xmlContent := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cber="http://temenos.com/CBEREMITANCE" xmlns:fun="http://temenos.com/FUNDSTRANSFERCBEREMITANCE">
	<soapenv:Header/>
	<soapenv:Body>
	   <cber:RMTFundtransfer>
		  <WebRequestCommon>
			 <company/>
			 <password>SamplePassword</password>
			 <userName>SampleUser</userName>
		  </WebRequestCommon>
		  <OfsFunction>
		   
		  </OfsFunction>
		  <fun:FUNDSTRANSFERCBEREMITANCEType id="">
			 <fun:DEBITAMOUNT>1</fun:DEBITAMOUNT>
			 <fun:DEBITTHEIRREF>123450216</fun:DEBITTHEIRREF>
			 <fun:CREDITTHEIRREF></fun:CREDITTHEIRREF>
			 <fun:CREDITACCTNO>1000263499216</fun:CREDITACCTNO>
			 <fun:CREDITCURRENCY>ETB</fun:CREDITCURRENCY>
			 <fun:CREDITAMOUNT></fun:CREDITAMOUNT>
			 <fun:CREDITVALUEDATE></fun:CREDITVALUEDATE>            
			 <fun:RemitterName>ANDUALE</fun:RemitterName>
			 <fun:BeneficiaryName>ABIY</fun:BeneficiaryName>
			 <fun:BENCUST></fun:BENCUST>
			 <fun:ORDCUST></fun:ORDCUST>
		  </fun:FUNDSTRANSFERCBEREMITANCEType>
	   </cber:RMTFundtransfer>
	</soapenv:Body>
</soapenv:Envelope>`

	var envelope Envelope
	err := xml.Unmarshal([]byte(xmlContent), &envelope)
	if err != nil {
		fmt.Println("Failed to parse XML:", err)
		return
	}

	// Access the extracted data
	fmt.Println("Password:", envelope.Body.RMTFundTransfer.WebRequestCommon.Password)
	fmt.Println("UserName:", envelope.Body.RMTFundTransfer.WebRequestCommon.UserName)
	fmt.Println("DEBITAMOUNT:", envelope.Body.RMTFundTransfer.FundTransferType.DEBITAMOUNT)
	fmt.Println("DEBITTHEIRREF:", envelope.Body.RMTFundTransfer.FundTransferType.DEBITTHEIRREF)
	// Access other fields as needed
}