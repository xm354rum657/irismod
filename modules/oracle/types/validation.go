package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/service/exported"
)

const (
	//MaxLatestHistory define the the maximum number of feed value saved
	MaxLatestHistory = 100
	//MaxAggregateFuncNameLen define the the maximum lenght of the ggregate function name
	MaxAggregateFuncNameLen = 10
	//MaxDescriptionLen define the the maximum lenght of the description
	MaxDescriptionLen = 280
)

var (
	// the feed name only accepts alphanumeric characters, _ and - /
	regexpFeedName = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9/_-]*$`)
)

// ValidateFeedName verify that the feedName is legal
func ValidateFeedName(feedName string) error {
	if !regexpFeedName.MatchString(feedName) {
		return sdkerrors.Wrap(ErrInvalidFeedName, feedName)
	}
	return nil
}

// ValidateDescription verify that the desc is legal
func ValidateDescription(desc string) error {
	if len(desc) > MaxDescriptionLen {
		return sdkerrors.Wrap(ErrInvalidDescription, desc)
	}
	return nil
}

// ValidateAggregateFunc verify that the aggregateFunc is legal
func ValidateAggregateFunc(aggregateFunc string) error {
	if len(aggregateFunc) == 0 || len(aggregateFunc) > MaxAggregateFuncNameLen {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "aggregate func must between [1, %d], got: %d", MaxAggregateFuncNameLen, len(aggregateFunc))
	}

	if _, err := GetAggregateFunc(aggregateFunc); err != nil {
		return err
	}
	return nil
}

// ValidateLatestHistory verify that the latestHistory is legal
func ValidateLatestHistory(latestHistory uint64) error {
	if latestHistory < 1 || latestHistory > MaxLatestHistory {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "latest history is invalid, should be between 1 and %d", MaxLatestHistory)
	}
	return nil
}

// ValidateCreator verify that the creator is legal
func ValidateCreator(creator string) error {
	if _, err := sdk.AccAddressFromBech32(creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator")
	}
	return nil
}

// ValidateServiceName verifies whether the  parameters are legal
func ValidateServiceName(serviceName string) error {
	return exported.ValidateServiceName(serviceName)
}

// ValidateResponseThreshold verifies whether the  parameters are legal
func ValidateResponseThreshold(responseThreshold uint32, maxCnt int) error {
	if (maxCnt != 0 && int(responseThreshold) > maxCnt) || responseThreshold < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "response threshold should be between 1 and %d", maxCnt)
	}
	return nil
}

// ValidateTimeout verifies whether the  parameters are legal
func ValidateTimeout(timeout int64, frequency uint64) error {
	if frequency < uint64(timeout) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "timeout [%d] should be no more than frequency [%d]", timeout, frequency)
	}
	return nil
}

// ValidateServiceFeeCap verifies whether the  parameters are legal
func ValidateServiceFeeCap(serviceFeeCap sdk.Coins) error {
	if !serviceFeeCap.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidServiceFeeCap, serviceFeeCap.String())
	}
	return nil
}

// Modified return whether the  parameters are modified
func Modified(target string) bool {
	return target != DoNotModify
}
