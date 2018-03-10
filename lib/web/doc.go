// Convention:
//	- each method of a handler is a GET handler
//	- if not, the method (POST, PUT, DELETE, ...) must be added in the method name!
//		- DeletePOST
//		- EditPUT
//		- ...
//

// 	- if possible use the parsing middleware funcs, parseID, parseForm, plainTemplate
//		- and ajust the proper func signature
//		- parseForm: func(http.ResponseWriter, *http.Request)
//		- plainTemplate: func(templatefilename, *Renderer)
//		- parseID: func(http.ResponseWriter, id int)
package web
