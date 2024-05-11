package example

import exampleServer "server-fiber/service/example"

type FileUploadAndDownloadApi struct{}
type ExcelApi struct{}

var excelService = new(exampleServer.ExcelService)

type CustomerApi struct{}

var customerService = new(exampleServer.CustomerService)
var fileUploadAndDownloadService = new(exampleServer.FileUploadAndDownloadService)
