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

package infra_repository

import (
	"context"
	"database/sql"

	"github.com/school-by-hiit/quizz-app/internal/back/domain/model"
	"github.com/school-by-hiit/quizz-app/internal/back/domain/repository"
	"github.com/school-by-hiit/quizz-app/internal/back/infrastructure/db"
	"github.com/school-by-hiit/quizz-app/internal/back/infrastructure/sqlc"
)

type QuizzDBRepository struct {
	repository.QuizzRepository

	c *sql.DB
	q *sqlc.Queries
}

func New() *QuizzDBRepository {
	connection := db.Connect(false, "data")
	db.New(connection).Migrate()

	return &QuizzDBRepository{q: sqlc.New(connection), c: connection}
}

func (r *QuizzDBRepository) Close() {
	r.c.Close()
}

func (r *QuizzDBRepository) Create(ctx context.Context, quizz model.Quizz) error {

	err := r.q.CreateOrReplaceQuizz(ctx, sqlc.CreateOrReplaceQuizzParams{
		Sha1:     quizz.Sha1,
		Name:     quizz.Name,
		Filename: quizz.Filename,
		Version:  1,
	})
	if err != nil {
		return err
	}

	for _, question := range quizz.Questions {
		err := r.q.CreateOrReplaceQuestion(ctx, sqlc.CreateOrReplaceQuestionParams{
			Sha1:    question.Sha1,
			Content: question.Content,
		})
		if err != nil {
			return err
		}

		err = r.q.LinkQuestion(ctx, sqlc.LinkQuestionParams{
			QuizzSha1:    quizz.Sha1,
			QuestionSha1: question.Sha1,
		})
		if err != nil {
			return err
		}

		for _, answer := range question.Answers {

			var valid int64
			if answer.Valid {
				valid = 1
			} else {
				valid = 0
			}

			err := r.q.CreateOrReplaceAnswer(ctx, sqlc.CreateOrReplaceAnswerParams{
				Sha1:    answer.Sha1,
				Content: answer.Content,
				Valid:   valid,
			})
			if err != nil {
				return err
			}

			err = r.q.LinkAnswer(ctx, sqlc.LinkAnswerParams{
				QuestionSha1: question.Sha1,
				AnswerSha1:   answer.Sha1,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
