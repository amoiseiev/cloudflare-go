package cloudflare

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAcessCACertificate(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "result": {
    "id": "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
    "aud": "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
    "public_key": "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org"
  },
  "success": true,
  "errors": [],
  "messages": []
}
		`)
	}

	want := AccessCACertificate{
		ID:        "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
		Aud:       "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
		PublicKey: "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org",
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	actual, err := client.AccessCACertificate(accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	actual, err = client.ZoneLevelAccessCACertificate(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAcessCACertificates(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "result": [{
    "id": "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
    "aud": "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
    "public_key": "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org"
  }],
  "success": true,
  "errors": [],
  "messages": []
}
		`)
	}

	want := []AccessCACertificate{{
		ID:        "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
		Aud:       "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
		PublicKey: "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org",
	}}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/ca", handler)

	actual, err := client.AccessCACertificates(accountID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/ca", handler)

	actual, err = client.ZoneLevelAccessCACertificates(zoneID)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateAcessCACertificates(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "result": {
    "id": "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
    "aud": "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
    "public_key": "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org"
  },
  "success": true,
  "errors": [],
  "messages": []
}
		`)
	}

	want := AccessCACertificate{
		ID:        "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4",
		Aud:       "7d1996154eb606c19e31dd777fe6981f57a5ab66735c5c00fefd01b1200ba9d0",
		PublicKey: "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTI...3urg/XpGMdgaSs5ZdptUPw= open-ssh-ca@cloudflareaccess.org",
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	actual, err := client.CreateAccessCACertificate(accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	actual, err = client.CreateZoneLevelAccessCACertificate(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteAcessCACertificates(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "result": {
    "id": "4f74df465b2a53271d4219ac2ce2598e24b5e2c60c7924f4"
  },
  "success": true,
  "errors": [],
  "messages": []
}
		`)
	}

	mux.HandleFunc("/accounts/"+accountID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	err := client.DeleteAccessCACertificate(accountID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	assert.NoError(t, err)

	mux.HandleFunc("/zones/"+zoneID+"/access/apps/f174e90a-fafe-4643-bbbc-4a0ed4fc8415/ca", handler)

	err = client.DeleteZoneLevelAccessCACertificate(zoneID, "f174e90a-fafe-4643-bbbc-4a0ed4fc8415")

	assert.NoError(t, err)
}
