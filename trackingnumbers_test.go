package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const TrackingInput = `{
    "distributionModality": 33170,
    "quantity": 5
}`

const TackingInputPut = `{
    "senderName": "Paulo.COM ",
    "senderAddress": "Rua Remetente",
    "senderAddressNumber": "123",
    "senderAddressComplement": "comp",
    "senderZipCode": "12345678",
    "senderCityName": "Grecia",
    "senderState": "",
    "senderCountryCode": "GB",
    "senderEmail": "emaildoremetente@fsdsdf.com",
    "senderWebsite": "www.remetente.com",
    "recipientName": "Destinatario da Silva",
    "recipientDocumentType": "CPF",
    "recipientDocumentNumber": "12345678909",
    "recipientAddress": "Rua do Destinatario, Bairro do Destinatario",
    "recipientAddressNumber": "123",
    "recipientAddressComplement": "ED. SAO JOAO, APTO 900",
    "recipientCityName": "Belo Horizonte",
    "recipientState": "MG",
    "recipientZipCode": "31748320",
    "recipientEmail": "CUSTOMER@CORREIOS.COM.BR",
    "recipientPhoneNumber": "1499999999",
    "totalWeight": 200.15,
    "packagingLength": 35.5,
    "packagingWidth": 15.45,
    "packagingHeight": 50,
    "taxPaymentMethod": "DDU",
    "currency": "USD",
    "freightPaidValue": 7.54,
    "insurancePaidValue": 0,
    "items": [
        {
            "hsCode": "401120",
            "description": "Tste de Itens",
            "quantity": 1,
            "value": 14.61
        }
    ]
}`

const packageasyncPut = `{
    "packageList": [
        {
            "trackingNumber": "NX002551329BR",
            "senderName": "Remetente de Teste",
            "senderAddress": "S. SILVERSTONE WAY ",
            "senderAddressNumber": "201654",
            "senderAddressComplement": "comp ",
            "senderZipCode": "83642555887858953222",
            "senderCityName": "MERIDIAN",
            "senderState": "",
            "senderCountryCode": "US",
            "senderEmail": "sdfsdfsdfsdf@fsdsdf.com",
            "senderWebsite": "www.bodybuilding.comdfasdf",
            "recipientName": "Destinatario de Teste 1",
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
            "trackingNumber": "IX000682556BR",
            "senderName": "Remetente de Teste 2",
            "senderAddress": "S. SILVERSTONE WAY ",
            "senderAddressNumber": "201654",
            "senderAddressComplement": "comp ",
            "senderZipCode": "83642555887858953222",
            "senderCityName": "MERIDIAN",
            "senderState": "",
            "senderCountryCode": "US",
            "senderEmail": "sdfsdfsdfsdf@fsdsdf.com",
            "senderWebsite": "www.bodybuilding.comdfasdf",
            "recipientName": "Destinatario de Teste 2",
            "recipientDocumentType": "CPF",
            "recipientDocumentNumber": "61385307056",
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

func TestTrackingNumbersGet(t *testing.T) {

	packList, err := client.GetTrackingNumbers()
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestTrackingNumbersPost(t *testing.T) {

	trackingNumbers := &TrackingNumbersInput{}
	var bPack = []byte(TrackingInput)
	var err error
	err = json.Unmarshal(bPack, trackingNumbers)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PostTrackingNumbers(trackingNumbers)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestTrackingNumbersPut(t *testing.T) {

	trackingNumbers := &PackageTrack{}
	var bPack = []byte(TackingInputPut)
	var err error
	err = json.Unmarshal(bPack, trackingNumbers)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PutTrackingNumbers(trackingNumbers, "IX000478440BR")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestPackageAsyncPut(t *testing.T) {

	trackingNumbers := &PackList{}
	var bPack = []byte(packageasyncPut)
	var err error
	err = json.Unmarshal(bPack, trackingNumbers)
	if err != nil {
		spew.Dump(err)
	}

	packList, err := client.PutPackageAsync(trackingNumbers)
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestPackageAsyncGet(t *testing.T) {

	requestReturn, err := client.GetPackageAsync("854bc67f-1fef-45a8-8bc6-7f1fefa5a812")
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestPackageTrackingNumbersGet(t *testing.T) {

	packList, err := client.GetPackageTrackingNumbersId("NX000928682BR", "NX000928696BR")
	spew.Dump(packList)
	assert.NoError(t, err)
}

func TestPackageTrackingNumbersListGet(t *testing.T) {

	requestReturn, err := client.GetPackageTrackingNumbersList(time.Now().AddDate(-1, 0, 0), time.Now())
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}
