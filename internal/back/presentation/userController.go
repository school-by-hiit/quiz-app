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

package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *ApiController) userList(ctx *gin.Context) {
	users, err := c.userService.FindAllUser(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	dtos := make([]*User, len(users))

	for i, user := range users {
		dto := User{}
		dtos[i] = dto.fromDomain(user)
	}

	ctx.JSON(http.StatusOK, dtos)
}

func (c *ApiController) deactivateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.userService.DeactivateUser(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deactivated"})
}

func (c *ApiController) activateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.userService.ActivateUser(ctx, id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user activated"})
}

func (c *ApiController) assignUserToClass(ctx *gin.Context) {
	id := ctx.Param("id")

	classIdStr := ctx.Param("uuid")

	classId, err := uuid.Parse(classIdStr)
	if err != nil {
		handleHttpError(ctx, http.StatusBadRequest, "invalid classId")
		return
	}

	err = c.userService.AssignUserToClass(ctx, id, classId)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user assigned"})
}
