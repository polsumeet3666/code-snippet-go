package main

import (
	"fmt"
	"log"
	"time"

	"github.com/linxGnu/gosmpp"
	"github.com/linxGnu/gosmpp/data"
	"github.com/linxGnu/gosmpp/pdu"
)

func main() {
	// auth
	auth := gosmpp.Auth{
		SMSC:       "smscsim.melroselabs.com:2775",
		SystemID:   "670188",
		Password:   "OGRVFF",
		SystemType: "",
	}

	// transrecier
	trans, err := gosmpp.NewTransceiverSession(gosmpp.NonTLSDialer, auth, gosmpp.TransceiveSettings{
		EnquireLink: 5 * time.Second,

		OnSubmitError: func(p pdu.PDU, err error) {
			log.Fatal(err)
		},

		OnReceivingError: func(err error) {
			fmt.Println(err)
		},

		OnRebindingError: func(err error) {
			fmt.Println(err)
		},

		OnPDU: handlePDU(),

		OnClosed: func(state gosmpp.State) {
			fmt.Println(state)
		},
	}, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	err = trans.Transceiver().Submit(newSubmitSM())
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		_ = trans.Close()
	}()
}
func handlePDU() func(pdu.PDU, bool) {
	return func(p pdu.PDU, responded bool) {
		switch pd := p.(type) {
		case *pdu.SubmitSMResp:
			fmt.Printf("SubmitSMResp:%+v\n", pd)

		case *pdu.GenerickNack:
			fmt.Println("GenericNack Received")

		case *pdu.EnquireLinkResp:
			fmt.Println("EnquireLinkResp Received")

		case *pdu.DataSM:
			fmt.Printf("DataSM:%+v\n", pd)

		case *pdu.DeliverSM:
			fmt.Printf("DeliverSM:%+v\n", pd)
			fmt.Println(pd.Message.GetMessage())
		}
	}
}
func newSubmitSM() *pdu.SubmitSM {
	// build up submitSM
	srcAddr := pdu.NewAddress()
	srcAddr.SetTon(5)
	srcAddr.SetNpi(0)
	_ = srcAddr.SetAddress("00" + "522241")

	destAddr := pdu.NewAddress()
	destAddr.SetTon(1)
	destAddr.SetNpi(1)
	_ = destAddr.SetAddress("99" + "522241")

	submitSM := pdu.NewSubmitSM().(*pdu.SubmitSM)
	submitSM.SourceAddr = srcAddr
	submitSM.DestAddr = destAddr
	_ = submitSM.Message.SetMessageWithEncoding("Đừng buồn thế dù ngoài kia vẫn mưa nghiễng rợi tý tỵ", data.UCS2)
	submitSM.ProtocolID = 0
	submitSM.RegisteredDelivery = 1
	submitSM.ReplaceIfPresentFlag = 0
	submitSM.EsmClass = 0

	return submitSM
}
