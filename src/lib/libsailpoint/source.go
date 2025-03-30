package libsailpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"root/src/lib/libauth"
	"root/src/lib/libcache"
)

func GetSources() []Source {
	cid, csec := libcache.LoadCredentials()

	token, err := libauth.GetBearerToken("https://tmf-group-sb.api.identitynow.com/oauth/token", cid, csec)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodGet, "https://tmf-group-sb.api.identitynow.com/v2024/sources", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var sources []Source
	err = json.Unmarshal(body, &sources)
	if err != nil {
		fmt.Println(err)
	}

	return sources

}
