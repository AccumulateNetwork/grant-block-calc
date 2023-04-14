package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/AccumulateNetwork/grant-block-calc/accumulate"
)

const API_URL = "https://mainnet.accumulatenetwork.io/v2"

// accounts
const GRANT_POOL = "acc://accumulate.acme/grant-block"
const BUSINESS_COMMITTEE = "acc://accumulate.acme/business/grants"
const GOVERNANCE_COMMITTEE = "acc://accumulate.acme/governance/grants"
const ECOSYSTEM_COMMITTEE = "acc://accumulate.acme/ecosystem/grants"
const CORE_DEV_COMMITTEE = "acc://accumulate.acme/core-dev/grants"

// shares (in bps)
const BUSINESS_SHARE = 2660
const GOVERNANCE_SHARE = 2020
const ECOSYSTEM_SHARE = 2130
const CORE_DEV_SHARE = 3190

type Output struct {
	URL          string `json:"url"`
	Amount       int64  `json:"-"`
	AmountString string `json:"amount"`
	Share        int    `json:"-"`
}

type Outputs struct {
	Items []*Output
}

type Balance struct {
	Balance int64 `json:"balance"`
}

// FromString parses balance from string
func (b *Balance) FromString(s string) {
	b.Balance, _ = strconv.ParseInt(s, 10, 64)
}

// String converts balance into human readable format
func (b *Balance) Human() string {
	hr := float64(b.Balance) * math.Pow10(-8)
	return fmt.Sprintf("%.8f", hr)
}

// String converts balance into string
func (b *Balance) String() string {
	return strconv.FormatInt(b.Balance, 10)
}

// FromBalance fills output from balance
func (o *Output) FromBalance(b *Balance) {
	amount := math.Floor(float64(b.Balance) / 10000 * float64(o.Share))
	o.Amount = int64(amount)
	o.AmountString = strconv.FormatInt(o.Amount, 10)
}

// String converts output into human readable format
func (o *Output) String() string {
	hr := float64(o.Amount) * math.Pow10(-8)
	return fmt.Sprintf("%d%% => %s : %.8f ACME", o.Share/100, o.URL, hr)
}

func main() {

	// set distribution for liquid staking rewards
	// https://docs.accumulated.finance/accumulated-finance/fees
	// share is in bps (1% = 100)
	outputs := &Outputs{}
	outputs.Items = append(outputs.Items, &Output{
		URL:   BUSINESS_COMMITTEE,
		Share: BUSINESS_SHARE,
	})
	outputs.Items = append(outputs.Items, &Output{
		URL:   GOVERNANCE_COMMITTEE,
		Share: GOVERNANCE_SHARE,
	})
	outputs.Items = append(outputs.Items, &Output{
		URL:   ECOSYSTEM_COMMITTEE,
		Share: ECOSYSTEM_SHARE,
	})
	outputs.Items = append(outputs.Items, &Output{
		URL:   CORE_DEV_COMMITTEE,
		Share: CORE_DEV_SHARE,
	})

	// validate shares
	var totalShare int
	for _, item := range outputs.Items {
		totalShare += item.Share
	}
	if totalShare != 10000 {
		log.Fatal("Expected total shares: ", 10000, ", received: ", totalShare)
	}

	client := accumulate.NewAccumulateClient(API_URL, 5)

	// grant distribution calculator
	fmt.Println("Calculating grant pool distribution...")
	fmt.Println("Getting account balance:", GRANT_POOL)

	tokenAccount, err := client.QueryTokenAccount(&accumulate.Params{URL: GRANT_POOL})
	if err != nil {
		log.Fatal(err)
	}

	balance := &Balance{}
	balance.FromString(tokenAccount.Data.Balance)

	fmt.Println("Balance:", balance.Human(), "ACME")

	for _, item := range outputs.Items {

		item.FromBalance(balance)
		fmt.Println(item)

	}

	fmt.Println("Generating CLI params...")

	jsonPayload, err := json.Marshal(outputs.Items)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("{ type: sendTokens, to: %v }", string(jsonPayload))
	fmt.Println("")

}
