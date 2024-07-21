package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	GetTaskResponse struct {
		ID       int64  `json:"id"`
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	CreateTaskResponse struct {
		ID int64 `json:"id"`
	}

	UpdateTaskRequest struct {
		Desc     string `json:"description"`
		Deadline int64  `json:"deadline"`
	}

	Task struct {
		ID       int64
		Desc     string
		Deadline int64
	}
)

var (
	taskIDCounter int64 = 1
	tasks               = make(map[int64]Task)
)

func CreateTask(c *fiber.Ctx) error {
	var req CreateTaskRequest

	if err := c.BodyParser(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	task := Task{
		ID:       taskIDCounter,
		Desc:     req.Desc,
		Deadline: req.Deadline,
	}

	tasks[taskIDCounter] = task
	taskIDCounter++

	response := CreateTaskResponse{
		ID: task.ID,
	}

	return c.JSON(response)
}

func GetTask(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if task, ok := tasks[id]; !ok {
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.JSON(GetTaskResponse{
			ID:       task.ID,
			Desc:     task.Desc,
			Deadline: task.Deadline,
		})
	}
}

func UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	var req UpdateTaskRequest

	if err := c.BodyParser(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	task, ok := tasks[id]
	if !ok {
		return c.SendStatus(http.StatusNotFound)
	}

	task.Desc = req.Desc
	task.Deadline = req.Deadline
	tasks[id] = task

	return c.SendStatus(http.StatusOK)
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	if _, ok := tasks[id]; !ok {
		return c.SendStatus(http.StatusNotFound)
	} else {
		delete(tasks, id)
		return c.SendStatus(http.StatusOK)
	}
}
