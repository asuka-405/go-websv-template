package libsailpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"root/src/lib/libauth"
	"root/src/lib/libcache"
)

func CreateAccessProfile(name, ownerId, sourceName, sourceId, entitlements string) {
	var entitlementList []APReqEntitlement
	err := json.Unmarshal([]byte(entitlements), &entitlementList)
	if err != nil {
		log.Println(err)
	}

	cid, csec := libcache.LoadCredentials()
	token, err := libauth.GetBearerToken("https://tmf-group-sb.api.identitynow.com/oauth/token", cid, csec)
	if err != nil {
		log.Println(err)
	}

	sourceType := "SOURCE"
	ownerType := "IDENTITY"

	ap_req_body := AccessProfileReq{
		Name: name,
		Owner: &APReqOwner{
			ID:   ownerId,
			Type: &ownerType,
		},
		Source: &APReqSource{
			ID:   sourceId,
			Name: sourceName,
			Type: &sourceType,
		},
		Requestable:  true,
		Enabled:      true,
		Entitlements: entitlementList,
	}

	ap_req_body_json, err := json.Marshal(ap_req_body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(ap_req_body_json))

	req, err := http.NewRequest(http.MethodPost, "https://tmf-group-sb.api.identitynow.com/v2024/access-profiles", bytes.NewBuffer(ap_req_body_json))

	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Status)
}
