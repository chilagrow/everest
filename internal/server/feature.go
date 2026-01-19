// everest
// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server contains the API server implementation.
package server

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// ListDataStoreFeatures lists features of the data store.
func (e *EverestServer) ListDataStoreFeatures(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// CreateDataStoreFeature creates a feature for the data store in the Everest server.
func (e *EverestServer) CreateDataStoreFeature(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// DeleteDataStoreFeature deletes a feature of the data store in the Everest server.
func (e *EverestServer) DeleteDataStoreFeature(c echo.Context, cluster, namespace, dataStore, feature string) error {
	return errors.New("not implemented")
}

// GetDataStoreFeature gets a feature of the data store in the Everest server.
func (e *EverestServer) GetDataStoreFeature(c echo.Context, cluster, namespace, dataStore, feature string) error {
	return errors.New("not implemented")
}

// UpdateDataStoreFeature updates a feature of the data store in the Everest server.
func (e *EverestServer) UpdateDataStoreFeature(c echo.Context, cluster, namespace, dataStore, feature string) error {
	return errors.New("not implemented")
}

// UpdateProviderFeatures lists features of the provider in the Everest server.
func (e *EverestServer) ListProviderFeatures(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// GetFeatureSchema gets the schema of a feature in the Everest server.
func (e *EverestServer) GetFeatureSchema(c echo.Context, cluster, namespace, dataStore, feature string) error {
	return errors.New("not implemented")
}
