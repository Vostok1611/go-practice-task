package handlers

import (
	"context"
	// УБРАТЬ: "strconv" - больше не нужен!

	taskservice "gomeWork/internal/taskService"
	"gomeWork/internal/web/api"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request api.GetTasksRequestObject) (api.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := api.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := api.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request api.PostTasksRequestObject) (api.PostTasksResponseObject, error) {
	if request.Body == nil || request.Body.Task == "" {
		return nil, echo.NewHTTPError(400, "Task is required")
	}

	isDone := "false"
	if request.Body.IsDone != nil && *request.Body.IsDone != "" {
		isDone = *request.Body.IsDone
	}

	createdTask, err := h.service.CreateTask(request.Body.Task, isDone)
	if err != nil {
		return nil, err
	}

	return api.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request api.PatchTasksIdRequestObject) (api.PatchTasksIdResponseObject, error) {
	// ID теперь string, не нужно конвертировать!
	idStr := request.Id // ← уже string!

	if request.Body == nil || (request.Body.Task == nil && request.Body.IsDone == nil) {
		return nil, echo.NewHTTPError(400, "Request body is required")
	}

	var taskText, isDone string

	if request.Body.Task != nil {
		taskText = *request.Body.Task
	}

	if request.Body.IsDone != nil {
		isDone = *request.Body.IsDone
	}

	currentTask, err := h.service.GetTaskByID(idStr) // ← исправлена переменная
	if err != nil {
		return api.PatchTasksId404Response{}, nil
	}

	if request.Body.Task == nil {
		taskText = currentTask.Task
	}

	if request.Body.IsDone == nil {
		isDone = currentTask.IsDone
	}

	updatedTask, err := h.service.UpdateTask(idStr, taskText, isDone) // ← исправлена переменная
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return api.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request api.DeleteTasksIdRequestObject) (api.DeleteTasksIdResponseObject, error) {
	// ID теперь string, не нужно конвертировать!
	idStr := request.Id // ← уже string!

	err := h.service.DeleteTask(idStr) // ← исправлена переменная
	if err != nil {
		return api.DeleteTasksId404Response{}, nil
	}
	return api.DeleteTasksId204Response{}, nil
}
