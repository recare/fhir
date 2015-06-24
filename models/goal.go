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

type Goal struct {
	Id          string                 `json:"-" bson:"_id"`
	Identifier  []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient     *Reference             `bson:"patient,omitempty" json:"patient,omitempty"`
	TargetDate  *FHIRDateTime          `bson:"targetDate,omitempty" json:"targetDate,omitempty"`
	Description string                 `bson:"description,omitempty" json:"description,omitempty"`
	Status      string                 `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate  *FHIRDateTime          `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	Author      *Reference             `bson:"author,omitempty" json:"author,omitempty"`
	Priority    *CodeableConcept       `bson:"priority,omitempty" json:"priority,omitempty"`
	Concern     []Reference            `bson:"concern,omitempty" json:"concern,omitempty"`
	Notes       string                 `bson:"notes,omitempty" json:"notes,omitempty"`
	Outcome     []GoalOutcomeComponent `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

type GoalOutcomeComponent struct {
	ResultCodeableConcept *CodeableConcept `bson:"resultCodeableConcept,omitempty" json:"resultCodeableConcept,omitempty"`
	ResultReference       *Reference       `bson:"resultReference,omitempty" json:"resultReference,omitempty"`
}

type GoalBundle struct {
	Type         string            `json:"resourceType,omitempty"`
	Title        string            `json:"title,omitempty"`
	Id           string            `json:"id,omitempty"`
	Updated      time.Time         `json:"updated,omitempty"`
	TotalResults int               `json:"totalResults,omitempty"`
	Entry        []GoalBundleEntry `json:"entry,omitempty"`
	Category     GoalCategory      `json:"category,omitempty"`
}

type GoalBundleEntry struct {
	Title    string       `json:"title,omitempty"`
	Id       string       `json:"id,omitempty"`
	Content  Goal         `json:"content,omitempty"`
	Category GoalCategory `json:"category,omitempty"`
}

type GoalCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *Goal) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		Goal
	}{
		ResourceType: "Goal",
		Goal:         *resource,
	}
	return json.Marshal(x)
}
