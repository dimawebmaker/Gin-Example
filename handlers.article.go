package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  article := getAllArticles()

  // Call the render function with the name of the template to render
  render(c, gin.H{
    "title": "Home Page",
    "payload": articles}, "index.html")
}

func show Article()

