package instacart

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Instacart interface {
	DeliveryAvailable(storeName string) (bool, error)
}

func NewInstacart(email, password string) (Instacart, error) {
	const login = "https://www.instacart.com/v3/dynamic_data/authenticate/login?source=mobile_web&cache_key=undefined"
	const cookieSessionName = "_instacart_session="

	sessionToken := ""
	reqBody, err := json.Marshal(map[string]string{
		"email":      email,
		"password":   password,
		"grant_type": "password",
	})

	if err != nil {
		return &instacart{}, err
	}

	resp, err := http.Post(login, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return &instacart{}, err
	}

	defer resp.Body.Close()

	values := resp.Header.Values("set-cookie")
	exists := startsWith(values, cookieSessionName)

	if !exists {
		return &instacart{}, errors.New("couldn't login.Unable to retrieve session token")
	}

	for _, c := range values {
		if strings.HasPrefix(c, cookieSessionName) {
			sessionToken = strings.Split(c, cookieSessionName)[1]
			break
		}
	}
	return &instacart{SessionToken: sessionToken}, nil
}

type instacart struct {
	SessionToken string
}

func (i *instacart) DeliveryAvailable(storeName string) (bool, error) {
	url := fmt.Sprintf(
		"https://www.instacart.com/v3/containers/%s/next_gen/retailer_information/content/delivery?source=web",
		storeName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Cookie", fmt.Sprintf("_instacart_session=%s", i.SessionToken))

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	d := delivery{}
	err = json.Unmarshal(body, &d)

	if err != nil {
		return false, err
	}

	var e []string
	for _, module := range d.Container.Modules {
		e = append(e, module.ID)
	}

	noAvailability := startsWith(e, "errors_no_availability")
	deliverySlots := startsWith(e, "delivery_option_list")

	if noAvailability {
		return false, nil
	}
	if deliverySlots {
		return true, nil
	}

	return false, nil
}

// startsWith tells whether a startsWith x.
func startsWith(a []string, x string) bool {
	for _, n := range a {
		if strings.HasPrefix(n, x) {
			return true
		}
	}
	return false
}
