package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func main() { 
	
	filePath := flag.String("file", "", "Path to the PDF file to split")
	every := flag.Int("every", 1, "Number of pages in each split PDF")
	outputPath := flag.String("output" , "", "Path to save the new pdfs")


    // Parse the flags
    flag.Parse()

	if *filePath == "" {
		fmt.Println("You must specify a file path using the --file flag.")
        return
	}

	if *outputPath == "" {
		fmt.Println("You must specify output to store your new files using --output flag.")
        return
	}
	
	if *every < 5 {
		fmt.Println("Every should be greater than 5")
		return;
	}

	err := validateDocument(*filePath) 

	if err != nil {
		fmt.Println("Error :" ,err)
	}

	licenseErr := license.SetMeteredKey("efb06ab3d5a559b529caff0482d5b148b41dbb3f2c6463a9fa1458e4e32321c4")

	if licenseErr != nil{
		fmt.Printf("Error setting license key: %v\n", err)
	}

	splitErr := splitPdf(filePath,outputPath,every)
	
	if splitErr != nil {
		fmt.Println(splitErr)
	}
	
}

func validateDocument(filePath string) error {

	if _,err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("File doesn't exists in the provided path : " + filePath)
	}
	return nil
}

func splitPdf(inputPath *string ,outputPath *string, every *int) error {
	file,err := os.Open(*inputPath)

	if err != nil {
		return err
	}

	defer file.Close(); // this will executed at the end

	pdfReader, err := model.NewPdfReader(file) // get pdf reader object of this file

	if err != nil {
		return err
	}

	pagesNumber,err := pdfReader.GetNumPages()

	fmt.Println("The number of the book pages is " , pagesNumber)

	if err != nil {
		return err
	}

	//loop through pages of pdf and get the every pages and create new pdf
	pageCount := 0

    c := creator.New()
	resetor := 1;
	//1360 
	for i := 1 ; i <= pagesNumber ; i++ {
		page, err := pdfReader.GetPage(i)

		if err != nil {
			return err
		}
		//add to new pdf the page
		c.AddPage(page)
		pageCount++

		if *every == pageCount || i == pagesNumber  {
			//save to file with the numbered 1,2,3
			newFilename := fmt.Sprintf("kumar_7h_edition_medical_%d.pdf",resetor)
			newPath := fmt.Sprintf("%v/%v" , *outputPath,newFilename)
			fmt.Println(newPath)
		

			resetor++
			
			err := c.WriteToFile(newPath)

			if err != nil {
				return err
			}
			fmt.Println("File Added")
			pageCount = 0;
			c = creator.New();
		}
		
	}

	return nil
}