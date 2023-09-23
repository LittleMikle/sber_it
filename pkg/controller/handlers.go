package controller

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary CreateList
// @Tags create
// @Description create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "todolist info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/create [post]
func (h *Controller) createList(c *gin.Context) {
	var input todo.TodoList
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = validateInput(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get List
// @Tags get
// @Description get list
// @ID get-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists [get]
func (h *Controller) getLists(c *gin.Context) {
	var params todo.TodoParams
	var pageNum int
	params.Page = c.Query("page")
	params.Date = c.Query("date")
	params.Status = c.Query("status")

	if params.Page != "" {
		id, err := strconv.Atoi(params.Page)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}
		pageNum = id
	}

	lists, err := h.services.TodoList.GetLists(pageNum, params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Update List
// @Tags update
// @Description update list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body todo.UpdateListInput true "update todolist info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/:id [put]
func (h *Controller) updateList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary DeleteList
// @Tags delete
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/:id [delete]
func (h *Controller) deleteList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.services.TodoList.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
