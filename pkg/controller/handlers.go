package controller

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func (h *Controller) getLists(c *gin.Context) {
	var params todo.TodoParams
	params.Date = c.Query("date")
	params.Status = c.Query("status")

	lists, err := h.services.TodoList.GetLists(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

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
