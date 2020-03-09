package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/node_tooling/Celo/cmd"
)

func (v *validator) getBalance(msg tgbotapi.MessageConfig) {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS", msg)
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS", msg)
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorGr) getBalance(msg tgbotapi.MessageConfig) {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS", msg)
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS", msg)
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func UpdateBalance(p perform, msg tgbotapi.MessageConfig) {
	p.getBalance(msg)
}
