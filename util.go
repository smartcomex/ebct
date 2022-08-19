package ebct

import (
	"github.com/satori/go.uuid"
	"time"
)

type State struct {
	CdState string
	NmState string
}

var States = [...]State{
	{CdState: "AC", NmState: "Acre"},
	{CdState: "AL", NmState: "Alagoas"},
	{CdState: "AM", NmState: "Amazonas"},
	{CdState: "AP", NmState: "Amapá"},
	{CdState: "BA", NmState: "Bahia"},
	{CdState: "CE", NmState: "Ceará"},
	{CdState: "DF", NmState: "Distrito Federal"},
	{CdState: "ES", NmState: "Espírito Santo"},
	{CdState: "GO", NmState: "Goiás"},
	{CdState: "MA", NmState: "Maranhão"},
	{CdState: "MG", NmState: "Minas Gerais"},
	{CdState: "MS", NmState: "Mato Grosso do Sul"},
	{CdState: "MT", NmState: "Mato Grosso"},
	{CdState: "PA", NmState: "Pará"},
	{CdState: "PB", NmState: "Paraíba"},
	{CdState: "PE", NmState: "Pernambuco"},
	{CdState: "PI", NmState: "Piauí"},
	{CdState: "PR", NmState: "Paraná"},
	{CdState: "RJ", NmState: "Rio de Janeiro"},
	{CdState: "RN", NmState: "Rio Grande do Norte"},
	{CdState: "RO", NmState: "Rondônia"},
	{CdState: "RR", NmState: "Roraima"},
	{CdState: "RS", NmState: "Rio Grande do Sul"},
	{CdState: "SC", NmState: "Santa Catarina"},
	{CdState: "SE", NmState: "Sergipe"},
	{CdState: "SP", NmState: "São Paulo"},
	{CdState: "TO", NmState: "Tocantin"},
}

func In(value string, list ...string) bool {

	for _, element := range list {
		if value == element {
			return true
		}
	}

	return false
}

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}

func StringPtr(value string) *string {
	return &value
}

func IntPtr(value int) *int {
	return &value
}

func FloatPtr(value float64) *float64 {
	return &value
}

func BoolPtr(value bool) *bool {
	return &value
}

func IsUUID(value string) bool {
	_, err := uuid.FromString(value)
	return err == nil
}

func TimePtr(value time.Time) *time.Time {
	return &value
}
