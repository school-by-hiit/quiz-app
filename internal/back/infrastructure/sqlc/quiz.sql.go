// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: quiz.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const activateOnlyVersion = `-- name: ActivateOnlyVersion :exec
UPDATE quiz
SET active = 0
WHERE filename = ?
  AND version <> ?
`

type ActivateOnlyVersionParams struct {
	Filename string `db:"filename"`
	Version  int    `db:"version"`
}

func (q *Queries) ActivateOnlyVersion(ctx context.Context, arg ActivateOnlyVersionParams) error {
	_, err := q.db.ExecContext(ctx, activateOnlyVersion, arg.Filename, arg.Version)
	return err
}

const countAllActiveQuiz = `-- name: CountAllActiveQuiz :one
SELECT COUNT(1)
FROM quiz q
WHERE q.active = 1
`

func (q *Queries) CountAllActiveQuiz(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllActiveQuiz)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countAllActiveQuizForUser = `-- name: CountAllActiveQuizForUser :one
SELECT COUNT(1)
FROM quiz q
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.active = 1
  AND u.id = ?
`

func (q *Queries) CountAllActiveQuizForUser(ctx context.Context, id string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllActiveQuizForUser, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOrReplaceAnswer = `-- name: CreateOrReplaceAnswer :exec
REPLACE INTO quiz_answer (sha1, content, valid)
VALUES (?, ?, ?)
`

type CreateOrReplaceAnswerParams struct {
	Sha1    string `db:"sha1"`
	Content string `db:"content"`
	Valid   bool   `db:"valid"`
}

func (q *Queries) CreateOrReplaceAnswer(ctx context.Context, arg CreateOrReplaceAnswerParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceAnswer, arg.Sha1, arg.Content, arg.Valid)
	return err
}

const createOrReplaceQuestion = `-- name: CreateOrReplaceQuestion :exec
REPLACE INTO quiz_question (sha1, position, content)
VALUES (?, ?, ?)
`

type CreateOrReplaceQuestionParams struct {
	Sha1     string `db:"sha1"`
	Position int64  `db:"position"`
	Content  string `db:"content"`
}

func (q *Queries) CreateOrReplaceQuestion(ctx context.Context, arg CreateOrReplaceQuestionParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceQuestion, arg.Sha1, arg.Position, arg.Content)
	return err
}

const createOrReplaceQuiz = `-- name: CreateOrReplaceQuiz :exec
REPLACE INTO quiz (sha1, name, filename, version, duration, created_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateOrReplaceQuizParams struct {
	Sha1      string `db:"sha1"`
	Name      string `db:"name"`
	Filename  string `db:"filename"`
	Version   int    `db:"version"`
	Duration  int    `db:"duration"`
	CreatedAt string `db:"created_at"`
}

func (q *Queries) CreateOrReplaceQuiz(ctx context.Context, arg CreateOrReplaceQuizParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceQuiz,
		arg.Sha1,
		arg.Name,
		arg.Filename,
		arg.Version,
		arg.Duration,
		arg.CreatedAt,
	)
	return err
}

const findAllActiveQuiz = `-- name: FindAllActiveQuiz :many
SELECT sha1, q.name, filename, version, q.active, created_at, duration, qcv.class_uuid, quiz_sha1, uuid, sc.name, id, login, u.name, picture, u.active, role_id, u.class_uuid
FROM quiz q
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.active = 1
    AND u.id = ''
   OR u.id = ?
LIMIT ? OFFSET ?
`

type FindAllActiveQuizParams struct {
	ID     string `db:"id"`
	Limit  int64  `db:"limit"`
	Offset int64  `db:"offset"`
}

type FindAllActiveQuizRow struct {
	Sha1        string    `db:"sha1"`
	Name        string    `db:"name"`
	Filename    string    `db:"filename"`
	Version     int       `db:"version"`
	Active      bool      `db:"active"`
	CreatedAt   string    `db:"created_at"`
	Duration    int       `db:"duration"`
	ClassUuid   uuid.UUID `db:"class_uuid"`
	QuizSha1    string    `db:"quiz_sha1"`
	Uuid        uuid.UUID `db:"uuid"`
	Name_2      string    `db:"name_2"`
	ID          string    `db:"id"`
	Login       string    `db:"login"`
	Name_3      string    `db:"name_3"`
	Picture     string    `db:"picture"`
	Active_2    bool      `db:"active_2"`
	RoleID      int8      `db:"role_id"`
	ClassUuid_2 uuid.UUID `db:"class_uuid_2"`
}

func (q *Queries) FindAllActiveQuiz(ctx context.Context, arg FindAllActiveQuizParams) ([]FindAllActiveQuizRow, error) {
	rows, err := q.db.QueryContext(ctx, findAllActiveQuiz, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindAllActiveQuizRow{}
	for rows.Next() {
		var i FindAllActiveQuizRow
		if err := rows.Scan(
			&i.Sha1,
			&i.Name,
			&i.Filename,
			&i.Version,
			&i.Active,
			&i.CreatedAt,
			&i.Duration,
			&i.ClassUuid,
			&i.QuizSha1,
			&i.Uuid,
			&i.Name_2,
			&i.ID,
			&i.Login,
			&i.Name_3,
			&i.Picture,
			&i.Active_2,
			&i.RoleID,
			&i.ClassUuid_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllQuizSessions = `-- name: FindAllQuizSessions :many
SELECT quiz_sha1, quiz_name, quiz_filename, quiz_version, quiz_duration, quiz_created_at, session_uuid, user_id, user_name, user_picture, remaining_sec, checked_answers, results
FROM quiz_session_view
LIMIT ? OFFSET ?
`

type FindAllQuizSessionsParams struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func (q *Queries) FindAllQuizSessions(ctx context.Context, arg FindAllQuizSessionsParams) ([]QuizSessionView, error) {
	rows, err := q.db.QueryContext(ctx, findAllQuizSessions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuizSessionView{}
	for rows.Next() {
		var i QuizSessionView
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizName,
			&i.QuizFilename,
			&i.QuizVersion,
			&i.QuizDuration,
			&i.QuizCreatedAt,
			&i.SessionUuid,
			&i.UserID,
			&i.UserName,
			&i.UserPicture,
			&i.RemainingSec,
			&i.CheckedAnswers,
			&i.Results,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllQuizSessionsForUser = `-- name: FindAllQuizSessionsForUser :many
SELECT qsv.quiz_sha1, qsv.quiz_name, qsv.quiz_filename, qsv.quiz_version, qsv.quiz_duration, qsv.quiz_created_at, qsv.session_uuid, qsv.user_id, qsv.user_name, qsv.user_picture, qsv.remaining_sec, qsv.checked_answers, qsv.results
FROM quiz_session_view qsv
         JOIN quiz_class_visibility qcv ON qsv.quiz_sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid AND qsv.user_id = u.id
WHERE user_id = ?
LIMIT ? OFFSET ?
`

type FindAllQuizSessionsForUserParams struct {
	UserID string `db:"user_id"`
	Limit  int64  `db:"limit"`
	Offset int64  `db:"offset"`
}

func (q *Queries) FindAllQuizSessionsForUser(ctx context.Context, arg FindAllQuizSessionsForUserParams) ([]QuizSessionView, error) {
	rows, err := q.db.QueryContext(ctx, findAllQuizSessionsForUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuizSessionView{}
	for rows.Next() {
		var i QuizSessionView
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizName,
			&i.QuizFilename,
			&i.QuizVersion,
			&i.QuizDuration,
			&i.QuizCreatedAt,
			&i.SessionUuid,
			&i.UserID,
			&i.UserName,
			&i.UserPicture,
			&i.RemainingSec,
			&i.CheckedAnswers,
			&i.Results,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findQuizByFilenameAndLatestVersion = `-- name: FindQuizByFilenameAndLatestVersion :one
SELECT sha1, name, filename, version, active, created_at, duration
FROM quiz
WHERE filename = ?
ORDER BY version DESC
LIMIT 1
`

func (q *Queries) FindQuizByFilenameAndLatestVersion(ctx context.Context, filename string) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, findQuizByFilenameAndLatestVersion, filename)
	var i Quiz
	err := row.Scan(
		&i.Sha1,
		&i.Name,
		&i.Filename,
		&i.Version,
		&i.Active,
		&i.CreatedAt,
		&i.Duration,
	)
	return i, err
}

const findQuizFullBySha1 = `-- name: FindQuizFullBySha1 :many
SELECT q.sha1       AS quiz_sha1,
       q.filename   AS quiz_filename,
       q.name       AS quiz_name,
       q.version    AS quiz_version,
       q.created_at AS quiz_created_at,
       q.duration   AS quiz_duration,
       q.active     AS quiz_active,
       qq.sha1      AS question_sha1,
       qq.content   AS question_content,
       qq.position  AS question_position,
       qa.sha1      AS answer_sha1,
       qa.content   AS answer_content,
       qa.valid     AS answer_valid
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question qq ON qq.sha1 = qqq.question_sha1
         JOIN quiz_question_answer qqa ON qq.sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
         JOIN quiz_class_visibility qcv ON q.sha1 = qcv.quiz_sha1
         JOIN student_class sc ON sc.uuid = qcv.class_uuid
         JOIN user u ON sc.uuid = u.class_uuid
WHERE q.sha1 = ?
    AND u.id = ''
   OR u.id = ?
`

type FindQuizFullBySha1Params struct {
	Sha1 string `db:"sha1"`
	ID   string `db:"id"`
}

type FindQuizFullBySha1Row struct {
	QuizSha1         string `db:"quiz_sha1"`
	QuizFilename     string `db:"quiz_filename"`
	QuizName         string `db:"quiz_name"`
	QuizVersion      int    `db:"quiz_version"`
	QuizCreatedAt    string `db:"quiz_created_at"`
	QuizDuration     int    `db:"quiz_duration"`
	QuizActive       bool   `db:"quiz_active"`
	QuestionSha1     string `db:"question_sha1"`
	QuestionContent  string `db:"question_content"`
	QuestionPosition int    `db:"question_position"`
	AnswerSha1       string `db:"answer_sha1"`
	AnswerContent    string `db:"answer_content"`
	AnswerValid      bool   `db:"answer_valid"`
}

func (q *Queries) FindQuizFullBySha1(ctx context.Context, arg FindQuizFullBySha1Params) ([]FindQuizFullBySha1Row, error) {
	rows, err := q.db.QueryContext(ctx, findQuizFullBySha1, arg.Sha1, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindQuizFullBySha1Row{}
	for rows.Next() {
		var i FindQuizFullBySha1Row
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizFilename,
			&i.QuizName,
			&i.QuizVersion,
			&i.QuizCreatedAt,
			&i.QuizDuration,
			&i.QuizActive,
			&i.QuestionSha1,
			&i.QuestionContent,
			&i.QuestionPosition,
			&i.AnswerSha1,
			&i.AnswerContent,
			&i.AnswerValid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findQuizSessionByUuid = `-- name: FindQuizSessionByUuid :many
SELECT session_uuid, user_id, remaining_sec, quiz_sha1, quiz_name, quiz_duration, checked_answers, results, question_sha1, question_content, question_position, answer_sha1, answer_content, answer_checked, answer_valid
FROM quiz_session_detail_view
WHERE session_uuid = ?
`

func (q *Queries) FindQuizSessionByUuid(ctx context.Context, sessionUuid uuid.UUID) ([]QuizSessionDetailView, error) {
	rows, err := q.db.QueryContext(ctx, findQuizSessionByUuid, sessionUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuizSessionDetailView{}
	for rows.Next() {
		var i QuizSessionDetailView
		if err := rows.Scan(
			&i.SessionUuid,
			&i.UserID,
			&i.RemainingSec,
			&i.QuizSha1,
			&i.QuizName,
			&i.QuizDuration,
			&i.CheckedAnswers,
			&i.Results,
			&i.QuestionSha1,
			&i.QuestionContent,
			&i.QuestionPosition,
			&i.AnswerSha1,
			&i.AnswerContent,
			&i.AnswerChecked,
			&i.AnswerValid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const linkAnswer = `-- name: LinkAnswer :exec
REPLACE INTO quiz_question_answer (question_sha1, answer_sha1)
VALUES (?, ?)
`

type LinkAnswerParams struct {
	QuestionSha1 string `db:"question_sha1"`
	AnswerSha1   string `db:"answer_sha1"`
}

func (q *Queries) LinkAnswer(ctx context.Context, arg LinkAnswerParams) error {
	_, err := q.db.ExecContext(ctx, linkAnswer, arg.QuestionSha1, arg.AnswerSha1)
	return err
}

const linkQuestion = `-- name: LinkQuestion :exec
REPLACE INTO quiz_question_quiz (quiz_sha1, question_sha1)
VALUES (?, ?)
`

type LinkQuestionParams struct {
	QuizSha1     string `db:"quiz_sha1"`
	QuestionSha1 string `db:"question_sha1"`
}

func (q *Queries) LinkQuestion(ctx context.Context, arg LinkQuestionParams) error {
	_, err := q.db.ExecContext(ctx, linkQuestion, arg.QuizSha1, arg.QuestionSha1)
	return err
}
