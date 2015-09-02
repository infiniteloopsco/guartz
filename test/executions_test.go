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
			resp, _ := client.CallRequest("POST", fmt.Sprintf("/tasks/%s/executions", task.ID), bytes.NewReader(executionJSON))
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

	})

	Context("after an execution is created", func() {

		BeforeEach(func() {
			execution = createExecution(task)
		})

		Describe("GET /tasks/:task_id/executions", func() {

			It("create an execution", func() {
				resp, _ := client.CallRequestNoBody("GET", fmt.Sprintf("/tasks/%s/executions", task.ID))
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				var executionsResp []models.Execution
				defer resp.Body.Close()
				getBodyJSON(resp, &executionsResp)
				Expect(len(executionsResp)).To(BeEquivalentTo(1))
			})

		})

	})

})
