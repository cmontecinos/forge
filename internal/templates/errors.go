package templates

import "fmt"

// TemplateError wraps errors with template operation context
type TemplateError struct {
	Op  string // Operation that failed (clone, process, cleanup)
	Err error  // Underlying error
}

func (e *TemplateError) Error() string {
	return fmt.Sprintf("template %s failed: %v", e.Op, e.Err)
}

func (e *TemplateError) Unwrap() error {
	return e.Err
}

// WrapTemplateError wraps an error with operation context
func WrapTemplateError(op string, err error) error {
	if err == nil {
		return nil
	}
	return &TemplateError{
		Op:  op,
		Err: err,
	}
}

// Common error types
var (
	ErrCloneFailed   = fmt.Errorf("failed to clone template repository")
	ErrProcessFailed = fmt.Errorf("failed to process template files")
	ErrCleanupFailed = fmt.Errorf("failed to cleanup git directory")
)

// IsCloneError checks if the error is a clone operation error
func IsCloneError(err error) bool {
	if te, ok := err.(*TemplateError); ok {
		return te.Op == "clone"
	}
	return false
}

// IsProcessError checks if the error is a process operation error
func IsProcessError(err error) bool {
	if te, ok := err.(*TemplateError); ok {
		return te.Op == "process"
	}
	return false
}

// IsCleanupError checks if the error is a cleanup operation error
func IsCleanupError(err error) bool {
	if te, ok := err.(*TemplateError); ok {
		return te.Op == "cleanup"
	}
	return false
}
