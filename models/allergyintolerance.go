// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"time"
)

type AllergyIntolerance struct {
	Id            string                             `json:"-" bson:"_id"`
	Identifier    []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	RecordedDate  *FHIRDateTime                      `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder      *Reference                         `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Patient       *Reference                         `bson:"patient,omitempty" json:"patient,omitempty"`
	Reporter      *Reference                         `bson:"reporter,omitempty" json:"reporter,omitempty"`
	Substance     *CodeableConcept                   `bson:"substance,omitempty" json:"substance,omitempty"`
	Status        string                             `bson:"status,omitempty" json:"status,omitempty"`
	Criticality   string                             `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Type          string                             `bson:"type,omitempty" json:"type,omitempty"`
	Category      string                             `bson:"category,omitempty" json:"category,omitempty"`
	LastOccurence *FHIRDateTime                      `bson:"lastOccurence,omitempty" json:"lastOccurence,omitempty"`
	Comment       string                             `bson:"comment,omitempty" json:"comment,omitempty"`
	Event         []AllergyIntoleranceEventComponent `bson:"event,omitempty" json:"event,omitempty"`
}

type AllergyIntoleranceEventComponent struct {
	Substance     *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Certainty     string            `bson:"certainty,omitempty" json:"certainty,omitempty"`
	Manifestation []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description   string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset         *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Duration      *Quantity         `bson:"duration,omitempty" json:"duration,omitempty"`
	Severity      string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Comment       string            `bson:"comment,omitempty" json:"comment,omitempty"`
}

type AllergyIntoleranceBundle struct {
	Type         string                          `json:"resourceType,omitempty"`
	Title        string                          `json:"title,omitempty"`
	Id           string                          `json:"id,omitempty"`
	Updated      time.Time                       `json:"updated,omitempty"`
	TotalResults int                             `json:"totalResults,omitempty"`
	Entry        []AllergyIntoleranceBundleEntry `json:"entry,omitempty"`
	Category     AllergyIntoleranceCategory      `json:"category,omitempty"`
}

type AllergyIntoleranceBundleEntry struct {
	Title    string                     `json:"title,omitempty"`
	Id       string                     `json:"id,omitempty"`
	Content  AllergyIntolerance         `json:"content,omitempty"`
	Category AllergyIntoleranceCategory `json:"category,omitempty"`
}

type AllergyIntoleranceCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		AllergyIntolerance
	}{
		ResourceType:       "AllergyIntolerance",
		AllergyIntolerance: *resource,
	}
	return json.Marshal(x)
}
