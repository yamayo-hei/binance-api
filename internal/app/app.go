package app

import (
	"fmt"
	"github.com/yamayo-hei/binance-api/internal/pkg/account"
	"log"
	"time"
)

type Product struct {
	ProductID int `json:"ProductID"`
	Amount  int    `json:"Amount"`
}

func App(account *account.Account, id int64) {
	defer fmt.Scanf("\n")

	orderCreateBody, err := createOrderCreateBody(id)
	if err != nil {
		log.Fatalf("error marshal buy box: %s\n", err.Error())
	}
	orderCreateRequest := account.Auth.orderCreateRequest(orderCreateBody)

	// log.Println("Waiting started successfully")
	// wait(box.Information.StartTime)
	log.Println("Start buy")

	go func() {
		for {
			resp, err := account.Auth.NFTMysteryBoxBuy(orderCreateRequest)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(string(resp.Body()))
			return
		}
	}()

	// time.Sleep(6 * time.Second)
	// box.Status = true
	// time.Sleep(1 * time.Second)

	fmt.Println("Purchases are completed")
}

func wait(s int64) {
	t := time.Unix(s, 0).UTC().Add(-3 * time.Second).Unix()
	for {
		if time.Now().UTC().Unix() >= t {
			return
		}
	}
}

func createOrderCreateBody(id int64) ([]byte, error) {
	// fmt.Print("Enter the product id: ")
	// fmt.Fscan(os.Stdin, &number)
	b, err := json.Marshal(Product{ProductID: id, Amount: 0.5})
	if err != nil {
		return nil, err
	}
	return b, nil
}
