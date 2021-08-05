package forms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	jtd "github.com/jsontypedef/json-typedef-go"
)

// RESTTransport is the structure for the HTTP REST transport
type RESTTransport struct {
	Addr    string
	Service *Service
}

// NewRESTTransport creates a new transport using the REST protocol
func NewRESTTransport(addr string) *RESTTransport {
	return &RESTTransport{
		Addr: addr,
	}
}

// ServeHTTP handles HTTP requests
func (t RESTTransport) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//Check if method is supported
	if req.Method != "GET" && req.Method != "POST" && req.Method != "PUT" && req.Method != "DELETE" && req.Method != "OPTIONS" {
		respond(rw, "method not allowed", http.StatusMethodNotAllowed)
		return

	}

	if req.Method == "OPTIONS" {
		respondCORS(rw, http.StatusOK)
		return
	}

	p := req.URL.Path
	m := req.Method

	var formId string

	switch {
	case match(p, "/"):
		switch {
		case m == "GET":
			t.home(rw, req)
		default:
			http.NotFound(rw, req)
			return
		}
	case match(p, "/forms/"):
		switch {
		case m == "GET":
			t.getForms(rw, req)
			return
		case m == "POST":
			t.createForm(rw, req)
			return
		default:
			http.NotFound(rw, req)
			return
		}
	case match(p, "/forms/([^/]+)", &formId):
		switch {
		case m == "GET":
			t.getForm(formId, rw, req)
			return
		case m == "DELETE":
			t.deleteForm(formId, rw, req)
			return
		case m == "POST":
			t.postFormResponse(formId, rw, req)
			return
		case m == "PUT":
			t.updateForm(formId, rw, req)
			return
		default:
			http.NotFound(rw, req)
			return
		}
	case match(p, "/forms/([^/]+)/responses", &formId):
		switch {
		case m == "GET":
			t.getFormResponses(formId, rw, req)
			return
		case m == "POST":
			t.postFormResponse(formId, rw, req)
			return
		default:
			http.NotFound(rw, req)
			return
		}
	default:
		http.NotFound(rw, req)
		return
	}
}

// Run runs the HTTP server
func (t RESTTransport) Run(s *Service) {
	t.Service = s
	//http.Handle("/forms/", t)
	if err := http.ListenAndServe(t.Addr, t); err != nil {
		panic(err)
	}
}

// home handles the home endpoint ("/")
func (t RESTTransport) home(rw http.ResponseWriter, req *http.Request) {
	respond(rw, "welcome to coopgo/forms", http.StatusOK)
}

// getForms handles the endpoint to retrieve the list of forms
func (t RESTTransport) getForms(rw http.ResponseWriter, req *http.Request) {
	forms, err := t.Service.GetForms()
	if err != nil {
		respond(rw, errJson(err), http.StatusInternalServerError)
		return
	}
	formsResponse := []RESTFormDefinition{}
	for _, f := range forms {
		formsResponse = append(formsResponse, RESTForm(f))
	}
	respond(rw, formsResponse, http.StatusOK)
}

// createForms handles the endpoint to create a new form
func (t RESTTransport) createForm(rw http.ResponseWriter, req *http.Request) {
	var formDefinition RESTFormDefinition

	decoder := json.NewDecoder(req.Body)
	// decoder.DisallowUnknownFields()
	if err := decoder.Decode(&formDefinition); err != nil {
		respond(rw, "error decoding json", http.StatusBadRequest)
		return
	}

	if formDefinition.ID == "" {
		respond(rw, "missing form ID", http.StatusBadRequest)
		return
	}
	var form Form

	switch {
	case formDefinition.Type == "Unstructured":
		form = NewUnstructuredForm(formDefinition.ID)
		if err := t.Service.AddForm(form); err != nil {
			respond(rw, errJson(err), http.StatusInternalServerError)
			return
		}
	case formDefinition.Type == "Structured":
		if formDefinition.Schema == nil {
			respond(rw, "schema is mandatory for structured forms", http.StatusBadRequest)
			return
		}
		form = NewStructuredForm(formDefinition.ID, formDefinition.Schema)
		if err := t.Service.AddForm(form); err != nil {
			respond(rw, errJson(err), http.StatusInternalServerError)
			return
		}
	default:
		respond(rw, "form type unknown", http.StatusBadRequest)
		return
	}
	t.Service.AddForm(form)
	respond(rw, RESTForm(form), http.StatusOK)
}

// getForm handles the endpoint to retrieve a form from its ID
func (t RESTTransport) getForm(formId string, rw http.ResponseWriter, req *http.Request) {
	form, err := t.Service.GetForm(formId)
	if err != nil {
		respond(rw, errJson(err), http.StatusNotFound)
		return
	}

	respond(rw, RESTForm(form), http.StatusOK)
}

// deleteForm handles the endpoint for form removal with the givern formID
func (t RESTTransport) deleteForm(formId string, rw http.ResponseWriter, req *http.Request) {
	err := t.Service.DeleteForm(formId)
	if err != nil {
		respond(rw, errJson(err), http.StatusNotFound)
		return
	}

	respond(rw, "", http.StatusOK)
}

// updateForm updates a form with the given ID
func (t RESTTransport) updateForm(formId string, rw http.ResponseWriter, req *http.Request) {
	fmt.Println("UPDATE FORM ", formId)
}

// postFormResponse creates a new response for the form with the given ID
func (t RESTTransport) postFormResponse(formId string, rw http.ResponseWriter, req *http.Request) {
	var response Response

	decoder := json.NewDecoder(req.Body)
	// decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		respond(rw, errJson(err), http.StatusBadRequest)
		return
	}
	if err := t.Service.SubmitResponse(formId, response); err != nil {
		respond(rw, errJson(err), http.StatusBadRequest)
		return
	}

	respond(rw, response, http.StatusOK)
}

// getFormResponses retrieve form responses for the given form id
func (t RESTTransport) getFormResponses(formId string, rw http.ResponseWriter, req *http.Request) {
	responses, err := t.Service.GetFormResponses(formId)
	if err != nil {
		respond(rw, errJson(err), http.StatusBadRequest)
		return
	}
	respond(rw, responses, http.StatusOK)
}

// RESTFormDefintion is the form form definition
//
// The type of the form is "Unstructured" if there is no specific schema, or structured. Default value (if unknown or any other value) is "Unstructured".
type RESTFormDefinition struct {
	ID                 string                  `json:"id"`
	Type               string                  `json:"type"` // "Unstructured" or "Structured"
	Schema             *jtd.Schema             `json:"schema,omitempty"`
	AdditionalBackends []RESTBackendDefinition `json:"additional_backends,omitempty"`
}

func RESTForm(f Form) RESTFormDefinition {
	result := RESTFormDefinition{
		ID:                 f.ID(),
		Type:               f.Type(),
		AdditionalBackends: []RESTBackendDefinition{},
	}

	if f.Type() == "Structured" {
		result.Schema = f.GetSchema()
	}

	for _, b := range f.GetAdditionalBackends() {
		result.AdditionalBackends = append(result.AdditionalBackends, RESTBackend(b))
	}

	return result
}

// RESTBackendDefinition defines the storage type and backend configuration
//
// Configuration keys depend on the storage type
type RESTBackendDefinition struct {
	StorageType   string
	Configuration map[string]string
}

// RESTBackend transforms a Backend to the RESTBackendDefinition type for correct JSON marshalling
func RESTBackend(b Backend) RESTBackendDefinition {
	return RESTBackendDefinition{
		StorageType: b.Type(),
	}
}

// match reports whether path matches regex ^pattern$, and if it matches,
// assigns any capture groups to the *string or *int vars.
func match(path, pattern string, vars ...interface{}) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
		default:
			panic("vars must be *string or *int")
		}
	}
	return true
}

// respond is a facility function to create HTTP responses
func respond(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	json, _ := json.Marshal(body)
	_, _ = w.Write(json)
}

// respondCORS is a facility function to create HTTP response for to enable CORS
func respondCORS(w http.ResponseWriter, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Origin,Access,token,Authorization,appid")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")

	w.WriteHeader(code)
}

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}

//errJson transform an error to a JSON response
func errJson(err error) string {
	return fmt.Sprintf(`{error: "%s"}`, err)
}
