package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type CheckIdResponse struct {
	Check struct {
		CheckID string `json:"check_id"`
	} `json:"check"`
}

type GetResponse struct {
	Check struct {
		CheckID        string `json:"check_id"`
		CompanySummary struct {
			CompanyStatus string `json:"company_status"`
			Result        string `json:"result"`
		} `json:"company_summary"`
		Country      string    `json:"country"`
		CreationDate time.Time `json:"creation_date"`
		NameScore    int       `json:"name_score"`
		IDScore      int       `json:"id_score"`
		Score        int       `json:"score"`
	}
}

func requestApi(method, url string, payload *strings.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Truora-API-key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0kzMThiNTZkZWJjMGI4YjgzYTA4OTM0YjdhNTgzMzFkOCIsImV4cCI6MzI2MTU5NDUzNiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ3OTQ1MzYsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX2NSWWFpaVo5VCIsImp0aSI6ImRhOGZkZDIxLTU0ZTktNGNlMy1iNDM4LTFhMDVmMTNlZGQ4NiIsImtleV9uYW1lIjoiYXBpdHJ1b3JhIiwia2V5X3R5cGUiOiJiYWNrZW5kIiwidXNlcm5hbWUiOiJnbWFpbGFsZWpqYXJpdmVyYS1hcGl0cnVvcmEifQ.JqbV837beUzgpu5OdOBbl504B7e9_yvfINKIobmEox4")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func RetrieveAccessToken(national_id, country, person_type string, user_authorized bool) (*url.URL, url.Values) {
	// The data we are going to send with our request
	url := "https://api.checks.truora.com"
	resource := "/v1/checks"
	data := url.Values{}
	data.Set("national_id", national_id)
	data.Add("country", country)
	data.Add("person_type", person_type)
	data.Add("user_authorized", strconv.FormatBool(user_authorized))

	fullurl, _ := url.ParseRequestURI(url)
	fullurl.Path = resource

	return fullurl, data
}

// para obtener el check_id
func PostRequestApiTruora(national_id, country, person_type string, user_authorized bool) (string, error) {

	fullurl, data := RetrieveAccessToken(national_id, country, person_type, user_authorized)
	urlToStr := fullurl.String()
	payload := strings.NewReader(data.Encode())
	method := "POST"

	body, err := requestApi(method, urlToStr, payload)
	if err != nil {
		fmt.Println(err)
		return "h", err
	}

	var ObtainResponse CheckIdResponse
	err = json.Unmarshal(body, &ObtainResponse)
	if err != nil {
		fmt.Println(err)
		return "h", err
	}

	// Get the ID of the object
	checkID := ObtainResponse.Check.CheckID

	return checkID, nil
}

func GetRequestApiTruora(check_id string) (int, error) {
	url := "https://api.checks.truora.com/v1/checks/" + check_id
	method := "GET"
	payload := strings.NewReader("national_id=74900799&country=PE&type=person&user_authorized=true")

	req, err := requestApi(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return 1, err
	}
	var GetApiResponse GetResponse
	err = json.Unmarshal(req, &GetApiResponse)
	if err != nil {
		fmt.Println(err)
		return 1, err
	}

	score := GetApiResponse.Check.Score

	return score, nil
}

func TruoraResponse(national_id, country, person_type string, user_authorized bool) (bool, error) {

	check_id, err := PostRequestApiTruora(national_id, country, person_type, user_authorized)
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	time.Sleep(10 * time.Second)
	score, err := GetRequestApiTruora(check_id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return score == 1, nil
}

func validation(national_id, country, person_type string, user_authorized bool) (bool, error) {
	score, err := TruoraResponse(national_id, country, person_type, user_authorized)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return score == 1, nil
}
