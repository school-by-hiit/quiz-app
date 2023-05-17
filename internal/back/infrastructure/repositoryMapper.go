/*
 * Copyright (c) 2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package infrastructure

import (
	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

func toDomain(entity sqlc.Quiz) domain.Quiz {
	return domain.Quiz{
		Sha1:      entity.Sha1,
		Filename:  entity.Filename,
		Name:      entity.Name,
		Version:   int(entity.Version),
		Active:    intToBool(entity.Active),
		CreatedAt: entity.CreatedAt,
	}
}

func toDomainArray(entities []sqlc.Quiz) []domain.Quiz {
	domains := make([]domain.Quiz, len(entities))

	for i, entity := range entities {
		domains[i] = toDomain(entity)
	}

	return domains
}

func intToBool(value int64) bool {
	if value == 1 {
		return true
	}

	return false
}

func boolToInt(value bool) int64 {
	if value {
		return 1
	}

	return 0
}
