package request

type CreateProjectRequest struct {
	Name string `binding:"required"`
	Url  string `binding:"optional"`
}