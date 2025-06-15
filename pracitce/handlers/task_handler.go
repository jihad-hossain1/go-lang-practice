package handlers

import (
	"go-todo-fiber/models"
	"go-todo-fiber/utils"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var tasks []models.Task
var dataFile = "tasks.json"

func LoadTasks() {
	utils.LoadFromFile(dataFile, &tasks)
}

func saveTasks() {
	utils.SaveToFile(dataFile, tasks)
}

func GetTasks(c *fiber.Ctx) error {
	search := c.Query("search")
	sortOrder := c.Query("sort", "asc")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// filter
	filtered := []models.Task{}

	for _, t := range tasks {
		if search == "" || utils.ContainsIgnoreCase(t.Title, search) {
			filtered = append(filtered, t)
		}
	}

	// sort
	if sortOrder == "desc" {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].Title > filtered[j].Title
		})
	} else {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].Title < filtered[j].Title
		})
	}

	// Paginate
	start := (page - 1) * limit
	end := start + limit
	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	paged := filtered[start:end]

	return c.JSON(fiber.Map{
		"data":  paged,
		"page":  page,
		"limit": limit,
		"total": len(filtered),
	})
}

func AddTask(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid body")
	}
	newTask := models.Task{
		ID:    len(tasks) + 1,
		Title: body.Title,
		Done:  false,
	}
	tasks = append(tasks, newTask)
	saveTasks()
	return c.JSON(newTask)
}

func MarkTaskDone(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			saveTasks()
			return c.JSON(tasks[i])
		}
	}
	return c.Status(404).SendString("Task not found")
}
