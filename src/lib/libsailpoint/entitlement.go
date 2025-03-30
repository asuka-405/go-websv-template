package libsailpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"root/src/lib/libauth"
	"root/src/lib/libcache"
)

func GetEntitlements(sourceID string) []Entitlement {
	cid, csec := libcache.LoadCredentials()

	token, err := libauth.GetBearerToken("https://tmf-group-sb.api.identitynow.com/oauth/token", cid, csec)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	offset := 0
	limit := 250
	var allEntitlements []Entitlement

	filter := fmt.Sprintf("source.id eq \"%s\"", sourceID)
	encodedFilter := url.QueryEscape(filter)

	for {
		url := fmt.Sprintf("https://tmf-group-sb.api.identitynow.com/v2024/entitlements?offset=%d&limit=%d&filters=%s", offset, limit, encodedFilter)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println("entitlement request build err: %s", err.Error())
			break
		}
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-SailPoint-Experimental", "true")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("entitlement response err: %s", err.Error())
			break
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("entitlement body read err: %s", err.Error())
			break
		}
		fmt.Println(string(body))
		var entitlements []Entitlement
		err = json.Unmarshal(body, &entitlements)
		if err != nil {
			fmt.Println("request url, headers: %s, %s", url, req.Header)
			fmt.Println("entitlement unmarshall err: %s", err.Error())
			break
		}

		// Stop when an empty array is received
		if len(entitlements) == 0 {
			break
		}

		allEntitlements = append(allEntitlements, entitlements...)
		offset += limit
	}

	return allEntitlements

}
