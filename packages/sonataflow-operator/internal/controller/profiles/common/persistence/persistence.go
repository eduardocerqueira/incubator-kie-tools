// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package persistence

import (
	"github.com/magiconair/properties"

	operatorapi "github.com/apache/incubator-kie-tools/packages/sonataflow-operator/api/v1alpha08"
)

const (
	QuarkusFlywayMigrateAtStart         string = "quarkus.flyway.migrate-at-start"
	QuarkusDatasourceDBKind             string = "quarkus.datasource.db-kind"
	QuarkusDatasourceJDBCURL            string = "quarkus.datasource.jdbc.url"
	KogitoPersistenceType               string = "kogito.persistence.type"
	JDBCPersistenceType                 string = "jdbc"
	KogitoPersistenceQueryTimeoutMillis string = "kogito.persistence.query.timeout.millis"
	KogitoPersistenceProtoMarshaller    string = "kogito.persistence.proto.marshaller"
	PostgreSQLDBKind                    string = "postgresql"
)

// ResolveWorkflowPersistenceProperties returns the set of application properties required for the workflow persistence.
// Never nil.
func ResolveWorkflowPersistenceProperties(workflow *operatorapi.SonataFlow, platform *operatorapi.SonataFlowPlatform) (*properties.Properties, error) {
	if UsesPostgreSQLPersistence(workflow, platform) {
		return GetPostgreSQLWorkflowProperties(workflow), nil
	}
	return properties.NewProperties(), nil
}
