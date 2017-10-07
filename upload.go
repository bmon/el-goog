package main

import (
	"fmt"
	//"mime/multipart"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	/*
	 * Dont make the same mistake i did:
	 * Documentation for fine uploader is on their website
	 * but not in Docs, in API.
	 *
	 */

	file, header, err := r.FormFile("qqfile")

	// TODO appropriately utilise these
	fmt.Println(file)   // multipart.File, which happens to point to an interface?
	fmt.Println(header) // also has the body
	fmt.Println(err)

	if err == nil {
		fmt.Fprintf(w, "{\"success\":true}")
	} else {
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
}
