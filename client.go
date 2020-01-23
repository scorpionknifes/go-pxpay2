package pxpay2

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func NewClient(pxPayUserId string, pxPayKey string, APIBase string) (*Client, error) {
	if pxPayUserId == "" || pxPayKey == "" || APIBase == "" {
		return nil, errors.New("pxPayUserId, pxPayKey and APIBase are required to create a Client")
	}

	return &Client{
		Client:      &http.Client{},
		PxPayUserId: pxPayUserId,
		PxPayKey:    pxPayKey,
		APIBase:     APIBase,
	}, nil
}

func (c *Client) CreatePurchase(p Purchase) (string, error) {
	payload := XmlPurchase{
		PxPayUserId:       c.PxPayUserId,
		PxPayKey:          c.PxPayKey,
		AmountInput:       p.Amount.AmountInput,
		CurrencyInput:     p.Amount.CurrencyInput,
		EmailAddress:      p.EmailAddress,
		MerchantReference: p.MechantReference,
		TxnData1:          p.Txn.TxnData1,
		TxnData2:          p.Txn.TxnData2,
		TxnData3:          p.Txn.TxnData3,
		TxnType:           "Purchase",
		UrlFail:           p.UrlFail,
		UrlSuccess:        p.UrlSuccess,
	}
	resp, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/pxaccess/pxpay.aspx"), payload)
	if err != nil {
		log.Panicln(err)
		return "", err
	}
	var res XmlRequest
	data, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(data, &res)
	if err != nil {
		log.Println(err)
		return "", err

	}

	if res.Valid == 0 {
		return "", errors.New(res.URI)
	}

	return res.URI, nil
}

func (c *Client) ProcessResponse(p ProcessResponse) (XmlResponse, error) {
	payload := XmlProcessResponse{
		PxPayUserId: c.PxPayUserId,
		PxPayKey:    c.PxPayKey,
		Response:    p.Response,
	}
	resp, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/pxaccess/pxpay.aspx"), payload)
	if err != nil {
		log.Panicln(err)
		return XmlResponse{}, err
	}

	var res XmlResponse
	data, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(data, &res)
	if err != nil {
		log.Println(err)
		return XmlResponse{}, err
	}

	if res.Success == 0 {
		return XmlResponse{}, errors.New(res.CardHolderName)
	}

	return res, nil
}

// NewRequest constructs a request
// Convert payload to a XML
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Response, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := xml.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml")
	resp, err := c.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
