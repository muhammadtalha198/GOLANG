package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegistered struct {
	ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Registered           bool               `json:"registered,omitempty" bson:"registered,omitempty"`
	Blocked              bool               `json:"blocked" bson:"blocked"`
	ReceivedAmount       uint64             `json:"receivedAmount" bson:"receivedAmount"`
	WithdrawAmount       uint64             `json:"withdrawAmount" bson:"withdrawAmount"`
	ReceiveFromTreasury  uint64             `json:"receiveFromTreasury" bson:"receiveFromTreasury"`
	ReceiveFromOwnerShip uint64             `json:"receiveFromOwnerShip" bson:"receiveFromOwnerShip"`
	TotalStakedAmount    uint64             `json:"totalStakedAmount" bson:"totalStakedAmount"`
}

type Stake struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"` // Foreign Key to UserRegistered
	Amount    string             `json:"amount" bson:"amount"`
	Restake   bool               `json:"restake" bson:"restake"`
	TimeStamp time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Withdraw struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"` // Foreign Key to UserRegistered
	Amount    string             `json:"amount" bson:"amount"`
	TimeStamp time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Contract struct {
	ID                        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TreasuryPoolAmount        uint64             `json:"treasuryPoolAmount" bson:"treasuryPoolAmount"`
	OwnerShipPoolAmount       uint64             `json:"ownerShipPoolAmount" bson:"ownerShipPoolAmount"`
	TotalStakedAmount         uint64             `json:"totalStakedAmount" bson:"totalStakedAmount"`
	TDividentPayoutPercentage uint64             `json:"tdividentPayoutPercentage" bson:"tdividentPayoutPercentage"`
	ODividentPayoutPercentage uint64             `json:"odividentPayoutPercentage" bson:"odividentPayoutPercentage"`
	FlowToTreasuryPercentage  uint64             `json:"flowToTreasuryPercentage" bson:"flowToTreasuryPercentage"`
	MaintainceFeePercentage   uint64             `json:"maintainceFeePercentage" bson:"maintainceFeePercentage"`
	OwnerRemainingPercentage  uint64             `json:"ownerRemainingPercentage" bson:"ownerRemainingPercentage"`
	NoOfUsers                 uint64             `json:"noOfUsers" bson:"noOfUsers"`
	TotalProjects             uint64             `json:"totalProjects" bson:"totalProjects"`
	Locked                    bool               `json:"locked" bson:"locked"`
	Permission                bool               `json:"permission" bson:"permission"`
	ThirdOwner                string             `json:"thirdOwner" bson:"thirdOwner"`
	RewardSender              string             `json:"rewardSender" bson:"rewardSender"`
	SecondryOwner             string             `json:"secondryOwner" bson:"secondryOwner"`
	MultisigAddress           string             `json:"multisigAddress" bson:"multisigAddress"`
	WhiteListed               map[string]bool    `json:"whiteListed" bson:"whiteListed"`
	AlreadyAdded              map[string]bool    `json:"alreadyAdded" bson:"alreadyAdded"`
	TotalUsers                map[uint64]string  `json:"totalUsers" bson:"totalUsers"`
	TPPercentages             map[uint64]uint64  `json:"tPPercentages" bson:"tPPercentages"`
}
