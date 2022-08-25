<!--
order: 2
-->

# Messages

In this section we describe the processing of the crisis messages and the
corresponding updates to the state.

## MsgVerifyInvariant

Blockchain invariants can be checked using the `MsgVerifyInvariant` message.

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/proto/cosmos/crisis/v1beta1/tx.proto#L16-L26

This message is expected to fail if:

* the sender does not have enough coins for the constant fee
* the invariant route is not registered

This message checks the invariant provided, and if the invariant is broken it
panics, potentially halting the blockchain if it is configured according to **Configuring the Chain to Halt** below.
If the invariant is broken, the constant fee is never deducted as the transaction is never committed to a block (equivalent to
being refunded). However, if the invariant is not broken, the constant fee will not be refunded.

## Configuring the Chain to Halt on MsgVerifyInvariant
By default, cosmos-sdk chains not allow a panic to halt the chain while executing a Tx. To configure the chain
to halt when any registered invariant returns an error, there are two options: forcing all nodes to halt or allowing node operators to decide.

### Making all nodes halt
```go
// in app/app.go when creating your chain's App, after creating the baseapp add the following
app.BaseApp.AddRunTxRecoveryHandler(baseapp.NewCrisisInvariantsHaltNodeRecoveryHandler())
```

### Allowing node operators to decide
First, add a flag to your application's start command to give node operators the option to halt:
```go
// under the cmd/ directory in root.go, add the following to addModuleInitFlags() after crisis.AddModuleInitFlags(startCmd)
crisis.AddOptionalModuleInitFlags(startCmd)
```
Next, handle the new flag that has been added after creating the BaseApp:
```go
// note: baseapp is usually created by calling baseapp.NewBaseApp() or appBuilder.Build()
// in app/app.go when creating your chain's App, after creating the baseapp add the following
app.BaseApp.ApplyCrisisVerifyHaltsNodeFlag(cast.ToBool(appOpts.Get(crisis.OptionalFlagVerifyHaltsNode)))
```