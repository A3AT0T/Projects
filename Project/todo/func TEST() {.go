func TEST() {
	// Register a handler function for the "/user" endpoint with POST method
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Process the request body (in this example, we'll just echo it)
		fmt.Fprintf(w, "Received POST request with data: %s\n", string(body))
	})

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println(err)
	}
}