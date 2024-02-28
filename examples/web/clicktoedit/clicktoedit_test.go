package clicktoedit_test

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit"
)

func TestDemo(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	clicktoedit.NewHandler(false).ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("expected status code 200 got %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	text := doc.Text()

	if !strings.Contains(text, "First Name") {
		t.Errorf("expected response to contain 'First Name' got %s", text)
	}
	if !strings.Contains(text, "Last Name") {
		t.Errorf("expected response to contain 'Last Name' got %s", text)
	}
	if !strings.Contains(text, "Email") {
		t.Errorf("expected response to contain 'Email' got %s", text)
	}

	link, ok := doc.Find("button").Attr("hx-get")
	if !ok {
		t.Errorf("expected button to have hx-get attribute")
	}
	if link != "/examples/templ/click-to-edit/edit" {
		t.Errorf("expected button to have hx-get attribute with value '/examples/click-to-edit/edit' got %s", link)
	}
}

func TestEditPost(t *testing.T) {
	post := func(firstName, lastName, email string) *httptest.ResponseRecorder {
		form := url.Values{}
		form.Add("firstName", firstName)
		form.Add("lastName", lastName)
		form.Add("email", email)
		body := strings.NewReader(form.Encode())

		req := httptest.NewRequest("POST", "/edit", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()

		clicktoedit.NewHandler(false).ServeHTTP(w, req)

		return w
	}

	t.Run("should go back to editing from good post", func(t *testing.T) {
		w := post("John", "Smith", "john@smith.com")

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatalf("expected status code 200 got %d", res.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// check that the form is back to the view state.
		link, ok := doc.Find("button").Attr("hx-get")
		if !ok {
			t.Errorf("expected button to have hx-get attribute")
		}
		if link != "/examples/templ/click-to-edit/edit" {
			t.Errorf("expected button to have hx-get attribute with value '/examples/click-to-edit/templ/edit' got %s", link)
		}

		// Should include form data in view response
		if !strings.Contains(doc.Text(), "John") {
			t.Errorf("expected response to contain 'John' got %s", doc.Text())
		}
		if !strings.Contains(doc.Text(), "Smith") {
			t.Errorf("expected response to contain 'Smith' got %s", doc.Text())
		}
		if !strings.Contains(doc.Text(), "john@smith.com") {
			t.Errorf("expected response to contain 'john@smith.com' got %s", doc.Text())
		}
	})

	t.Run("should return error for invalid email", func(t *testing.T) {
		w := post("John", "Smith", "bad_email")

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 422 {
			t.Fatalf("expected status code 422 got %d", res.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		emailInput := doc.Find("input[name=email]")
		if emailInput.Length() == 0 {
			t.Fatalf("expected response to contain input[name=email]")
		}

		invalid, exists := emailInput.Attr("aria-invalid")
		if !exists {
			t.Fatalf("expected email to have aria-invalid attribute")
		}
		if invalid != "true" {
			t.Errorf("expected email to have aria-invalid attribute with value 'true' got %s", invalid)
		}
		emailError := emailInput.Next().First()
		if emailError.Length() != 1 {
			t.Fatalf("expected response to contain email error message")
		}
		if emailError.First().Nodes[0].Data != "small" {
			t.Fatalf("expected response to contain email error message")
		}
		if emailError.Text() != "Invalid email address" {
			t.Errorf("expected response to contain 'Invalid email address' got %s", emailError.Text())
		}
	})
}
