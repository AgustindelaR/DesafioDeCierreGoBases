package tickets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets_NoConuntryEqualsCero(t *testing.T) {

	expectedResult := 0

	obtainedResult, _ := GetTotalTickets("")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestGetTotalTickets_BrazsilHas45Trips(t *testing.T) {

	expectedResult := 45

	obtainedResult, _ := GetTotalTickets("Brazil")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestGetMornings_MorngingsEquals160(t *testing.T) {

	expectedResult := 160

	obtainedResult, _ := GetMornings("Morning")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestGetMornings_AfternoonEquals289(t *testing.T) {
	expectedResult := 289

	obtainedResult, _ := GetMornings("Afternoon")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestGetMornings_NoHourEqualsCero(t *testing.T) {
	expectedResult := 0

	obtainedResult, _ := GetMornings("")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestAverageDestination_BrazilEqualsFourPointFive(t *testing.T) {
	expectedResult := 4.5

	obtainedResult, _ := AverageDestination("Brazil")

	require.Equal(t, expectedResult, obtainedResult)
}

func TestAverageDestination_NoCountryEqualsCero(t *testing.T) {
	expectedResult := 0.0

	obtainedResult, _ := AverageDestination("")

	require.Equal(t, expectedResult, obtainedResult)
}
