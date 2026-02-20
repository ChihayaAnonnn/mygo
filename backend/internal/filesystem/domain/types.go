package domain

import (
	"errors"
	"regexp"
	"strings"
)

// ==================== 领域错误 ====================

var (
	ErrObjectNotFound     = errors.New("object not found")
	ErrObjectDeleted      = errors.New("object deleted")
	ErrInvalidInput       = errors.New("invalid input")
	ErrPreconditionFailed = errors.New("precondition failed")
)

// ==================== 值对象：Namespace ====================

// Namespace 逻辑隔离域（不透明字符串），用于区分不同对象集合。
type Namespace string

var namespaceRe = regexp.MustCompile(`^[a-z0-9][a-z0-9._-]{0,63}$`)

func ParseNamespace(s string) (Namespace, error) {
	if !namespaceRe.MatchString(s) {
		return "", ErrInvalidInput
	}
	return Namespace(s), nil
}

func (n Namespace) String() string { return string(n) }

// ==================== 值对象：Key ====================

// Key 对象键（相对路径风格字符串），由上层 schema 生成。
type Key string

func ParseKey(s string) (Key, error) {
	if s == "" {
		return "", ErrInvalidInput
	}
	if strings.Contains(s, `\`) {
		return "", ErrInvalidInput
	}
	if strings.HasPrefix(s, "/") || strings.HasSuffix(s, "/") {
		return "", ErrInvalidInput
	}
	if strings.Contains(s, "//") {
		return "", ErrInvalidInput
	}
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
		case r >= 'A' && r <= 'Z':
		case r >= '0' && r <= '9':
		case r == '/' || r == '.' || r == '_' || r == '-':
		default:
			return "", ErrInvalidInput
		}
	}

	for p := range strings.SplitSeq(s, "/") {
		if p == "" || p == "." || p == ".." {
			return "", ErrInvalidInput
		}
	}

	return Key(s), nil
}

func (k Key) String() string { return string(k) }
