package handler

import (
	"net/http"
	helper "pdam/Helper"
	"pdam/service"

	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	groupService service.GroupService
}

func NewGroupHandler(groupService service.GroupService) *groupHandler {
	return &groupHandler{groupService}
}

func (h *groupHandler) GetGroups(c *gin.Context) {
	groups, err := h.groupService.GetGroups()
	if err != nil {
		response := helper.ApiResponse("error to get groups", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of groups", 200, "success", groups)

	c.JSON(http.StatusOK, response)
}

func (h *groupHandler) GetGroupByID(c *gin.Context) {
	var input helper.InputKodeGetGroup
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("error to get group", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	group, err := h.groupService.GetGroupByID(input)
	if err != nil {
		response := helper.ApiResponse("error to get groups", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of groups", 200, "success", group)

	c.JSON(http.StatusOK, response)
}

func (h *groupHandler) GetGroupByNama(c *gin.Context) {
	var input helper.InputNamaGetGroup
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("error to get group", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	group, err := h.groupService.GetGroupByNama(input)
	if err != nil {
		response := helper.ApiResponse("error to get groups", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of groups", 200, "success", group)

	c.JSON(http.StatusOK, response)
}

func (h *groupHandler) CreateGroup(c *gin.Context) {
	var input helper.Inputgroup

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errormessagge := gin.H{"errors": err}

		response := helper.ApiResponse("Create group failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newGroups, err := h.groupService.SaveGroup(input)
	if err != nil {
		response := helper.ApiResponse("Create group failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Group successfully created", 200, "success", newGroups)

	c.JSON(http.StatusOK, response)
}

func (h *groupHandler) UpdateGroup(c *gin.Context) {
	var input helper.Inputgroup

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errormessagge := gin.H{"errors": err}

		response := helper.ApiResponse("update group failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	editGroup, err := h.groupService.EditGroup(input)
	if err != nil {
		response := helper.ApiResponse("update group failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Group successfully updated", 200, "success", editGroup)

	c.JSON(http.StatusOK, response)
}

func (h *groupHandler) DeleteGroup(c *gin.Context) {
	var input helper.InputKodeGetGroup
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("error to delete group", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.groupService.DeleteGroup(input)
	if err != nil {
		response := helper.ApiResponse("delete group failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Group successfully deleted", 200, "success", nil)

	c.JSON(http.StatusOK, response)
}
