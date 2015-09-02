package test

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/infiniteloopsco/guartz/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tasks", func() {

	Describe("POST /tasks", func() {

		It("creates a task", func() {
			task = models.Task{
				Periodicity: "@every 2s",
				Command:     "curl -X POST --data payload={\"channel\":\"#general\",\"text\":\"EOOO\"} https://hooks.slack.com/services/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",
			}
			taskJSON, _ := json.Marshal(task)
			var taskResp models.Task
			client.CallRequest("POST", "/tasks", bytes.NewReader(taskJSON)).WithResponseJSON(taskResp, func(resp *http.Response) error {
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				Expect(taskResp.ID).NotTo(BeEmpty())
				return nil
			})
		})

	})

	Context("After the task is created", func() {

		BeforeEach(func() {
			task = createTask()
		})

		Describe("GET /tasks", func() {

			It("gets a list with one element", func() {
				var tasksResp []models.Task
				client.CallRequestNoBody("GET", "/tasks").WithResponseJSON(tasksResp, func(resp *http.Response) error {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					Expect(len(tasksResp)).To(BeEquivalentTo(1))
					return nil
				})
			})

		})

		Describe("GET /tasks/:id", func() {

			It("gets a task by id", func() {
				var taskResp models.Task
				client.CallRequestNoBody("GET", "/tasks/"+task.ID).WithResponseJSON(taskResp, func(resp *http.Response) error {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					Expect(taskResp.Command).To(BeEquivalentTo(task.Command))
					return nil
				})
			})

		})

		Describe("DELETE /tasks/:id", func() {

			It("deletes a task by id", func() {
				client.CallRequestNoBody("DELETE", "/tasks/"+task.ID).WithResponse(func(resp *http.Response) error {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					return nil
				})
			})

		})

		Describe("POST /tasks", func() {

			It("reschedule the task", func() {
				updateTask := models.Task{
					ID:          task.ID,
					Periodicity: "@every 1m",
					Command:     "curl -X POST --data payload={\"channel\":\"#general\",\"text\":\"EOOO\"} https://hooks.slack.com/services/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",
				}
				taskJSON, _ := json.Marshal(updateTask)
				client.CallRequest("POST", "/tasks", bytes.NewReader(taskJSON)).WithResponse(func(resp *http.Response) error {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					return nil
				})
			})

		})

	})

})
