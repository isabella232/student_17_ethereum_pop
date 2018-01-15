# Proof of Personhood on the Ethereum blockchain
Concile anonymity and accountability using proof of personhood tokens and the ethereum blockchain.

## Setting up the configuration of the party

### Step 1, setting up the paramaters :

Organizer 1 (admin) deploys the popcontract and sends the configuration of the party using a private function, registering the different data needed : name, place, duration and the Ethereum public addresses of the n other organizers.

### Step 2, signing the configuration:

Each organizer individually signs off individually the configuration using is public key (the contract matches the address with the initial addresses of the configuration).
The admin then calls a function to sign the whole configuration.
Note : Might be possible to modify the function such that the other organizers can sign using their private key, removing the need of the admin to sign it manually.

### Step 3, deposit of public keys of the participants :

Every organizer sends a (sorted) list of public attendees keys to the contract. Until the limit time/date, they are free to update the list.

### Step 4, reach consensus :

The contract decides which set of keys to register by taking the one submitted by most organizers. Other ways to reach consensus are also possible. The organizer could include the one choosen in the configuration of the party. Once this is done, the contract enters the state "Locked".

## How to use using Go :

You will need to have geth and to run a full node. Organizers will need to create accounts using geth and will need their private key.
To deploy the contract run :

`go build`

`./popcontract org link "your private key" "your geth path" "your account nonce (in string format)" `

To get your account nonce, either use a block explorer like etherscan.io or :

`geth attach ipc:"path/to/your/geth.ipc"`

`web3.eth.getTransactionCount("your public address")`

To set the configuration :

`./popcontract org config "dest.toml"`

To sign :

`./popcontract org sign "private key"`

To sign whole configuration :

`./popcontract org signAdmin`

To deposit new key set :

`./popcontract org public "private key" "keyset"`

To reach consensus :

`./popcontract org final`

To apply modification to popcontract.sol :

`abigen --sol=popcontract.sol --pkg=main --out=pop.go`

![Image](/doc/fsm.jpg)


## To do :

* Change consensus function
* Add final statement function to final
* Add different config for organizers

## Notes :

* The gas prices and limits are hard-coded in app.go. They are very high so it might be
a good idea to lower them if you want to test it on the main network.
* If you prefer interacting with the contract using a GUI you can use myetherwallet with the address of the contract
and it's ABI that you can find using remix.ethereum.org
* Always wait for previous transaction to confirm or your next one will fail as the contract won't be in the correct state
