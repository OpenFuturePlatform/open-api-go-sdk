# Open Platform

`opengo` is a Go library for interacting with Open Platform API

### Content
* [Installing](#installing)
* [Get started](#get-started)
* [Entity](#entity)
* [Usage](#usage)

#### Installing
```bash
go get github.com/openfutureplatform/open-api-go-sdk/opengo
```

#### Get started

`token` - your Open Platform key

```Go
op := opengo.NewOpenGo(token)
ctx := context.Background()
scaffolds, err := op.GetScaffolds(ctx)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffolds:", scaffolds)
```

#### Entities
##### Scaffold

###### Scaffold attributes

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
address         |String                                               | Scaffold address
abi             |String                                               | Scaffold json interface
description     |String                                               | Scaffold description
fiatAmount      |String                                               | Scaffold fiat amount
currency        |String                                               | Fiat amount currency
conversionAmount|String                                               | Fiat amount converted to ethereum 
developerAddress|String                                               | Scaffold developer address
webHook         |String                                               | Scaffold webhook for events
properties      |[ScaffoldProperty](#scaffold-properties-attributes)[]| Scaffold properties

###### Scaffold properties attributes

Attribute   | Type       | Description
------------|------------|-----------
name        |String      | Property name
type        |PropertyType| Property type
defaultValue|String      | Property default value

###### Scaffold summary attributes

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
scaffold        |[Scaffold](#scaffold-attributes)                     | Scaffold
transactionIndex|BigInteger                                           | Transaction index
tokenBalance    |BigInteger                                           | Scaffold token balance
enabled         |Boolean                                              | Scaffold enabled
currency        |String                                               | Fiat amount currency
shareHolders    |[ShareHolder](#share-holder-attributes)              | Scaffold share holders



##### Transaction

###### Transaction attributes

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
scaffold        |[Scaffold](#scaffold-attributes)                     | Scaffold
event           |[Event](#event)                                      | Event


##### Event

Each event has field `type` which define an event type

###### Activate scaffold event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
activated       |Boolean                                              |Scaffold state
type            |String                                               |"ACTIVATED_SCAFFOLD"

###### Added share holder event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
userAddress     |String                                               |Share holder address
partnerShare    |BigInteger                                           |Share holder share
type            |String                                               |"ADDED_SHARE_HOLDER"

###### Delete share holder event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
userAddress     |String                                               |Share holder address
type            |String                                               |"DELETED_SHARE_HOLDER"

###### Edit share holder event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
userAddress     |String                                               |Shareholder address
partnerShare    |BigInteger                                           |Shareholder share
type            |String                                               |"EDITED_SHARE_HOLDER"

###### Funds deposited event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
amount          |BigInteger                                           |Funds amount
toAddress       |String                                               |To address
type            |String                                               |"FUNDS_DEPOSITED"

###### Paid for shareholder event

Attribute       | Type                                                | Description
----------------|-----------------------------------------------------|-----------
userAddress     |String                                               |Shareholder address
amount          |BigInteger                                           |Paid amount
type            |String                                               |"PAID_FOR_SHARE_HOLDER"

###### Payment completed event

Attribute               | Type                                                | Description
------------------------|-----------------------------------------------------|-----------
customerAddress         |String                                               |Customer address
transactionAmount       |BigInteger                                           |Paid amount
scaffoldTransactionIndex|BigInteger                                           |Transaction index
properties              |Map<String, Any>                                     |Transaction properties
type                    |String                                               |"PAYMENT_COMPLETED"




#### Usage

`OpenGo` - an object that produces requests for each entity type

##### Get All Scaffolds

```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
scaffolds, err := op.GetScaffolds(ctx)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffolds:", scaffolds)
```

##### Get a Scaffold
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
address := "0x0000000000000000000000000000000000000000"
scaffold, err := op.GetScaffold(ctx, address)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffold:", scaffold)
```

##### Deploy a Scaffold
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()

scaffoldForDeployment := opengo.NewScaffold(
    token,
    "0x0000000000000000000000000000000000000000",
    "ABI",
    "Super Description3",
    "10",
    "USD",
    "5",
    "0x0000000000000000000000000000000000000001",
    "http://example.com",
    true,
)
scaffold.AddProperty(
    "PropertyName",
    "STRING",
    "PropertyDefaultValue",
)

deployedScaffold, err := op.DeployScaffold(ctx, *scaffold)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Deployed Scaffold", deployedScaffold)
}
```

##### Update a Scaffold
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
address := "0x0000000000000000000000000000000000000000"
scaffoldForUpdate, err := op.GetScaffold(ctx, address)
scaffoldForUpdate.WebHook = "https://www.example.com/webhook"
if err != nil {
    fmt.Println(err)
}
updatedScaffold, _ := op.UpdateScaffold(ctx, *scaffoldForUpdate)
fmt.Println("Scaffold: ", updatedScaffold)
```

##### Get a Scaffold Summary
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
address := "0x0000000000000000000000000000000000000000"
summary, err := op.GetScaffoldSummary(ctx, address)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffold Summary:", summary)
```

##### Get Scaffold Transactions
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
address := "0x0000000000000000000000000000000000000000"
transactions, err := op.GetTransactions(ctx, address)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffold Transactions:", transactions)
```

##### Deactivate a Scaffold
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
scaffoldAddress := "0x0000000000000000000000000000000000000000"
deactivatedScaffold, err := op.DeactivateScaffold(ctx, scaffoldAddress)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Scaffold:", deactivatedScaffold)
```

##### Get quota
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
quota, err := op.GetQuota(ctx)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Quota:", quota.CurrentCount)
```

##### Add Share Holders
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
quota, err := op.GetQuota(ctx)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Quota:", quota.CurrentCount)
```

##### Update Share Holders
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
scaffoldAddress := "0x0000000000000000000000000000000000000000"
shareholder := opengo.NewShareHolder("0x0000000000000000000000000000000000000000", 9)
response, err := op.UpdateShareHolder(ctx, scaffoldAddress, *shareholder)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Updated Share Holder Response:", response)
```

##### Delete Share Holders
```go
op := opengo.NewOpenGo(token)
ctx := context.Background()
scaffoldAddress := "0x0000000000000000000000000000000000000000"
holderAddress := "0x0000000000000000000000000000000000000000"
response, err := op.DeleteShareHolder(ctx, scaffoldAddress, holderAddress)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Share Holder Deleted response:", response)
```


