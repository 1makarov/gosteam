package gosteam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
	"time"
)

type Confirmation struct {
	ID        uint64
	Key       uint64
	Title     string
	Receiving string
	Since     string
	OfferID   uint64
}

var (
	//ErrConfirmationsUnknownError = errors.New("unknown error occurred finding confirmation")
	ErrCannotFindConfirmations   = errors.New("unable to find confirmation")
	ErrCannotFindDescriptions    = errors.New("unable to find confirmation descriptions")
	ErrConfirmationsDescMismatch = errors.New("cannot match confirmation with their respective descriptions")
	ErrWGTokenExpired            = errors.New("WGToken expired")
)

func (s *session) execConfirmationRequest(request, key, tag string, current int64, values map[string]interface{}) (*fasthttp.Response, error) {
	params := url.Values{
		"p":   {s.deviceID},
		"a":   {s.oauth.SteamID.ToString()},
		"k":   {key},
		"t":   {strconv.FormatInt(current, 10)},
		"m":   {"android"},
		"tag": {tag},
	}

	for k, v := range values {
		switch v := v.(type) {
		case string:
			params.Add(k, v)
		case uint64:
			params.Add(k, strconv.FormatUint(v, 10))
		default:
			return nil, fmt.Errorf("execConfirmationRequest: missing implementation for type %v", v)
		}
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetRequestURI("https://steamcommunity.com/mobileconf/" + request + params.Encode())
	req.Header.SetMethod("GET")
	s.cookieClient.FillRequest(req)
	resp := fasthttp.AcquireResponse()
	if err := s.getRedirect(req, resp, 200, "execConfirmationRequest"); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *session) GetConfirmations(identitySecret string) ([]*Confirmation, error) {
	current := time.Now().Unix()

	key, err := GenerateConfirmationCode(identitySecret, "conf", current)
	if err != nil {
		return nil, err
	}

	resp, err := s.execConfirmationRequest("conf?", key, "conf", current, nil)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, err
	}

	/* FIXME: broken
	if empty := doc.Find(".mobileconf_empty"); empty != nil {
		if done := doc.Find(".mobileconf_done"); done != nil {
			return nil, nil
		}

		return nil, ErrConfirmationsUnknownError // FIXME
	}
	*/

	entries := doc.Find(".mobileconf_list_entry")
	if entries == nil {
		return nil, ErrCannotFindConfirmations
	}

	descriptions := doc.Find(".mobileconf_list_entry_description")
	if descriptions == nil {
		return nil, ErrCannotFindDescriptions
	}

	if len(entries.Nodes) != len(descriptions.Nodes) {
		return nil, ErrConfirmationsDescMismatch
	}

	var confirmations []*Confirmation
	for k, sel := range entries.Nodes {
		confirmation := &Confirmation{}
		for _, attr := range sel.Attr {
			if attr.Key == "data-confid" {
				confirmation.ID, _ = strconv.ParseUint(attr.Val, 10, 64)
			} else if attr.Key == "data-key" {
				confirmation.Key, _ = strconv.ParseUint(attr.Val, 10, 64)
			} else if attr.Key == "data-creator" {
				confirmation.OfferID, _ = strconv.ParseUint(attr.Val, 10, 64)
			}
		}

		descSel := descriptions.Nodes[k]
		depth := 0
		for child := descSel.FirstChild; child != nil; child = child.NextSibling {
			for n := child.FirstChild; n != nil; n = n.NextSibling {
				switch depth {
				case 0:
					confirmation.Title = n.Data
				case 1:
					confirmation.Receiving = n.Data
				case 2:
					confirmation.Since = n.Data
				}
				depth++
			}
		}

		confirmations = append(confirmations, confirmation)
	}

	return confirmations, nil
}

func (s *session) AnswerConfirmation(confirmation *Confirmation, identitySecret string) error {
	current := time.Now().Unix()
	answer := "allow"

	key, err := GenerateConfirmationCode(identitySecret, answer, current)
	if err != nil {
		return err
	}

	op := map[string]interface{}{
		"op":  answer,
		"cid": confirmation.ID,
		"ck":  confirmation.Key,
	}

	resp, err := s.execConfirmationRequest("ajaxop?", key, answer, current, op)

	if err != nil {
		return err
	}

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	var response Response
	if err := json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return errors.New(response.Message)
	}

	return nil
}
