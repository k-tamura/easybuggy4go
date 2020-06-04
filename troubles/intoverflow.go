package troubles

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostIntOverflow(ctx *gin.Context) {
	description := ""
	strTimes := ctx.PostForm("times")
	if strTimes != "" {
		times, _ := strconv.Atoi(strTimes)
		if times > 0 {
			multipleNumber := 1
			for i := 0; i < times; i++ {
				multipleNumber = multipleNumber * 2
			}
			thickness := float64(multipleNumber) / 10.0 // mm
			thicknessM := thickness / 1000.0            // m
			thicknessKm := thicknessM / 1000.0          // km

			description += strconv.FormatFloat(thickness, 'f', 1, 64) + " mm"
			if thicknessM >= 1 && thicknessKm < 1 {
				description += " = " + strconv.FormatFloat(thicknessM, 'f', 0, 64) + " m"
			}
			if thicknessKm >= 1 {
				description += " = " + strconv.FormatFloat(thicknessKm, 'f', 0, 64) + " km"
			}
		}
	}
	ctx.HTML(200, "intoverflow.html", gin.H{"times": strTimes, "description": description,
		"key_title": "title.intoverflow", "key_note": "note.intoverflow"})
}

func GetIntOverflow(ctx *gin.Context) {
	ctx.HTML(200, "intoverflow.html", gin.H{"times": "", "description": "",
		"key_title": "title.intoverflow", "key_note": "note.intoverflow"})
}
