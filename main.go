package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/hashicorp/go-retryablehttp"
)

type Response struct {
	TokenId   string
	Timestamp string `json:"timestamp"`
	Balances  []struct {
		Account string `json:"account"`
		Balance int    `json:"balance"`
	} `json:"balances"`
	Links struct {
		Next string `json:"next"`
	} `json:"links"`
}

func main() {
	// token, amount and file flags
	tokenPtr := flag.String("token", "", "token id to query")
	amountPtr := flag.Int64("amount", 0, "amount to set for airdrop")
	balancePtr := flag.Int64("balance", 0, "the exact balance an account should have")
	filePtr := flag.String("file", "results.csv", "filename to save results as")
	flag.Parse()

	// if token is empty print defaults and exit
	if *tokenPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check token id matches regex
	match, _ := regexp.MatchString("^[0-9]+.[0-9]+.[0-9]+$", *tokenPtr)
	if !match {
		fmt.Println("Bad token ID format: should be shard.realm.num")
		os.Exit(1)
	}

	// query mirror node
	url := "https://mainnet-public.mirrornode.hedera.com"
	endpoint := fmt.Sprintf("/api/v1/tokens/%v/balances?account.balance=eq:%v&limit=100", *tokenPtr, *balancePtr)

	done := false
	var balances []string

	for !done {
		url := fmt.Sprintf("%v%v", url, endpoint)

		resp, err := retryablehttp.Get(url)

		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var response Response
		err = json.Unmarshal(data, &response)
		if err != nil {
			panic(err)
		}

		for _, a := range response.Balances {
			balances = append(balances, a.Account)
		}

		if len(response.Links.Next) == 0 {
			done = true
			break
		}

		endpoint = response.Links.Next
	}

	// create csv file
	file, err := os.Create(*filePtr)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	// loop through balances and write to csv
	for _, x := range balances {
		row := []string{x, strconv.Itoa(int(*amountPtr))}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	fmt.Printf("%v accounts with %v balance\n", len(balances), *balancePtr)
}
