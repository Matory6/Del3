package Rivercrossing_test

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
    wanted := "---\\ \\_Fox_/ _________________/---]"
    state := ViewState();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}