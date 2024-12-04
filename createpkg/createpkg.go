package createpkg

import (
	"fmt"
	"os"
)

func CreatePackage(packageName string) error {
	baseDir := fmt.Sprintf("./internal/%s", packageName)

	files := map[string]string{
		"handler.go": fmt.Sprintf(`package %s
import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler for the %s package
type Handler struct {
	s ServiceInterface
}

func NewHandler(s ServiceInterface) *Handler {
	return &Handler{s}
}

func (h *Handler) ExampleHandler(c echo.Context) error {
	err := h.service.ExampleService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "ok")
}`, packageName, packageName),

		"service.go": fmt.Sprintf(`package %s

import "context"

type ServiceInterface interface {
	ExampleService(context.Context) error
}

// Service for the %s package
type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo}
}

func (s *Service) ExampleService(ctx context.Context) error {
	return s.repo.ExampleRepository(ctx)
}`, packageName, packageName),

		"repository.go": fmt.Sprintf(`package %s

import "context"

type RepositoryInterface interface {
	ExampleRepository(context.Context) error
}

// Repository for the %s package
type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) ExampleRepository(ctx context.Context) error {
	return nil
}`, packageName, packageName),

		"model.go": fmt.Sprintf(`package %s

// Model represents an example data structure
type Model struct {
	ID   int
	Name string
}`, packageName),
	}

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return err
	}

	for fileName, content := range files {
		filePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			return err
		}
	}

	fmt.Printf("Package %s created successfully at %s\n", packageName, baseDir)

	return nil
}
