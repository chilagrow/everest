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

// ListDataStoreExtensions lists extensions of the data store.
func (e *EverestServer) ListDataStoreExtensions(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// CreateDataStoreExtension creates an extension for the data store in the Everest server.
func (e *EverestServer) CreateDataStoreExtension(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// DeleteDataStoreExtension deletes an extension of the data store in the Everest server.
func (e *EverestServer) DeleteDataStoreExtension(c echo.Context, cluster, namespace, dataStore, extension string) error {
	return errors.New("not implemented")
}

// GetDataStoreExtension gets an extension of the data store in the Everest server.
func (e *EverestServer) GetDataStoreExtension(c echo.Context, cluster, namespace, dataStore, extension string) error {
	return errors.New("not implemented")
}

// UpdateDataStoreExtension updates an extension of the data store in the Everest server.
func (e *EverestServer) UpdateDataStoreExtension(c echo.Context, cluster, namespace, dataStore, extension string) error {
	return errors.New("not implemented")
}

// UpdateProviderExtensions lists extensions of the provider in the Everest server.
func (e *EverestServer) ListProviderExtensions(c echo.Context, cluster, namespace, dataStore string) error {
	return errors.New("not implemented")
}

// GetExtensionSchema gets the schema of an extension in the Everest server.
func (e *EverestServer) GetExtensionSchema(c echo.Context, cluster, namespace, dataStore, extension string) error {
	return errors.New("not implemented")
}
