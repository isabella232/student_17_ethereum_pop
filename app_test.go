package main

/*
  var key = `{"address":"286485b3026d5d817f1f444060516b439b13dd2b","crypto":{"cipher":"aes-128-ctr","ciphertext":"d1a54d49808b658d9ea5a2c795c6a26741483699bf258d43a1d102dbfded867a","cipherparams":{"iv":"779603e70f888ee1496cbe19a7575cef"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3fcc04ce5dbbfcdcdadeff6d5a69b05bb1931e435459105883ac64de5aefe271"},"mac":"d170e43d5e67e747bf766a112f48f649f574e6938ceb8c361dd73f7e2a586c16"},"id":"cad9bc2c-89c2-401f-bc5e-de4b4a11161d","version":3}`

func Test_Rinkeby(t *testing.T) {

	s := make([]common.Address, 1)
	s[0] = common.HexToAddress("0x286485b3026d5d817f1f444060516b439b13dd2b")
	fmt.Println(s)

	keysetA := make([][32]byte, 1)
	keyA := []byte("AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQ")
	copy(keysetA[0][:], keyA)

	//keysetA.push("AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")
	//keyset[0][32]("AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")
	//attendeeKey := "AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P"
	//copy(keyset[0], attendeeKey)

	//	copy(keysetA[0][:], "AAAAC3NzaC1lZDI1NTE5AAAAIDD6LNlQsbLyr0Mp/6mJUCAILBbNGrJefiaVp+G6H97P")

	endPoint := "/home/hugo/.ethereum/rinkeby/geth.ipc"

	auth, err := bind.NewTransactor(strings.NewReader(key), "testpassword")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//Deploys contract
	contract := orgConfig(auth, endPoint, Config{"First Party", "Nazareth", 1, s, 60})

	//Link to deployed contract
	//contract := orgLink(endPoint, testAddress)

	time.Sleep(30 * time.Second)

	st, _ := contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//organisator signs contract
	sign(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//Administrator signs whole configuration
	signAdmin(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	//organizator deposit key set
	orgPublic(auth, contract, keysetA)

	time.Sleep(45 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

	fmt.Println(contract.AllSets(nil, big.NewInt(0)))

	//Administrator calls consensus function
	orgFinal(auth, contract)

	time.Sleep(30 * time.Second)

	st, _ = contract.CurrentState(nil)
	fmt.Println(returnState(uint(st)))
	fmt.Println(contract.CurrentState(nil))

}
*/
