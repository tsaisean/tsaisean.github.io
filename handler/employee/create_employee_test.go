package employee

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"iHR/db/models"
	"iHR/db/repositories/mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateEmployeeHandler", func() {
	var (
		router   *gin.Engine
		mockRepo *mocks.EmployeeRepository
		recorder *httptest.ResponseRecorder
	)

	// Shared setup for all tests
	BeforeEach(func() {
		mockRepo = new(mocks.EmployeeRepository)
		handler := NewEmployeeHandler(mockRepo)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.POST("/employee", handler.CreateEmployee)

		recorder = httptest.NewRecorder()
	})

	Context("When the request is valid", func() {
		It("should create an employee and return 201 status", func() {
			// Arrange
			inputEmployee := &models.Employee{FirstName: "John", LastName: "Doe"}
			mockRepo.On("CreateEmployee", inputEmployee).Return(&models.Employee{ID: 1, FirstName: "John", LastName: "Doe"}, nil)

			// Act
			executeRequest(router, inputEmployee, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusCreated))
			var responseBody map[string]interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
			gomega.Expect(err).To(gomega.BeNil())

			gomega.Expect(responseBody["ID"]).To(gomega.Equal(float64(1)))
			gomega.Expect(responseBody["FirstName"]).To(gomega.Equal("John"))
			gomega.Expect(responseBody["LastName"]).To(gomega.Equal("Doe"))

			mockRepo.AssertExpectations(GinkgoT())
		})
	})

	Context("When the request payload is invalid", func() {
		It("should return 400 status for malformed JSON", func() {
			// Arrange
			invalidJSON := `{"FirstName": "John", "LastName":"}`

			// Act
			executeRequest(router, []byte(invalidJSON), recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})

		It("should return 400 status for empty body", func() {
			// Act
			executeRequest(router, []byte{}, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})
	})

	Context("When the request payload is invalid", func() {
		It("should return 400 status for mistype field", func() {
			// Arrange
			invalidJSON := `{"FirstName": "John", "LastName":12222}`

			// Act
			executeRequest(router, []byte(invalidJSON), recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})

		It("should return 400 status for empty body", func() {
			// Act
			executeRequest(router, []byte{}, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})
	})

	Context("When the repository returns an error", func() {
		It("should return 500 status", func() {
			// Arrange
			inputEmployee := &models.Employee{FirstName: "Jane", LastName: "Doe"}
			mockRepo.On("CreateEmployee", inputEmployee).Return(nil, errors.New("repository error"))

			// Act
			executeRequest(router, inputEmployee, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusInternalServerError))

			mockRepo.AssertExpectations(GinkgoT())
		})
	})
})

func executeRequest(router *gin.Engine, body interface{}, recorder *httptest.ResponseRecorder) {
	var requestBody []byte
	var err error

	switch v := body.(type) {
	case []byte:
		requestBody = v
	case *models.Employee:
		requestBody, err = json.Marshal(v)
		if err != nil {
			panic("Failed to marshal request body")
		}
	default:
		panic("Unsupported body type")
	}

	// Create and execute the HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/employee", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
}
