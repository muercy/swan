package mesos

import (
	"encoding/json"
	"sync"

	"github.com/Dataman-Cloud/swan/mesosproto"
)

type Agent struct {
	id       string
	hostname string
	attrs    []*mesosproto.Attribute

	sync.RWMutex
	offers map[string]*mesosproto.Offer
}

func newAgent(id, hostname string, attrs []*mesosproto.Attribute) *Agent {
	return &Agent{
		id:       id,
		hostname: hostname,
		offers:   make(map[string]*mesosproto.Offer),
		attrs:    attrs,
	}
}

func (s *Agent) MarshalJSON() ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	m := map[string]interface{}{
		"id":         s.id,
		"hostname":   s.hostname,
		"attributes": s.attrs,
		"offers":     s.offers,
	}
	return json.Marshal(m)
}

func (s *Agent) addOffer(offer *mesosproto.Offer) {
	s.Lock()
	s.offers[offer.Id.GetValue()] = offer
	s.Unlock()
}

func (s *Agent) removeOffer(offerID string) bool {
	s.Lock()
	defer s.Unlock()

	found := false
	if _, found = s.offers[offerID]; found {
		delete(s.offers, offerID)
	}
	return found
}

func (s *Agent) empty() bool {
	s.RLock()
	defer s.RUnlock()

	return len(s.offers) == 0
}

func (s *Agent) getOffer(offerId string) *mesosproto.Offer {
	s.RLock()
	defer s.RUnlock()

	offer, ok := s.offers[offerId]
	if !ok {
		return nil
	}

	return offer
}

func (s *Agent) getOffers() map[string]*mesosproto.Offer {
	s.RLock()
	defer s.RUnlock()

	return s.offers
}

func (s *Agent) Resources() (cpus, mem, disk float64, ports []uint64) {
	for _, offer := range s.getOffers() {
		for _, resource := range offer.Resources {
			if *resource.Name == "cpus" {
				cpus += *resource.Scalar.Value
			}

			if *resource.Name == "mem" {
				mem += *resource.Scalar.Value
			}

			if *resource.Name == "disk" {
				disk += *resource.Scalar.Value
			}

			if *resource.Name == "ports" {
				for _, rang := range resource.GetRanges().GetRange() {
					for i := rang.GetBegin(); i <= rang.GetEnd(); i++ {
						ports = append(ports, i)
					}
				}
			}

		}
	}

	return
}

func (s *Agent) Attributes() map[string]string {
	attrs := make(map[string]string)

	for _, offer := range s.getOffers() {
		for _, attr := range offer.Attributes {
			if attr.GetType() == mesosproto.Value_TEXT {
				attrs[attr.GetName()] = attr.GetText().GetValue()
			}
		}
	}

	return attrs
}
