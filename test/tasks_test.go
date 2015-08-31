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
			resp, _ := client.CallRequest("POST", "/tasks", bytes.NewReader(taskJSON))
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			var taskResp models.Task
			defer resp.Body.Close()
			getBodyJSON(resp, &taskResp)
			Expect(taskResp.ID).NotTo(BeEmpty())
		})

	})

	Context("After the task is created", func() {

		BeforeEach(func() {
			task = createTask()
		})

		Describe("GET /tasks", func() {

			It("gets a list with one element", func() {
				resp, _ := client.CallRequestNoBody("GET", "/tasks")
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				defer resp.Body.Close()
				var tasksResp []models.Task
				getBodyJSON(resp, &tasksResp)
				Expect(len(tasksResp)).To(BeEquivalentTo(1))
			})

		})

		Describe("GET /tasks/:id", func() {

			It("gets a task by id", func() {
				resp, _ := client.CallRequestNoBody("GET", "/tasks/"+task.ID)
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				defer resp.Body.Close()
				var taskResp models.Task
				getBodyJSON(resp, &taskResp)
				Expect(taskResp.Command).To(BeEquivalentTo(task.Command))
			})

		})

		Describe("DELETE /tasks/:id", func() {

			It("deletes a task by id", func() {
				resp, _ := client.CallRequestNoBody("DELETE", "/tasks/"+task.ID)
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
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
				resp, _ := client.CallRequest("POST", "/tasks", bytes.NewReader(taskJSON))
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})

		})

	})

})
