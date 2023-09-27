package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/test-request", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("Invalid when we shoud get valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/test-request", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "test-request", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows does not have required fields when it has")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/test-request", nil)
	form := New(r.PostForm)

	has := form.Has("test-request")

	if has {
		t.Error("Form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("Shows form does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/test-request", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("Form Shows min length for non existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have and error but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("field", "value")
	form = New(postedValues)

	form.MinLength("field", 100)

	if form.Valid() {
		t.Error("Shows min length of 100 when length is not")
	}

	postedValues = url.Values{}

	postedValues.Add("another_field", "stringhas1")
	form = New(postedValues)

	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("Shows min length of 1 when length is not")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should not have and error but get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid for non existent fields")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc@email.com")
	form = New(postedValues)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("got an invalid email when we should get valid")
	}
}
