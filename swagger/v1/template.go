package v1

import "template/main/models/template"

// swagger:parameters GetTemplate
type GetTemplateRequest struct {
	// in: query
	template.Template
}

// 取得當前Template
//
// swagger:response GetTemplate
type GetTemplateResponse struct {
	// in: body
	Body []template.Template
}

// swagger:route GET /api/v1/template 設定檔 GetTemplate
//
// 取得Template
//
// 取得Template
//
//     Responses:
//       200: GetTemplate

// swagger:parameters PostTemplate
type PostTemplateRequest struct {
	// in: body
	Body template.Template
}

// 新增Template
//
// swagger:response PostTemplate
type PostTemplateResponse struct {
	// in: body
	Body template.Template
}

// swagger:route POST /api/v1/template 設定檔 PostTemplate
//
// 新增Template
//
// 新增Template
//
//     Responses:
//       200: PostTemplate

// swagger:parameters PutTemplate
type PutTemplateRequest struct {
	// in: body
	Body template.Template
}

// 更新Template
//
// swagger:response PutTemplate
type PutTemplateResponse struct {
	// in: body
	Body template.Template
}

// swagger:route PUT /api/v1/template 設定檔 PutTemplate
//
// 更新Template
//
// 更新Template
//
//     Responses:
//       200: PutTemplate

// swagger:parameters DeleteTemplate
type DeleteTemplateRequest struct {
	// in: path
	Id int `json:"id"`
}

// 更新Template
//
// swagger:response DeleteTemplate
type DeleteTemplateResponse struct {
	// in: body
	Body template.Template
}

// swagger:route Delete /api/v1/template/{id} 設定檔 DeleteTemplate
//
// 刪除Template
//
// 刪除Template
//
//     Responses:
//       200: DeleteTemplate
