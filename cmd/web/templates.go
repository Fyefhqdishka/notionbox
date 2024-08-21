package main

import "golangify.com/notionbox/pkg/models"

type templateData struct {
	Note  *models.Note
	Notes []*models.Note
}
