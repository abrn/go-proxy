package network

// ActivePetUpdatePacket

type ActivePetUpdateCommand byte

const (
	PetFollow   ActivePetUpdateCommand = 1
	PetUnfollow ActivePetUpdateCommand = 2
	PetRelease  ActivePetUpdateCommand = 3
)

func GetPetUpdateType(command byte) string {
	switch command {
	case 1:
		return "Follow"
	case 2:
		return "Unfollow"
	case 3:
		return "Release"
	default:
		return "Unknown: " + string(command)
	}
}

// Currency - used in most purchase based packets

type CurrencyType byte

const (
	CurrencyGold          CurrencyType = 0
	CurrencyFame          CurrencyType = 1
	CurrencyGuildFame     CurrencyType = 2
	CurrencyFortuneTokens CurrencyType = 3
)

func GetCurrencyType(currency byte) string {
	switch currency {
	case 0:
		return "Gold"
	case 1:
		return "Fame"
	case 2:
		return "Guild Fame"
	case 3:
		return "Fortune Tokens"
	default:
		return "Unknown: " + string(currency)
	}
}

// BuyResultPacket

type BuyResultCode int32

const (
	BuyUnknown        BuyResultCode = -1
	BuySuccess        BuyResultCode = 0
	BuyInvalidChar    BuyResultCode = 1
	BuyItemNotFound   BuyResultCode = 2
	BuyNotEnoughGold  BuyResultCode = 3
	BuyInventoryFull  BuyResultCode = 4
	BuyRankTooLow     BuyResultCode = 5
	BuyNotEnoughFame  BuyResultCode = 6
	BuyPetFeedSuccess BuyResultCode = 7
)

func GetBuyResultType(code int32) string {
	switch code {
	case -1:
		return "Unknown: -1"
	case 0:
		return "Success"
	case 1:
		return "Invalid Character"
	case 2:
		return "Item Not Found"
	case 3:
		return "Not Enough Gold"
	case 4:
		return "Inventory Full"
	case 5:
		return "Rank Too Low"
	case 6:
		return "Not Enough Fame"
	case 7:
		return "Pet Feed Success"
	default:
		return "Unknown: " + string(code)
	}
}

// ExaltationRedeemMessagePacket

func GetExaltationResultType(code int32) string {
	switch {
	case code > 0 && code <= 24:
		return "Success"
	case code > 24 && code < 29:
		return "Failed"
	case code == 29:
		return "Not Found"
	default:
		return "Unknown: " + string(code)
	}
}

// InvResultPacket

type InvResultCode int32

const (
	// todo: forgot what these are, double check
	InvSwapFailed  InvResultCode = 0
	InvSwapSuccess InvResultCode = 1
)

func GetInvResultType(code int32) string {
	switch code {
	case 0:
		return "Failed"
	case 1:
		return "Success"
	default:
		return "Unknown: " + string(code)
	}
}

// PetYardUpdatePacket

type PetYardType int32

const (
	YardCommon    PetYardType = 1
	YardUncommon  PetYardType = 2
	YardRare      PetYardType = 3
	YardLegendary PetYardType = 4
	YardDivine    PetYardType = 5
)

func GetPetYardType(code int32) string {
	switch code {
	case 1:
		return "Common"
	case 2:
		return "Uncommon"
	case 3:
		return "Rare"
	case 4:
		return "Legendary"
	case 5:
		return "Divine"
	default:
		return "Unknown: " + string(code)
	}
}

// TradeDonePacket

type TradeDoneResult int32

const (
	TradeSuccessful TradeDoneResult = 0
	TradeCancelled  TradeDoneResult = 1
	TradeError      TradeDoneResult = 2
)

func GetTradeDoneType(code int32) string {
	switch code {
	case 0:
		return "Trade Successful"
	case 1:
		return "Trade Cancelled"
	case 2:
		return "Trade Error"
	default:
		return "Unknown: " + string(code)
	}
}
