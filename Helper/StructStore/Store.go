package StructStore

type URLResponse struct {
	OriginalURL   string
	NewURL        string
	IsAnyError    bool
	ErrorMessages []string
}

type URLRequest struct {
	URL string
}

type DbNumericResponse struct {
	Count int
}

type DbNumericResponse_Adv struct {
	Count        int
	ShortenedURL string
}
