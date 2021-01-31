package util

import (
	"errors"
	"github.com/golang/freetype/truetype"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/golang/freetype"
)

const (
	top           = 400
	left          = 200
	interval      = 140
	baseFontSize  = 35
	backgroundURL = "./static/normal/background.png"
	sealURL       = "./static/normal/seal.png"
	fontURL       = "./static/normal/华文楷体.ttf"
	SaveURL       = "./static/certificate/"
)

type PictureInfo struct {
	CreatedAt         int64  `json:"created_at"`
	Name              string `json:"name"`
	Sex               uint   `json:"sex"`                // 0 男 1女
	Ethnic            string `json:"ethnic"`             // 民族
	Birthday          int64  `json:"birthday"`           // 生日
	CardCode          string `json:"card_code"`          // 身份证号
	EducationCategory string `json:"education_category"` // 学历类别
	Level             string `json:"level"`              // 层次
	University        string `json:"university"`         // 学校
	Professional      string `json:"professional"`       // 专业
	LearningFormat    string `json:"learning_format"`    // 学习形式
	EducationalSystem string `json:"educational_system"` // 学制
	AdmissionDate     string `json:"admission_date"`     // 入学日期
	GraduationDate    string `json:"graduation_date"`    // 毕业日期
	Status            string `json:"status"`             // 状态（是否结业）
	StudentAvatar     string `json:"student_avatar"`     // 照片
}

var backgroundImg image.Image
var sealImg image.Image
var font *truetype.Font

func getBackgroundImg() image.Image {
	if backgroundImg == nil {
		imgfile, _ := os.Open(backgroundURL)
		defer imgfile.Close()

		backgroundImg, _ = png.Decode(imgfile)
	}
	return backgroundImg
}
func getSealImg() image.Image {
	if sealImg == nil {
		seal, _ := os.Open(sealURL)
		sealimg, _ := png.Decode(seal)
		sealImg = resize.Resize(440, 0, sealimg, resize.Lanczos3)
		defer seal.Close()
	}
	return sealImg
}
func getFont() *truetype.Font {
	if font == nil {
		fontBytes, _ := ioutil.ReadFile(fontURL)
		font, _ = freetype.ParseFont(fontBytes)
	}
	return font
}
func (picInfo *PictureInfo) createImg() (string, error) {
	var err error
	pngimg := getBackgroundImg() //背景图
	font := getFont()            //字体文件
	sealimg := getSealImg()      //印章图

	img := image.NewNRGBA(pngimg.Bounds()) //
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, pngimg.At(x, y))
		}
	}

	drawHead(img, font)
	drawText(img, font, baseFontSize, left, top+0*interval, "姓        名："+picInfo.Name)
	drawText(img, font, baseFontSize, left, top+1*interval, "性        别："+IntToSex(picInfo.Sex))
	drawText(img, font, baseFontSize, left, top+2*interval, "民        族："+picInfo.Ethnic)
	drawText(img, font, baseFontSize, left, top+3*interval, "出生日期："+Int64ToStr(picInfo.Birthday))
	drawText(img, font, baseFontSize, left, top+4*interval, "身份证号："+picInfo.CardCode)
	drawText(img, font, baseFontSize, left, top+5*interval, "学历类别："+picInfo.EducationCategory)
	drawText(img, font, baseFontSize, left, top+6*interval, "层        次："+picInfo.Level)
	drawText(img, font, baseFontSize, left, top+7*interval, "专        业："+picInfo.Professional)
	drawText(img, font, baseFontSize, left, top+8*interval, "学习形式："+picInfo.LearningFormat)
	drawText(img, font, baseFontSize, left, top+9*interval, "学        制："+picInfo.EducationalSystem)
	drawText(img, font, baseFontSize, left, top+10*interval, "毕业院校："+picInfo.University)
	drawText(img, font, baseFontSize, left, top+11*interval, "入学日期："+picInfo.AdmissionDate)
	drawText(img, font, baseFontSize, left, top+12*interval, "毕业日期："+picInfo.GraduationDate)
	drawText(img, font, baseFontSize, left, top+13*interval, "状        态："+picInfo.Status)
	drawText(img, font, baseFontSize, left, top+14*interval, "报告生成日期："+Int64ToStr(time.Now().Unix()))
	drawText(img, font, baseFontSize+5, left, top+15*interval, "以上学历情况属实，专此认证")
	drawText(img, font, baseFontSize+7, 1550, 3150, "DBAS学历认证系统")

	// 图片处理
	tou, _ := os.Open("./static/file/" + picInfo.StudentAvatar)
	defer tou.Close()
	var touimg image.Image
	if path.Ext(picInfo.StudentAvatar) == ".png" {
		touimg, err = png.Decode(tou)
		if err != nil {
			return "", err
		}
	} else if path.Ext(picInfo.StudentAvatar) == ".jpg" {
		touimg, err = jpeg.Decode(tou)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("头像图片错误")
	}

	touimg = resize.Resize(380, 0, touimg, resize.Lanczos3)

	draw.Draw(img, img.Bounds(), touimg, image.Point{X: -1730, Y: -350}, draw.Over)

	draw.Draw(img, img.Bounds(), sealimg, image.Point{X: -1700, Y: -2800}, draw.Over)
	//保存到新文件中
	newFileName := uuid.Must(uuid.NewRandom()).String() + "jpg"
	newfile, _ := os.Create(SaveURL + newFileName)
	defer newfile.Close()

	err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 80})
	if err != nil {
		return "", err
	}
	return newFileName, err
}

func drawHead(img *image.NRGBA, font *truetype.Font) error {
	f := freetype.NewContext()
	f.SetDPI(144)
	f.SetFont(font)
	f.SetFontSize(55)
	f.SetClip(img.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))

	pt := freetype.Pt(900, 250)
	_, err := f.DrawString("学历认证报告", pt)
	if err != nil {
		return err
	}
	return nil
}

func drawText(img *image.NRGBA, font *truetype.Font, fontSize float64, x int, y int, text string) error {
	f := freetype.NewContext()
	f.SetDPI(144)
	f.SetFont(font)
	f.SetFontSize(fontSize)
	f.SetClip(img.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))

	pt := freetype.Pt(x, y)
	_, err := f.DrawString(text, pt)
	if err != nil {
		return err
	}
	return nil
}
