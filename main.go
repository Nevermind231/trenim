package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = make(map[int]Task)
var nextID = 1
var mu sync.Mutex

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{
		"error": msg,
	})
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		mu.Lock()
		list := make([]Task, 0, len(tasks))
		for _, task := range tasks {
			list = append(list, task)
		}
		mu.Unlock()
		writeJSON(w, http.StatusOK, list)

	case http.MethodPost:
		var input struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}

		mu.Lock()
		task := Task{
			ID:        nextID,
			Title:     input.Title,
			Completed: false,
		}
		tasks[nextID] = task
		nextID++
		mu.Unlock()

		writeJSON(w, http.StatusCreated, task)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	mu.Lock()
	task, ok := tasks[id]
	mu.Unlock()
	if !ok {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	switch r.Method {

	case http.MethodGet:
		writeJSON(w, http.StatusOK, task)

	case http.MethodPut:
		var input struct {
			Completed bool `json:"completed"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		mu.Lock()
		task.Completed = input.Completed
		tasks[id] = task
		mu.Unlock()

		writeJSON(w, http.StatusOK, task)

	case http.MethodDelete:
		mu.Lock()
		delete(tasks, id)
		mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]string{
			"status": "deleted",
		})

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	http.ListenAndServe(":8080", nil)
}
