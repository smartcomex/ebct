package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const PackageList = `{
    "packageList": [
        {
            "customerControlCode": "100052202",
            "senderName": "BODYBUILDING.COM ",
            "senderAddress": "S. SILVERSTONE WAY ",
            "senderAddressNumber": "201654",
            "senderAddressComplement": "comp ",
            "senderZipCode": "83642555887858953222",
            "senderCityName": "MERIDIAN",
            "senderState": "",
            "senderCountryCode": "US",
            "senderEmail": "sdfsdfsdfsdf@fsdsdf.com",
            "senderWebsite": "www.bodybuilding.comdfasdf",
            "recipientName": "John Park",
            "recipientDocumentType": "CPF",
            "recipientDocumentNumber": "44517000001",
            "recipientAddress": "QQAA 5 LOTE 5 BLOCO 8",
            "recipientAddressNumber": "4565456sad",
            "recipientAddressComplement": "ED. SAO JOAO, APTO 900",
            "recipientCityName": "SAO PAULO",
            "recipientState": "SP",
            "recipientZipCode": "17270001",
            "recipientEmail": "CUSTOMER@CORREIOS.COM.BR",
            "recipientPhoneNumber": "1499999999",
            "totalWeight": 120,
            "packagingLength": 16,
            "packagingWidth": 11,
            "packagingHeight": 2,
            "distributionModality": 33162,
            "taxPaymentMethod": "DDU",
            "currency": "USD",
            "freightPaidValue": 7.54,
            "insurancePaidValue": 0,
            "items": [
                {
                    "hsCode": "401120",
                    "description": "TOYS, GAMES AND SPORTS REQUISITES",
                    "quantity": 1,
                    "value": 140.61
                }
            ]
        },
        {
            "customerControlCode": "100052203",
            "senderName": "BODYBUILDING.COM ",
            "senderAddress": "S. SILVERSTONE WAY ",
            "senderAddressNumber": "201654",
            "senderAddressComplement": "comp ",
            "senderZipCode": "83642555887858953222",
            "senderCityName": "MERIDIAN",
            "senderState": "",
            "senderCountryCode": "US",
            "senderEmail": "sdfsdfsdfsdf@fsdsdf.com",
            "senderWebsite": "www.bodybuilding.comdfasdf",
            "recipientName": "Jean Albuquerque",
            "recipientDocumentType": "CPF",
            "recipientDocumentNumber": "44517000001",
            "recipientAddress": "QMSW 5 LOTE 2 BLOCO B",
            "recipientAddressNumber": "4565456sad",
            "recipientAddressComplement": "ED. SAN JUAN, APTO 600",
            "recipientCityName": "SAO PAULO",
            "recipientState": "SP",
            "recipientZipCode": "17270001",
            "recipientEmail": "JEANCS@CORREIOS.COM.BR",
            "recipientPhoneNumber": "1499999999",
            "totalWeight": 120,
            "packagingLength": 16,
            "packagingWidth": 11,
            "packagingHeight": 2,
            "distributionModality": 33170,
            "taxPaymentMethod": "DDU",
            "currency": "USD",
            "freightPaidValue": 7.54,
            "insurancePaidValue": 0,
            "items": [
                {
                    "hsCode": "401120",
                    "description": "TOYS, GAMES AND SPORTS REQUISITES",
                    "quantity": 1,
                    "value": 140.61
                }
            ]
        }
    ]
}`

func TestPackageRquest(t *testing.T) {

	aPackageList := &PackageRequestList{}
	var bPack = []byte(PackageList)
	var err error
	err = json.Unmarshal(bPack, aPackageList)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostPackage(aPackageList)
	spew.Dump(packList)
	assert.NoError(t, err)
}
