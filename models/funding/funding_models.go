package funding

import "github.com/amir-the-h/okex"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransId string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     okex.JsonFloat64 `json:"amt"`
		From    okex.AccountType `json:"from"`
		To      okex.AccountType `json:"to"`
	}
	Bill struct {
		BillId string           `json:"billId"`
		Ccy    string           `json:"ccy"`
		Bal    okex.JsonFloat64 `json:"bal"`
		BalChg okex.JsonFloat64 `json:"balChg"`
		Type   okex.BillType    `json:"type"`
		Ts     okex.JsonTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtId    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to"`
		Ts       okex.JsonTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxId  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   okex.JsonFloat64  `json:"amt"`
		State okex.DepositState `json:"state"`
		Ts    okex.JsonTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		WdId  okex.JsonInt64   `json:"wdId"`
		Amt   okex.JsonFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		TxId  string               `json:"txId"`
		From  string               `json:"from"`
		To    string               `json:"to"`
		Tag   string               `json:"tag,omitempty"`
		PmtId string               `json:"pmtId,omitempty"`
		Memo  string               `json:"memo,omitempty"`
		Amt   okex.JsonFloat64     `json:"amt"`
		Fee   okex.JsonInt64       `json:"fee"`
		WdId  okex.JsonInt64       `json:"wdId"`
		State okex.WithdrawalState `json:"state"`
		Ts    okex.JsonTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string           `json:"ccy"`
		Amt  okex.JsonFloat64 `json:"amt"`
		Side okex.ActionType  `json:"side"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy"`
		Amt      okex.JsonFloat64 `json:"amt"`
		Earnings okex.JsonFloat64 `json:"earnings"`
	}
)