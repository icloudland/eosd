package token

import "github.com/icloudland/eosd"

func init() {
	eos.RegisterAction(AN("eosio.token"), ActN("transfer"), Transfer{})
	eos.RegisterAction(AN("eosio.token"), ActN("issue"), Issue{})
	eos.RegisterAction(AN("eosio.token"), ActN("create"), Create{})
}
