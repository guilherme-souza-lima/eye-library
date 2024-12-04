package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: createpkg <package_name>")
		return
	}
	packageName := os.Args[1]
	baseDir := fmt.Sprintf("./internal/%s", packageName)

	files := map[string]string{
		"handler.go": fmt.Sprintf(`package %s

// Handler for the %s package
type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(s ServiceInterface) *Handler {
	return &Handler{s}
}

func (h *Handler) HandleExample() error {
	return nil
}
`, packageName, packageName),

		"service.go": fmt.Sprintf(`package %s

type ServiceInterface interface {
	ServiceExample() string
}

// Service for the %s package
type Service struct {
	RepositoryInterface RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{r}
}

func (s *Service) ServiceExample() string {
	return "Hello from the service!"
}`, packageName, packageName),

		"repository.go": fmt.Sprintf(`package %s

type RepositoryInterface interface {
	RepositoryExample() error
}

// Repository for the %s package
type Repository struct {}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) RepositoryExample() error {
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
		return
	}

	for fileName, content := range files {
		filePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			return
		}
	}

	fmt.Printf("Package %s created successfully at %s\n", packageName, baseDir)
}

func CreatePackage(packageName string) error {
	baseDir := fmt.Sprintf("./internal/%s", packageName)

	files := map[string]string{
		"handler.go": fmt.Sprintf(`package %s

// Handler for the %s package
type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(s ServiceInterface) *Handler {
	return &Handler{s}
}

func (h *Handler) HandleExample() error {
	return nil
}
`, packageName, packageName),

		"service.go": fmt.Sprintf(`package %s

type ServiceInterface interface {
	ServiceExample() string
}

// Service for the %s package
type Service struct {
	RepositoryInterface RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{r}
}

func (s *Service) ServiceExample() string {
	return "Hello from the service!"
}`, packageName, packageName),

		"repository.go": fmt.Sprintf(`package %s

type RepositoryInterface interface {
	RepositoryExample() error
}

// Repository for the %s package
type Repository struct {}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) RepositoryExample() error {
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
