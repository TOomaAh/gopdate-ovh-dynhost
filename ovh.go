package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

func UpdateDynHost(ip WanIP, dynHost DynHost) (bool, error) {
	var url string = fmt.Sprintf("http://www.ovh.com/nic/update?system=dyndns&hostname=%s&myip=%s", dynHost.Domain, ip.IP)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	credential := fmt.Sprintf("%s:%s", dynHost.Username, dynHost.Password)
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(credential)))

	res, err := client.Do(req)

	if err != nil {
		return false, errors.New("Domain: %s => Cannot make request\n")
	}

	if res.StatusCode == 401 {
		return false, errors.New("Domain: %s => Username of Password incorrect\n")
	}

	defer res.Body.Close()

	return true, nil
}
