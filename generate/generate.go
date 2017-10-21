package generate

type Render interface {
	// Single 渲染单个内容：文章/单页面
	// @param src 源markdown 路径
	// @return output html
	// 		   error -
	Single(src []byte) (output []byte, err error)



}
