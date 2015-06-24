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

type VisionPrescription struct {
	Id                    string                                `json:"-" bson:"_id"`
	Identifier            []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DateWritten           *FHIRDateTime                         `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Patient               *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Prescriber            *Reference                            `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Encounter             *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReasonCodeableConcept *CodeableConcept                      `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []VisionPrescriptionDispenseComponent `bson:"dispense,omitempty" json:"dispense,omitempty"`
}

type VisionPrescriptionDispenseComponent struct {
	Product   *Coding   `bson:"product,omitempty" json:"product,omitempty"`
	Eye       string    `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere    *float64  `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder  *float64  `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis      *int32    `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism     *float64  `bson:"prism,omitempty" json:"prism,omitempty"`
	Base      string    `bson:"base,omitempty" json:"base,omitempty"`
	Add       *float64  `bson:"add,omitempty" json:"add,omitempty"`
	Power     *float64  `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve *float64  `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter  *float64  `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration  *Quantity `bson:"duration,omitempty" json:"duration,omitempty"`
	Color     string    `bson:"color,omitempty" json:"color,omitempty"`
	Brand     string    `bson:"brand,omitempty" json:"brand,omitempty"`
	Notes     string    `bson:"notes,omitempty" json:"notes,omitempty"`
}

type VisionPrescriptionBundle struct {
	Type         string                          `json:"resourceType,omitempty"`
	Title        string                          `json:"title,omitempty"`
	Id           string                          `json:"id,omitempty"`
	Updated      time.Time                       `json:"updated,omitempty"`
	TotalResults int                             `json:"totalResults,omitempty"`
	Entry        []VisionPrescriptionBundleEntry `json:"entry,omitempty"`
	Category     VisionPrescriptionCategory      `json:"category,omitempty"`
}

type VisionPrescriptionBundleEntry struct {
	Title    string                     `json:"title,omitempty"`
	Id       string                     `json:"id,omitempty"`
	Content  VisionPrescription         `json:"content,omitempty"`
	Category VisionPrescriptionCategory `json:"category,omitempty"`
}

type VisionPrescriptionCategory struct {
	Term   string `json:"term,omitempty"`
	Label  string `json:"label,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

func (resource *VisionPrescription) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		VisionPrescription
	}{
		ResourceType:       "VisionPrescription",
		VisionPrescription: *resource,
	}
	return json.Marshal(x)
}
