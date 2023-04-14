# Grant Block Calculator
An utility tool to calculate grant pool distributions to committees.

1. Fill in accounts and shares in `main.go`.
Shares are in bps (1% = 100), total shares 100% = 10000.

```go
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
```

2. Run `go run main.go` to calculate distribution and generate Accumulate CLI output.
```bash
go run main.go
Calculating grant pool distribution...
Getting account balance: acc://accumulate.acme/grant-block
Balance: 1242083.13194200 ACME
26% => acc://accumulate.acme/business/grants : 330394.11309657 ACME
20% => acc://accumulate.acme/governance/grants : 250900.79265228 ACME
21% => acc://accumulate.acme/ecosystem/grants : 264563.70710364 ACME
31% => acc://accumulate.acme/core-dev/grants : 396224.51908949 ACME
Generating CLI params...
{ type: sendTokens, to: [{"url":"acc://accumulate.acme/business/grants","amount":"33039411309657"},{"url":"acc://accumulate.acme/governance/grants","amount":"25090079265228"},{"url":"acc://accumulate.acme/ecosystem/grants","amount":"26456370710364"},{"url":"acc://accumulate.acme/core-dev/grants","amount":"39622451908949"}] }
```
