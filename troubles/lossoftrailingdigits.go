package troubles

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLossOfTrailingDigits(ctx *gin.Context) {
	ctx.HTML(200, "lossoftrailingdigits.html", gin.H{"number": "", "result": "",
		"key_title": "title.lossoftrailingdigits", "key_note": "note.lossoftrailingdigits"})
}

func PostLossOfTrailingDigits(ctx *gin.Context) {
	strNumber := ctx.PostForm("number")
	number, err := strconv.ParseFloat(strNumber, 64)
	if err == nil && -1 < number && number < 1 {
		ctx.HTML(200, "lossoftrailingdigits.html", gin.H{"number": strNumber, "result": number + 1,
			"key_title": "title.lossoftrailingdigits", "key_note": "note.lossoftrailingdigits"})
	} else {
		GetLossOfTrailingDigits(ctx)
	}
}
