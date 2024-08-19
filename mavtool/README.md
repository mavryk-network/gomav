# mavtool

An operation injection helper

## Example

```go
package example

import (
    "context"
    "fmt"
    "math/big"

    "github.com/mavryk-network/gomav/v2"
    "github.com/mavryk-network/gomav/v2/b58"
    "github.com/mavryk-network/gomav/v2/client"
    "github.com/mavryk-network/gomav/v2/crypt"
    "github.com/mavryk-network/gomav/v2/protocol/core"
    "github.com/mavryk-network/gomav/v2/protocol/latest"
    "github.com/mavryk-network/gomav/v2/mavtool"
)

type logger struct{}

func (logger) Printf(format string, a ...any) {
    fmt.Printf(format, a...)
    fmt.Printf("\n")
}

func TransferToWallet(url, chainID, address, privateKey string, amount *big.Int) (core.OperationsGroup, error) {
    chain, err := b58.ParseChainID([]byte(chainID))
    if err != nil {
        return nil, err
    }
    addr, err := b58.ParsePublicKeyHash([]byte(address))
    if err != nil {
        return nil, err
    }
    c := client.Client{
        URL: url,
    }
    pk, err := b58.ParsePrivateKey([]byte(privateKey))
    if err != nil {
        return nil, err
    }
    priv, err := crypt.NewPrivateKey(pk)
    if err != nil {
        return nil, err
    }

    // initialize mavool
    tool := mavtool.New(&c, chain)
    tool.DebugLogger = logger{}

    // initialize signer
    signer := mavtool.NewLocalSigner(priv)

    // make a transaction
    val, err := gomav.NewBigUint(amount)
    if err != nil {
        // amount is negative
        return nil, err
    }
    tx := latest.Transaction{
        ManagerOperation: latest.ManagerOperation{
            Source: priv.Public().Hash(),
        },
        Amount:      val,
        Destination: core.ImplicitContract{PublicKeyHash: addr},
    }

    return tool.FillSignAndInjectWait(context.Background(), signer, []latest.OperationContents{&tx}, client.MetadataAlways, mavtool.FillAll)
}
```
