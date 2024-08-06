package ethereum_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/codecodeorg/opgeth/ethclient"
)

func TestEthereum(t *testing.T) {

	client, err := ethclient.Dial("https://base-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(block)

	// receipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(block.Hash(), true))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, receipt := range receipts {
	// 	log.Println(receipt.TxHash.String())
	// }

	// log.Println(receipts)
}
