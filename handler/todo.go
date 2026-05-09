package handler

import (
	"context"
	"fmt"
	"grpc-go-server/grpcapi"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoHandler struct {
	mu    sync.Mutex
	todos map[string]*grpcapi.Todo
	seq   int
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos: make(map[string]*grpcapi.Todo),
	}
}

func (h *TodoHandler) CreateTodo(_ context.Context, req *grpcapi.CreateTodoRequest) (*grpcapi.CreateTodoResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.seq++
	todo := &grpcapi.Todo{
		Id:          fmt.Sprintf("%d", h.seq),
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
	}
	h.todos[todo.Id] = todo

	return &grpcapi.CreateTodoResponse{Todo: todo}, nil
}

func (h *TodoHandler) GetTodo(_ context.Context, req *grpcapi.GetTodoRequest) (*grpcapi.GetTodoResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	todo, ok := h.todos[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %s not found", req.Id)
	}

	return &grpcapi.GetTodoResponse{Todo: todo}, nil
}

func (h *TodoHandler) ListTodos(_ *grpcapi.ListTodosRequest, stream grpcapi.TodoService_ListTodosServer) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, todo := range h.todos {
		if err := stream.Send(&grpcapi.ListTodosResponse{Todo: todo}); err != nil {
			return err
		}
	}

	return nil
}

func (h *TodoHandler) UpdateTodo(_ context.Context, req *grpcapi.UpdateTodoRequest) (*grpcapi.UpdateTodoResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	todo, ok := h.todos[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %s not found", req.Id)
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Done = req.Done

	return &grpcapi.UpdateTodoResponse{Todo: todo}, nil
}

func (h *TodoHandler) DeleteTodo(_ context.Context, req *grpcapi.DeleteTodoRequest) (*grpcapi.DeleteTodoResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.todos[req.Id]; !ok {
		return nil, status.Errorf(codes.NotFound, "todo %s not found", req.Id)
	}

	delete(h.todos, req.Id)

	return &grpcapi.DeleteTodoResponse{Id: req.Id}, nil
}
