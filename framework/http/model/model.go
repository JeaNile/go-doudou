package model

import (
	"net/http"
)

type Annotation struct {
	Name   string
	Params []string
}

type AnnotationStore map[string][]Annotation

func (receiver AnnotationStore) HasAnnotation(key string, annotationName string) bool {
	for _, item := range receiver[key] {
		if item.Name == annotationName {
			return true
		}
	}
	return false
}

func (receiver AnnotationStore) GetParams(key string, annotationName string) []string {
	for _, item := range receiver[key] {
		if item.Name == annotationName {
			return item.Params
		}
	}
	return nil
}

// Route wraps config for route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
