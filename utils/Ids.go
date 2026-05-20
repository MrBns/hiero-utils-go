package hutils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	hiero "github.com/hiero-ledger/hiero-sdk-go/v2/sdk"
)

func IsNftId(s string) error {

	split := strings.Split(s, "@")
	if len(split) < 2 {
		return errors.New("wrong NftID format. expected {token}@{serial}")
	}

	tokenId := split[0]

	if err := IsValidTokenId(tokenId); err != nil {
		return err
	}

	serial := split[1]
	if err := IsValidNftSerial(serial); err != nil {
		return err
	}

	return nil
}

func IsValidNftSerial(s string) error {
	serial, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	if serial <= 0 {
		return fmt.Errorf("serial cannot be zero or lower than zero")
	}
	return nil

}

func IsValidTokenId(s string) error {

	if s == "" {
		return fmt.Errorf("invalid nft id")
	}

	var values []string
	if strings.Contains(s, "-") {
		values = strings.SplitN(s, "-", 2)

		if len(values) > 2 {
			return fmt.Errorf("expected {shard}.{realm}.{num}-{checksum}")
		}
	}

	values = strings.SplitN(s, ".", 3)
	if len(values) != 3 {
		// Was not three values separated by periods
		return fmt.Errorf("expected {shard}.{realm}.{num}")
	}

	return nil
}

func IsAllNftIds(values ...string) error {

	for index, v := range values {
		if err := IsNftId(v); err != nil {
			return fmt.Errorf("%v item [%v] is not a valid nft id; %v", index+1, v, err)
		}
	}
	return nil
}

func NftIdFromTokenAndSerial(tokenId string, serial int64) (hiero.NftID, error) {
	tokenID, err := hiero.TokenIDFromString(tokenId)
	if err != nil {
		return hiero.NftID{}, nil
	}

	return hiero.NftID{
		TokenID:      tokenID,
		SerialNumber: serial,
	}, nil
}
