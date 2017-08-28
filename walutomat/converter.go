package walutomat

import (
	"strconv"
	"log"
	"encoding/json"
)

func Convert(response []byte) ([] Offer) {
	var offerResponseDecoded OffersResponse

	err := json.Unmarshal(response, &offerResponseDecoded)

	if err != nil {
		log.Fatal(err)
	}

	offers :=  offerResponseToOffer(offerResponseDecoded)

	log.Printf("Offers in response: %d", len(offers))

	return offers
}

func offerResponseToOffer(offerResponse OffersResponse) ([] Offer) {
	var offers []Offer

	for _, offerResponse := range offerResponse.Offers {
		offers = append(offers, Offer{
			Pair: offerResponse.Pair,
			Buy: stringToFloat(offerResponse.Buy),
			BuyOld: stringToFloat(offerResponse.BuyOld),
			Sell: stringToFloat(offerResponse.Sell),
			SellOld: stringToFloat(offerResponse.SellOld),
			CountBuy: offerResponse.CountBuy,
			CountSell: offerResponse.CountSell,
			Avg: stringToFloat(offerResponse.Avg),
			AvgOld: stringToFloat(offerResponse.AvgOld),
		})
	}

	return offers
}

func stringToFloat(floatAsString string) (float32)  {
	result, err := strconv.ParseFloat(floatAsString, 32)

	if nil != err {
		log.Printf("Conversion error %s\n", err)
	}

	return float32(result)
}