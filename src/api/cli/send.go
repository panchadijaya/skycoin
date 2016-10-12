package cli

import (
	gcli "github.com/urfave/cli"
)

func init() {
	cmd := gcli.Command{
		Name: "send",
		Description: `Send Skycoins from a wallet or an address to another address. 
                      If you are sending from a wallet the coins will be taken recursively 
                      from all addresses within the wallet starting with the first wallet 
                      until the amount of the transaction is met.
                      
                      Use caution when using the “-p” command. If you have command history 
                      enabled your wallet encryption password can be recovered from the history log. 
                      If you do not include the “-p” option you will be prompted to enter your password 
                      after you enter your command. 
                      `,
		Usage: "skycoin send  [option] [from wallet or address] [to address] [amount]",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "w",
				Usage: "[wallet file or path], From wallet. If no path is specified your default wallet path will be used.",
			},
			gcli.StringFlag{
				Name:  "a",
				Usage: "[address] From address.",
			},
			gcli.StringFlag{
				Name:  "c",
				Usage: "[changeAddress] Specify different change address. By default the from address or a wallets coinbase address will be used.",
			},
			gcli.StringFlag{
				Name:  "p",
				Usage: "[password] Password for address or wallet.",
			},
			gcli.StringFlag{
				Name:  "j,json",
				Usage: "Returns the results in JSON format.",
			},
		},
		Action: func(c *gcli.Context) error {
			return nil
		},
	}
	Commands = append(Commands, cmd)
}
