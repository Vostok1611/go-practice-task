package handlers

import (
	"context"

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
	var allTasks []taskservice.Task
	var err error

	if request.Params.UserId != nil && *request.Params.UserId != "" {
		allTasks, err = h.service.GetTasksByUserID(*request.Params.UserId)
	} else {
		allTasks, err = h.service.GetAllTasks()
	}
	if err != nil {
		return nil, err
	}

	response := api.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := api.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request api.PostTasksRequestObject) (api.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(400, "Task is required")
	}

	if request.Body.Task == "" {
		return nil, echo.NewHTTPError(400, "Task is required")
	}
	if request.Body.UserId == "" {
		return nil, echo.NewHTTPError(400, "User is required")
	}

	isDone := "false"
	if request.Body.IsDone != nil && *request.Body.IsDone != "" {
		isDone = *request.Body.IsDone
	}

	createdTask, err := h.service.CreateTask(request.Body.Task, isDone, request.Body.UserId)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return api.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request api.PatchTasksIdRequestObject) (api.PatchTasksIdResponseObject, error) {
	if request.Body == nil || (request.Body.Task == nil && request.Body.IsDone == nil && request.Body.UserId == nil) {
		return nil, echo.NewHTTPError(400, "At least one field (task, is_done, or user_id) must be provided for update")
	}

	currentTask, err := h.service.GetTaskByID(request.Id) // ← исправлена переменная
	if err != nil {
		return api.PatchTasksId404Response{}, nil
	}

	taskText := currentTask.Task
	isDone := currentTask.IsDone
	userID := currentTask.UserID

	if request.Body.Task != nil {
		taskText = *request.Body.Task
	}

	if request.Body.IsDone != nil {
		isDone = *request.Body.IsDone
	}

	if request.Body.UserId != nil {
		userID = *request.Body.UserId
	}

	updatedTask, err := h.service.UpdateTask(request.Id, taskText, isDone, userID) // ← исправлена переменная
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return api.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
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
