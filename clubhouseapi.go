// Package clubhouseapi handle work with clubhouse
package clubhouseapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lafin/http"
)

var url = "https://www.clubhouseapi.com/api"
var headers = map[string]string{
	"Content-Type": "application/json",
	"User-Agent":   "clubhouse/269 (iPhone; iOS 14.1; Scale/3.00)",
	"CH-AppBuild":  "269",
}

// StartPhoneNumberAuth is a method of logging in by phone
func StartPhoneNumberAuth(phoneNumber string) (StartPhoneNumberAuthResponse, error) {
	var data StartPhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/start_phone_number_auth", url), strings.NewReader(fmt.Sprintf(`{"phone_number":"%s"}`, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// CompletePhoneNumberAuth is a method of completing phone authentication
func CompletePhoneNumberAuth(phoneNumber, verificationCode string) (CompletePhoneNumberAuthResponse, error) {
	var data CompletePhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/complete_phone_number_auth", url), strings.NewReader(fmt.Sprintf(`{"verification_code":"%s","phone_number":"%s"}`, verificationCode, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetChannels is a method of getting all channels
func GetChannels() (GetChannelsResponse, error) {
	var data GetChannelsResponse
	response, err := http.Post(fmt.Sprintf("%s/get_channels", url), strings.NewReader("{}"), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// RefreshToken is a method for getting a new access token
func RefreshToken(refreshToken string) (RefreshTokenResponse, error) {
	var data RefreshTokenResponse
	response, err := http.Post(fmt.Sprintf("%s/refresh_token", url), strings.NewReader(fmt.Sprintf(`{"refresh":"%s"}`, refreshToken)), headers)
	if err != nil {
		return data, err
	}
	fmt.Println(string(response))
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// AddCredentials is a method of adding credentials such as User ID or Access Token
func AddCredentials(credentials map[string]string) {
	for k, v := range credentials {
		headers[k] = v
	}
}
