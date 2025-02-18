package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// bank module event types
const (
	EventTypeTransfer = "transfer"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = sdk.AttributeKeySender

	// supply and balance tracking events name and attributes
	EventTypeCoinSpent    = "coin_spent"
	EventTypeCoinReceived = "coin_received"
	EventTypeCoinMint     = "coinbase" // NOTE(fdymylja): using mint clashes with mint module event
	EventTypeCoinBurn     = "burn"

	AttributeKeySpender  = "spender"
	AttributeKeyReceiver = "receiver"
	AttributeKeyMinter   = "minter"
	AttributeKeyBurner   = "burner"
)

// NewCoinSpentEvent constructs a new coin spent sdk.Event
func NewCoinSpentEvent(spender string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinSpent,
		sdk.NewAttribute(AttributeKeySpender, spender),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinReceivedEvent constructs a new coin received sdk.Event
func NewCoinReceivedEvent(receiver string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinReceived,
		sdk.NewAttribute(AttributeKeyReceiver, receiver),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinMintEvent construct a new coin minted sdk.Event
func NewCoinMintEvent(minter string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinMint,
		sdk.NewAttribute(AttributeKeyMinter, minter),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinBurnEvent constructs a new coin burned sdk.Event
func NewCoinBurnEvent(burner string, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinBurn,
		sdk.NewAttribute(AttributeKeyBurner, burner),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}
