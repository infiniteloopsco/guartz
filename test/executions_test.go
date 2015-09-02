package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/infiniteloopsco/guartz/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Execution", func() {

	BeforeEach(func() {
		task = createTask()
	})

	Describe("POST /tasks/:task_id/executions", func() {

		It("create an execution", func() {
			execution := models.Execution{}
			executionJSON, _ := json.Marshal(execution)
			client.CallRequest("POST", fmt.Sprintf("/tasks/%s/executions", task.ID), bytes.NewReader(executionJSON)).WithResponse(func(resp *http.Response) error {
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				return nil
			})
		})

	})

	Context("after an execution is created", func() {

		BeforeEach(func() {
			execution = createExecution(task)
		})

		Describe("GET /tasks/:task_id/executions", func() {

			It("create an execution", func() {
				var executionsResp []models.Execution
				client.CallRequestNoBody("GET", fmt.Sprintf("/tasks/%s/executions", task.ID)).WithResponseJSON(&executionsResp, func(resp *http.Response) error {
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
					Expect(len(executionsResp)).To(BeEquivalentTo(1))
					return nil
				})
			})

		})

	})

})
