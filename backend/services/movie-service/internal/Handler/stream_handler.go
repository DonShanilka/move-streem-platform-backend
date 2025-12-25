package Handler

// import (
// 	"io"
// 	"net/http"
// )

// type StreamHandler struct{}

// func NewStreamHandler() *StreamHandler {
// 	return &StreamHandler{}
// }

// func (h *StreamHandler) StreamMovie(w http.ResponseWriter, r *http.Request) {
// 	videoURL := r.URL.Query().Get("movie")
// 	if videoURL == "" {
// 		http.Error(w, "Missing video URL", http.StatusBadRequest)
// 		return
// 	}

// 	req, err := http.NewRequest("GET", videoURL, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Forward range header (IMPORTANT)
// 	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
// 		req.Header.Set("Range", rangeHeader)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		http.Error(w, "Failed to fetch video", http.StatusBadGateway)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Copy headers
// 	for k, v := range resp.Header {
// 		w.Header()[k] = v
// 	}

// 	w.WriteHeader(resp.StatusCode)
// 	io.Copy(w, resp.Body)
// }
