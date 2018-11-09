package main

import (
	"testing"
	"math/big"
)

func TestRetrieveMsg(t *testing.T) {
	contract := retrieve("0x5492cd102a3a7c7a85f607699c720b822a1d443d")
	if contract == nil {
		t.Error("Contract not found")
	}
	rating, err := contract.Rating(nil)

	if err != nil {
		t.Errorf("Error while retrieving Rating %s", err)
	}
	if rating.Value.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Value don't match %s", rating.Value)
	}
}