package controllers

import "github.com/gofiber/fiber"

func AgentPhotoUpload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""
	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/agents/"+filename); err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/agents/" + filename,
	})
}
