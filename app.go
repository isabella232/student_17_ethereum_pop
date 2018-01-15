package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/dedis/onet.v1/app"
	"gopkg.in/dedis/onet.v1/log"
	"gopkg.in/dedis/onet.v1/network"
	cli "gopkg.in/urfave/cli.v1"
)

var gasL int64 = 3000000
var gasP int64 = 90000000000

func main() {
	network.RegisterMessage(&Config{})
	appCli := cli.NewApp()
	appCli.Name = "Proof-of-personhood party"
	appCli.Usage = "Handles party-creation, finalizing, pop-token creation, and verification"
	appCli.Version = "0.1"
	appCli.Commands = []cli.Command{}
	appCli.Commands = []cli.Command{
		{
			Name:    "organizer",
			Aliases: []string{"org"},
			Usage:   "Organising a PoParty",
			Subcommands: []cli.Command{
				{
					Name:      "link",
					Aliases:   []string{"l"},
					Usage:     "deploy new pop contract",
					ArgsUsage: "private key, network path and account nonce",
					Action:    orgLink,
				},
				{
					Name:      "config",
					Aliases:   []string{"c"},
					Usage:     "stores the configuration",
					ArgsUsage: "private key and pop_desc.toml",
					Action:    orgConfig,
				},
				{
					Name:      "public",
					Aliases:   []string{"p"},
					Usage:     "stores one or more public keys during the party",
					ArgsUsage: "key1,key2,key3 party_hash",
					Action:    orgPublic,
				},
				{
					Name:      "sign",
					Aliases:   []string{"s"},
					Usage:     "sign the configuration for organizers",
					ArgsUsage: "private key",
					Action:    sign,
				},
				{
					Name:      "signAdmin",
					Aliases:   []string{"sA"},
					Usage:     "sign the whoel config",
					ArgsUsage: "private key of administrator",
					Action:    signAdmin,
				},
				{
					Name:      "final",
					Aliases:   []string{"f"},
					Usage:     "reach consensus",
					ArgsUsage: "private key",
					Action:    orgFinal,
				},
			},
		}}
	appCli.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "debug,d",
			Value: 0,
			Usage: "debug-level: 1 for terse, 5 for maximal",
		},
		cli.StringFlag{
			Name:  "config,c",
			Value: "~/.config/cothority/pop",
			Usage: "The configuration-directory of pop",
		},
		cli.StringFlag{
			Name:  "Orgconfig, Oc",
			Value: "~/.config/cothority/pop",
			Usage: "The configuration-directory of pop",
		},
	}
	appCli.Before = func(c *cli.Context) error {
		log.SetDebugVisible(c.Int("debug"))
		return nil
	}
	appCli.Run(os.Args)

}

type Config struct {
	//name of config-file
	Name string
	//path to ipc
	Network string
	//Address of contract
	Address string
	//private key
	Private string
	//Nonce of account
	Nonce int
}

type popDescToml struct {
	// config-file name
	Name string
	//Location of party
	Location string
	//number of organizers
	NumberOfOrganizers int64
	//organizers Adresses in Ethereum format
	OrganizersAddresses []string
	//Duration of party in minutes
	Deadline int64
}

type PopDesc struct {
	// config-file name
	Name string
	//Location of party
	Location string
	//number of organizers
	NumberOfOrganizers int64
	//organizers Adresses in Ethereum format
	OrganizersAddresses []common.Address
	//Duration of party in minutes
	Deadline int64
}

//connect to contract
func orgLink(c *cli.Context) error {
	log.Lvl3("Org: Link")
	if c.NArg() < 3 {
		log.Fatal("Please provide valid private key, geth.ipc path and account Nonce")
	}
	key, _ := crypto.HexToECDSA(c.Args().Get(0))
	auth := bind.NewKeyedTransactor(key)
	network := c.Args().Get(1)
	conn, err := ethclient.Dial(network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	cfg := getConfig(c)
	cfg.Nonce, err = strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return err
	}
	addr, txe, _, err := DeployPopcontract(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	}, conn)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	cfg.Nonce = cfg.Nonce + 1
	fmt.Printf("Successfully linked with : %x \n", addr)
	fmt.Printf("Transaction receipt : %x \n", txe.Hash())
	cfg.Address = addr.String()
	cfg.Network = network
	cfg.Private = c.Args().First()
	cfg.write()
	return nil
}

//Deploy & configure a new pop-party
func orgConfig(c *cli.Context) error {
	log.Lvl3("Org: Config")
	if c.NArg() < 1 {
		log.Fatal(`Please give valid pop_desc.toml `)
	}
	cfg := getConfig(c)
	if cfg.Address == "" {
		log.Fatal("No address")
		return errors.New("No address found - please link first")
	}
	var desc popDescToml
	desc1, err := ioutil.ReadFile(c.Args().Get(0))
	if err != nil {
		return err
	}
	if _, err := toml.Decode(string(desc1), &desc); err != nil {
		log.Fatal("Error decoding toml file.", err)
	}
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	fmt.Println("Connected to contract.")
	key, err := crypto.HexToECDSA(cfg.Private)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(key)
	addresses := make([]common.Address, desc.NumberOfOrganizers)
	var i int64 = 0
	for i < desc.NumberOfOrganizers {
		addresses[i] = common.HexToAddress(desc.OrganizersAddresses[i])
		i++
	}
	txe, err := contract.SetConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	}, desc.Name, desc.Location, big.NewInt(desc.NumberOfOrganizers), addresses, big.NewInt(desc.Deadline))
	if err != nil {
		log.Fatalf("could not set configuration: %v", err)
	}
	cfg.Nonce = cfg.Nonce + 1
	fmt.Printf("Configuration set. Transaction : %x \n", txe.Hash())
	cfg.write()
	return nil
}

// adds a public key to the list
func orgPublic(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal("Please give the public keys to add")
	}
	log.Lvl3("Org: Adding public keys", c.Args().Get(1))
	str := c.Args().Get(1)
	if !strings.HasPrefix(str, "[") {
		str = "[" + str + "]"
	}
	// TODO: better cleanup rules
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	str = strings.Replace(str, "\\", "", -1)
	log.Lvl3("Niceified public keys are:\n", str)
	keys := strings.Split(str, ",")
	cfg := getConfig(c)
	keyset := make([][32]byte, len(keys))
	for i := 0; i < len(keys); i++ {
		keyA := []byte(keys[i])
		copy(keyset[i][:], keyA)
	}
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(cfg.Private)
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.DepositPublicKeys(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	}, keyset)
	if err != nil {
		log.Fatalf("could not deposit keyset: %v", err)
	}
	cfg.Nonce = cfg.Nonce + 1
	fmt.Printf("Keyset Added. Transaction : %x \n", txe.Hash())
	cfg.write()
	return nil
}

//Organizators sign configuration
func sign(c *cli.Context) error {
	if c.NArg() > 0 {
		log.Fatal(`No argument needed`)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(cfg.Private)
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.ConfigSignOrganizers(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	})
	if err != nil {
		log.Fatalf("could not sign contract: %v", err)
	}
	cfg.Nonce = cfg.Nonce + 1
	fmt.Printf("Configuration signed. Transaction : %x \n", txe.Hash())
	cfg.write()
	return nil
}

//Administrator sign whole configuration
func signAdmin(c *cli.Context) error {
	if c.NArg() > 0 {
		log.Fatal(`No arguments needed `)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(cfg.Private)
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v \n", err)
	}
	txe, err := contract.SignWholeConfiguration(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	})
	if err != nil {
		log.Fatalf("could not sign contract: %v", err)
	}
	cfg.Nonce = cfg.Nonce + 1
	fmt.Printf("Whole configuration signed. Transaction : %x \n", txe.Hash())
	cfg.write()
	return nil
}

//reach consensus
func orgFinal(c *cli.Context) error {
	if c.NArg() > 0 {
		log.Fatal(`No arguments needed `)
	}
	cfg := getConfig(c)
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		log.Fatalf("could not connect to network: %v", err)
	}
	key, _ := crypto.HexToECDSA(cfg.Private)
	auth := bind.NewKeyedTransactor(key)
	contract, err := NewPopcontract(common.HexToAddress(cfg.Address), conn)
	if err != nil {
		log.Fatalf("could not connect to contract: %v \n", err)
	}
	txe, err := contract.PublicKeyConsensus(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(gasL),
		GasPrice: big.NewInt(gasP),
		Value:    big.NewInt(0),
		Nonce:    big.NewInt(int64(cfg.Nonce)),
	})
	if err != nil {
		log.Fatalf("Could not reach consensus: %v", err)
	}
	cfg.Nonce++
	fmt.Printf("Asking for consensus. Transaction : %x \n", txe.Hash())
	cfg.write()

	return nil
}

// getConfigClient returns the configuration and a client-structure.
func getConfig(c *cli.Context) *Config {
	cfg, err := newConfig(path.Join(c.GlobalString("config"), "config.bin"))
	log.ErrFatal(err)
	return cfg
}

// newConfig tries to read the config and returns an organizer-
// config if it doesn't find anything.
func newConfig(fileConfig string) (*Config, error) {
	name := app.TildeToHome(fileConfig)
	if _, err := os.Stat(name); err != nil {
		return &Config{
			Name: name,
		}, nil
	}
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s: %s - please remove it",
			name, err)
	}
	//blocks here
	_, msg, err := network.Unmarshal(buf)
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %s",
			name, err)
	}
	cfg, ok := msg.(*Config)
	if !ok {
		log.Fatal("Wrong data-structure in file", name)
	}
	cfg.Name = name
	return cfg, nil
}

// write saves the config to the given file.
func (cfg *Config) write() {
	buf, err := network.Marshal(cfg)
	log.ErrFatal(err)
	log.ErrFatal(ioutil.WriteFile(cfg.Name, buf, 0660))
}
