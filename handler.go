package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Amirhos-esm/portfolio/models"
	"github.com/Amirhos-esm/portfolio/util"
	"github.com/Amirhos-esm/portfolio/views"
	"github.com/Amirhos-esm/portfolio/views/pages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Application) loginHandler(ctx *gin.Context) {
	// type Authentication struct {
	// 	Password string `json:"password"`
	// }
	// // read json payload
	// input := Authentication{}
	// // Parse and validate the JSON input
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	ctx.String(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if input.Password != app.password {
	// 	ctx.String(http.StatusUnauthorized, "")
	// 	return
	// }

	// create a jwt user
	jwt_user := jwtUser{
		ID:        1,
		FirstName: "admin",
		LastName:  "admin",
	}
	tokens, err := app.auth.GenerateTokenPair(&jwt_user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	refreshCokie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(ctx.Writer, refreshCokie)

	ctx.JSON(http.StatusOK, tokens)

}

func (app *Application) LandingPageHandler(ctx *gin.Context) {
	data, err := app.repo.GetAllData()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	views.Render(ctx.Writer, ctx.Request, pages.MainPage(data))
}

func (app *Application) ProjectHandler(ctx *gin.Context) {
	var project *models.Project
	projectId, err := util.GetPathParam[uint](ctx, "id", nil)
	if err != nil {
		p, _ := util.GetPathParam[string](ctx, "id", nil)
		if p == "" {
			ctx.String(http.StatusNotFound, "")
			return
		}
		projects, err := app.repo.GetAllProjects()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		for _, val := range projects {
			if val.Title == p {
				project = val
				goto fetch
			}
		}
		ctx.String(http.StatusNotFound, "")
		return
	}

	project, err = app.repo.GetProject(projectId)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
fetch:
	if project == nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}
	views.Render(ctx.Writer, ctx.Request, pages.Project(project))
}
func (app *Application) addProjectGallery(ctx *gin.Context) {
	// 1. Get project ID
	projectId, err := util.GetPathParam[uint](ctx, "id", nil)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	project, err := app.repo.GetProject(projectId)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if project == nil {
		ctx.String(http.StatusNotFound, "")
		return
	}

	// Expect exactly ONE file with key "file"
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// Optional: size limit (e.g. 500kB)
	if file.Size > 1000*500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	// Optional: mime type check
	contentType := file.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	name, err := uuid.NewRandom()
	if err != nil {
		return
	}
	ext := filepath.Ext(file.Filename)
	if len(ext) <= 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad file name"})
		return
	}
	filename := fmt.Sprintf(
		// "%s_%d_%s",
		// projectID,
		// time.Now().UnixNano(),
		// file.Filename,
		"%s%s",
		name,
		ext,
	)

	dst := filepath.Join(app.staticFolder, "project", filename)

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create directory"})
		return
	}

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	project.AddGalleryImage(filename)

	if err := app.repo.UpdateProject(projectId, project); err != nil {
		os.Remove(dst)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file uploaded",
		"path":    dst,
	})
}
func (app *Application) resumeUploadHandler(ctx *gin.Context) {

	// Expect exactly ONE file with key "file"
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	// Optional: size limit (e.g. 3.5MB)
	if file.Size > 1000*3500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	// Optional: mime type check
	contentType := file.Header.Get("Content-Type")
	if contentType != "application/pdf" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	dst := filepath.Join(app.staticFolder, "resume.pdf")

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create directory"})
		return
	}

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file uploaded",
		"path":    dst,
	})
}

func (app *Application) profileUploadHandler(ctx *gin.Context) {

	// Expect exactly ONE file with key "file"
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	// Optional: size limit (e.g. 3.5MB)
	if file.Size > 1000*3500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	contentType := file.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}
	ext := filepath.Ext(file.Filename)
	if len(ext) <= 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad file name"})
		return
	}
	dst := filepath.Join(app.staticFolder, "profile"+ext)

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create directory"})
		return
	}

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file uploaded",
		"path":    dst,
	})
}
