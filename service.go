package forms

import (
	"errors"
	"fmt"
	"sync"
)

// Service is the web service for the forms handling. It needs transports to communicate with the outside world, and backends to manage the data
//
// A service can have multiple transports and backends. For data consistency, in case of multiple backends, the first one in the list is the reference.
// If you mix backends with data retention (database for example) and backends without (email for example), set a backend with data retention as the first one in the list
type Service struct {
	Transports []Transport
	Backends   []Backend
}

// NewService creates a new service with transports and backends
func NewService(transports []Transport, backends []Backend) *Service {
	return &Service{
		Backends:   backends,
		Transports: transports,
	}
}

// Run runs the service, launching with backends and transports
func (s *Service) Run() {
	var wg sync.WaitGroup

	if len(s.Backends) == 0 || len(s.Transports) == 0 {
		panic(errors.New("You must define at least one backend and one transport"))
	}

	for _, t := range s.Transports {
		wg.Add(1)
		go func(t Transport) {
			defer wg.Done()
			t.Run(s)
		}(t)
	}

	wg.Wait()
}

// AddForm creates a new form
func (s *Service) AddForm(f Form) error {
	for _, b := range s.Backends {
		if err := b.AddForm(f); err != nil {
			return err
		}
	}
	return nil
}

// DeleteForm deletes a form
func (s *Service) DeleteForm(formId string) error {
	for _, b := range s.Backends {
		if err := b.DeleteForm(formId); err != nil {
			return err
		}
	}
	return nil
}

// GetForms retrieves the forms list
func (s *Service) GetForms() ([]Form, error) {
	if len(s.Backends) == 0 {
		return []Form{}, errors.New("not enough backends")
	}

	return s.Backends[0].GetForms()
}

// GetForm retrieves a single form given its ID
func (s *Service) GetForm(formId string) (Form, error) {
	if len(s.Backends) == 0 {
		return nil, errors.New("not enough backends")
	}

	form, err := s.Backends[0].GetForm(formId)

	if err != nil {
		return nil, err
	}

	return form, nil
}

// SubmitResponse sends a response to a form
func (s *Service) SubmitResponse(formId string, response Response) error {
	if len(s.Backends) == 0 {
		return errors.New("not enough backends")
	}

	form, err := s.Backends[0].GetForm(formId)

	if err != nil {
		return err
	}

	if !form.Validate(response) {
		return errors.New("bad response format")
	}

	for _, b := range s.Backends {
		if err := b.SubmitResponse(form, response); err != nil {
			return err
		}
	}

	additionalBackends, err := s.GetFormBackends(formId)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Form backends
	for _, b := range additionalBackends {
		if err := b.SubmitResponse(form, response); err != nil {
			return err
		}
	}
	return nil
}

// GetFormResponses sends a response to a form
func (s *Service) GetFormResponses(formId string) ([]Response, error) {
	if len(s.Backends) == 0 {
		return nil, errors.New("not enough backends")
	}

	responses, err := s.Backends[0].GetFormResponses(formId)

	if err != nil {
		return nil, err
	}

	return responses, nil
}

// AddFormBackend adds a backend to a form
func (s *Service) AddFormBackend(formid string, b Backend) error {
	if len(s.Backends) == 0 {
		return errors.New("not enough backends")
	}

	if err := s.Backends[0].AddFormBackend(formid, b); err != nil {
		return err
	}

	return nil
}

// AddFormBackend adds a backend to a form
func (s *Service) GetFormBackends(formid string) ([]Backend, error) {
	if len(s.Backends) == 0 {
		return nil, errors.New("not enough backends")
	}

	backends, err := s.Backends[0].GetFormBackends(formid)

	if err != nil {
		return nil, err
	}

	return backends, nil
}
